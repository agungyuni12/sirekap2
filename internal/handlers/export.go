package handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/xuri/excelize/v2"
	"sirekap/internal/models"
)

// ExportRekapExcelHandler GET /api/rekap/export
// Exports rekap honor (filtered the same way as /api/rekap/filter) to an .xlsx
// file with a BPS-style header, ready to print.
func ExportRekapExcelHandler(w http.ResponseWriter, r *http.Request) {
	filter := models.RekapFilter{
		TahunAnggaran: r.URL.Query().Get("tahunAnggaran"),
		Bulan:         r.URL.Query().Get("bulan"),
		Search:        r.URL.Query().Get("search"),
		Page:          1,
		PerPage:       1000000, // export is unpaginated: pull every matching row
	}

	result, err := models.GetRekapFiltered(filter)
	if err != nil {
		http.Error(w, "Gagal mengambil data rekap", http.StatusInternalServerError)
		return
	}

	const sbmlLimit = 3346000

	f := excelize.NewFile()
	defer f.Close()

	sheet := "Rekap Honor"
	f.SetSheetName("Sheet1", sheet)

	// BPS-style report header
	f.SetCellValue(sheet, "A1", "BADAN PUSAT STATISTIK KABUPATEN DOMPU")
	f.SetCellValue(sheet, "A2", "REKAPITULASI HONOR MITRA STATISTIK")
	period := "Semua Periode"
	if filter.Bulan != "" || filter.TahunAnggaran != "" {
		period = fmt.Sprintf("Bulan %s Tahun Anggaran %s", filter.Bulan, filter.TahunAnggaran)
	}
	f.SetCellValue(sheet, "A3", period)
	f.SetCellValue(sheet, "A4", fmt.Sprintf("Diunduh: %s", time.Now().Format("02 January 2006 15:04")))
	titleStyle, _ := f.NewStyle(&excelize.Style{Font: &excelize.Font{Bold: true, Size: 13}})
	f.SetCellStyle(sheet, "A1", "A1", titleStyle)
	f.MergeCell(sheet, "A1", "H1")
	f.MergeCell(sheet, "A2", "H2")
	f.MergeCell(sheet, "A3", "H3")
	f.MergeCell(sheet, "A4", "H4")

	headers := []string{"No", "ID Sobat", "Nama Mitra", "Kegiatan", "Bulan", "Tahun", "Honor (Rp)", "Status SBML"}
	headerRow := 6
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, headerRow)
		f.SetCellValue(sheet, cell, h)
	}
	headerStyle, _ := f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Bold: true, Color: "FFFFFF"},
		Fill:      excelize.Fill{Type: "pattern", Color: []string{"0284C7"}, Pattern: 1},
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

	warnStyle, _ := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{Type: "pattern", Color: []string{"FEF3C7"}, Pattern: 1},
		Font: &excelize.Font{Color: "92400E"},
	})

	// Accumulate honor per mitra+bulan+tahun so "Status SBML" reflects the same
	// accumulated-limit rule used everywhere else in the app, not a per-row check.
	accumulated := map[string]float64{}
	for _, rec := range result.Data {
		key := fmt.Sprintf("%s|%s|%s", rec.IDSobat, rec.Bulan, rec.Tahun)
		accumulated[key] += rec.Honor
	}

	row := headerRow + 1
	for i, rec := range result.Data {
		key := fmt.Sprintf("%s|%s|%s", rec.IDSobat, rec.Bulan, rec.Tahun)
		status := "Normal"
		exceeded := accumulated[key] > sbmlLimit
		if exceeded {
			status = "Melebihi SBML"
		}

		values := []interface{}{i + 1, rec.IDSobat, rec.NamaMitra, rec.Kegiatan, rec.Bulan, rec.Tahun, rec.Honor, status}
		for j, v := range values {
			cell, _ := excelize.CoordinatesToCellName(j+1, row)
			f.SetCellValue(sheet, cell, v)
			if exceeded {
				f.SetCellStyle(sheet, cell, cell, warnStyle)
			}
		}
		row++
	}

	f.SetColWidth(sheet, "A", "A", 6)
	f.SetColWidth(sheet, "B", "B", 16)
	f.SetColWidth(sheet, "C", "C", 30)
	f.SetColWidth(sheet, "D", "D", 30)
	f.SetColWidth(sheet, "E", "F", 10)
	f.SetColWidth(sheet, "G", "G", 16)
	f.SetColWidth(sheet, "H", "H", 16)

	var buf bytes.Buffer
	if err := f.Write(&buf); err != nil {
		http.Error(w, "Gagal membuat file Excel", http.StatusInternalServerError)
		return
	}

	filename := fmt.Sprintf("Rekap_Honor_%s_%s.xlsx", filter.Bulan, filter.TahunAnggaran)
	if filter.Bulan == "" && filter.TahunAnggaran == "" {
		filename = "Rekap_Honor.xlsx"
	}

	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, filename))
	w.Header().Set("Content-Length", fmt.Sprintf("%d", buf.Len()))
	w.Write(buf.Bytes())
}
