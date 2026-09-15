package handlers

import (
	"archive/zip"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"sirekap/internal/database"
)

// BAST SE2026 SENGAJA custom (bukan pakai alur generik spk_bast.go / tabel `surat`):
// mitra SE2026 nomor SPK-nya di-import langsung ke rekap.id_spk (bukan lewat
// models.CreateSurat), jadi alur generik yang butuh baris `surat` stujuan=SPK gagal
// total. Nomornya juga beda format ("BAST-SE2026", bukan "BAST.P") & beda urutan
// (PPL 1-233 dulu baru PML 234-265, sama seperti BAPP II - lihat seq_all_termin2).
// BAST = SATU dokumen per orang utk SELURUH kontrak (Termin 1+2 gabung, "gabungan
// total" arahan user) - makanya kolom rekap.id_bast_se2026/tgl_bast_se2026 terpisah
// dari id_bapp1/id_bapp2 (yang per-termin), dan data target/realisasi/wilayah kerja
// diambil dari lk_ppk_termin2_se2026 (state final/terkini per petugas).

type bastSE2026Data struct {
	Nomor          string
	NomorSPK       string
	NamaPetugas    string
	NIKPetugas     string
	TargetSLS      int
	RealisasiSLS   int
	HonorAngka     string
	HonorTerbilang string
}

type wilayahKerjaRow struct {
	Nama      string // kosong utk PPL (baris milik diri sendiri) - diisi nama PPL binaan utk PML
	KodeKec   string
	NamaKec   string
	KodeDesa  string
	NamaDesa  string
	JumlahSLS int
}

// fmtKodeNama gabungkan kode+nama wilayah spt "[050] Manggelewa". Kalau nama
// tidak ditemukan (mis. tidak ada padanan di se2026.sls) tampilkan kode saja.
func fmtKodeNama(kode, nama string) string {
	if nama == "" {
		return kode
	}
	return fmt.Sprintf("[%s] %s", kode, nama)
}

