package models

import (
	"database/sql"
	"fmt"
	"strings"

	"sirekap/internal/database"
)

// PredikatFromSkor classifies a final score into the 4-tier predikat used by
// the Dashboard & Rekap Penilaian Kinerja Mitra screen. This is a finer-grained
// classification than the 2-tier lulus/perlu_perhatian used by GetRekapPenilaian
// (Tahap 2 scoring screen) — the two are independent and intentionally don't share
// a threshold constant, since the PRD defines them as separate features.
func PredikatFromSkor(skor float64) string {
	switch {
	case skor >= 90:
		return "Sangat Baik"
	case skor >= 80:
		return "Baik"
	case skor >= 70:
		return "Cukup"
	default:
		return "Perlu Pembinaan"
	}
}

// DistribusiKelulusan is the predikat breakdown, keyed the way the API contract
// specifies (snake_case, one field per predikat).
type DistribusiKelulusan struct {
	SangatBaik     int `json:"sangat_baik"`
	Baik           int `json:"baik"`
	Cukup          int `json:"cukup"`
	PerluPembinaan int `json:"perlu_pembinaan"`
}

// RataPerAspek is the average score per one of the 5 penilaian aspects.
type RataPerAspek struct {
	Kualitas       float64 `json:"kualitas"`
	KetepatanWaktu float64 `json:"ketepatan_waktu"`
	KepatuhanSOP   float64 `json:"kepatuhan_sop"`
	Komunikasi     float64 `json:"komunikasi"`
	Sikap          float64 `json:"sikap"`
}

// DashboardPenilaianStats is the aggregate payload for GET /api/penilaian/dashboard.
type DashboardPenilaianStats struct {
	TotalDinilai        int                 `json:"total_dinilai"`
	RataRataSkor        float64             `json:"rata_rata_skor"`
	DistribusiKelulusan DistribusiKelulusan `json:"distribusi_kelulusan"`
	RataPerAspek        RataPerAspek        `json:"rata_per_aspek"`
}

// dashboardFilterClause builds the shared WHERE clause + args for filtering
// evaluasi_petugas by kegiatan_id and kegiatan.tanggaran (tahun anggaran).
// evaluasi_petugas has no FK to kegiatan (see create_evaluasi_petugas_table.sql),
// so the join is a plain equality on kegiatan.id.
func dashboardFilterClause(kegiatanID int, tahun string) (string, []interface{}) {
	clauses := []string{"1=1"}
	var args []interface{}
	if kegiatanID > 0 {
		clauses = append(clauses, "e.kegiatan_id = ?")
		args = append(args, kegiatanID)
	}
	if tahun != "" {
		clauses = append(clauses, "k.tanggaran = ?")
		args = append(args, tahun)
	}
	return strings.Join(clauses, " AND "), args
}

