package handlers

import (
	"archive/zip"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"sirekap/internal/database"
)

// Identitas tetap PPK & Ketua Tim Pelaksana SE2026 BPS Kabupaten Dompu (dikonfirmasi user).
const (
	ppkNamaSE2026      = "Emalia Septiani Hirma"
	ppkNIPSE2026       = "199809032022012002"
	ketuaTimNamaSE2026 = "Abdul Farid, SE"
	ketuaTimNIPSE2026  = "19720404 199803 1 005"
)

type bappSE2026Data struct {
	Nomor          string
	NomorSPK       string
	Hari           string
	TglTeks        string
	BlnTeks        string
	TglAngka       string
	NamaPetugas    string
	NIKPetugas     string
	TargetSLS      int
	RealisasiSLS   int
	HonorAngka     string
	HonorTerbilang string
}

var bappNomorSeqRe = regexp.MustCompile(`B-(\d+)/SPK-SE2026`)

// bappNomorFromSPK menurunkan nomor BAPP/Surat Pernyataan dari nomor SPK yang sudah
// dimiliki petugas (id_spk), memakai ulang sequence yang sama supaya dokumen-dokumen
// ini tetap terhubung 1:1 ke perjanjian kerja yang sama, bukan sequence baru.
func bappNomorFromSPK(idSpk, kind string, termin int, tahun string) (string, error) {
	m := bappNomorSeqRe.FindStringSubmatch(idSpk)
	if m == nil {
		return "", fmt.Errorf("nomor SPK tidak dikenali: %s", idSpk)
	}
	seq := m[1]
	switch kind {
	case "bapp":
		roman := "I"
		if termin == 2 {
			roman = "II"
		}
		return fmt.Sprintf("B-%s/BAPP-%s-SE2026/5205/PL.200/%s", seq, roman, tahun), nil
	case "pernyataan":
		return fmt.Sprintf("B-%s/SE2026/5205/PL.200/%s", seq, tahun), nil
	default:
		return "", fmt.Errorf("kind tidak dikenal: %s", kind)
	}
}

// jenisPetugasSE2026 menentukan pcl/pml dari teks kegiatan, sama seperti buildSPKSE2026.
func jenisPetugasSE2026(kegiatan string) string {
	if strings.Contains(strings.ToLower(kegiatan), "pemeriksa") {
		return "pml"
	}
	return "pcl"
}

// buildBAPPSE2026 mengambil data rekap+NIK dan menyusun map penggantian placeholder.
func buildBAPPSE2026(rekapID, termin int, tanggalBAPP string) (bappSE2026Data, string, error) {
	var d bappSE2026Data
	var honorStr, idSpk, kegiatan, tahun, jumlahSLSStr string

	row := database.DB.QueryRow(`
		SELECT r.namamitra, r.honor, COALESCE(r.id_spk,''), r.kegiatan, r.tahun,
		       COALESCE(NULLIF(r.jumlah_sls,''),'0'), r.realisasi_sls
		FROM rekap r WHERE r.id = ?`, rekapID)
	if err := row.Scan(&d.NamaPetugas, &honorStr, &idSpk, &kegiatan, &tahun, &jumlahSLSStr, &d.RealisasiSLS); err != nil {
		return d, "", fmt.Errorf("rekap id %d tidak ditemukan", rekapID)
	}
	d.NamaPetugas = properCase(d.NamaPetugas)
	d.TargetSLS, _ = strconv.Atoi(strings.TrimSpace(jumlahSLSStr))

	if idSpk == "" {
		return d, "", fmt.Errorf("petugas ini belum memiliki nomor SPK")
	}
	jenis := jenisPetugasSE2026(kegiatan)

	nomor, err := bappNomorFromSPK(idSpk, "bapp", termin, tahun)
	if err != nil {
		return d, jenis, err
	}
	d.Nomor = nomor
	d.NomorSPK = idSpk

	// NIK asli ada di seleksimitra.pegawai (mitra.nik kosong utk baris SE2026), dicari
	// live by nama tiap generate (bukan disimpan) - lihat internal/handlers/se2026_import.go.
	database.DB.QueryRow(`
		SELECT nip FROM pegawai
		WHERE UPPER(CONVERT(nama USING utf8mb4)) COLLATE utf8mb4_unicode_ci = UPPER(?) COLLATE utf8mb4_unicode_ci
		LIMIT 1`, d.NamaPetugas).Scan(&d.NIKPetugas)

	tgl, err := time.Parse("2006-01-02", tanggalBAPP)
	if err != nil {
		tgl, _ = time.Parse("2006-01-02", "2026-08-15")
	}
	d.Hari = hariIndonesia[tgl.Weekday().String()]
	d.TglTeks = dayToTeks(tgl.Day())
	d.BlnTeks = bulanTeksSE[tgl.Month()]
	d.TglAngka = fmt.Sprintf("%02d-%02d", tgl.Day(), int(tgl.Month()))

	honorVal, _ := strconv.ParseFloat(honorStr, 64)
	honorTermin := honorVal * 0.4
	if termin == 2 {
		honorTermin = honorVal * 0.6
	}
	honorInt := int64(honorTermin)
	d.HonorAngka = formatAngkaSE(honorInt)
	d.HonorTerbilang = strings.Title(Terbilang(honorInt))

	return d, jenis, nil
}