// buildBASTSE2026 mengambil data petugas dari rekap + hitung target/realisasi/wilayah
// kerja dari lk_ppk_termin2_se2026 (state final per petugas, gabungan seluruh SLS yg
// pernah/masih ditangani s.d. akhir kontrak).
func buildBASTSE2026(rekapID int, tanggal string) (bastSE2026Data, string, []wilayahKerjaRow, error) {
	var d bastSE2026Data
	var idsobat, honorStr, idSpk, kegiatan, tahun string

	row := database.DB.QueryRow(`
		SELECT r.idsobat, r.namamitra, r.honor, COALESCE(r.id_spk,''), r.kegiatan, r.tahun
		FROM rekap r WHERE r.id = ?`, rekapID)
	if err := row.Scan(&idsobat, &d.NamaPetugas, &honorStr, &idSpk, &kegiatan, &tahun); err != nil {
		return d, "", nil, fmt.Errorf("rekap id %d tidak ditemukan", rekapID)
	}
	d.NamaPetugas = properCase(d.NamaPetugas)
	if idSpk == "" {
		return d, "", nil, fmt.Errorf("petugas ini belum memiliki nomor SPK")
	}
	jenis := jenisPetugasSE2026(kegiatan)
	isPML := jenis == "pml"

	_, seqAll, _, ok := getPayableSeq(idsobat, isPML, 2)
	if !ok {
		return d, jenis, nil, fmt.Errorf("'%s' tidak ada di daftar yang bisa dibayarkan termin 2 (BAST digabung dgn termin 2)", d.NamaPetugas)
	}
	_, _, targetSLS, realisasiSLS, err := computeUsahaKeluargaSE2026(idsobat, isPML, 2)
	if err != nil {
		return d, jenis, nil, fmt.Errorf("gagal menghitung SLS: %v", err)
	}
	d.TargetSLS, d.RealisasiSLS = targetSLS, realisasiSLS

	if tanggal == "" {
		tanggal = batchTanggalDefaultTermin2
	}
	tgl, err := time.Parse("2006-01-02", tanggal)
	if err != nil {
		tgl, _ = time.Parse("2006-01-02", batchTanggalDefaultTermin2)
	}

	nomor, err := suratNomorSE2026(seqAll, "bast", 2, tahun, tgl)
	if err != nil {
		return d, jenis, nil, err
	}
	d.Nomor = nomor
	d.NomorSPK = idSpk

	database.DB.QueryRow(`SELECT nik FROM mitra WHERE idsobat = ? AND nik IS NOT NULL AND nik != '' LIMIT 1`, idsobat).Scan(&d.NIKPetugas)

	// Nilai Perjanjian BAST tetap (flat) sesuai honor SPK per peran, BUKAN
	// honor per-orang di rekap.honor (arahan user - lihat template asli).
	_ = honorStr
	if isPML {
		d.HonorAngka = "12.192.500"
		d.HonorTerbilang = "Dua Belas Juta Seratus Sembilan Puluh Dua Ribu Lima Ratus Rupiah"
	} else {
		d.HonorAngka = "11.572.500"
		d.HonorTerbilang = "Sebelas Juta Lima Ratus Tujuh Puluh Dua Ribu Lima Ratus Rupiah"
	}

	// PENTING: hitung COUNT(*) di subquery DULU sebelum JOIN ke se2026.sls -
	// se2026.sls punya banyak baris per (kode_kec,kode_desa) (satu per SLS se-
	// kabupaten), jadi JOIN langsung tanpa subquery akan fan-out dan bikin
	// Jumlah SLS meledak/salah (bug yg dilaporkan user sbg "error bast").
	var wilayah []wilayahKerjaRow
	if isPML {
		rows, err := database.DB.Query(`
			SELECT g.ppl_nama, g.kode_kec, MAX(s.nama_kec), g.kode_desa, MAX(s.nama_desa), g.jml
			FROM (
				SELECT ppl_idsobat, ppl_nama, kode_kec, kode_desa, COUNT(*) AS jml
				FROM lk_ppk_termin2_se2026
				WHERE pml_idsobat = ?
				GROUP BY ppl_idsobat, ppl_nama, kode_kec, kode_desa
			) g
			LEFT JOIN se2026.sls s ON s.kode_kec = g.kode_kec AND s.kode_desa = g.kode_desa
			GROUP BY g.ppl_idsobat, g.ppl_nama, g.kode_kec, g.kode_desa, g.jml
			ORDER BY g.ppl_nama, g.kode_kec, g.kode_desa`, idsobat)
		if err != nil {
			return d, jenis, nil, fmt.Errorf("gagal mengambil wilayah kerja: %v", err)
		}
		defer rows.Close()
		for rows.Next() {
			var w wilayahKerjaRow
			if err := rows.Scan(&w.Nama, &w.KodeKec, &w.NamaKec, &w.KodeDesa, &w.NamaDesa, &w.JumlahSLS); err == nil {
				w.Nama = properCase(w.Nama)
				wilayah = append(wilayah, w)
			}
		}
	} else {
		rows, err := database.DB.Query(`
			SELECT g.kode_kec, MAX(s.nama_kec), g.kode_desa, MAX(s.nama_desa), g.jml
			FROM (
				SELECT kode_kec, kode_desa, COUNT(*) AS jml
				FROM lk_ppk_termin2_se2026
				WHERE ppl_idsobat = ?
				GROUP BY kode_kec, kode_desa
			) g
			LEFT JOIN se2026.sls s ON s.kode_kec = g.kode_kec AND s.kode_desa = g.kode_desa
			GROUP BY g.kode_kec, g.kode_desa, g.jml
			ORDER BY g.kode_kec, g.kode_desa`, idsobat)
		if err != nil {
			return d, jenis, nil, fmt.Errorf("gagal mengambil wilayah kerja: %v", err)
		}
		defer rows.Close()
		for rows.Next() {
			var w wilayahKerjaRow
			if err := rows.Scan(&w.KodeKec, &w.NamaKec, &w.KodeDesa, &w.NamaDesa, &w.JumlahSLS); err == nil {
				wilayah = append(wilayah, w)
			}
		}
	}

	return d, jenis, wilayah, nil
}

