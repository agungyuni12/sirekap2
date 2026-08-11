package handlers

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/xuri/excelize/v2"
	"sirekap/internal/models"
)

// DashboardPenilaianHandler GET /api/penilaian/dashboard?kegiatan_id=&tahun=
func DashboardPenilaianHandler(w http.ResponseWriter, r *http.Request) {
	kegiatanID, _ := strconv.Atoi(r.URL.Query().Get("kegiatan_id"))
	tahun := r.URL.Query().Get("tahun")

	stats, err := models.GetDashboardPenilaianStats(kegiatanID, tahun)
	if err != nil {
		log.Printf("Error loading dashboard penilaian stats: %v", err)
		apiError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Gagal memuat statistik penilaian")
		return
	}
	writeJSON(w, http.StatusOK, stats)
}

// rekapDashboardFilters reads the shared kegiatan_id/tahun/peran/predikat/search
// query params used by both the JSON listing and the Excel export.
func rekapDashboardFilters(r *http.Request) (kegiatanID int, tahun, peran, predikat, search string) {
	kegiatanID, _ = strconv.Atoi(r.URL.Query().Get("kegiatan_id"))
	tahun = r.URL.Query().Get("tahun")
	peran = r.URL.Query().Get("peran")
	predikat = r.URL.Query().Get("predikat")
	search = r.URL.Query().Get("search")
	return
}

// RekapDashboardHandler GET /api/penilaian/rekap-dashboard?kegiatan_id=&tahun=&peran=&predikat=&search=
func RekapDashboardHandler(w http.ResponseWriter, r *http.Request) {
	kegiatanID, tahun, peran, predikat, search := rekapDashboardFilters(r)

	items, err := models.GetRekapDashboard(kegiatanID, tahun, peran, predikat, search)
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

// ExportRekapDashboardHandler GET /api/penilaian/rekap-dashboard/export?kegiatan_id=&tahun=&peran=&predikat=
func ExportRekapDashboardHandler(w http.ResponseWriter, r *http.Request) {
	kegiatanID, tahun, peran, predikat, search := rekapDashboardFilters(r)

	items, err := models.GetRekapDashboard(kegiatanID, tahun, peran, predikat, search)
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
	if tahun != "" {
		period = fmt.Sprintf("Tahun Anggaran %s", tahun)
	}
	f.SetCellValue(sheet, "A3", period)
	f.SetCellValue(sheet, "A4", fmt.Sprintf("Diunduh: %s", time.Now().Format("02 January 2006 15:04")))
	titleStyle, _ := f.NewStyle(&excelize.Style{Font: &excelize.Font{Bold: true, Size: 13}})
	f.SetCellStyle(sheet, "A1", "A1", titleStyle)
	f.MergeCell(sheet, "A1", "J1")
	f.MergeCell(sheet, "A2", "J2")
	f.MergeCell(sheet, "A3", "J3")
	f.MergeCell(sheet, "A4", "J4")

	headers := []string{"No", "ID Sobat", "Nama Mitra", "Kegiatan", "Peran", "Skor Tahap 1", "Skor Tahap 2", "Skor Akhir", "Predikat", "Catatan"}
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
		"Sangat Baik":     "D1FAE5",
		"Baik":            "DBEAFE",
		"Cukup":           "FEF3C7",
		"Perlu Pembinaan": "FEE2E2",
	}
	for predikatLabel, color := range predikatColors {
		style, _ := f.NewStyle(&excelize.Style{Fill: excelize.Fill{Type: "pattern", Color: []string{color}, Pattern: 1}})
		predikatStyles[predikatLabel] = style
	}

	row := headerRow + 1
	for i, item := range items {
		skorTahap1 := interface{}("-")
		if item.SkorTahap1 != nil {
			skorTahap1 = *item.SkorTahap1
		}
		skorTahap2 := interface{}("-")
		if item.SkorTahap2 != nil {
			skorTahap2 = *item.SkorTahap2
		}

		values := []interface{}{
			i + 1, item.IDSobat, item.NamaMitra, item.Kegiatan, item.Peran,
			skorTahap1, skorTahap2, item.SkorAkhir, item.Predikat, item.Catatan,
		}
		for j, v := range values {
			cell, _ := excelize.CoordinatesToCellName(j+1, row)
			f.SetCellValue(sheet, cell, v)
		}
		if style, ok := predikatStyles[item.Predikat]; ok {
			predikatCell, _ := excelize.CoordinatesToCellName(9, row)
			f.SetCellStyle(sheet, predikatCell, predikatCell, style)
		}
		row++
	}

	f.SetColWidth(sheet, "A", "A", 6)
	f.SetColWidth(sheet, "B", "B", 16)
	f.SetColWidth(sheet, "C", "D", 30)
	f.SetColWidth(sheet, "E", "E", 10)
	f.SetColWidth(sheet, "F", "I", 14)
	f.SetColWidth(sheet, "J", "J", 30)

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
