package models

import (
	"database/sql"
	"fmt"
	"strings"

	"sirekap/internal/database"
)

// Kegiatan represents a kegiatan (activity) record
type Kegiatan struct {
	ID        int     `json:"id"`
	Nama      string  `json:"nama"`
	Vol       float64 `json:"vol"`
	Mak       string  `json:"mak"`
	Satuan    string  `json:"satuan"`
	Harga     float64 `json:"harga"`
	Total     float64 `json:"total"`
	Tanggaran string  `json:"tanggaran"` // tahun anggaran
}

type KegiatanFilter struct {
	Search string
	Tahun  string
}

// KegiatanSearchResult represents a simplified result for autocomplete
type KegiatanSearchResult struct {
	Nama string `json:"nama"`
}

// KegiatanDetail represents detailed kegiatan info with sisa calculation
type KegiatanDetail struct {
	Mak    string  `json:"mak"`
	Satuan string  `json:"satuan"`
	Harga  float64 `json:"harga"`
	Sisa   float64 `json:"sisa"`
}

// SearchKegiatan searches kegiatan by name and tahun anggaran
func SearchKegiatan(query, tahun string) ([]KegiatanSearchResult, error) {
	var results []KegiatanSearchResult
	sqlQuery := `SELECT nama FROM kegiatan WHERE 1 = 1`
	args := make([]interface{}, 0, 8)

	if tahun != "" {
		sqlQuery += " AND tanggaran = ?"
		args = append(args, tahun)
	}

	searchClause, searchArgs := buildFlexibleSearchClause([]string{"nama"}, query)
	sqlQuery += searchClause + " LIMIT 20"
	args = append(args, searchArgs...)

	rows, err := database.DB.Query(sqlQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var result KegiatanSearchResult
		if err := rows.Scan(&result.Nama); err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

// GetKegiatanDetail gets kegiatan details including calculated sisa honor
func GetKegiatanDetail(nama, tahun string) (*KegiatanDetail, error) {
	detail := &KegiatanDetail{}
	nama = strings.TrimSpace(nama)
	tahun = strings.TrimSpace(tahun)

	// Get kegiatan basic details
	var sqlQuery string
	var row *sql.Row

	if tahun != "" {
		sqlQuery = `SELECT nama, mak, satuan, harga, total FROM kegiatan WHERE TRIM(nama) = TRIM(?) AND TRIM(tanggaran) = TRIM(?) LIMIT 1`
		row = database.DB.QueryRow(sqlQuery, nama, tahun)
	} else {
		sqlQuery = `SELECT nama, mak, satuan, harga, total FROM kegiatan WHERE TRIM(nama) = TRIM(?) ORDER BY tanggaran DESC LIMIT 1`
		row = database.DB.QueryRow(sqlQuery, nama)
	}

	var total float64
	var matchedNama string
	var mak, satuan sql.NullString
	var harga sql.NullFloat64

	err := row.Scan(&matchedNama, &mak, &satuan, &harga, &total)
	if err == sql.ErrNoRows {
		matchedNama, mak, satuan, harga, total, err = findKegiatanDetailFallback(nama, tahun)
	}
	if err != nil {
		return nil, err
	}

	if mak.Valid {
		detail.Mak = mak.String
	}
	if satuan.Valid {
		detail.Satuan = satuan.String
	}
	if harga.Valid {
		detail.Harga = harga.Float64
	}

	// Calculate sisa honor
	// Get sum of honor from rekap for this kegiatan
	var sisaQuery string
	var sisaRow *sql.Row

	if tahun != "" {
		sisaQuery = `
			SELECT p.total AS harga, COALESCE(SUM(sp.honor), 0) AS jumlah 
			FROM kegiatan AS p 
			LEFT JOIN rekap AS sp ON p.nama = sp.kegiatan AND p.tanggaran = sp.tanggaran 
			WHERE p.nama = ? AND TRIM(p.tanggaran) = TRIM(?) 
			GROUP BY p.nama, p.tanggaran`
		sisaRow = database.DB.QueryRow(sisaQuery, matchedNama, tahun)
	} else {
		sisaQuery = `
			SELECT p.total AS harga, COALESCE(SUM(sp.honor), 0) AS jumlah 
			FROM kegiatan AS p 
			LEFT JOIN rekap AS sp ON p.nama = sp.kegiatan 
			WHERE p.nama = ? 
			GROUP BY p.nama 
			LIMIT 1`
		sisaRow = database.DB.QueryRow(sisaQuery, matchedNama)
	}

	var hargaTotal, honorTerbayar float64
	err = sisaRow.Scan(&hargaTotal, &honorTerbayar)
	if err != nil {
		if err == sql.ErrNoRows {
			// No rekap records yet, sisa = total
			detail.Sisa = total
		} else {
			return nil, err
		}
	} else {
		detail.Sisa = hargaTotal - honorTerbayar
	}

	return detail, nil
}

func findKegiatanDetailFallback(nama, tahun string) (string, sql.NullString, sql.NullString, sql.NullFloat64, float64, error) {
	query := `SELECT nama, mak, satuan, harga, total FROM kegiatan WHERE 1 = 1`
	args := make([]interface{}, 0, 8)

	if tahun != "" {
		query += " AND TRIM(tanggaran) = TRIM(?)"
		args = append(args, tahun)
	}

	searchClause, searchArgs := buildFlexibleSearchClause([]string{"nama"}, nama)
	if searchClause == "" {
		return "", sql.NullString{}, sql.NullString{}, sql.NullFloat64{}, 0, sql.ErrNoRows
	}

	query += searchClause + " ORDER BY tanggaran DESC, id ASC LIMIT 1"
	args = append(args, searchArgs...)

	var matchedNama string
	var mak, satuan sql.NullString
	var harga sql.NullFloat64
	var total float64
	err := database.DB.QueryRow(query, args...).Scan(&matchedNama, &mak, &satuan, &harga, &total)
	return matchedNama, mak, satuan, harga, total, err
}

func ListKegiatan(filter KegiatanFilter) ([]Kegiatan, error) {
	query := `
		SELECT id, nama, vol, satuan, harga, total, mak, tanggaran
		FROM kegiatan
		WHERE 1 = 1`

	args := make([]interface{}, 0, 2)

	if filter.Tahun != "" {
		query += " AND tanggaran = ?"
		args = append(args, filter.Tahun)
	}

	if filter.Search != "" {
		searchClause, searchArgs := buildFlexibleSearchClause([]string{"nama"}, filter.Search)
		query += searchClause
		args = append(args, searchArgs...)
	}

	query += " ORDER BY tanggaran DESC, id ASC"

	rows, err := database.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Kegiatan
	for rows.Next() {
		item, err := scanKegiatan(rows)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func GetKegiatanByID(id int) (*Kegiatan, error) {
	query := `
		SELECT id, nama, vol, satuan, harga, total, mak, tanggaran
		FROM kegiatan
		WHERE id = ?`

	row := database.DB.QueryRow(query, id)
	item, err := scanKegiatan(row)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func CreateKegiatan(k *Kegiatan) (int64, error) {
	total := calculateKegiatanTotal(k.Vol, k.Harga)

	result, err := database.DB.Exec(
		`INSERT INTO kegiatan (nama, vol, satuan, harga, total, mak, tanggaran)
		 VALUES (?, ?, ?, ?, ?, ?, ?)`,
		strings.TrimSpace(k.Nama),
		k.Vol,
		strings.TrimSpace(k.Satuan),
		k.Harga,
		total,
		strings.TrimSpace(k.Mak),
		strings.TrimSpace(k.Tanggaran),
	)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func UpdateKegiatanMaster(id int, k *Kegiatan) error {
	total := calculateKegiatanTotal(k.Vol, k.Harga)

	_, err := database.DB.Exec(
		`UPDATE kegiatan
		 SET nama = ?, vol = ?, satuan = ?, harga = ?, total = ?, mak = ?, tanggaran = ?
		 WHERE id = ?`,
		strings.TrimSpace(k.Nama),
		k.Vol,
		strings.TrimSpace(k.Satuan),
		k.Harga,
		total,
		strings.TrimSpace(k.Mak),
		strings.TrimSpace(k.Tanggaran),
		id,
	)
	return err
}

func DeleteKegiatanMaster(id int) error {
	_, err := database.DB.Exec("DELETE FROM kegiatan WHERE id = ?", id)
	return err
}

func GetKegiatanTahunOptions() ([]string, error) {
	rows, err := database.DB.Query(`
		SELECT DISTINCT tanggaran
		FROM kegiatan
		WHERE tanggaran IS NOT NULL
		ORDER BY tanggaran DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var years []string
	for rows.Next() {
		var tahun sql.NullInt64
		if err := rows.Scan(&tahun); err != nil {
			return nil, err
		}
		if tahun.Valid {
			years = append(years, fmt.Sprintf("%d", tahun.Int64))
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return years, nil
}

type kegiatanScanner interface {
	Scan(dest ...interface{}) error
}

func scanKegiatan(scanner kegiatanScanner) (Kegiatan, error) {
	var item Kegiatan
	var nama, satuan, mak sql.NullString
	var vol, harga, total sql.NullFloat64
	var tanggaran sql.NullInt64

	err := scanner.Scan(&item.ID, &nama, &vol, &satuan, &harga, &total, &mak, &tanggaran)
	if err != nil {
		return Kegiatan{}, err
	}

	if nama.Valid {
		item.Nama = nama.String
	}
	if vol.Valid {
		item.Vol = vol.Float64
	}
	if satuan.Valid {
		item.Satuan = satuan.String
	}
	if harga.Valid {
		item.Harga = harga.Float64
	}
	if total.Valid {
		item.Total = total.Float64
	} else {
		item.Total = calculateKegiatanTotal(item.Vol, item.Harga)
	}
	if mak.Valid {
		item.Mak = mak.String
	}
	if tanggaran.Valid {
		item.Tanggaran = fmt.Sprintf("%d", tanggaran.Int64)
	}

	return item, nil
}

func calculateKegiatanTotal(vol, harga float64) float64 {
	return vol * harga
}