func generateBAPPDocx(d bappSE2026Data, jenis string, termin int) ([]byte, error) {
	var tpl string
	switch {
	case jenis == "pml" && termin == 2:
		tpl = "static/templates-se2026/bapp_termin2_pml_se2026.docx"
	case jenis == "pml":
		tpl = "static/templates-se2026/bapp_termin1_pml_se2026.docx"
	case termin == 2:
		tpl = "static/templates-se2026/bapp_termin2_pcl_se2026.docx"
	default:
		tpl = "static/templates-se2026/bapp_termin1_pcl_se2026.docx"
	}
	return replaceInDocx(tpl, map[string]string{
		"nomor":           d.Nomor,
		"nomor_spk":       d.NomorSPK,
		"hari":            d.Hari,
		"tgl_teks":        d.TglTeks,
		"bln_teks":        d.BlnTeks,
		"tgl_angka":       d.TglAngka,
		"nama_petugas":    d.NamaPetugas,
		"nik_petugas":     d.NIKPetugas,
		"nama_ppk":        ppkNamaSE2026,
		"nip_ppk":         ppkNIPSE2026,
		"nama_ketua_tim":  ketuaTimNamaSE2026,
		"nip_ketua_tim":   ketuaTimNIPSE2026,
		"target_sls":      strconv.Itoa(d.TargetSLS),
		"realisasi_sls":   strconv.Itoa(d.RealisasiSLS),
		"honor_angka":     d.HonorAngka,
		"honor_terbilang": d.HonorTerbilang,
	})
}

// assignBAPPNumber menetapkan id_bappN/tgl_bappN kalau belum ada (idempotent - tidak
// menimpa nomor yang sudah pernah dibuat, sama seperti CreateSurat untuk kegiatan biasa).
func assignBAPPNumber(rekapID, termin int, tanggal string) error {
	if termin != 1 && termin != 2 {
		return fmt.Errorf("termin harus 1 atau 2")
	}
	idCol, tglCol := "id_bapp1", "tgl_bapp1"
	if termin == 2 {
		idCol, tglCol = "id_bapp2", "tgl_bapp2"
	}

	var idSpk, tahun string
	var existing sql.NullString
	err := database.DB.QueryRow(fmt.Sprintf(
		`SELECT COALESCE(id_spk,''), tahun, %s FROM rekap WHERE id=?`, idCol), rekapID,
	).Scan(&idSpk, &tahun, &existing)
	if err != nil {
		return fmt.Errorf("rekap id %d tidak ditemukan", rekapID)
	}
	if existing.Valid && existing.String != "" {
		return nil
	}
	if idSpk == "" {
		return fmt.Errorf("petugas belum memiliki nomor SPK")
	}

	nomor, err := bappNomorFromSPK(idSpk, "bapp", termin, tahun)
	if err != nil {
		return err
	}
	tgl, err := time.Parse("2006-01-02", tanggal)
	if err != nil {
		return fmt.Errorf("format tanggal salah, gunakan YYYY-MM-DD")
	}
	_, err = database.DB.Exec(fmt.Sprintf(`UPDATE rekap SET %s=?, %s=? WHERE id=?`, idCol, tglCol),
		nomor, tgl.Format("2006-01-02"), rekapID)
	return err
}

type createBAPPPayload struct {
	RekapID int    `json:"rekap_id"`
	Termin  int    `json:"termin"`
	Tanggal string `json:"tanggal"`
}

