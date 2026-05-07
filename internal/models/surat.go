package models

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"sirekap/internal/database"
)

// Surat represents a letter record for SPK/BAST tracking
type Surat struct {
	ID         int    `json:"id"`
	IDSobat    string `json:"idsobat"`
	NamaMitra  string `json:"nmitra"`
	Bulan      string `json:"sbulan"`
	Tahun      string `json:"stahun"`
	Tujuan     string `json:"stujuan"` // SPK or BAST
	KSurat     int    `json:"ksurat"`  // Letter sequence number
	NSurat     string `json:"nsurat"`  // Full letter number
	TglSurat   string `json:"tglsurat"`
	NmrSurat   string `json:"nmrsurat"`
	Keterangan string `json:"keterangan"` // For BAST: kegiatan name
}

type SuratFilter struct {
	Search string
	Tahun  string
	Tujuan string
}

// GetExistingSurat checks if a letter already exists for this mitra/period/type
func GetExistingSurat(idsobat, bulan, tahun, tujuan, keterangan string) (*Surat, error) {
	return getExistingSurat(database.DB, idsobat, bulan, tahun, tujuan, keterangan)
}

// CreateSurat creates a new letter record and returns the letter number
func CreateSurat(idsobat, namaMitra, bulan, tahun, tujuan, tglSurat, keterangan string) (string, error) {
	idsobat = strings.TrimSpace(idsobat)
	namaMitra = strings.TrimSpace(namaMitra)
	bulan = strings.TrimSpace(bulan)
	tahun = strings.TrimSpace(tahun)
	tujuan = strings.TrimSpace(strings.ToUpper(tujuan))
	tglSurat = strings.TrimSpace(tglSurat)
	keterangan = strings.TrimSpace(keterangan)

	tx, err := database.DB.Begin()
	if err != nil {
		return "", err
	}
	defer tx.Rollback()

	existing, err := getExistingSurat(tx, idsobat, bulan, tahun, tujuan, keterangan)
	if err != nil {
		return "", err
	}

	nextK := 0
	if existing != nil && existing.TglSurat == tglSurat {
		nextK = existing.KSurat
	} else {
		excludeID := 0
		if existing != nil {
			excludeID = existing.ID
		}
		nextK, err = getNextAvailableSuratSequenceTx(tx, tglSurat, tujuan, excludeID)
		if err != nil {
			return "", err
		}
	}

	nmrSurat, nsurat, err := buildSuratNumbers(tglSurat, nextK, tujuan, tahun)
	if err != nil {
		return "", err
	}

	if existing == nil {
		_, err = tx.Exec(
			`INSERT INTO surat (idsobat, nmitra, sbulan, stahun, stujuan, ksurat, nsurat, tglsurat, nmrsurat, keterangan) 
			 VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			idsobat, namaMitra, bulan, tahun, tujuan, nextK, nsurat, tglSurat, nmrSurat, keterangan,
		)
	} else {
		_, err = tx.Exec(
			`UPDATE surat
			 SET idsobat = ?, nmitra = ?, sbulan = ?, stahun = ?, stujuan = ?, ksurat = ?, nsurat = ?, tglsurat = ?, nmrsurat = ?, keterangan = ?
			 WHERE id = ?`,
			idsobat, namaMitra, bulan, tahun, tujuan, nextK, nsurat, tglSurat, nmrSurat, keterangan, existing.ID,
		)
	}
	if err != nil {
		return "", err
	}

	if err := tx.Commit(); err != nil {
		return "", err
	}

	return nsurat, nil
}

// GetSPKSurat gets SPK letter for a mitra/period
func GetSPKSurat(idsobat, bulan, tahun string) (*Surat, error) {
	return GetExistingSurat(idsobat, bulan, tahun, "SPK", "")
}

// GetBASTSurat gets BAST letter for a specific kegiatan
func GetBASTSurat(idsobat, bulan, tahun, kegiatan string) (*Surat, error) {
	return GetExistingSurat(idsobat, bulan, tahun, "BAST", kegiatan)
}

func ListSurat(filter SuratFilter) ([]Surat, error) {
	query := `
		SELECT id, idsobat, nmitra, sbulan, stahun, stujuan, ksurat, nsurat, tglsurat, nmrsurat, COALESCE(keterangan, '')
		FROM surat
		WHERE 1 = 1`

	args := make([]interface{}, 0, 3)

	if filter.Tahun != "" {
		query += " AND stahun = ?"
		args = append(args, filter.Tahun)
	}

	if filter.Tujuan != "" {
		query += " AND stujuan = ?"
		args = append(args, filter.Tujuan)
	}

	if filter.Search != "" {
		searchClause, searchArgs := buildFlexibleSearchClause(
			[]string{"idsobat", "nmitra", "nsurat", "nmrsurat", "COALESCE(keterangan, '')"},
			filter.Search,
		)
		query += searchClause
		args = append(args, searchArgs...)
	}

	query += " ORDER BY tglsurat DESC, id DESC"

	rows, err := database.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Surat
	for rows.Next() {
		item, err := scanSurat(rows)
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

func GetSuratTahunOptions() ([]string, error) {
	rows, err := database.DB.Query(`
		SELECT DISTINCT stahun
		FROM surat
		WHERE stahun IS NOT NULL AND stahun <> ''
		ORDER BY stahun DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []string
	for rows.Next() {
		var tahun sql.NullString
		if err := rows.Scan(&tahun); err != nil {
			return nil, err
		}
		if tahun.Valid {
			items = append(items, tahun.String)
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func CreateSuratManual(idsobat, namaMitra, bulan, tahun, tujuan, tglSurat, keterangan string) (string, error) {
	return CreateSurat(
		strings.TrimSpace(idsobat),
		strings.TrimSpace(namaMitra),
		strings.TrimSpace(bulan),
		strings.TrimSpace(tahun),
		strings.TrimSpace(strings.ToUpper(tujuan)),
		strings.TrimSpace(tglSurat),
		strings.TrimSpace(keterangan),
	)
}

type suratScanner interface {
	Scan(dest ...interface{}) error
}

type suratQueryer interface {
	QueryRow(query string, args ...interface{}) *sql.Row
}

func scanSurat(scanner suratScanner) (Surat, error) {
	var item Surat
	var tgl sql.NullTime

	err := scanner.Scan(
		&item.ID,
		&item.IDSobat,
		&item.NamaMitra,
		&item.Bulan,
		&item.Tahun,
		&item.Tujuan,
		&item.KSurat,
		&item.NSurat,
		&tgl,
		&item.NmrSurat,
		&item.Keterangan,
	)
	if err != nil {
		return Surat{}, err
	}

	if tgl.Valid {
		item.TglSurat = tgl.Time.Format("2006-01-02")
	}

	return item, nil
}

func getExistingSurat(q suratQueryer, idsobat, bulan, tahun, tujuan, keterangan string) (*Surat, error) {
	var query string
	var args []interface{}

	if tujuan == "BAST" && keterangan != "" {
		query = `SELECT id, idsobat, nmitra, sbulan, stahun, stujuan, ksurat, nsurat, tglsurat, nmrsurat, keterangan
				 FROM surat WHERE idsobat = ? AND sbulan = ? AND stahun = ? AND stujuan = ? AND keterangan = ?`
		args = []interface{}{idsobat, bulan, tahun, tujuan, keterangan}
	} else {
		query = `SELECT id, idsobat, nmitra, sbulan, stahun, stujuan, ksurat, nsurat, tglsurat, nmrsurat, COALESCE(keterangan, '')
				 FROM surat WHERE idsobat = ? AND sbulan = ? AND stahun = ? AND stujuan = ?`
		args = []interface{}{idsobat, bulan, tahun, tujuan}
	}

	item, err := scanSurat(q.QueryRow(query, args...))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func getNextAvailableSuratSequenceTx(tx *sql.Tx, tglSurat, tujuan string, excludeID int) (int, error) {
	rows, err := tx.Query(
		`SELECT ksurat
		 FROM surat
		 WHERE tglsurat = ? AND stujuan = ? AND id <> ?
		 ORDER BY ksurat ASC`,
		tglSurat, tujuan, excludeID,
	)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	next := 1
	for rows.Next() {
		var current sql.NullInt64
		if err := rows.Scan(&current); err != nil {
			return 0, err
		}
		if !current.Valid || current.Int64 < int64(next) {
			continue
		}
		if current.Int64 == int64(next) {
			next++
		}
	}

	if err := rows.Err(); err != nil {
		return 0, err
	}

	return next, nil
}

func buildSuratNumbers(tglSurat string, sequence int, tujuan, tahun string) (string, string, error) {
	t, err := time.Parse("2006-01-02", tglSurat)
	if err != nil {
		return "", "", err
	}

	nmrSurat := fmt.Sprintf("%02d.%02d.%d", t.Day(), int(t.Month()), sequence)
	nsurat := fmt.Sprintf("B-%s/5205.PPK/%s.P/%s", nmrSurat, tujuan, tahun)
	return nmrSurat, nsurat, nil
}
