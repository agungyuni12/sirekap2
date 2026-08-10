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

// Identitas tetap Kepala BPS, PPK & Ketua Tim Pelaksana SE2026 BPS Kabupaten Dompu (dikonfirmasi user).
const (
	kepalaNamaSE2026   = "Ahwan Hadi, S.ST., M.Ak"
	kepalaNIPSE2026    = "19780223 200012 1 002"
	ppkNamaSE2026      = "Emalia Septiani Hirma, S.Tr.Stat."
	ppkNIPSE2026       = "19980903 202201 2 002"
	ketuaTimNamaSE2026 = "Abdul Farid, SE"
	ketuaTimNIPSE2026  = "19720404 199803 1 005"
)

type bappSE2026Data struct {
	Nomor          string
	NomorSPK       string
	Hari           string
	TglTeks        string
	TglHari        string
	BlnTeks        string
	TglAngka       string
	NamaPetugas    string
	NIKPetugas     string
	TargetSLS      int
	RealisasiSLS   int
	HonorAngka     string
	HonorTerbilang string
}

// suratNomorSE2026 membangun nomor surat (BAPP / Surat Pernyataan) dari seq (nomor urut
// baru, alfabetis per role, lihat lk_ppk_payable_se2026 & getPayableSeq) digabung dengan
// bulan+tanggal terbit dokumen ini sendiri. Sequence ini SENGAJA independen dari nomor SPK
// (rekap.id_spk) - direnumber ulang khusus utk daftar petugas yang bisa dicairkan termin ini.
// Format: B-{bulan(2)}.{tanggal(2)}.{seq(3)}/....
// kind: "pernyataan_ppl" | "pernyataan_pml" | "bapp"
func suratNomorSE2026(seq int, kind string, termin int, tahun string, tgl time.Time) (string, error) {
	tgd := fmt.Sprintf("%02d.%02d.%03d", int(tgl.Month()), tgl.Day(), seq)
	switch kind {
	case "bapp":
		roman := "I"
		if termin == 2 {
			roman = "II"
		}
		return fmt.Sprintf("B-%s/BAPP-%s-SE2026/5205.PPK/BA/%s", tgd, roman, tahun), nil
	case "pernyataan_ppl":
		return fmt.Sprintf("B-%s/SE2026/5205/Super.PPL/%s", tgd, tahun), nil
	case "pernyataan_pml":
		return fmt.Sprintf("B-%s/SE2026/5205/Super.PML/%s", tgd, tahun), nil
	default:
		return "", fmt.Errorf("kind tidak dikenal: %s", kind)
	}
}

// batchTanggalDefault mengembalikan tanggal default per batch & jenis dokumen (dipakai
// kalau tanggal tidak dikirim eksplisit dari UI).
//
//	batch 1: Pernyataan 16 Juli 2026, BAPP/Kepala 17 Juli 2026
//	batch 2: Pernyataan 18 Juli 2026, BAPP/Kepala 19 Juli 2026
//	batch 3: Pernyataan 20 Juli 2026, BAPP/Kepala 21 Juli 2026
//	batch 4: Pernyataan 8 Agustus 2026, BAPP/Kepala 9 Agustus 2026
func batchTanggalDefault(batch int, kind string) string {
	pernyataan, bappKepala := "2026-07-16", "2026-07-17"
	switch batch {
	case 2:
		pernyataan, bappKepala = "2026-07-18", "2026-07-19"
	case 3:
		pernyataan, bappKepala = "2026-07-20", "2026-07-21"
	case 4:
		pernyataan, bappKepala = "2026-08-08", "2026-08-09"
	}
	if kind == "pernyataan" {
		return pernyataan
	}
	return bappKepala
}

