package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"sirekap/internal/database"
)

// ---------------------------------------------------------------------------------
// Surat Pernyataan Evaluasi Pelaksanaan Lapangan SE2026 Termin I - Kepala BPS
// Dokumen tunggal (bukan per-petugas): rekap seluruh PPL+PML dalam satu lampiran tabel.
// ---------------------------------------------------------------------------------

type kepalaSE2026Data struct {
	Nomor    string
	Hari     string
	TglTeks  string
	TglHari  string
	BlnTeks  string
	TglAngka string
}

// buildPernyataanKepalaSE2026 mengagregasi SEMUA petugas SE2026 (PML lalu PPL) yang ada
// di lk_ppk_payable_se2026 - daftar resmi "yang bisa dicairkan" termin ini (215 PPL +
// 14 PML terkonfirmasi, lihat migrations/add_lk_ppk_payable_se2026.sql), diurutkan
// sesuai seq (alfabetis per role).
func buildPernyataanKepalaSE2026(tanggal string) (kepalaSE2026Data, []usahaKeluargaLampiranRow, error) {
	var d kepalaSE2026Data

	tgl, err := time.Parse("2006-01-02", tanggal)
	if err != nil {
		tgl, _ = time.Parse("2006-01-02", "2026-07-17")
	}
	d.Hari = hariIndonesia[tgl.Weekday().String()]
	d.TglTeks = dayToTeks(tgl.Day())
	d.TglHari = strconv.Itoa(tgl.Day())
	d.BlnTeks = bulanTeksSE[tgl.Month()]
	d.TglAngka = fmt.Sprintf("%02d-%02d", tgl.Day(), int(tgl.Month()))
	d.Nomor = kepalaNomorSE2026("2026", tgl)

	var lampiran []usahaKeluargaLampiranRow

	pmlRows, err := database.DB.Query(`SELECT idsobat, nama FROM lk_ppk_payable_se2026 WHERE role='pml' ORDER BY seq`)
	if err != nil {
		return d, nil, fmt.Errorf("gagal mengambil daftar PML: %v", err)
	}
	for pmlRows.Next() {
		var idsobat, nama string
		if err := pmlRows.Scan(&idsobat, &nama); err != nil {
			continue
		}
		target, realisasi, _, _, err := computeUsahaKeluargaSE2026(idsobat, true)
		if err != nil {
			continue
		}
		lampiran = append(lampiran, usahaKeluargaLampiranRow{
			Nama: properCase(nama), Jabatan: "Pemeriksa Lapangan", TargetPrelist: target, Realisasi: realisasi,
		})
	}
	pmlRows.Close()

	pplRows, err := database.DB.Query(`SELECT idsobat, nama FROM lk_ppk_payable_se2026 WHERE role='ppl' ORDER BY seq`)
	if err != nil {
		return d, nil, fmt.Errorf("gagal mengambil daftar PPL: %v", err)
	}
	for pplRows.Next() {
		var idsobat, nama string
		if err := pplRows.Scan(&idsobat, &nama); err != nil {
			continue
		}
		target, realisasi, _, _, err := computeUsahaKeluargaSE2026(idsobat, false)
		if err != nil {
			continue
		}
		lampiran = append(lampiran, usahaKeluargaLampiranRow{
			Nama: properCase(nama), Jabatan: "Petugas Lapangan", TargetPrelist: target, Realisasi: realisasi,
		})
	}
	pplRows.Close()

	return d, lampiran, nil
}

// kepalaColW adalah lebar kolom (twips/dxa) persis seperti tabel contoh di
// surat_pernyataan_kepala_se2026.docx asli (total 13410 dxa = penuh selebar halaman
// landscape lampiran tersebut).
var kepalaColW = [6]int{1123, 2692, 2690, 2424, 2426, 2055}

