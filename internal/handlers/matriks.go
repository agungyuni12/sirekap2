package handlers

import (
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
	"sirekap/internal/models"
)

type matriksConflictPayload struct {
	NIP            string `json:"nip"`
	TanggalMulai   string `json:"tanggalMulai"`
	TanggalSelesai string `json:"tanggalSelesai"`
	IDMatriks      int    `json:"idMatriks"`
}

var matriksImportHeaders = []string{
	"jenis_perjalanan",
	"kegiatan",
	"mak",
	"no_fp",
	"perihal_fp",
	"tgl_surtug_utama",
	"tgl_spd_utama",
	"nip",
	"nama",
	"golongan",
	"asal",
	"tujuan",
	"tanggal_mulai",
	"tanggal_selesai",
	"no_surtug",
	"no_spd",
	"jumlah",
}

// ListMatriksAdministrasiHandler handles GET /api/matriks/administrasi
func ListMatriksAdministrasiHandler(w http.ResponseWriter, r *http.Request) {
	filter := models.MatriksAdministrasiFilter{
		Search: strings.TrimSpace(r.URL.Query().Get("search")),
	}

	items, err := models.ListMatriksAdministrasi(filter)
	if err != nil {
		log.Printf("Error listing matriks administrasi: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	totalPeserta := 0
	totalLaporan := 0
	for _, item := range items {
		totalPeserta += item.JumlahPeserta
		totalLaporan += item.JumlahLaporan
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data":         items,
		"totalData":    len(items),
		"totalPeserta": totalPeserta,
		"totalLaporan": totalLaporan,
	})
}

// ListMatriksAdministrasiPesertaHandler handles GET /api/matriks/administrasi/peserta
func ListMatriksAdministrasiPesertaHandler(w http.ResponseWriter, r *http.Request) {
	noFP := strings.TrimSpace(r.URL.Query().Get("no_fp"))

	items, err := models.ListMatriksAdministrasiPeserta(noFP)
	if err != nil {
		var validationErr models.ValidationError
		if errors.As(err, &validationErr) {
			http.Error(w, validationErr.Error(), http.StatusBadRequest)
			return
		}

		log.Printf("Error listing matriks administrasi peserta: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data":      items,
		"totalData": len(items),
	})
}

// ListMatriksCalendarHandler handles GET /api/matriks/calendar
func ListMatriksCalendarHandler(w http.ResponseWriter, r *http.Request) {
	year, _ := strconv.Atoi(strings.TrimSpace(r.URL.Query().Get("year")))
	month, _ := strconv.Atoi(strings.TrimSpace(r.URL.Query().Get("month")))

	filter := models.MatriksCalendarFilter{
		Search: strings.TrimSpace(r.URL.Query().Get("search")),
		Jenis:  strings.TrimSpace(r.URL.Query().Get("jenis")),
		Year:   year,
		Month:  month,
	}

	items, err := models.ListMatriksCalendar(filter)
	if err != nil {
		var validationErr models.ValidationError
		if errors.As(err, &validationErr) {
			http.Error(w, validationErr.Error(), http.StatusBadRequest)
			return
		}

		log.Printf("Error listing matriks calendar: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	tahunOptions, err := models.GetMatriksTahunOptions()
	if err != nil {
		log.Printf("Error loading matriks calendar year options: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data":         items,
		"tahunOptions": tahunOptions,
		"totalData":    len(items),
	})
}

// ListMatriksHandler handles GET /api/matriks
func ListMatriksHandler(w http.ResponseWriter, r *http.Request) {
	filter := models.MatriksFilter{
		Search: strings.TrimSpace(r.URL.Query().Get("search")),
		Jenis:  strings.TrimSpace(r.URL.Query().Get("jenis")),
		Tahun:  strings.TrimSpace(r.URL.Query().Get("tahun")),
	}

	items, err := models.ListMatriks(filter)
	if err != nil {
		log.Printf("Error listing matriks: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	tahunOptions, err := models.GetMatriksTahunOptions()
	if err != nil {
		log.Printf("Error loading matriks year options: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	totalTransportLokal := 0
	totalPerjalananDinas := 0
	for _, item := range items {
		switch item.JenisPerjalanan {
		case "Transport Lokal":
			totalTransportLokal++
		case "Perjalanan Dinas":
			totalPerjalananDinas++
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data":                 items,
		"tahunOptions":         tahunOptions,
		"totalData":            len(items),
		"totalTransportLokal":  totalTransportLokal,
		"totalPerjalananDinas": totalPerjalananDinas,
	})
}

// SearchDaftarTranslokHandler handles GET /api/matriks/kegiatan/search
func SearchDaftarTranslokHandler(w http.ResponseWriter, r *http.Request) {
	query := strings.TrimSpace(r.URL.Query().Get("query"))
	if query == "" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]models.DaftarTranslokSearchResult{})
		return
	}

	results, err := models.SearchDaftarTranslok(query)
	if err != nil {
		log.Printf("Error searching daftartranslok: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if results == nil {
		results = []models.DaftarTranslokSearchResult{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

// SearchPegawaiHandler handles GET /api/pegawai/search
func SearchPegawaiHandler(w http.ResponseWriter, r *http.Request) {
	query := strings.TrimSpace(r.URL.Query().Get("query"))
	if query == "" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]models.PegawaiSearchResult{})
		return
	}

	results, err := models.SearchPegawai(query)
	if err != nil {
		log.Printf("Error searching pegawai: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if results == nil {
		results = []models.PegawaiSearchResult{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func DownloadMatriksImportTemplateHandler(w http.ResponseWriter, r *http.Request) {
	file := excelize.NewFile()
	sheetName := file.GetSheetName(file.GetActiveSheetIndex())

	headerCells := make([]interface{}, len(matriksImportHeaders))
	for i, header := range matriksImportHeaders {
		headerCells[i] = header
	}
	if err := file.SetSheetRow(sheetName, "A1", &headerCells); err != nil {
		http.Error(w, "Gagal membuat template Excel", http.StatusInternalServerError)
		return
	}

	sampleRow := []interface{}{
		"Transport Lokal",
		"Transport lokal petugas pemeriksaan lapangan updating listing (susenas maret organik)",
		"054.01.GG.2906.BMA.006.005.A.524113.1",
		"FP-2026-429310-92800-001",
		"Transport lokal petugas pemeriksaan lapangan Susenas Maret",
		"2026-03-01",
		"2026-03-01",
		"198001012006041001",
		"Nama Pegawai 1",
		"III/c",
		"Kabupaten Dompu",
		"Kecamatan Woja",
		"2026-03-03",
		"2026-03-03",
		"094/5205/III/2026",
		"001/SPD/III/2026",
		"170000",
	}
	if err := file.SetSheetRow(sheetName, "A2", &sampleRow); err != nil {
		http.Error(w, "Gagal membuat template Excel", http.StatusInternalServerError)
		return
	}

	if err := file.SetPanes(sheetName, &excelize.Panes{
		Freeze:      true,
		Split:       false,
		XSplit:      0,
		YSplit:      1,
		TopLeftCell: "A2",
		ActivePane:  "bottomLeft",
	}); err != nil {
		http.Error(w, "Gagal membuat template Excel", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Header().Set("Content-Disposition", `attachment; filename="template_import_matriks.xlsx"`)
	if err := file.Write(w); err != nil {
		http.Error(w, "Gagal mengirim template Excel", http.StatusInternalServerError)
		return
	}
}

func ImportMatriksHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "File import tidak valid", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "File CSV wajib dipilih", http.StatusBadRequest)
		return
	}
	defer file.Close()

	entries, rowErrors, err := parseImportedMatriksEntries(file, r.MultipartForm.File["file"][0].Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if len(rowErrors) > 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Terdapat baris import yang tidak valid",
			"errors":  rowErrors,
		})
		return
	}

	inserted, rowErrors, err := models.ImportMatriksEntries(entries)
	if err != nil {
		log.Printf("Error importing matriks: %v", err)
		http.Error(w, "Gagal memproses import matriks", http.StatusInternalServerError)
		return
	}

	if len(rowErrors) > 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Terdapat baris import yang tidak valid",
			"errors":  rowErrors,
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":  true,
		"message":  "Import matriks berhasil diproses",
		"inserted": inserted,
	})
}

// CheckMatriksConflictHandler handles POST /api/matriks/check-conflict
func CheckMatriksConflictHandler(w http.ResponseWriter, r *http.Request) {
	var payload matriksConflictPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Data JSON tidak valid", http.StatusBadRequest)
		return
	}

	payload.NIP = strings.TrimSpace(payload.NIP)
	payload.TanggalMulai = strings.TrimSpace(payload.TanggalMulai)
	payload.TanggalSelesai = strings.TrimSpace(payload.TanggalSelesai)

	if payload.NIP == "" || payload.TanggalMulai == "" || payload.TanggalSelesai == "" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(models.MatriksConflictResult{Bentrok: false, Pesan: ""})
		return
	}

	tanggalMulai, err := models.ParseMatriksDate(payload.TanggalMulai)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tanggalSelesai, err := models.ParseMatriksDate(payload.TanggalSelesai)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if tanggalSelesai.Before(tanggalMulai) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(models.MatriksConflictResult{
			Bentrok: true,
			Pesan:   "Tanggal selesai tidak boleh lebih awal dari tanggal mulai",
		})
		return
	}

	result, err := models.CheckMatriksConflict(payload.NIP, tanggalMulai, tanggalSelesai, payload.IDMatriks)
	if err != nil {
		log.Printf("Error checking matriks conflict: %v", err)
		http.Error(w, "Gagal mengecek jadwal", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// GetMatriksHandler handles GET /api/matriks/{id}
func GetMatriksHandler(w http.ResponseWriter, r *http.Request) {
	id, ok := parseIDParam(w, r)
	if !ok {
		return
	}

	item, err := models.GetMatriksByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Data matriks tidak ditemukan", http.StatusNotFound)
			return
		}
		log.Printf("Error getting matriks: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

// CreateMatriksHandler handles POST /api/matriks
func CreateMatriksHandler(w http.ResponseWriter, r *http.Request) {
	var payload models.MatriksCreatePayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Data JSON tidak valid", http.StatusBadRequest)
		return
	}

	inserted, err := models.CreateMatriksBatch(&payload)
	if err != nil {
		var validationErr models.ValidationError
		if errors.As(err, &validationErr) {
			http.Error(w, validationErr.Error(), http.StatusBadRequest)
			return
		}

		log.Printf("Error creating matriks: %v", err)
		http.Error(w, "Gagal menyimpan data matriks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":  true,
		"message":  "Data matriks berhasil disimpan",
		"inserted": inserted,
	})
}

// UpdateMatriksHandler handles PUT /api/matriks/{id}
func UpdateMatriksHandler(w http.ResponseWriter, r *http.Request) {
	id, ok := parseIDParam(w, r)
	if !ok {
		return
	}

	var payload models.MatriksEntryPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Data JSON tidak valid", http.StatusBadRequest)
		return
	}

	err := models.UpdateMatriks(id, &payload)
	if err != nil {
		var validationErr models.ValidationError
		if errors.As(err, &validationErr) {
			http.Error(w, validationErr.Error(), http.StatusBadRequest)
			return
		}

		log.Printf("Error updating matriks: %v", err)
		http.Error(w, "Gagal memperbarui data matriks", http.StatusInternalServerError)
		return
	}

	item, err := models.GetMatriksByID(id)
	if err != nil {
		log.Printf("Error loading updated matriks: %v", err)
		http.Error(w, "Data berhasil diperbarui tetapi gagal dimuat ulang", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Data matriks berhasil diperbarui",
		"data":    item,
	})
}

// DeleteMatriksHandler handles DELETE /api/matriks/{id}
func DeleteMatriksHandler(w http.ResponseWriter, r *http.Request) {
	id, ok := parseIDParam(w, r)
	if !ok {
		return
	}

	if err := models.DeleteMatriks(id); err != nil {
		log.Printf("Error deleting matriks: %v", err)
		http.Error(w, "Gagal menghapus data matriks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Data matriks berhasil dihapus",
	})
}

func normalizeImportHeader(value string) string {
	value = strings.TrimSpace(strings.ToLower(value))
	value = strings.TrimPrefix(value, "\ufeff")
	replacer := strings.NewReplacer(" ", "_", "-", "_")
	return replacer.Replace(value)
}

func csvValue(record []string, headerMap map[string]int, key string) string {
	index, ok := headerMap[key]
	if !ok || index >= len(record) {
		return ""
	}
	return strings.TrimSpace(record[index])
}

func isImportRecordEmpty(record []string) bool {
	for _, value := range record {
		if strings.TrimSpace(value) != "" {
			return false
		}
	}
	return true
}

func parseImportJumlah(value string) (int, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return 0, errors.New("empty jumlah")
	}
	return strconv.Atoi(value)
}

func parseImportedMatriksEntries(file io.Reader, filename string) ([]models.MatriksEntryPayload, []models.MatriksImportRowError, error) {
	ext := strings.ToLower(filepath.Ext(filename))
	switch ext {
	case ".xlsx":
		return parseMatriksExcelEntries(file)
	default:
		return parseMatriksCSVEntries(file)
	}
}

func parseMatriksCSVEntries(file io.Reader) ([]models.MatriksEntryPayload, []models.MatriksImportRowError, error) {
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	reader.TrimLeadingSpace = true

	header, err := reader.Read()
	if err != nil {
		if errors.Is(err, io.EOF) {
			return nil, nil, errors.New("file import kosong")
		}
		return nil, nil, errors.New("gagal membaca header CSV")
	}

	headerMap, err := buildImportHeaderMap(header)
	if err != nil {
		return nil, nil, err
	}

	entries := make([]models.MatriksEntryPayload, 0)
	rowErrors := make([]models.MatriksImportRowError, 0)
	rowNumber := 1
	for {
		record, err := reader.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, nil, fmt.Errorf("gagal membaca CSV di baris %d", rowNumber+1)
		}
		rowNumber++

		entry, rowErr, skip := buildImportEntryFromRecord(record, headerMap, rowNumber)
		if skip {
			continue
		}
		if rowErr != nil {
			rowErrors = append(rowErrors, *rowErr)
			continue
		}
		entries = append(entries, entry)
	}

	return entries, rowErrors, nil
}

func parseMatriksExcelEntries(file io.Reader) ([]models.MatriksEntryPayload, []models.MatriksImportRowError, error) {
	workbook, err := excelize.OpenReader(file)
	if err != nil {
		return nil, nil, errors.New("gagal membaca file Excel")
	}
	defer workbook.Close()

	sheets := workbook.GetSheetList()
	if len(sheets) == 0 {
		return nil, nil, errors.New("file Excel tidak memiliki sheet")
	}

	rows, err := workbook.GetRows(sheets[0])
	if err != nil {
		return nil, nil, errors.New("gagal membaca sheet Excel")
	}
	if len(rows) == 0 {
		return nil, nil, errors.New("file Excel kosong")
	}

	headerMap, err := buildImportHeaderMap(rows[0])
	if err != nil {
		return nil, nil, err
	}

	entries := make([]models.MatriksEntryPayload, 0)
	rowErrors := make([]models.MatriksImportRowError, 0)
	for rowIndex := 1; rowIndex < len(rows); rowIndex++ {
		record := rows[rowIndex]
		entry, rowErr, skip := buildImportEntryFromRecord(record, headerMap, rowIndex+1)
		if skip {
			continue
		}
		if rowErr != nil {
			rowErrors = append(rowErrors, *rowErr)
			continue
		}
		entries = append(entries, entry)
	}

	return entries, rowErrors, nil
}

func buildImportHeaderMap(header []string) (map[string]int, error) {
	headerMap := make(map[string]int, len(header))
	for index, column := range header {
		normalized := normalizeImportHeader(column)
		headerMap[normalized] = index
	}

	for _, required := range matriksImportHeaders {
		if _, ok := headerMap[required]; !ok {
			return nil, fmt.Errorf("kolom wajib tidak ditemukan: %s", required)
		}
	}

	return headerMap, nil
}

func buildImportEntryFromRecord(record []string, headerMap map[string]int, rowNumber int) (models.MatriksEntryPayload, *models.MatriksImportRowError, bool) {
	if isImportRecordEmpty(record) {
		return models.MatriksEntryPayload{}, nil, true
	}

	jumlah, parseErr := parseImportJumlah(csvValue(record, headerMap, "jumlah"))
	if parseErr != nil {
		return models.MatriksEntryPayload{}, &models.MatriksImportRowError{
			Row:     rowNumber,
			NIP:     csvValue(record, headerMap, "nip"),
			Nama:    csvValue(record, headerMap, "nama"),
			Message: "Jumlah harus berupa angka bulat",
		}, false
	}

	return models.MatriksEntryPayload{
		JenisPerjalanan: csvValue(record, headerMap, "jenis_perjalanan"),
		Kegiatan:        csvValue(record, headerMap, "kegiatan"),
		Mak:             csvValue(record, headerMap, "mak"),
		NoFP:            csvValue(record, headerMap, "no_fp"),
		PerihalFP:       csvValue(record, headerMap, "perihal_fp"),
		TglSurtugUtama:  csvValue(record, headerMap, "tgl_surtug_utama"),
		TglSpdUtama:     csvValue(record, headerMap, "tgl_spd_utama"),
		NIP:             csvValue(record, headerMap, "nip"),
		Nama:            csvValue(record, headerMap, "nama"),
		Golongan:        csvValue(record, headerMap, "golongan"),
		Asal:            csvValue(record, headerMap, "asal"),
		Tujuan:          csvValue(record, headerMap, "tujuan"),
		TanggalMulai:    csvValue(record, headerMap, "tanggal_mulai"),
		TanggalSelesai:  csvValue(record, headerMap, "tanggal_selesai"),
		NoSurtug:        csvValue(record, headerMap, "no_surtug"),
		NoSpd:           csvValue(record, headerMap, "no_spd"),
		Jumlah:          jumlah,
	}, nil, false
}