// getPayableSeq mengembalikan nomor urut petugas di lk_ppk_payable_se2026 - juga dipakai
// sebagai whitelist "bisa dicairkan": ok=false berarti petugas ini TIDAK ada di daftar
// resmi yang bisa dibayarkan termin ini.
// seq = nomor urut per role, LANJUT antar batch (1..208 PPL batch 1, 209..225 batch 2 dst /
// 1..14 PML batch 1, 15..20 batch 2 dst), dipakai utk nomor Surat Pernyataan (formatnya
// beda suffix Super.PPL/Super.PML jadi aman walau seq sama antar role).
// seqAll = nomor urut gabungan PML lalu PPL, LANJUT antar batch juga, WAJIB dipakai utk
// nomor BAPP karena format BAPP sama persis utk PPL & PML (harus unik silang role & batch).
// batch = 1, 2, atau 3, dipakai utk pilih tanggal default dokumen (lihat batchTanggalDefault).
func getPayableSeq(idsobat string, isPML bool) (seq int, seqAll int, batch int, ok bool) {
	role := "ppl"
	if isPML {
		role = "pml"
	}
	err := database.DB.QueryRow(`SELECT seq, seq_all, batch FROM lk_ppk_payable_se2026 WHERE idsobat = ? AND role = ?`, idsobat, role).Scan(&seq, &seqAll, &batch)
	return seq, seqAll, batch, err == nil
}

// kepalaNomorSE2026 membangun nomor Surat Pernyataan Kepala BPS - dokumen tunggal per
// batch (bukan per-petugas), jadi sequence-nya = nomor batch (001 utk batch 1, 002 utk
// batch 2, dst), bukan diturunkan dari SPK.
func kepalaNomorSE2026(batch int, tahun string, tgl time.Time) string {
	return fmt.Sprintf("B-%02d.%02d.%03d/SE2026/5205/Super.KPL/%s", int(tgl.Month()), tgl.Day(), batch, tahun)
}

// jenisPetugasSE2026 menentukan pcl/pml dari teks kegiatan, sama seperti buildSPKSE2026.
func jenisPetugasSE2026(kegiatan string) string {
	if strings.Contains(strings.ToLower(kegiatan), "pemeriksa") {
		return "pml"
	}
	return "pcl"
}