// wilayahColW4 (PPL: No/Kec/Desa/Jumlah) & wilayahColW5 (PML: +Nama Petugas) - lebar
// kolom (dxa) disamakan proporsional dgn tabel contoh di template asli.
var wilayahColW4 = [4]int{900, 3000, 3000, 3000}
var wilayahColW5 = [5]int{700, 3200, 2600, 2600, 2600}

func generateWilayahKerjaTableXML(rows []wilayahKerjaRow, withNama bool) string {
	var c []int
	var headers []string
	if withNama {
		c = wilayahColW5[:]
		headers = []string{"No", "Nama Petugas Lapangan Sensus", "[Kode] KECAMATAN", "[Kode] DESA", "Jumlah SLS/Sub-SLS"}
	} else {
		c = wilayahColW4[:]
		headers = []string{"No", "[Kode] KECAMATAN/DISTRIK", "[Kode] DESA", "Jumlah SLS/sub-SLS"}
	}

	var rowXML strings.Builder
	for i, r := range rows {
		var cells []string
		if withNama {
			cells = []string{strconv.Itoa(i + 1), escapeXML(r.Nama), fmtKodeNama(r.KodeKec, r.NamaKec), fmtKodeNama(r.KodeDesa, r.NamaDesa), strconv.Itoa(r.JumlahSLS)}
		} else {
			cells = []string{strconv.Itoa(i + 1), fmtKodeNama(r.KodeKec, r.NamaKec), fmtKodeNama(r.KodeDesa, r.NamaDesa), strconv.Itoa(r.JumlahSLS)}
		}
		rowXML.WriteString("<w:tr>")
		for ci, cell := range cells {
			rowXML.WriteString(tcOpen(c[ci]))
			rowXML.WriteString(cellP(cell, false, ci != 1 || !withNama))
			rowXML.WriteString("</w:tc>")
		}
		rowXML.WriteString("</w:tr>")
	}
	if len(rows) == 0 {
		empty := make([]string, len(c))
		for i := range empty {
			empty[i] = "-"
		}
		rowXML.WriteString("<w:tr>")
		for ci, cell := range empty {
			rowXML.WriteString(tcOpen(c[ci]))
			rowXML.WriteString(cellP(cell, false, true))
			rowXML.WriteString("</w:tc>")
		}
		rowXML.WriteString("</w:tr>")
	}

	total := 0
	for _, w := range c {
		total += w
	}
	var gridXML, headerXML strings.Builder
	for i, w := range c {
		gridXML.WriteString(fmt.Sprintf(`<w:gridCol w:w="%d"/>`, w))
		headerXML.WriteString(tcOpen(w))
		headerXML.WriteString(cellP(headers[i], true, true))
		headerXML.WriteString("</w:tc>")
	}

	return fmt.Sprintf(`<w:tbl>
<w:tblPr><w:tblStyle w:val="TableGrid"/><w:tblW w:w="%d" w:type="dxa"/><w:jc w:val="center"/><w:tblLayout w:type="fixed"/>
<w:tblBorders>
<w:top w:val="single" w:sz="4" w:space="0" w:color="000000"/>
<w:left w:val="single" w:sz="4" w:space="0" w:color="000000"/>
<w:bottom w:val="single" w:sz="4" w:space="0" w:color="000000"/>
<w:right w:val="single" w:sz="4" w:space="0" w:color="000000"/>
<w:insideH w:val="single" w:sz="4" w:space="0" w:color="000000"/>
<w:insideV w:val="single" w:sz="4" w:space="0" w:color="000000"/>
</w:tblBorders>
</w:tblPr>
<w:tblGrid>%s</w:tblGrid>
<w:tr>%s</w:tr>%s</w:tbl>`, total, gridXML.String(), headerXML.String(), rowXML.String())
}