// CreateBAPPHandler POST /api/rekap/spk/se2026/bapp/create
func CreateBAPPHandler(w http.ResponseWriter, r *http.Request) {
	var req createBAPPPayload
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	if err := assignBAPPNumber(req.RekapID, req.Termin, req.Tanggal); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

type createBulkBAPPPayload struct {
	Termin  int    `json:"termin"`
	Tanggal string `json:"tanggal"`
	Items   []struct {
		RekapID int `json:"rekap_id"`
	} `json:"items"`
}

// CreateBulkBAPPHandler POST /api/rekap/spk/se2026/bapp/create-bulk
func CreateBulkBAPPHandler(w http.ResponseWriter, r *http.Request) {
	var req createBulkBAPPPayload
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	ok, failed := 0, 0
	var errs []string
	for _, item := range req.Items {
		if err := assignBAPPNumber(item.RekapID, req.Termin, req.Tanggal); err != nil {
			failed++
			errs = append(errs, fmt.Sprintf("rekap_id %d: %v", item.RekapID, err))
			continue
		}
		ok++
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"ok": ok, "failed": failed, "errors": errs})
}

// DownloadBAPPSE2026Handler GET /api/rekap/spk/se2026/bapp/download?rekap_id=X&termin=1&tanggal=2026-08-15
func DownloadBAPPSE2026Handler(w http.ResponseWriter, r *http.Request) {
	rekapID, _ := strconv.Atoi(r.URL.Query().Get("rekap_id"))
	termin, _ := strconv.Atoi(r.URL.Query().Get("termin"))
	if rekapID == 0 || (termin != 1 && termin != 2) {
		http.Error(w, "rekap_id dan termin (1/2) wajib", http.StatusBadRequest)
		return
	}
	tanggal := r.URL.Query().Get("tanggal")
	if tanggal == "" {
		tanggal = "2026-08-15"
	}

	d, jenis, err := buildBAPPSE2026(rekapID, termin, tanggal)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if d.Nomor == "" {
		http.Error(w, "Petugas ini belum memiliki nomor BAPP - buat nomornya dulu", http.StatusBadRequest)
		return
	}

	docx, err := generateBAPPDocx(d, jenis, termin)
	if err != nil {
		http.Error(w, "Gagal generate BAPP: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fname := fmt.Sprintf("BAPP_Termin%d_SE2026_%s.docx", termin, safeName(d.NamaPetugas))
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, fname))
	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.wordprocessingml.document")
	w.Header().Set("Content-Length", strconv.Itoa(len(docx)))
	w.Write(docx)
}

// DownloadAllBAPPSE2026Handler GET /api/rekap/spk/se2026/bapp/download-all?jenis=pcl&termin=1&tanggal=2026-08-15
func DownloadAllBAPPSE2026Handler(w http.ResponseWriter, r *http.Request) {
	jenisFlt := strings.TrimSpace(r.URL.Query().Get("jenis"))
	termin, _ := strconv.Atoi(r.URL.Query().Get("termin"))
	if termin != 1 && termin != 2 {
		http.Error(w, "termin (1/2) wajib", http.StatusBadRequest)
		return
	}
	tanggal := r.URL.Query().Get("tanggal")
	if tanggal == "" {
		tanggal = "2026-08-15"
	}

	var kegiatanFilter string
	switch jenisFlt {
	case "pml":
		kegiatanFilter = "%Pemeriksa Lapangan Sensus Ekonomi%"
	case "pcl":
		kegiatanFilter = "%Pendataan Sensus Ekonomi%"
	default:
		kegiatanFilter = "%Sensus Ekonomi 2026%"
	}

	idCol := "id_bapp1"
	if termin == 2 {
		idCol = "id_bapp2"
	}
	rows, err := database.DB.Query(fmt.Sprintf(`
		SELECT r.id FROM rekap r
		WHERE r.kegiatan LIKE ? AND r.%s IS NOT NULL AND r.%s != ''
		  AND (r.tanggaran = 2026 OR r.tahun = '2026')`, idCol, idCol), kegiatanFilter)
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
		http.Error(w, "Tidak ada BAPP dengan nomor surat", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="BAPP_Termin%d_SE2026_%s.zip"`, termin, jenisFlt))

	zw := zip.NewWriter(w)
	defer zw.Close()

	for _, id := range ids {
		d, jenis, err := buildBAPPSE2026(id, termin, tanggal)
		if err != nil {
			continue
		}
		docx, err := generateBAPPDocx(d, jenis, termin)
		if err != nil {
			continue
		}
		fname := fmt.Sprintf("BAPP_Termin%d_%s_%s.docx", termin, safeName(d.Nomor), safeName(d.NamaPetugas))
		f, err := zw.Create(fname)
		if err != nil {
			continue
		}
		f.Write(docx)
	}
}
