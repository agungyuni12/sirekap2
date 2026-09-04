package models

import (
	"database/sql"

	"sirekap/internal/database"
)

// Mitra represents a mitra (partner) record
type Mitra struct {
	IDSobat string `json:"idsobat"`
	Nama    string `json:"nama"` // nmitra column in database
	Tahun   string `json:"tahun"`
}

// MitraSearchResult represents a simplified result for autocomplete
type MitraSearchResult struct {
	IDSobat   string `json:"idsobat"`
	Nama      string `json:"nama"`
	Kecamatan string `json:"kecamatan,omitempty"`
}

// SearchMitra searches mitra by idsobat or name
func SearchMitra(query, tahun string) ([]MitraSearchResult, error) {
	var results []MitraSearchResult
	sqlQuery := `
		SELECT idsobat, nmitra, kecamatan
		FROM mitra
		WHERE 1 = 1`
	args := make([]interface{}, 0, 12)

	if tahun != "" {
		sqlQuery += " AND tahun = ?"
		args = append(args, tahun)
	}

	searchClause, searchArgs := buildFlexibleSearchClause(
		[]string{"idsobat", "nmitra", "CONCAT(idsobat,'/', nmitra)"},
		query,
	)
	sqlQuery += searchClause + " LIMIT 20"
	args = append(args, searchArgs...)

	rows, err := database.DB.Query(sqlQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var result MitraSearchResult
		var idsobat, nmitra, kecamatan sql.NullString

		if err := rows.Scan(&idsobat, &nmitra, &kecamatan); err != nil {
			return nil, err
		}

		if idsobat.Valid {
			result.IDSobat = idsobat.String
		}
		if nmitra.Valid {
			result.Nama = nmitra.String
		}
		if kecamatan.Valid {
			result.Kecamatan = kecamatan.String
		}

		results = append(results, result)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

// MitraInfo contains mitra information
type MitraInfo struct {
	IDSobat   string
	NamaMitra string
	NIK       string
	Alamat    string
	Kecamatan string
}

// GetMitraByIDSobat retrieves mitra info by IDSobat
func GetMitraByIDSobat(idsobat string) (*MitraInfo, error) {
	var info MitraInfo
	var nik, alamat, kecamatan sql.NullString
	err := database.DB.QueryRow(
		"SELECT idsobat, nmitra, nik, alamat, kecamatan FROM mitra WHERE idsobat = ? LIMIT 1",
		idsobat,
	).Scan(&info.IDSobat, &info.NamaMitra, &nik, &alamat, &kecamatan)
	if err != nil {
		return nil, err
	}
	info.NIK = nik.String
	info.Alamat = alamat.String
	info.Kecamatan = kecamatan.String
	return &info, nil
}