// generateKepalaLampiranTableXML sama seperti generateLampiranTableXML tapi dengan
// tambahan kolom Jabatan (PML/PPL campur dalam satu tabel) - lebar kolom & tblLayout=fixed
// disamakan dengan tabel contoh di template aslinya.
func generateKepalaLampiranTableXML(rows []usahaKeluargaLampiranRow) string {
	c := kepalaColW
	var rowXML string
	totalTarget, totalRealisasi := 0, 0
	for i, r := range rows {
		totalTarget += r.TargetPrelist
		totalRealisasi += r.Realisasi
		rowXML += fmt.Sprintf(`
<w:tr>
%s%s</w:tc>
%s%s</w:tc>
%s%s</w:tc>
%s%s</w:tc>
%s%s</w:tc>
%s%s</w:tc>
</w:tr>`,
			tcOpen(c[0]), cellP(fmt.Sprintf("%d", i+1), false, true),
			tcOpen(c[1]), cellP(escapeXML(r.Nama), false, false),
			tcOpen(c[2]), cellP(escapeXML(r.Jabatan), false, true),
			tcOpen(c[3]), cellP(fmt.Sprintf("%d", r.TargetPrelist), false, true),
			tcOpen(c[4]), cellP(fmt.Sprintf("%d", r.Realisasi), false, true),
			tcOpen(c[5]), cellP(fmt.Sprintf("%.1f", persentase(r.Realisasi, r.TargetPrelist)), false, true))
	}
	rowXML += fmt.Sprintf(`
<w:tr>
<w:tc><w:tcPr><w:tcW w:w="%d" w:type="dxa"/><w:gridSpan w:val="3"/>`+tcBordersXML+`</w:tcPr>%s</w:tc>
%s%s</w:tc>
%s%s</w:tc>
%s%s</w:tc>
</w:tr>`,
		c[0]+c[1]+c[2], cellP("Jumlah", true, false),
		tcOpen(c[3]), cellP(fmt.Sprintf("%d", totalTarget), true, true),
		tcOpen(c[4]), cellP(fmt.Sprintf("%d", totalRealisasi), true, true),
		tcOpen(c[5]), cellP(fmt.Sprintf("%.1f", persentase(totalRealisasi, totalTarget)), true, true))

	return fmt.Sprintf(`<w:tbl>
<w:tblPr>
<w:tblStyle w:val="TableGrid"/>
<w:tblW w:w="%d" w:type="dxa"/>
<w:jc w:val="center"/>
<w:tblLayout w:type="fixed"/>
<w:tblBorders>
<w:top w:val="single" w:sz="4" w:space="0" w:color="000000"/>
<w:left w:val="single" w:sz="4" w:space="0" w:color="000000"/>
<w:bottom w:val="single" w:sz="4" w:space="0" w:color="000000"/>
<w:right w:val="single" w:sz="4" w:space="0" w:color="000000"/>
<w:insideH w:val="single" w:sz="4" w:space="0" w:color="000000"/>
<w:insideV w:val="single" w:sz="4" w:space="0" w:color="000000"/>
</w:tblBorders>
</w:tblPr>
<w:tblGrid>
<w:gridCol w:w="%d"/>
<w:gridCol w:w="%d"/>
<w:gridCol w:w="%d"/>
<w:gridCol w:w="%d"/>
<w:gridCol w:w="%d"/>
<w:gridCol w:w="%d"/>
</w:tblGrid>
<w:tr>
%s%s</w:tc>
%s%s</w:tc>
%s%s</w:tc>
%s%s</w:tc>
%s%s</w:tc>
%s%s</w:tc>
</w:tr>`,
		c[0]+c[1]+c[2]+c[3]+c[4]+c[5], c[0], c[1], c[2], c[3], c[4], c[5],
		tcOpen(c[0]), cellP("No", true, true),
		tcOpen(c[1]), cellP("Nama Petugas", true, true),
		tcOpen(c[2]), cellP("Jabatan", true, true),
		tcOpen(c[3]), cellP("Target Prelist", true, true),
		tcOpen(c[4]), cellP("Realisasi Hasil Pendataan (Usaha+Keluarga)", true, true),
		tcOpen(c[5]), cellP("Presentase (%)", true, true),
	) + rowXML + `</w:tbl>`
}

