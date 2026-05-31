package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/xuri/excelize/v2"
	"sirekap/internal/database"
)

// DownloadTemplateSE2026Handler GET /api/rekap/spk/se2026/import/template
// Download template Excel untuk import petugas SE2026.
func DownloadTemplateSE2026Handler(w http.ResponseWriter, r *http.Request) {
	f := excelize.NewFile()
	defer f.Close()

	sheet := "Petugas SE2026"
	f.SetSheetName("Sheet1", sheet)

	headers := []string{
		"Nama Lengkap", "SOBAT ID", "Posisi (PCL/PML)",
		"Alamat", "Kecamatan",
	}
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, h)
	}

	// Style header
	style, _ := f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Bold: true, Color: "FFFFFF"},
		Fill:      excelize.Fill{Type: "pattern", Color: []string{"F97316"}, Pattern: 1},
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center", WrapText: true},
		Border: []excelize.Border{
			{Type: "left", Color: "E2E8F0", Style: 1},
			{Type: "right", Color: "E2E8F0", Style: 1},
			{Type: "bottom", Color: "E2E8F0", Style: 2},
		},
	})
	f.SetRowHeight(sheet, 1, 30)
	f.SetCellStyle(sheet, "A1", "E1", style)

	// Lebar kolom
	f.SetColWidth(sheet, "A", "A", 35)
	f.SetColWidth(sheet, "B", "B", 18)
	f.SetColWidth(sheet, "C", "C", 18)
	f.SetColWidth(sheet, "D", "D", 40)
	f.SetColWidth(sheet, "E", "E", 20)

	// Contoh data
	contoh := [][]interface{}{
		{"NAMA LENGKAP PETUGAS", "520522010001", "PCL", "Jl. Contoh No. 1, Desa ABC", "Dompu"},
		{"NAMA PEMERIKSA LAPANGAN", "520522010002", "PML", "Jl. Contoh No. 2, Desa DEF", "Woja"},
	}
	for i, row := range contoh {
		for j, val := range row {
			cell, _ := excelize.CoordinatesToCellName(j+1, i+2)
			f.SetCellValue(sheet, cell, val)
		}
	}

	// Validasi dropdown Posisi
	dv := excelize.NewDataValidation(true)
	dv.SetSqref("C2:C500")
	dv.SetDropList([]string{"PCL", "PML"})
	f.AddDataValidation(sheet, dv)

	var buf bytes.Buffer
	if err := f.Write(&buf); err != nil {
		http.Error(w, "Gagal buat template", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Header().Set("Content-Disposition", `attachment; filename="Template_Import_Petugas_SE2026.xlsx"`)
	w.Header().Set("Content-Length", fmt.Sprintf("%d", buf.Len()))
	w.Write(buf.Bytes())
}

// ImportPetugasSE2026Handler POST /api/rekap/spk/se2026/import
// Upload Excel berisi daftar petugas baru, insert ke mitra+rekap.
func ImportPetugasSE2026Handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "File terlalu besar (maks 10MB)", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "File tidak ditemukan", http.StatusBadRequest)
		return
	}
	defer file.Close()

	f, err := excelize.OpenReader(file)
	if err != nil {
		http.Error(w, "Format file tidak valid", http.StatusBadRequest)
		return
	}
	defer f.Close()

	sheetName := f.GetSheetName(0)
	rows, err := f.GetRows(sheetName)
	if err != nil || len(rows) < 2 {
		http.Error(w, "Sheet kosong atau format salah", http.StatusBadRequest)
		return
	}

	type Result struct {
		Baris  int    `json:"baris"`
		Nama   string `json:"nama"`
		Sobat  string `json:"sobat"`
		Status string `json:"status"`
		Pesan  string `json:"pesan"`
	}

	var results []Result
	inserted, skipped, errors := 0, 0, 0

	for i, row := range rows[1:] { // skip header
		baris := i + 2

		get := func(idx int) string {
			if idx < len(row) {
				return strings.TrimSpace(row[idx])
			}
			return ""
		}

		nama   := strings.ToUpper(get(0))
		sobat  := get(1)
		posisi := strings.ToUpper(get(2))
		alamat := get(3)
		kec    := get(4)

		if nama == "" || sobat == "" {
			continue
		}
		if posisi != "PCL" && posisi != "PML" {
			results = append(results, Result{baris, nama, sobat, "error", "Posisi harus PCL atau PML"})
			errors++
			continue
		}

		// Escape SQL
		esc := func(s string) string { return strings.ReplaceAll(s, "'", "''") }

		// Insert mitra
		_, err := database.DB.Exec(`
			INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
			SELECT ?, ?, '', ?, ?, '2026'
			WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat=? AND tahun='2026')`,
			sobat, nama, alamat, kec, sobat)
		if err != nil {
			results = append(results, Result{baris, nama, sobat, "error", err.Error()})
			errors++
			continue
		}

		kegiatan := "Pendataan Sensus Ekonomi 2026"
		honor    := 11572500
		hsatuan  := 4629000
		mak      := "2902.BMA.006.005.B.521213"

		if posisi == "PML" {
			kegiatan = "Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)"
			honor    = 12192500
			hsatuan  = 4877000
			mak      = "2902.FAN.ZZ1.051.A.521213"
		}

		// Insert rekap
		res, err := database.DB.Exec(`
			INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
			                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls)
			SELECT ?, 'ipds5205', ?, ?, 'Juni-Agustus', '2026', 2026,
			       ?, '2.5', 'O-B', ?, ?, '2026-06-15', '2026-08-31', 0
			WHERE NOT EXISTS (
			    SELECT 1 FROM rekap WHERE idsobat=? AND kegiatan=? AND tahun='2026'
			)`,
			sobat, esc(nama), kegiatan, honor, hsatuan, mak, sobat, kegiatan)

		if err != nil {
			results = append(results, Result{baris, nama, sobat, "error", err.Error()})
			errors++
			continue
		}

		affected, _ := res.RowsAffected()
		if affected == 0 {
			results = append(results, Result{baris, nama, sobat, "skip", "Sudah ada di rekap"})
			skipped++
		} else {
			results = append(results, Result{baris, nama, sobat, "ok", fmt.Sprintf("%s berhasil diimport", posisi)})
			inserted++
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"inserted": inserted,
		"skipped":  skipped,
		"errors":   errors,
		"results":  results,
	})
}
