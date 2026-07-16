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
	BlnTeks  string
	TglAngka string
}

// buildPernyataanKepalaSE2026 mengagregasi seluruh petugas SE2026 (PPL+PML) tahun 2026
// dari rekap, digabung ke se2026.sls/progress untuk target/realisasi usaha+keluarga -
// nama yang tidak ketemu tunggal di se2026.users dilewati (sama seperti sync-realisasi).
func buildPernyataanKepalaSE2026(tanggal string) (kepalaSE2026Data, []usahaKeluargaLampiranRow, error) {
	var d kepalaSE2026Data

	tgl, err := time.Parse("2006-01-02", tanggal)
	if err != nil {
		tgl, _ = time.Parse("2006-01-02", "2026-08-15")
	}
	d.Hari = hariIndonesia[tgl.Weekday().String()]
	d.TglTeks = dayToTeks(tgl.Day())
	d.BlnTeks = bulanTeksSE[tgl.Month()]
	d.TglAngka = fmt.Sprintf("%02d-%02d", tgl.Day(), int(tgl.Month()))
	d.Nomor = kepalaNomorSE2026("2026", tgl)

	rows, err := database.DB.Query(`
		SELECT r.namamitra, r.kegiatan
		FROM rekap r
		WHERE r.kegiatan LIKE '%Sensus Ekonomi 2026%'
		  AND (r.tanggaran = 2026 OR r.tahun = '2026')
		ORDER BY r.kegiatan, r.namamitra`)
	if err != nil {
		return d, nil, fmt.Errorf("gagal mengambil daftar petugas: %v", err)
	}
	defer rows.Close()

	var lampiran []usahaKeluargaLampiranRow
	for rows.Next() {
		var nama, kegiatan string
		if err := rows.Scan(&nama, &kegiatan); err != nil {
			continue
		}
		nama = properCase(nama)
		jenis := jenisPetugasSE2026(kegiatan)
		jabatan, seRole := "Petugas Lapangan", "ppl"
		if jenis == "pml" {
			jabatan, seRole = "Pemeriksa Lapangan", "pml"
		}
		uid, cnt := findSE2026UserID(nama, seRole)
		if cnt != 1 {
			continue
		}
		target, realisasi, err := computeUsahaKeluargaSE2026(uid, jenis == "pml")
		if err != nil {
			continue
		}
		lampiran = append(lampiran, usahaKeluargaLampiranRow{
			Nama: nama, Jabatan: jabatan, TargetPrelist: target, Realisasi: realisasi,
		})
	}
	return d, lampiran, nil
}

// generateKepalaLampiranTableXML sama seperti generateLampiranTableXML tapi dengan
// tambahan kolom Jabatan (PML/PPL campur dalam satu tabel).
func generateKepalaLampiranTableXML(rows []usahaKeluargaLampiranRow) string {
	var rowXML string
	totalTarget, totalRealisasi := 0, 0
	for i, r := range rows {
		totalTarget += r.TargetPrelist
		totalRealisasi += r.Realisasi
		rowXML += fmt.Sprintf(`
<w:tr>
<w:tc><w:p><w:pPr><w:jc w:val="center"/></w:pPr><w:r><w:rPr>`+docFontRPr+`</w:rPr><w:t>%d</w:t></w:r></w:p></w:tc>
<w:tc><w:p><w:r><w:rPr>`+docFontRPr+`</w:rPr><w:t>%s</w:t></w:r></w:p></w:tc>
<w:tc><w:p><w:pPr><w:jc w:val="center"/></w:pPr><w:r><w:rPr>`+docFontRPr+`</w:rPr><w:t>%s</w:t></w:r></w:p></w:tc>
<w:tc><w:p><w:pPr><w:jc w:val="center"/></w:pPr><w:r><w:rPr>`+docFontRPr+`</w:rPr><w:t>%d</w:t></w:r></w:p></w:tc>
<w:tc><w:p><w:pPr><w:jc w:val="center"/></w:pPr><w:r><w:rPr>`+docFontRPr+`</w:rPr><w:t>%d</w:t></w:r></w:p></w:tc>
<w:tc><w:p><w:pPr><w:jc w:val="center"/></w:pPr><w:r><w:rPr>`+docFontRPr+`</w:rPr><w:t>%.1f</w:t></w:r></w:p></w:tc>
</w:tr>`, i+1, escapeXML(r.Nama), escapeXML(r.Jabatan), r.TargetPrelist, r.Realisasi, persentase(r.Realisasi, r.TargetPrelist))
	}
	rowXML += fmt.Sprintf(`
<w:tr>
<w:tc><w:tcPr><w:gridSpan w:val="3"/></w:tcPr><w:p><w:r><w:rPr>`+docFontRPrBold+`</w:rPr><w:t>Jumlah</w:t></w:r></w:p></w:tc>
<w:tc><w:p><w:pPr><w:jc w:val="center"/></w:pPr><w:r><w:rPr>`+docFontRPrBold+`</w:rPr><w:t>%d</w:t></w:r></w:p></w:tc>
<w:tc><w:p><w:pPr><w:jc w:val="center"/></w:pPr><w:r><w:rPr>`+docFontRPrBold+`</w:rPr><w:t>%d</w:t></w:r></w:p></w:tc>
<w:tc><w:p><w:pPr><w:jc w:val="center"/></w:pPr><w:r><w:rPr>`+docFontRPrBold+`</w:rPr><w:t>%.1f</w:t></w:r></w:p></w:tc>
</w:tr>`, totalTarget, totalRealisasi, persentase(totalRealisasi, totalTarget))

	return `<w:tbl>
<w:tblPr>
<w:tblStyle w:val="TableGrid"/>
<w:tblW w:w="0" w:type="auto"/>
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
<w:gridCol w:w="600"/>
<w:gridCol w:w="2600"/>
<w:gridCol w:w="1600"/>
<w:gridCol w:w="1600"/>
<w:gridCol w:w="2000"/>
<w:gridCol w:w="1300"/>
</w:tblGrid>
<w:tr>
<w:tc><w:p><w:r><w:rPr>` + docFontRPrBold + `</w:rPr><w:t>No</w:t></w:r></w:p></w:tc>
<w:tc><w:p><w:r><w:rPr>` + docFontRPrBold + `</w:rPr><w:t>Nama Petugas</w:t></w:r></w:p></w:tc>
<w:tc><w:p><w:r><w:rPr>` + docFontRPrBold + `</w:rPr><w:t>Jabatan</w:t></w:r></w:p></w:tc>
<w:tc><w:p><w:r><w:rPr>` + docFontRPrBold + `</w:rPr><w:t>Target Prelist</w:t></w:r></w:p></w:tc>
<w:tc><w:p><w:r><w:rPr>` + docFontRPrBold + `</w:rPr><w:t>Realisasi Hasil Pendataan (Usaha+Keluarga)</w:t></w:r></w:p></w:tc>
<w:tc><w:p><w:r><w:rPr>` + docFontRPrBold + `</w:rPr><w:t>Presentase (%)</w:t></w:r></w:p></w:tc>
</w:tr>` + rowXML + `</w:tbl>`
}

func generatePernyataanKepalaDocx(d kepalaSE2026Data, lampiran []usahaKeluargaLampiranRow) ([]byte, error) {
	return replaceInDocx("static/templates-se2026/surat_pernyataan_kepala_se2026.docx", map[string]string{
		"nomor":       d.Nomor,
		"hari":        d.Hari,
		"tgl_teks":    d.TglTeks,
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
		tanggal = "2026-08-15"
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