// GetDashboardPenilaianStats aggregates evaluasi_petugas into the summary
// statistics shown on the Dashboard Penilaian Kinerja Mitra screen.
func GetDashboardPenilaianStats(kegiatanID int, tahun string) (*DashboardPenilaianStats, error) {
	where, args := dashboardFilterClause(kegiatanID, tahun)

	// Per (kegiatan, mitra) skor akhir = average across whichever tahap(s) exist,
	// so a mitra scored on both tahap for a kegiatan collapses to one row —
	// the same rule used for the rekap-dashboard table.
	rows, err := database.DB.Query(fmt.Sprintf(`
		SELECT AVG((e.skor_kualitas + e.skor_ketepatan_waktu + e.skor_kepatuhan_sop + e.skor_komunikasi + e.skor_sikap) / 5) AS skor_akhir
		FROM evaluasi_petugas e
		JOIN kegiatan k ON k.id = e.kegiatan_id
		WHERE %s
		GROUP BY e.kegiatan_id, e.yang_dinilai_idsobat
	`, where), args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	stats := &DashboardPenilaianStats{}
	var sumSkor float64
	for rows.Next() {
		var skor float64
		if err := rows.Scan(&skor); err != nil {
			return nil, err
		}
		stats.TotalDinilai++
		sumSkor += skor
		switch PredikatFromSkor(skor) {
		case "Sangat Baik":
			stats.DistribusiKelulusan.SangatBaik++
		case "Baik":
			stats.DistribusiKelulusan.Baik++
		case "Cukup":
			stats.DistribusiKelulusan.Cukup++
		default:
			stats.DistribusiKelulusan.PerluPembinaan++
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	if stats.TotalDinilai > 0 {
		stats.RataRataSkor = sumSkor / float64(stats.TotalDinilai)
	}

	aspekWhere, aspekArgs := dashboardFilterClause(kegiatanID, tahun)
	var kualitas, ketepatan, kepatuhan, komunikasi, sikap sql.NullFloat64
	err = database.DB.QueryRow(fmt.Sprintf(`
		SELECT AVG(e.skor_kualitas), AVG(e.skor_ketepatan_waktu), AVG(e.skor_kepatuhan_sop), AVG(e.skor_komunikasi), AVG(e.skor_sikap)
		FROM evaluasi_petugas e
		JOIN kegiatan k ON k.id = e.kegiatan_id
		WHERE %s
	`, aspekWhere), aspekArgs...).Scan(&kualitas, &ketepatan, &kepatuhan, &komunikasi, &sikap)
	if err != nil {
		return nil, err
	}
	stats.RataPerAspek = RataPerAspek{
		Kualitas:       kualitas.Float64,
		KetepatanWaktu: ketepatan.Float64,
		KepatuhanSOP:   kepatuhan.Float64,
		Komunikasi:     komunikasi.Float64,
		Sikap:          sikap.Float64,
	}

	return stats, nil
}

// RekapDashboardItem is one merged Tahap 1 + Tahap 2 row for the rekap-dashboard table.
type RekapDashboardItem struct {
	IDSobat    string   `json:"idsobat"`
	NamaMitra  string   `json:"nama_mitra"`
	Kegiatan   string   `json:"kegiatan"`
	Peran      string   `json:"peran"`
	SkorTahap1 *float64 `json:"skor_tahap1"`
	SkorTahap2 *float64 `json:"skor_tahap2"`
	SkorAkhir  float64  `json:"skor_akhir"`
	Predikat   string   `json:"predikat"`
	Catatan    string   `json:"catatan"`
}

// GetRekapDashboard merges each mitra's Tahap 1 and Tahap 2 scores (per kegiatan)
// into one row with a computed skor_akhir and predikat, then applies the
// peran/predikat/search filters — those three are derived values, so they're
// evaluated in Go after the per-mitra rows are assembled, not pushed into SQL.
func GetRekapDashboard(kegiatanID int, tahun, peran, predikat, search string) ([]RekapDashboardItem, error) {
	where, args := dashboardFilterClause(kegiatanID, tahun)

	rows, err := database.DB.Query(fmt.Sprintf(`
		SELECT e.kegiatan_id, k.nama, e.yang_dinilai_idsobat, COALESCE(r.namamitra, ''),
			e.tahap, (e.skor_kualitas + e.skor_ketepatan_waktu + e.skor_kepatuhan_sop + e.skor_komunikasi + e.skor_sikap) / 5,
			e.catatan
		FROM evaluasi_petugas e
		JOIN kegiatan k ON k.id = e.kegiatan_id
		LEFT JOIN (SELECT DISTINCT kegiatan, idsobat, namamitra FROM rekap) r
			ON r.kegiatan COLLATE utf8mb4_general_ci = k.nama
			AND r.idsobat COLLATE utf8mb4_general_ci = e.yang_dinilai_idsobat
		WHERE %s
		ORDER BY e.kegiatan_id, e.yang_dinilai_idsobat, e.tahap
	`, where), args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	type key struct {
		kegiatanID int
		idsobat    string
	}
	order := []key{}
	byKey := map[key]*RekapDashboardItem{}
	sumSkor := map[key]float64{}
	countSkor := map[key]int{}

	for rows.Next() {
		var kID int
		var kegiatanNama, idsobat, namaMitra, catatan string
		var tahap int
		var skor float64
		if err := rows.Scan(&kID, &kegiatanNama, &idsobat, &namaMitra, &tahap, &skor, &catatan); err != nil {
			return nil, err
		}

		k := key{kID, idsobat}
		item, ok := byKey[k]
		if !ok {
			item = &RekapDashboardItem{
				IDSobat:   idsobat,
				NamaMitra: namaMitra,
				Kegiatan:  kegiatanNama,
				Peran:     strings.ToUpper(classifyPeranFromKegiatan(kegiatanNama)),
			}
			byKey[k] = item
			order = append(order, k)
		}

		skorCopy := skor
		if tahap == 1 {
			item.SkorTahap1 = &skorCopy
		} else if tahap == 2 {
			item.SkorTahap2 = &skorCopy
		}
		if catatan != "" {
			item.Catatan = catatan // tahap 2 (Subject Matter) sorts after tahap 1, so it wins when both have notes
		}
		sumSkor[k] += skor
		countSkor[k]++
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	peranFilter := strings.ToLower(peran)
	items := make([]RekapDashboardItem, 0, len(order))
	for _, k := range order {
		item := *byKey[k]
		item.SkorAkhir = sumSkor[k] / float64(countSkor[k])
		item.Predikat = PredikatFromSkor(item.SkorAkhir)

		if peranFilter != "" && !strings.EqualFold(item.Peran, peranFilter) {
			continue
		}
		if predikat != "" && !strings.EqualFold(item.Predikat, predikat) {
			continue
		}
		if search != "" {
			q := strings.ToLower(search)
			if !strings.Contains(strings.ToLower(item.NamaMitra), q) && !strings.Contains(strings.ToLower(item.IDSobat), q) {
				continue
			}
		}
		items = append(items, item)
	}

	return items, nil
}
