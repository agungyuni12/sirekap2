package models

import (
	"strings"
	"time"

	"sirekap/internal/database"
)

// Rekap represents a rekap (summary) record for mitra payments
type Rekap struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	IDSobat   string    `json:"idsobat"`
	NamaMitra string    `json:"namamitra"`
	Kegiatan  string    `json:"kegiatan"`
	Bulan     string    `json:"bulan"`
	Tahun     string    `json:"tahun"`
	Tanggaran string    `json:"tanggaran"` // tahun anggaran
	Volume    float64   `json:"volume"`
	Satuan    string    `json:"satuan"`
	HSatuan   float64   `json:"hsatuan"`
	Honor     float64   `json:"honor"`
	Mak       string    `json:"mak"`
	PMWaktu   time.Time `json:"pmwaktu"` // start date
	PSWaktu   time.Time `json:"pswaktu"` // end date
}

// RekapFormData represents form data submitted for rekap
type RekapFormData struct {
	Username      string   `json:"username2"`
	Kegiatan      string   `json:"kegiatan"`
	Mak           string   `json:"mak"`
	Bulan         string   `json:"bulan"`
	Tahun         string   `json:"tahun"`
	TahunAnggaran string   `json:"tahunAnggaran"`
	Reservation   string   `json:"reservation"` // date range "DD/MM/YYYY - DD/MM/YYYY"
	NamaMitra     []string `json:"namamitra[]"`
	Volume        []string `json:"volume[]"`
	Satuan        []string `json:"satuan[]"`
	HSatuan       []string `json:"hsatuan[]"`
}

