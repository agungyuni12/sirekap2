package handlers

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/xuri/excelize/v2"
	"sirekap/internal/models"
)

// DashboardPenilaianHandler GET /api/penilaian/dashboard?kegiatan_id=&tahun=&peran=&kecamatan=
func DashboardPenilaianHandler(w http.ResponseWriter, r *http.Request) {
	stats, err := models.GetDashboardPenilaianStats(daftarPenilaianFilters(r))
	if err != nil {
		log.Printf("Error loading dashboard penilaian stats: %v", err)
		apiError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Gagal memuat statistik penilaian")
		return
	}
	writeJSON(w, http.StatusOK, stats)
}

// RekapDashboardHandler GET /api/penilaian/rekap-dashboard?kegiatan_id=&tahun=&peran=&kecamatan=&predikat=&search=
func RekapDashboardHandler(w http.ResponseWriter, r *http.Request) {
	items, err := models.GetRekapDashboard(daftarPenilaianFilters(r))
	if err != nil {
		log.Printf("Error loading rekap dashboard: %v", err)
		apiError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Gagal memuat rekapitulasi penilaian")
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{
		"data":       items,
		"total_data": len(items),
	})
}

// ExportRekapDashboardHandler GET /api/penilaian/rekap-dashboard/export?kegiatan_id=&tahun=&peran=&kecamatan=&predikat=
func ExportRekapDashboardHandler(w http.ResponseWriter, r *http.Request) {
	filter := daftarPenilaianFilters(r)
	items, err := models.GetRekapDashboard(filter)
	if err != nil {
		log.Printf("Error loading rekap dashboard for export: %v", err)
		http.Error(w, "Gagal mengambil data rekapitulasi penilaian", http.StatusInternalServerError)
		return
	}

	f := excelize.NewFile()
	defer f.Close()

	sheet := "Rekap Penilaian"
	f.SetSheetName("Sheet1", sheet)

	f.SetCellValue(sheet, "A1", "BADAN PUSAT STATISTIK KABUPATEN DOMPU")
	f.SetCellValue(sheet, "A2", "REKAPITULASI PENILAIAN KINERJA MITRA STATISTIK")
	period := "Semua Periode"
	if filter.Tahun != "" {
		period = fmt.Sprintf("Tahun Anggaran %s", filter.Tahun)
	}
	f.SetCellValue(sheet, "A3", period)
	f.SetCellValue(sheet, "A4", fmt.Sprintf("Diunduh: %s", time.Now().Format("02 January 2006 15:04")))
	titleStyle, _ := f.NewStyle(&excelize.Style{Font: &excelize.Font{Bold: true, Size: 13}})
	f.SetCellStyle(sheet, "A1", "A1", titleStyle)
	f.MergeCell(sheet, "A1", "L1")
	f.MergeCell(sheet, "A2", "L2")
	f.MergeCell(sheet, "A3", "L3")
	f.MergeCell(sheet, "A4", "L4")

	headers := []string{"No", "ID Sobat", "Nama Mitra", "Kegiatan", "Peran", "Kecamatan", "Skor Awal", "Status Konfirmasi", "Skor Ulang SM", "Skor Akhir", "Predikat", "Catatan"}
	headerRow := 6
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, headerRow)
		f.SetCellValue(sheet, cell, h)
	}
	headerStyle, _ := f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Bold: true, Color: "FFFFFF"},
		Fill:      excelize.Fill{Type: "pattern", Color: []string{"D97706"}, Pattern: 1},
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center", WrapText: true},
		Border: []excelize.Border{
			{Type: "left", Color: "E2E8F0", Style: 1},
			{Type: "right", Color: "E2E8F0", Style: 1},
			{Type: "bottom", Color: "E2E8F0", Style: 2},
		},
	})
	f.SetRowHeight(sheet, headerRow, 26)
	firstHeaderCell, _ := excelize.CoordinatesToCellName(1, headerRow)
	lastHeaderCell, _ := excelize.CoordinatesToCellName(len(headers), headerRow)
	f.SetCellStyle(sheet, firstHeaderCell, lastHeaderCell, headerStyle)

	predikatStyles := map[string]int{}
	predikatColors := map[string]string{
		"Sangat Baik":   "D1FAE5",
		"Baik":          "DBEAFE",
		"Cukup":         "FEF3C7",
		"Kurang":        "FFE4CE",
		"Sangat Kurang": "FEE2E2",
	}
	for predikatLabel, color := range predikatColors {
		style, _ := f.NewStyle(&excelize.Style{Fill: excelize.Fill{Type: "pattern", Color: []string{color}, Pattern: 1}})
		predikatStyles[predikatLabel] = style
	}

	row := headerRow + 1
	for i, item := range items {
		skorAwal := interface{}("-")
		if item.SkorAwal != nil {
			skorAwal = *item.SkorAwal
		}
		skorUlang := interface{}("-")
		if item.SkorUlang != nil {
			skorUlang = *item.SkorUlang
		}
		skorAkhir := interface{}("-")
		if item.SkorAkhir != nil {
			skorAkhir = *item.SkorAkhir
		}
		statusKonfirmasi := item.StatusKonfirmasi
		if statusKonfirmasi == "" {
			statusKonfirmasi = "-"
		}

		values := []interface{}{
			i + 1, item.IDSobat, item.NamaMitra, item.Kegiatan, item.Peran, item.Kecamatan,
			skorAwal, statusKonfirmasi, skorUlang, skorAkhir, item.Predikat, item.Catatan,
		}
		for j, v := range values {
			cell, _ := excelize.CoordinatesToCellName(j+1, row)
			f.SetCellValue(sheet, cell, v)
		}
		if style, ok := predikatStyles[item.Predikat]; ok {
			predikatCell, _ := excelize.CoordinatesToCellName(11, row)
			f.SetCellStyle(sheet, predikatCell, predikatCell, style)
		}
		row++
	}

	f.SetColWidth(sheet, "A", "A", 6)
	f.SetColWidth(sheet, "B", "B", 16)
	f.SetColWidth(sheet, "C", "D", 30)
	f.SetColWidth(sheet, "E", "E", 10)
	f.SetColWidth(sheet, "F", "F", 18)
	f.SetColWidth(sheet, "G", "K", 14)
	f.SetColWidth(sheet, "L", "L", 30)

	var buf bytes.Buffer
	if err := f.Write(&buf); err != nil {
		http.Error(w, "Gagal membuat file Excel", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Header().Set("Content-Disposition", `attachment; filename="rekap_penilaian_mitra.xlsx"`)
	w.Header().Set("Content-Length", fmt.Sprintf("%d", buf.Len()))
	w.Write(buf.Bytes())
}