func generateBASTSE2026Docx(d bastSE2026Data, jenis string, wilayah []wilayahKerjaRow) ([]byte, error) {
	tpl := "static/templates-se2026/bast_termin2_ppl_se2026.docx"
	if jenis == "pml" {
		tpl = "static/templates-se2026/bast_termin2_pml_se2026.docx"
	}
	return replaceInDocx(tpl, map[string]string{
		"nomor":           d.Nomor,
		"nomor_spk":       d.NomorSPK,
		"nama_petugas":    d.NamaPetugas,
		"nik_petugas":     d.NIKPetugas,
		"nama_ketua_tim":  ketuaTimNamaSE2026,
		"target_sls":      strconv.Itoa(d.TargetSLS),
		"realisasi_sls":   strconv.Itoa(d.RealisasiSLS),
		"honor_angka":     d.HonorAngka,
		"honor_terbilang": d.HonorTerbilang,
		"table":           generateWilayahKerjaTableXML(wilayah, jenis == "pml"),
	})
}

// assignBASTSE2026Number idempotent, sama pola dgn assignBAPPNumber.
func assignBASTSE2026Number(rekapID int, tanggal string) error {
	var idsobat, idSpk, tahun, kegiatan string
	var existing sql.NullString
	err := database.DB.QueryRow(`SELECT idsobat, COALESCE(id_spk,''), tahun, kegiatan, id_bast_se2026 FROM rekap WHERE id=?`, rekapID).
		Scan(&idsobat, &idSpk, &tahun, &kegiatan, &existing)
	if err != nil {
		return fmt.Errorf("rekap id %d tidak ditemukan", rekapID)
	}
	if existing.Valid && existing.String != "" {
		return nil
	}
	if idSpk == "" {
		return fmt.Errorf("petugas belum memiliki nomor SPK")
	}
	isPML := jenisPetugasSE2026(kegiatan) == "pml"
	_, seqAll, _, ok := getPayableSeq(idsobat, isPML, 2)
	if !ok {
		return fmt.Errorf("petugas tidak ada di daftar yang bisa dibayarkan termin 2 (BAST digabung dgn termin 2)")
	}
	if tanggal == "" {
		tanggal = batchTanggalDefaultTermin2
	}
	tgl, err := time.Parse("2006-01-02", tanggal)
	if err != nil {
		return fmt.Errorf("format tanggal salah, gunakan YYYY-MM-DD")
	}
	nomor, err := suratNomorSE2026(seqAll, "bast", 2, tahun, tgl)
	if err != nil {
		return err
	}
	_, err = database.DB.Exec(`UPDATE rekap SET id_bast_se2026=?, tgl_bast_se2026=? WHERE id=?`,
		nomor, tgl.Format("2006-01-02"), rekapID)
	return err
}