func generatePernyataanKepalaDocx(d kepalaSE2026Data, lampiran []usahaKeluargaLampiranRow) ([]byte, error) {
	return replaceInDocx("static/templates-se2026/surat_pernyataan_kepala_se2026.docx", map[string]string{
		"nomor":       d.Nomor,
		"hari":        d.Hari,
		"tgl_teks":    d.TglTeks,
		"tgl_hari":    d.TglHari,
		"bln_teks":    d.BlnTeks,
		"tgl_angka":   d.TglAngka,
		"nama_kepala": kepalaNamaSE2026,
		"nip_kepala":  kepalaNIPSE2026,
		"table":       generateKepalaLampiranTableXML(lampiran),
	})
}

// assignKepalaNomor menetapkan (idempotent) nomor Surat Pernyataan Kepala BPS untuk
// termin/tahun tertentu - dokumen tunggal, disimpan di tabel surat_kepala_se2026, bukan
// di rekap (rekap itu per-petugas, dokumen ini bukan milik petugas manapun).
func assignKepalaNomor(termin int, tahun, tanggal string) (string, error) {
	var existing sql.NullString
	err := database.DB.QueryRow(`SELECT nomor FROM surat_kepala_se2026 WHERE termin=? AND tahun=?`, termin, tahun).Scan(&existing)
	if err == nil && existing.Valid && existing.String != "" {
		return existing.String, nil
	}
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}

	tgl, err := time.Parse("2006-01-02", tanggal)
	if err != nil {
		return "", fmt.Errorf("format tanggal salah, gunakan YYYY-MM-DD")
	}
	nomor := kepalaNomorSE2026(tahun, tgl)
	_, err = database.DB.Exec(`INSERT INTO surat_kepala_se2026 (termin, tahun, nomor, tanggal) VALUES (?, ?, ?, ?)`,
		termin, tahun, nomor, tgl.Format("2006-01-02"))
	if err != nil {
		return "", err
	}
	return nomor, nil
}

// CreatePernyataanKepalaHandler POST /api/rekap/spk/se2026/pernyataan-kepala/create
func CreatePernyataanKepalaHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Tanggal string `json:"tanggal"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	if req.Tanggal == "" {
		req.Tanggal = "2026-08-15"
	}
	nomor, err := assignKepalaNomor(1, "2026", req.Tanggal)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok", "nomor": nomor})
}

// DownloadPernyataanKepalaSE2026Handler GET /api/rekap/spk/se2026/pernyataan-kepala/download?tanggal=2026-08-15
func DownloadPernyataanKepalaSE2026Handler(w http.ResponseWriter, r *http.Request) {
	tanggal := r.URL.Query().Get("tanggal")
	if tanggal == "" {
		tanggal = "2026-07-17"
	}

	var existing sql.NullString
	database.DB.QueryRow(`SELECT nomor FROM surat_kepala_se2026 WHERE termin=1 AND tahun='2026'`).Scan(&existing)
	if !existing.Valid || existing.String == "" {
		http.Error(w, "Surat Pernyataan Kepala BPS belum memiliki nomor - buat nomornya dulu", http.StatusBadRequest)
		return
	}

	d, lampiran, err := buildPernyataanKepalaSE2026(tanggal)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	d.Nomor = existing.String

	docx, err := generatePernyataanKepalaDocx(d, lampiran)
	if err != nil {
		http.Error(w, "Gagal generate Surat Pernyataan Kepala BPS: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fname := "Surat_Pernyataan_Kepala_BPS_SE2026.docx"
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, fname))
	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.wordprocessingml.document")
	w.Header().Set("Content-Length", strconv.Itoa(len(docx)))
	w.Write(docx)
}

// StatusPernyataanKepalaSE2026Handler GET /api/rekap/spk/se2026/pernyataan-kepala/status
func StatusPernyataanKepalaSE2026Handler(w http.ResponseWriter, r *http.Request) {
	var nomor sql.NullString
	database.DB.QueryRow(`SELECT nomor FROM surat_kepala_se2026 WHERE termin=1 AND tahun='2026'`).Scan(&nomor)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"sudah_ada": nomor.Valid && nomor.String != "",
		"nomor":     nomor.String,
	})
}