// InsertRekap inserts a single rekap record
func InsertRekap(rekap *Rekap) error {
	sqlQuery := `
		INSERT INTO rekap (username, idsobat, namamitra, kegiatan, bulan, tahun, tanggaran, volume, satuan, hsatuan, honor, mak, pmwaktu, pswaktu) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := database.DB.Exec(sqlQuery,
		rekap.Username,
		rekap.IDSobat,
		rekap.NamaMitra,
		rekap.Kegiatan,
		rekap.Bulan,
		rekap.Tahun,
		rekap.Tanggaran,
		rekap.Volume,
		rekap.Satuan,
		rekap.HSatuan,
		rekap.Honor,
		rekap.Mak,
		rekap.PMWaktu,
		rekap.PSWaktu,
	)

	return err
}

// InsertRekapBatch inserts multiple rekap records from form data
func InsertRekapBatch(formData *RekapFormData) error {
	// Parse date range from reservation (format: "DD/MM/YYYY - DD/MM/YYYY")
	var pmwaktu, pswaktu time.Time
	if formData.Reservation != "" {
		dates := strings.Split(formData.Reservation, " - ")
		if len(dates) == 2 {
			pmwaktu, _ = time.Parse("02/01/2006", strings.TrimSpace(dates[0]))
			pswaktu, _ = time.Parse("02/01/2006", strings.TrimSpace(dates[1]))
		}
	}

	// Prepare statement for batch insert
	sqlQuery := `
		INSERT INTO rekap (username, idsobat, namamitra, kegiatan, bulan, tahun, tanggaran, volume, satuan, hsatuan, honor, mak, pmwaktu, pswaktu) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	stmt, err := database.DB.Prepare(sqlQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Loop through each mitra entry
	for i := 0; i < len(formData.NamaMitra); i++ {
		// Skip empty entries
		if formData.NamaMitra[i] == "" {
			continue
		}

		// Parse namamitra to extract idsobat and name
		// Format: "IDSOBAT - Name" or "IDSOBAT/Name"
		namamitraFull := formData.NamaMitra[i]
		var idsobat, nmitra string

		if strings.Contains(namamitraFull, " - ") {
			parts := strings.SplitN(namamitraFull, " - ", 2)
			idsobat = strings.TrimSpace(parts[0])
			if len(parts) > 1 {
				nmitra = strings.TrimSpace(parts[1])
			} else {
				nmitra = idsobat
			}
		} else if strings.Contains(namamitraFull, "/") {
			parts := strings.SplitN(namamitraFull, "/", 2)
			idsobat = strings.TrimSpace(parts[0])
			if len(parts) > 1 {
				nmitra = strings.TrimSpace(parts[1])
			} else {
				nmitra = idsobat
			}
		} else {
			idsobat = namamitraFull
			nmitra = namamitraFull
		}

		// Parse volume
		var volume float64
		if len(formData.Volume) > i {
			volume = parseFloat(formData.Volume[i])
		}

		// Parse hsatuan
		var hsatuan float64
		if len(formData.HSatuan) > i {
			hsatuan = parseFloat(formData.HSatuan[i])
		}

		// Get satuan
		var satuan string
		if len(formData.Satuan) > i {
			satuan = formData.Satuan[i]
		}

		// Calculate honor
		honor := volume * hsatuan

		// Execute insert
		_, err := stmt.Exec(
			formData.Username,
			idsobat,
			nmitra,
			formData.Kegiatan,
			formData.Bulan,
			formData.Tahun,
			formData.TahunAnggaran,
			volume,
			satuan,
			hsatuan,
			honor,
			formData.Mak,
			pmwaktu,
			pswaktu,
		)

		if err != nil {
			return err
		}
	}

	return nil
}

// parseFloat is a helper to parse string to float64
func parseFloat(s string) float64 {
	// Remove any currency formatting (commas, periods as thousands separator)
	s = strings.ReplaceAll(s, ".", "")
	s = strings.ReplaceAll(s, ",", "")
	s = strings.TrimSpace(s)

	if s == "" {
		return 0
	}

	var result float64
	// Simple parsing - handle potential errors gracefully
	for _, c := range s {
		if c >= '0' && c <= '9' {
			result = result*10 + float64(c-'0')
		}
	}
	return result
}

// RekapFilter represents filter parameters for rekap query
type RekapFilter struct {
	TahunAnggaran string
	Bulan         string
	Status        string
	Search        string
	Page          int
	PerPage       int
}

// RekapResult contains filtered rekap data and metadata
type RekapResult struct {
	Data       []Rekap
	Total      int
	TotalHonor float64
	Page       int
	PerPage    int
	TotalPages int
}

// GetRekapFiltered retrieves rekap data with filters
func GetRekapFiltered(filter RekapFilter) (*RekapResult, error) {
	// Set defaults
	if filter.Page < 1 {
		filter.Page = 1
	}
	if filter.PerPage < 1 {
		filter.PerPage = 10
	}

	// Build query
	baseQuery := "FROM rekap WHERE 1=1"
	args := []interface{}{}

	if filter.TahunAnggaran != "" {
		baseQuery += " AND tanggaran = ?"
		args = append(args, filter.TahunAnggaran)
	}
	if filter.Bulan != "" {
		baseQuery += " AND bulan = ?"
		args = append(args, filter.Bulan)
	}
	if filter.Search != "" {
		searchClause, searchArgs := buildFlexibleSearchClause(
			[]string{"namamitra", "kegiatan", "idsobat"},
			filter.Search,
		)
		baseQuery += searchClause
		args = append(args, searchArgs...)
	}

	// Get total count
	var total int
	countQuery := "SELECT COUNT(*) " + baseQuery
	err := database.DB.QueryRow(countQuery, args...).Scan(&total)
	if err != nil {
		return nil, err
	}

	// Get total honor
	var totalHonor float64
	sumQuery := "SELECT COALESCE(SUM(honor), 0) " + baseQuery
	err = database.DB.QueryRow(sumQuery, args...).Scan(&totalHonor)
	if err != nil {
		return nil, err
	}

	// Calculate pagination
	totalPages := (total + filter.PerPage - 1) / filter.PerPage
	offset := (filter.Page - 1) * filter.PerPage

	// Get data with pagination
	dataQuery := "SELECT id, username, idsobat, namamitra, kegiatan, bulan, tahun, tanggaran, volume, satuan, hsatuan, honor, mak, pmwaktu, pswaktu " + baseQuery + " ORDER BY id DESC LIMIT ? OFFSET ?"
	args = append(args, filter.PerPage, offset)

	rows, err := database.DB.Query(dataQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []Rekap
	for rows.Next() {
		var r Rekap
		err := rows.Scan(
			&r.ID, &r.Username, &r.IDSobat, &r.NamaMitra, &r.Kegiatan,
			&r.Bulan, &r.Tahun, &r.Tanggaran, &r.Volume, &r.Satuan,
			&r.HSatuan, &r.Honor, &r.Mak, &r.PMWaktu, &r.PSWaktu,
		)
		if err != nil {
			return nil, err
		}
		data = append(data, r)
	}

	return &RekapResult{
		Data:       data,
		Total:      total,
		TotalHonor: totalHonor,
		Page:       filter.Page,
		PerPage:    filter.PerPage,
		TotalPages: totalPages,
	}, nil
}

// GetRekapSummary returns summary statistics for rekap
func GetRekapSummary(tahunAnggaran, bulan string) (totalMitra int, totalHonor float64, totalKegiatan int, err error) {
	baseQuery := "FROM rekap WHERE 1=1"
	args := []interface{}{}

	if tahunAnggaran != "" {
		baseQuery += " AND tanggaran = ?"
		args = append(args, tahunAnggaran)
	}
	if bulan != "" {
		baseQuery += " AND bulan = ?"
		args = append(args, bulan)
	}

	// Total distinct mitra
	mitraQuery := "SELECT COUNT(DISTINCT idsobat) " + baseQuery
	err = database.DB.QueryRow(mitraQuery, args...).Scan(&totalMitra)
	if err != nil {
		return
	}

	// Total honor
	honorQuery := "SELECT COALESCE(SUM(honor), 0) " + baseQuery
	err = database.DB.QueryRow(honorQuery, args...).Scan(&totalHonor)
	if err != nil {
		return
	}

	// Total distinct kegiatan
	kegiatanQuery := "SELECT COUNT(DISTINCT kegiatan) " + baseQuery
	err = database.DB.QueryRow(kegiatanQuery, args...).Scan(&totalKegiatan)
	if err != nil {
		return
	}

	return
}

// MitraSummary represents grouped data by mitra with total honor
type MitraSummary struct {
	IDSobat    string  `json:"idsobat"`
	NamaMitra  string  `json:"namamitra"`
	Bulan      string  `json:"bulan"`
	Tahun      string  `json:"tahun"`
	TotalHonor float64 `json:"total_honor"`
}

// MitraSummaryResult contains grouped mitra data and metadata
type MitraSummaryResult struct {
	Data       []MitraSummary
	Total      int
	TotalHonor float64
	Page       int
	PerPage    int
	TotalPages int
}

type RekapDownloadGroup struct {
	IDSobat   string
	NamaMitra string
	Bulan     string
	Tahun     string
}

type RekapBASTDownloadItem struct {
	ID        int
	IDSobat   string
	NamaMitra string
	Kegiatan  string
	Bulan     string
	Tahun     string
}

func buildMitraSummaryWhereClause(filter RekapFilter) (string, []interface{}) {
	whereClause := " WHERE 1=1"
	args := []interface{}{}

	if filter.TahunAnggaran != "" {
		whereClause += " AND tanggaran = ?"
		args = append(args, filter.TahunAnggaran)
	}
	if filter.Bulan != "" {
		whereClause += " AND bulan = ?"
		args = append(args, filter.Bulan)
	}
	if filter.Search != "" {
		searchClause, searchArgs := buildFlexibleSearchClause(
			[]string{"namamitra", "idsobat"},
			filter.Search,
		)
		whereClause += searchClause
		args = append(args, searchArgs...)
	}

	return whereClause, args
}

// GetMitraSummary retrieves grouped rekap data by mitra (like PHP GROUP BY idsobat)
func GetMitraSummary(filter RekapFilter) (*MitraSummaryResult, error) {
	// Set defaults
	if filter.Page < 1 {
		filter.Page = 1
	}
	if filter.PerPage < 1 {
		filter.PerPage = 10
	}

	whereClause, args := buildMitraSummaryWhereClause(filter)

	// Count grouped cards exactly as rendered in the data query.
	var total int
	countQuery := `
		SELECT COUNT(*) FROM (
			SELECT idsobat, namamitra, bulan, tahun
			FROM rekap` + whereClause + `
			GROUP BY idsobat, namamitra, bulan, tahun
		) AS grouped_rekap`
	err := database.DB.QueryRow(countQuery, args...).Scan(&total)
	if err != nil {
		return nil, err
	}

	// Get total honor
	var totalHonor float64
	sumQuery := "SELECT COALESCE(SUM(honor), 0) FROM rekap" + whereClause
	err = database.DB.QueryRow(sumQuery, args...).Scan(&totalHonor)
	if err != nil {
		return nil, err
	}

	// Calculate pagination
	totalPages := (total + filter.PerPage - 1) / filter.PerPage
	offset := (filter.Page - 1) * filter.PerPage

	// Get grouped data with pagination
	dataQuery := `
		SELECT idsobat, namamitra, bulan, tahun, SUM(honor) as total_honor 
		FROM rekap` + whereClause + `
		GROUP BY idsobat, namamitra, bulan, tahun 
		ORDER BY namamitra
		LIMIT ? OFFSET ?`
	args = append(args, filter.PerPage, offset)

	rows, err := database.DB.Query(dataQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []MitraSummary
	for rows.Next() {
		var m MitraSummary
		err := rows.Scan(&m.IDSobat, &m.NamaMitra, &m.Bulan, &m.Tahun, &m.TotalHonor)
		if err != nil {
			return nil, err
		}
		data = append(data, m)
	}

	return &MitraSummaryResult{
		Data:       data,
		Total:      total,
		TotalHonor: totalHonor,
		Page:       filter.Page,
		PerPage:    filter.PerPage,
		TotalPages: totalPages,
	}, nil
}

func GetFilteredMitraGroups(filter RekapFilter) ([]RekapDownloadGroup, error) {
	whereClause, args := buildMitraSummaryWhereClause(filter)

	query := `
		SELECT idsobat, namamitra, bulan, tahun
		FROM rekap` + whereClause + `
		GROUP BY idsobat, namamitra, bulan, tahun
		ORDER BY namamitra, tahun, bulan`

	rows, err := database.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []RekapDownloadGroup
	for rows.Next() {
		var item RekapDownloadGroup
		if err := rows.Scan(&item.IDSobat, &item.NamaMitra, &item.Bulan, &item.Tahun); err != nil {
			return nil, err
		}
		data = append(data, item)
	}

	return data, nil
}

func GetFilteredBASTDownloadItems(filter RekapFilter) ([]RekapBASTDownloadItem, error) {
	whereClause, args := buildMitraSummaryWhereClause(filter)

	query := `
		SELECT id, idsobat, namamitra, kegiatan, bulan, tahun
		FROM rekap` + whereClause + `
		ORDER BY namamitra, tahun, bulan, kegiatan, id`

	rows, err := database.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []RekapBASTDownloadItem
	for rows.Next() {
		var item RekapBASTDownloadItem
		if err := rows.Scan(&item.ID, &item.IDSobat, &item.NamaMitra, &item.Kegiatan, &item.Bulan, &item.Tahun); err != nil {
			return nil, err
		}
		data = append(data, item)
	}

	return data, nil
}

// MitraKegiatan represents a single kegiatan for a mitra
type MitraKegiatan struct {
	ID       int       `json:"id"`
	Kegiatan string    `json:"kegiatan"`
	Volume   float64   `json:"volume"`
	Satuan   string    `json:"satuan"`
	HSatuan  float64   `json:"hsatuan"`
	Honor    float64   `json:"honor"`
	Mak      string    `json:"mak"`
	PMWaktu  time.Time `json:"pmwaktu"`
	PSWaktu  time.Time `json:"pswaktu"`
}

// GetMitraKegiatan retrieves all kegiatan for a specific mitra in a period
func GetMitraKegiatan(idsobat, bulan, tahun string) ([]MitraKegiatan, error) {
	query := `
		SELECT id, kegiatan, volume, satuan, hsatuan, honor, mak, pmwaktu, pswaktu 
		FROM rekap 
		WHERE idsobat = ? AND bulan = ? AND tahun = ?
		ORDER BY kegiatan`

	rows, err := database.DB.Query(query, idsobat, bulan, tahun)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []MitraKegiatan
	for rows.Next() {
		var k MitraKegiatan
		err := rows.Scan(&k.ID, &k.Kegiatan, &k.Volume, &k.Satuan, &k.HSatuan, &k.Honor, &k.Mak, &k.PMWaktu, &k.PSWaktu)
		if err != nil {
			return nil, err
		}
		data = append(data, k)
	}

	return data, nil
}

// DeleteRekapById deletes a rekap record by ID
func DeleteRekapById(id int) error {
	_, err := database.DB.Exec("DELETE FROM rekap WHERE id = ?", id)
	return err
}

// UpdateRekap updates a rekap record
func UpdateRekap(id int, kegiatan string, volume float64, satuan string, hsatuan, honor float64) error {
	_, err := database.DB.Exec(
		"UPDATE rekap SET kegiatan = ?, volume = ?, satuan = ?, hsatuan = ?, honor = ? WHERE id = ?",
		kegiatan, volume, satuan, hsatuan, honor, id,
	)
	return err
}

// GetRekapById retrieves a single rekap record by ID
func GetRekapById(id int) (*MitraKegiatan, error) {
	var k MitraKegiatan
	err := database.DB.QueryRow(
		"SELECT id, kegiatan, volume, satuan, hsatuan, honor, mak, pmwaktu, pswaktu FROM rekap WHERE id = ?",
		id,
	).Scan(&k.ID, &k.Kegiatan, &k.Volume, &k.Satuan, &k.HSatuan, &k.Honor, &k.Mak, &k.PMWaktu, &k.PSWaktu)
	if err != nil {
		return nil, err
	}
	return &k, nil
}

func GetRekapRecordByID(id int) (*Rekap, error) {
	var r Rekap
	err := database.DB.QueryRow(
		`SELECT id, username, idsobat, namamitra, kegiatan, bulan, tahun, tanggaran, volume, satuan, hsatuan, honor, mak, pmwaktu, pswaktu
		 FROM rekap
		 WHERE id = ?`,
		id,
	).Scan(
		&r.ID, &r.Username, &r.IDSobat, &r.NamaMitra, &r.Kegiatan,
		&r.Bulan, &r.Tahun, &r.Tanggaran, &r.Volume, &r.Satuan,
		&r.HSatuan, &r.Honor, &r.Mak, &r.PMWaktu, &r.PSWaktu,
	)
	if err != nil {
		return nil, err
	}
	return &r, nil
}