// CreateBASTSE2026Handler POST /api/rekap/spk/se2026/bast/create
func CreateBASTSE2026Handler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		RekapID int    `json:"rekap_id"`
		Tanggal string `json:"tanggal"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	if err := assignBASTSE2026Number(req.RekapID, req.Tanggal); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

type createBulkBASTSE2026Item struct {
	RekapID int `json:"rekap_id"`
}

type createBulkBASTSE2026Payload struct {
	Tanggal string                     `json:"tanggal"`
	Items   []createBulkBASTSE2026Item `json:"items"`
}

// CreateBulkBASTSE2026Handler POST /api/rekap/spk/se2026/bast/create-bulk
// Urutan Items MENENTUKAN urutan seqAll krn ditentukan dari lk_ppk_payable_se2026
// (bukan dari urutan create) - beda dgn alur generik. Tetap diproses satu-satu.
func CreateBulkBASTSE2026Handler(w http.ResponseWriter, r *http.Request) {
	var payload createBulkBASTSE2026Payload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Payload tidak valid", http.StatusBadRequest)
		return
	}
	if len(payload.Items) == 0 {
		http.Error(w, "Pilih minimal satu petugas", http.StatusBadRequest)
		return
	}
	ok, failed := 0, 0
	var errs []string
	for _, item := range payload.Items {
		if err := assignBASTSE2026Number(item.RekapID, payload.Tanggal); err != nil {
			failed++
			errs = append(errs, fmt.Sprintf("rekap_id %d: %v", item.RekapID, err))
			continue
		}
		ok++
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"ok": ok, "failed": failed, "errors": errs})
}

// DownloadBASTSE2026Handler GET /api/rekap/spk/se2026/bast/download?rekap_id=X&tanggal=2026-09-15
func DownloadBASTSE2026Handler(w http.ResponseWriter, r *http.Request) {
	rekapID, _ := strconv.Atoi(r.URL.Query().Get("rekap_id"))
	if rekapID == 0 {
		http.Error(w, "rekap_id wajib", http.StatusBadRequest)
		return
	}
	tanggal := r.URL.Query().Get("tanggal")

	d, jenis, wilayah, err := buildBASTSE2026(rekapID, tanggal)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if d.Nomor == "" {
		http.Error(w, "Petugas ini belum memiliki nomor BAST - buat nomornya dulu", http.StatusBadRequest)
		return
	}
	docx, err := generateBASTSE2026Docx(d, jenis, wilayah)
	if err != nil {
		http.Error(w, "Gagal generate BAST: "+err.Error(), http.StatusInternalServerError)
		return
	}
	fname := fmt.Sprintf("BAST_SE2026_%s.docx", safeName(d.NamaPetugas))
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, fname))
	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.wordprocessingml.document")
	w.Header().Set("Content-Length", strconv.Itoa(len(docx)))
	w.Write(docx)
}

// DownloadAllBASTSE2026Handler GET /api/rekap/spk/se2026/bast/download-all?jenis=pcl&tanggal=2026-09-15
func DownloadAllBASTSE2026Handler(w http.ResponseWriter, r *http.Request) {
	jenisFlt := strings.TrimSpace(r.URL.Query().Get("jenis"))
	tanggal := r.URL.Query().Get("tanggal")

	var kegiatanFilter string
	switch jenisFlt {
	case "pml":
		kegiatanFilter = "%Pemeriksa Lapangan Sensus Ekonomi%"
	case "pcl":
		kegiatanFilter = "%Pendataan Sensus Ekonomi%"
	default:
		kegiatanFilter = "%Sensus Ekonomi 2026%"
	}

	rows, err := database.DB.Query(`
		SELECT r.id FROM rekap r
		WHERE r.kegiatan LIKE ? AND r.id_bast_se2026 IS NOT NULL AND r.id_bast_se2026 != ''
		  AND (r.tanggaran = 2026 OR r.tahun = '2026')`, kegiatanFilter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var ids []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err == nil {
			ids = append(ids, id)
		}
	}
	if len(ids) == 0 {
		http.Error(w, "Tidak ada BAST dengan nomor surat", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="BAST_SE2026_%s.zip"`, jenisFlt))

	zw := zip.NewWriter(w)
	defer zw.Close()

	for _, id := range ids {
		d, jenis, wilayah, err := buildBASTSE2026(id, tanggal)
		if err != nil {
			continue
		}
		docx, err := generateBASTSE2026Docx(d, jenis, wilayah)
		if err != nil {
			continue
		}
		fname := fmt.Sprintf("BAST_%s_%s.docx", safeName(d.Nomor), safeName(d.NamaPetugas))
		f, err := zw.Create(fname)
		if err != nil {
			continue
		}
		f.Write(docx)
	}
}
