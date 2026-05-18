package handlers

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"sirekap/internal/models"
)

type kegiatanMasterPayload struct {
	Nama      string  `json:"nama"`
	Vol       float64 `json:"vol"`
	Satuan    string  `json:"satuan"`
	Harga     float64 `json:"harga"`
	Mak       string  `json:"mak"`
	Tanggaran string  `json:"tanggaran"`
}

type suratManualPayload struct {
	IDSobat    string `json:"idsobat"`
	NamaMitra  string `json:"nmitra"`
	Bulan      string `json:"sbulan"`
	Tahun      string `json:"stahun"`
	Tujuan     string `json:"stujuan"`
	TglSurat   string `json:"tglsurat"`
	Keterangan string `json:"keterangan"`
}

// SearchKegiatanHandler handles GET /api/kegiatan/search
// Query params: query (search text), tahun (tahun anggaran)
func SearchKegiatanHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	tahun := r.URL.Query().Get("tahun")

	if query == "" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]models.KegiatanSearchResult{})
		return
	}

	results, err := models.SearchKegiatan(query, tahun)
	if err != nil {
		log.Printf("Error searching kegiatan: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if results == nil {
		results = []models.KegiatanSearchResult{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

// GetKegiatanDetailHandler handles GET /api/kegiatan/detail
// Query params: kegiatan (kegiatan name), tahun (tahun anggaran)
func GetKegiatanDetailHandler(w http.ResponseWriter, r *http.Request) {
	kegiatan := r.URL.Query().Get("kegiatan")
	tahun := r.URL.Query().Get("tahun")

	if kegiatan == "" {
		http.Error(w, "Parameter kegiatan tidak diberikan", http.StatusBadRequest)
		return
	}

	detail, err := models.GetKegiatanDetail(kegiatan, tahun)
	if err != nil {
		log.Printf("Error getting kegiatan detail: %v", err)
		w.Header().Set("Content-Type", "application/json")
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "Kegiatan tidak ditemukan"})
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Gagal memuat detail kegiatan"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(detail)
}

// ListKegiatanMasterHandler handles GET /api/master-kegiatan
func ListKegiatanMasterHandler(w http.ResponseWriter, r *http.Request) {
	filter := models.KegiatanFilter{
		Search: strings.TrimSpace(r.URL.Query().Get("search")),
		Tahun:  strings.TrimSpace(r.URL.Query().Get("tahun")),
	}

	items, err := models.ListKegiatan(filter)
	if err != nil {
		log.Printf("Error listing kegiatan master: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	tahunOptions, err := models.GetKegiatanTahunOptions()
	if err != nil {
		log.Printf("Error listing kegiatan years: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	var totalPagu float64
	for _, item := range items {
		totalPagu += item.Total
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data":         items,
		"tahunOptions": tahunOptions,
		"totalData":    len(items),
		"totalPagu":    totalPagu,
	})
}

// CreateKegiatanMasterHandler handles POST /api/master-kegiatan
func CreateKegiatanMasterHandler(w http.ResponseWriter, r *http.Request) {
	payload, ok := decodeAndValidateKegiatanPayload(w, r)
	if !ok {
		return
	}

	id, err := models.CreateKegiatan(&models.Kegiatan{
		Nama:      payload.Nama,
		Vol:       payload.Vol,
		Satuan:    payload.Satuan,
		Harga:     payload.Harga,
		Mak:       payload.Mak,
		Tanggaran: payload.Tanggaran,
	})
	if err != nil {
		log.Printf("Error creating kegiatan master: %v", err)
		http.Error(w, "Gagal menyimpan kegiatan", http.StatusInternalServerError)
		return
	}

	created, err := models.GetKegiatanByID(int(id))
	if err != nil {
		log.Printf("Error loading created kegiatan master: %v", err)
		http.Error(w, "Kegiatan berhasil dibuat tetapi gagal dimuat ulang", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Kegiatan berhasil ditambahkan",
		"data":    created,
	})
}

// GetMasterKegiatanHandler handles GET /api/master-kegiatan/{id}
func GetMasterKegiatanHandler(w http.ResponseWriter, r *http.Request) {
	id, ok := parseIDParam(w, r)
	if !ok {
		return
	}

	item, err := models.GetKegiatanByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Kegiatan tidak ditemukan", http.StatusNotFound)
			return
		}
		log.Printf("Error getting kegiatan master: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

// UpdateMasterKegiatanHandler handles PUT /api/master-kegiatan/{id}
func UpdateMasterKegiatanHandler(w http.ResponseWriter, r *http.Request) {
	id, ok := parseIDParam(w, r)
	if !ok {
		return
	}

	payload, ok := decodeAndValidateKegiatanPayload(w, r)
	if !ok {
		return
	}

	err := models.UpdateKegiatanMaster(id, &models.Kegiatan{
		Nama:      payload.Nama,
		Vol:       payload.Vol,
		Satuan:    payload.Satuan,
		Harga:     payload.Harga,
		Mak:       payload.Mak,
		Tanggaran: payload.Tanggaran,
	})
	if err != nil {
		log.Printf("Error updating kegiatan master: %v", err)
		http.Error(w, "Gagal memperbarui kegiatan", http.StatusInternalServerError)
		return
	}

	updated, err := models.GetKegiatanByID(id)
	if err != nil {
		log.Printf("Error loading updated kegiatan master: %v", err)
		http.Error(w, "Kegiatan berhasil diperbarui tetapi gagal dimuat ulang", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Kegiatan berhasil diperbarui",
		"data":    updated,
	})
}

// DeleteMasterKegiatanHandler handles DELETE /api/master-kegiatan/{id}
func DeleteMasterKegiatanHandler(w http.ResponseWriter, r *http.Request) {
	id, ok := parseIDParam(w, r)
	if !ok {
		return
	}

	err := models.DeleteKegiatanMaster(id)
	if err != nil {
		log.Printf("Error deleting kegiatan master: %v", err)
		http.Error(w, "Gagal menghapus kegiatan", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Kegiatan berhasil dihapus",
	})
}

// ListSuratHandler handles GET /api/surat
func ListSuratHandler(w http.ResponseWriter, r *http.Request) {
	filter := models.SuratFilter{
		Search: strings.TrimSpace(r.URL.Query().Get("search")),
		Tahun:  strings.TrimSpace(r.URL.Query().Get("tahun")),
		Tujuan: strings.TrimSpace(strings.ToUpper(r.URL.Query().Get("tujuan"))),
	}

	items, err := models.ListSurat(filter)
	if err != nil {
		log.Printf("Error listing surat: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	tahunOptions, err := models.GetSuratTahunOptions()
	if err != nil {
		log.Printf("Error loading surat year options: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	totalSPK := 0
	totalBAST := 0
	for _, item := range items {
		switch strings.ToUpper(item.Tujuan) {
		case "SPK":
			totalSPK++
		case "BAST":
			totalBAST++
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data":         items,
		"tahunOptions": tahunOptions,
		"totalData":    len(items),
		"totalSPK":     totalSPK,
		"totalBAST":    totalBAST,
	})
}

// CreateSuratManualHandler handles POST /api/surat
func CreateSuratManualHandler(w http.ResponseWriter, r *http.Request) {
	var payload suratManualPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		log.Printf("Error decoding surat payload: %v", err)
		http.Error(w, "Data JSON tidak valid", http.StatusBadRequest)
		return
	}

	payload.IDSobat = strings.TrimSpace(payload.IDSobat)
	payload.NamaMitra = strings.TrimSpace(payload.NamaMitra)
	payload.Bulan = strings.TrimSpace(payload.Bulan)
	payload.Tahun = strings.TrimSpace(payload.Tahun)
	payload.Tujuan = strings.TrimSpace(strings.ToUpper(payload.Tujuan))
	payload.TglSurat = strings.TrimSpace(payload.TglSurat)
	payload.Keterangan = strings.TrimSpace(payload.Keterangan)

	if payload.IDSobat == "" {
		http.Error(w, "IDSobat wajib diisi", http.StatusBadRequest)
		return
	}
	if payload.NamaMitra == "" {
		http.Error(w, "Nama mitra wajib diisi", http.StatusBadRequest)
		return
	}
	if payload.Bulan == "" {
		http.Error(w, "Bulan wajib dipilih", http.StatusBadRequest)
		return
	}
	if payload.Tahun == "" {
		http.Error(w, "Tahun wajib dipilih", http.StatusBadRequest)
		return
	}
	if payload.Tujuan != "SPK" && payload.Tujuan != "BAST" {
		http.Error(w, "Tipe surat tidak valid", http.StatusBadRequest)
		return
	}
	if payload.TglSurat == "" {
		http.Error(w, "Tanggal surat wajib diisi", http.StatusBadRequest)
		return
	}
	if payload.Tujuan == "BAST" && payload.Keterangan == "" {
		http.Error(w, "Nama kegiatan/keterangan wajib diisi untuk BAST", http.StatusBadRequest)
		return
	}

	nsurat, err := models.CreateSuratManual(
		payload.IDSobat,
		payload.NamaMitra,
		payload.Bulan,
		payload.Tahun,
		payload.Tujuan,
		payload.TglSurat,
		payload.Keterangan,
	)
	if err != nil {
		log.Printf("Error creating surat manual: %v", err)
		http.Error(w, "Gagal menyimpan nomor surat", http.StatusInternalServerError)
		return
	}

	items, err := models.ListSurat(models.SuratFilter{
		Search: nsurat,
	})
	if err != nil {
		log.Printf("Error loading created surat: %v", err)
		http.Error(w, "Nomor surat berhasil dibuat tetapi gagal dimuat ulang", http.StatusInternalServerError)
		return
	}

	var created interface{} = nil
	if len(items) > 0 {
		created = items[0]
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Nomor surat berhasil disimpan",
		"nsurat":  nsurat,
		"data":    created,
	})
}

// SearchMitraHandler handles GET /api/mitra/search
// Query params: query (search text), tahun (tahun)
func SearchMitraHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	tahun := r.URL.Query().Get("tahun")

	if query == "" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]models.MitraSearchResult{})
		return
	}

	results, err := models.SearchMitra(query, tahun)
	if err != nil {
		log.Printf("Error searching mitra: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if results == nil {
		results = []models.MitraSearchResult{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

// SubmitMitraHandler handles POST /api/mitra/submit
// Accepts form data matching the PHP version (posmitra.php)
func SubmitMitraHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Printf("Error parsing form: %v", err)
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	formData := &models.RekapFormData{
		Username:      r.FormValue("username2"),
		Kegiatan:      r.FormValue("kegiatan"),
		Mak:           r.FormValue("mak"),
		Bulan:         r.FormValue("bulan"),
		Tahun:         r.FormValue("tahun"),
		TahunAnggaran: r.FormValue("tahunAnggaran"),
		Reservation:   r.FormValue("reservation"),
		NamaMitra:     r.Form["namamitra[]"],
		Volume:        r.Form["volume[]"],
		Satuan:        r.Form["satuan[]"],
		HSatuan:       r.Form["hsatuan[]"],
	}

	// Fallback to raw values if formatted ones are not available
	if len(r.Form["hsatuan_raw[]"]) > 0 {
		formData.HSatuan = r.Form["hsatuan_raw[]"]
	}

	// Validate required fields
	if formData.Kegiatan == "" {
		http.Error(w, "Kegiatan belum diisi", http.StatusBadRequest)
		return
	}
	if formData.Bulan == "" {
		http.Error(w, "Bulan belum dipilih", http.StatusBadRequest)
		return
	}
	if formData.Tahun == "" {
		http.Error(w, "Tahun belum dipilih", http.StatusBadRequest)
		return
	}
	if len(formData.NamaMitra) == 0 {
		http.Error(w, "Harap tambahkan setidaknya satu mitra", http.StatusBadRequest)
		return
	}

	// Insert records
	err := models.InsertRekapBatch(formData)
	if err != nil {
		log.Printf("Error inserting rekap: %v", err)
		http.Error(w, "Gagal menyimpan data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": "Data berhasil disimpan",
	})
}

// FilterRekapHandler handles GET /api/rekap/filter
// Returns HTML cards grouped by mitra for HTMX swap
func FilterRekapHandler(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	tahunAnggaran := r.URL.Query().Get("tahunAnggaran")
	bulan := r.URL.Query().Get("bulan")
	search := r.URL.Query().Get("search")
	pageStr := r.URL.Query().Get("page")

	page := 1
	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}

	// Build filter
	filter := models.RekapFilter{
		TahunAnggaran: tahunAnggaran,
		Bulan:         bulan,
		Search:        search,
		Page:          page,
		PerPage:       20,
	}

	// Get grouped data by mitra
	result, err := models.GetMitraSummary(filter)
	if err != nil {
		log.Printf("Error filtering rekap: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Get summary
	totalMitra, totalHonor, totalKegiatan, err := models.GetRekapSummary(tahunAnggaran, bulan)
	if err != nil {
		log.Printf("Error getting rekap summary: %v", err)
		// Continue with zeros
	}

	// Set headers for Alpine.js to read
	w.Header().Set("X-Total-Mitra", strconv.Itoa(totalMitra))
	w.Header().Set("X-Total-Honor", fmt.Sprintf("%.0f", totalHonor))
	w.Header().Set("X-Total-Kegiatan", strconv.Itoa(totalKegiatan))
	w.Header().Set("X-Total-Data", strconv.Itoa(result.Total))
	w.Header().Set("X-Total-Pages", strconv.Itoa(result.TotalPages))

	// Calculate showing range
	start := (result.Page-1)*result.PerPage + 1
	end := start + len(result.Data) - 1
	if result.Total == 0 {
		start = 0
		end = 0
	}
	w.Header().Set("X-Showing", fmt.Sprintf("%d-%d", start, end))

	// Render HTML cards
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// SBML Limit constant
	const SBML_LIMIT = 3346000

	// Count SBML exceeded first
	sbmlExceededCount := 0
	for _, item := range result.Data {
		if item.TotalHonor > SBML_LIMIT {
			sbmlExceededCount++
		}
	}
	w.Header().Set("X-Sbml-Exceeded", strconv.Itoa(sbmlExceededCount))

	if len(result.Data) == 0 {
		// Empty state
		w.Write([]byte(`
			<div class="p-8 md:p-12 text-center">
				<div class="w-16 h-16 md:w-20 md:h-20 rounded-full bg-gray-100 flex items-center justify-center mx-auto mb-4">
					<i class="fas fa-inbox text-gray-400 text-2xl md:text-3xl"></i>
				</div>
				<h3 class="text-lg font-semibold text-gray-700 mb-1">Tidak Ada Data</h3>
				<p class="text-gray-500 text-sm mb-4">Tidak ada data yang sesuai dengan filter</p>
			</div>
		`))
		return
	}

	// Start container
	w.Write([]byte(`<div class="divide-y divide-gray-100">`))

	// Render data cards (grouped by mitra)
	for i, item := range result.Data {
		rowNum := start + i
		initials := getInitials(item.NamaMitra)
		honorFormatted := formatCurrency(item.TotalHonor)

		// Check if exceeds SBML limit
		exceedsSbml := item.TotalHonor > SBML_LIMIT
		cardClass := "p-4 hover:bg-gray-50 transition-colors cursor-pointer"
		honorClass := "font-bold text-green-600"
		honorText := fmt.Sprintf("Rp %s", honorFormatted)

		if exceedsSbml {
			cardClass = "p-4 sbml-warning transition-colors cursor-pointer"
			honorClass = "font-bold text-amber-600"
			excess := formatCurrency(item.TotalHonor - SBML_LIMIT)
			honorText = fmt.Sprintf("Rp %s <span class=\"text-xs text-red-500 font-normal\">(+%s)</span>", honorFormatted, excess)
		}

		// Build link to detail page with encoded parameters
		encodedData := base64.URLEncoding.EncodeToString([]byte(fmt.Sprintf("%s|%s|%s", item.IDSobat, item.Bulan, item.Tahun)))
		detailURL := fmt.Sprintf("/dashboard/detail?d=%s", encodedData)

		card := fmt.Sprintf(`
			<a href="%s" class="block %s">
				<div class="flex items-center justify-between">
					<div class="flex items-center space-x-3 flex-1 min-w-0">
						<!-- Number -->
						<span class="w-7 h-7 rounded-full bg-gray-100 text-gray-600 text-xs font-bold flex items-center justify-center flex-shrink-0">%d</span>
						
						<!-- Avatar -->
						<div class="w-10 h-10 rounded-full bg-gradient-to-br from-blue-500 to-indigo-600 text-white flex items-center justify-center text-sm font-bold shadow flex-shrink-0">%s</div>
						
						<!-- Info -->
						<div class="flex-1 min-w-0">
							<div class="flex items-center space-x-2 flex-wrap">
								<span class="px-2 py-0.5 rounded text-xs font-mono font-semibold bg-blue-50 text-blue-700">%s</span>
								<span class="font-semibold text-gray-900 truncate">%s</span>
							</div>
							<p class="text-xs text-gray-500 mt-0.5">%s %s</p>
						</div>
					</div>
					
					<!-- Total Honor & Arrow -->
					<div class="flex items-center space-x-3">
						<div class="text-right">
							<p class="text-xs text-gray-500 hidden sm:block">Total Honor</p>
							<p class="%s">%s</p>
						</div>
						<i class="fas fa-chevron-right text-gray-400"></i>
					</div>
				</div>
			</a>
		`, detailURL, cardClass,
			rowNum, initials,
			template.HTMLEscapeString(item.IDSobat), template.HTMLEscapeString(item.NamaMitra),
			item.Bulan, item.Tahun,
			honorClass, honorText)

		w.Write([]byte(card))
	}

	// Close container
	w.Write([]byte(`</div>`))
}

// GetMitraKegiatanHandler handles GET /api/rekap/kegiatan
// Returns HTML list of kegiatan for a specific mitra
func GetMitraKegiatanHandler(w http.ResponseWriter, r *http.Request) {
	idsobat := r.URL.Query().Get("idsobat")
	bulan := r.URL.Query().Get("bulan")
	tahun := r.URL.Query().Get("tahun")

	if idsobat == "" || bulan == "" || tahun == "" {
		http.Error(w, "Missing parameters", http.StatusBadRequest)
		return
	}

	kegiatan, err := models.GetMitraKegiatan(idsobat, bulan, tahun)
	if err != nil {
		log.Printf("Error getting mitra kegiatan: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if len(kegiatan) == 0 {
		w.Write([]byte(`<p class="text-sm text-gray-500 text-center py-2">Tidak ada kegiatan</p>`))
		return
	}

	// Render kegiatan list
	for _, k := range kegiatan {
		honorFormatted := formatCurrency(k.Honor)
		volumeFormatted := fmt.Sprintf("%.0f", k.Volume)

		row := fmt.Sprintf(`
			<div class="flex items-center justify-between bg-gray-50 rounded-lg p-3">
				<div class="flex-1 min-w-0">
					<p class="font-medium text-gray-800 text-sm truncate">%s</p>
					<p class="text-xs text-gray-500">%s %s × Rp %s</p>
				</div>
				<div class="text-right ml-3">
					<p class="font-semibold text-green-600 text-sm">Rp %s</p>
				</div>
			</div>
		`, template.HTMLEscapeString(k.Kegiatan),
			volumeFormatted, template.HTMLEscapeString(k.Satuan), formatCurrency(k.HSatuan),
			honorFormatted)

		w.Write([]byte(row))
	}
}

// getInitials extracts initials from a name
func getInitials(name string) string {
	if name == "" {
		return "?"
	}
	parts := strings.Fields(name)
	if len(parts) >= 2 {
		return strings.ToUpper(string(parts[0][0])) + strings.ToUpper(string(parts[1][0]))
	}
	if len(name) >= 2 {
		return strings.ToUpper(name[:2])
	}
	return strings.ToUpper(string(name[0]))
}

// formatCurrency formats a number as Indonesian currency
func formatCurrency(n float64) string {
	// Convert to integer for whole numbers
	intVal := int64(n)
	str := fmt.Sprintf("%d", intVal)

	// Add thousand separators
	result := ""
	length := len(str)
	for i, c := range str {
		if i > 0 && (length-i)%3 == 0 {
			result += "."
		}
		result += string(c)
	}
	return result
}

// DeleteKegiatanHandler handles DELETE /api/rekap/{id}
func DeleteKegiatanHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = models.DeleteRekapById(id)
	if err != nil {
		log.Printf("Error deleting rekap: %v", err)
		http.Error(w, "Failed to delete", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("")) // Empty response - HTMX will remove the element
}

// UpdateKegiatanHandler handles PUT /api/rekap/{id}
func UpdateKegiatanHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Parse JSON body
	var req struct {
		Kegiatan string  `json:"kegiatan"`
		Volume   float64 `json:"volume"`
		Satuan   string  `json:"satuan"`
		HSatuan  float64 `json:"hsatuan"`
		Honor    float64 `json:"honor"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Error decoding JSON: %v", err)
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	log.Printf("Update Rekap - kegiatan: %s, volume: %.2f, satuan: %s, hsatuan: %.2f, honor: %.2f",
		req.Kegiatan, req.Volume, req.Satuan, req.HSatuan, req.Honor)

	// If honor is 0, calculate from volume * hsatuan
	honor := req.Honor
	if honor == 0 && req.Volume > 0 && req.HSatuan > 0 {
		honor = req.Volume * req.HSatuan
	}

	err = models.UpdateRekap(id, req.Kegiatan, req.Volume, req.Satuan, req.HSatuan, honor)
	if err != nil {
		log.Printf("Error updating rekap: %v", err)
		http.Error(w, "Failed to update", http.StatusInternalServerError)
		return
	}

	// Return success response
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"success": true}`))
}

// GetKegiatanHandler handles GET /api/rekap/{id} - get single kegiatan for editing
func GetKegiatanHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	kegiatan, err := models.GetRekapById(id)
	if err != nil {
		log.Printf("Error getting rekap: %v", err)
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(kegiatan)
}

func decodeAndValidateKegiatanPayload(w http.ResponseWriter, r *http.Request) (*kegiatanMasterPayload, bool) {
	var payload kegiatanMasterPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		log.Printf("Error decoding kegiatan payload: %v", err)
		http.Error(w, "Data JSON tidak valid", http.StatusBadRequest)
		return nil, false
	}

	payload.Nama = strings.TrimSpace(payload.Nama)
	payload.Satuan = strings.TrimSpace(payload.Satuan)
	payload.Mak = strings.TrimSpace(payload.Mak)
	payload.Tanggaran = strings.TrimSpace(payload.Tanggaran)

	if payload.Nama == "" {
		http.Error(w, "Nama kegiatan wajib diisi", http.StatusBadRequest)
		return nil, false
	}

	if payload.Tanggaran == "" {
		http.Error(w, "Tahun anggaran wajib diisi", http.StatusBadRequest)
		return nil, false
	}

	if payload.Vol < 0 || payload.Harga < 0 {
		http.Error(w, "Volume dan harga tidak boleh negatif", http.StatusBadRequest)
		return nil, false
	}

	return &payload, true
}

func parseIDParam(w http.ResponseWriter, r *http.Request) (int, bool) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return 0, false
	}

	return id, true
}