// buildBAPPSE2026 mengambil data rekap+NIK dan menyusun map penggantian placeholder -
// TargetSLS/RealisasiSLS dihitung dari lk_ppk_termin1_se2026 (SLS PRIORITAS saja).
func buildBAPPSE2026(rekapID, termin int, tanggalBAPP string) (bappSE2026Data, string, error) {
	var d bappSE2026Data
	var idsobat, honorStr, idSpk, kegiatan, tahun string

	row := database.DB.QueryRow(`
		SELECT r.idsobat, r.namamitra, r.honor, COALESCE(r.id_spk,''), r.kegiatan, r.tahun
		FROM rekap r WHERE r.id = ?`, rekapID)
	if err := row.Scan(&idsobat, &d.NamaPetugas, &honorStr, &idSpk, &kegiatan, &tahun); err != nil {
		return d, "", fmt.Errorf("rekap id %d tidak ditemukan", rekapID)
	}
	d.NamaPetugas = properCase(d.NamaPetugas)

	if idSpk == "" {
		return d, "", fmt.Errorf("petugas ini belum memiliki nomor SPK")
	}
	jenis := jenisPetugasSE2026(kegiatan)
	isPML := jenis == "pml"

	_, seqAll, batch, ok := getPayableSeq(idsobat, isPML)
	if !ok {
		return d, jenis, fmt.Errorf("'%s' tidak ada di daftar LK PPK Termin 1 (tidak bisa dibayarkan termin ini)", d.NamaPetugas)
	}
	_, _, targetSLS, realisasiSLS, err := computeUsahaKeluargaSE2026(idsobat, isPML)
	if err != nil {
		return d, jenis, fmt.Errorf("gagal menghitung SLS prioritas: %v", err)
	}
	d.TargetSLS, d.RealisasiSLS = targetSLS, realisasiSLS

	if tanggalBAPP == "" {
		tanggalBAPP = batchTanggalDefault(batch, "bapp")
	}
	tgl, err := time.Parse("2006-01-02", tanggalBAPP)
	if err != nil {
		tgl, _ = time.Parse("2006-01-02", batchTanggalDefault(batch, "bapp"))
	}

	nomor, err := suratNomorSE2026(seqAll, "bapp", termin, tahun, tgl)
	if err != nil {
		return d, jenis, err
	}
	d.Nomor = nomor
	d.NomorSPK = idSpk

	// NIK diambil dari mitra.nik by idsobat (diisi dari BPJS - Petugas SE.xlsx, lihat
	// migrations/update_nik_mitra_se2026.sql) - jauh lebih andal drpd cari by nama ke
	// tabel pegawai (89 dari 229 petugas gagal ketemu krn beda ejaan nama).
	database.DB.QueryRow(`SELECT nik FROM mitra WHERE idsobat = ? AND nik IS NOT NULL AND nik != '' LIMIT 1`, idsobat).Scan(&d.NIKPetugas)

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

	var idsobat, idSpk, tahun, kegiatan string
	var existing sql.NullString
	err := database.DB.QueryRow(fmt.Sprintf(
		`SELECT idsobat, COALESCE(id_spk,''), tahun, kegiatan, %s FROM rekap WHERE id=?`, idCol), rekapID,
	).Scan(&idsobat, &idSpk, &tahun, &kegiatan, &existing)
	if err != nil {
		return fmt.Errorf("rekap id %d tidak ditemukan", rekapID)
	}
	if existing.Valid && existing.String != "" {
		return nil
	}
	if idSpk == "" {
		return fmt.Errorf("petugas belum memiliki nomor SPK")
	}
	_, seqAll, batch, ok := getPayableSeq(idsobat, jenisPetugasSE2026(kegiatan) == "pml")
	if !ok {
		return fmt.Errorf("petugas tidak ada di daftar LK PPK Termin 1 (tidak bisa dibayarkan termin ini)")
	}
	if tanggal == "" && termin == 1 {
		tanggal = batchTanggalDefault(batch, "bapp")
	}

	tgl, err := time.Parse("2006-01-02", tanggal)
	if err != nil {
		return fmt.Errorf("format tanggal salah, gunakan YYYY-MM-DD")
	}
	nomor, err := suratNomorSE2026(seqAll, "bapp", termin, tahun, tgl)
	if err != nil {
		return err
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
	if tanggal == "" && termin == 2 {
		tanggal = "2026-08-15"
	}
	// termin==1: dibiarkan kosong kalau tidak dikirim eksplisit - buildBAPPSE2026 akan
	// pilih tanggal default sesuai batch (1 atau 2) petugas ybs.

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
	if tanggal == "" && termin == 2 {
		tanggal = "2026-08-15"
	}
	// termin==1: dibiarkan kosong kalau tidak dikirim eksplisit - buildBAPPSE2026 pilih
	// tanggal per petugas sesuai batch-nya masing-masing.
	batch, _ := strconv.Atoi(r.URL.Query().Get("batch"))

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
	query := fmt.Sprintf(`
		SELECT r.id FROM rekap r
		WHERE r.kegiatan LIKE ? AND r.%s IS NOT NULL AND r.%s != ''
		  AND (r.tanggaran = 2026 OR r.tahun = '2026')`, idCol, idCol)
	args := []interface{}{kegiatanFilter}
	if termin == 1 && (batch >= 1 && batch <= 4) {
		query += ` AND EXISTS (SELECT 1 FROM lk_ppk_payable_se2026 p WHERE p.idsobat COLLATE utf8mb4_general_ci = r.idsobat AND p.batch = ?)`
		args = append(args, batch)
	}

	rows, err := database.DB.Query(query, args...)
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

	batchSuffix := ""
	if termin == 1 && (batch >= 1 && batch <= 4) {
		batchSuffix = fmt.Sprintf("_Batch%d", batch)
	}
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="BAPP_Termin%d_SE2026_%s%s.zip"`, termin, jenisFlt, batchSuffix))

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
