package models

import (
	"strings"

	"sirekap/internal/database"
)

// PredikatFromSkor classifies a final score into the 5-tier predikat used by
// the Dashboard & Rekap Penilaian Kinerja Mitra screen: 91-100 Sangat Baik,
// 81-90 Baik, 71-80 Cukup, 61-70 Kurang, 1-60 Sangat Kurang.
func PredikatFromSkor(skor float64) string {
	switch {
	case skor >= 91:
		return "Sangat Baik"
	case skor >= 81:
		return "Baik"
	case skor >= 71:
		return "Cukup"
	case skor >= 61:
		return "Kurang"
	default:
		return "Sangat Kurang"
	}
}

// DistribusiKelulusan is the predikat breakdown, keyed the way the API contract
// specifies (snake_case, one field per predikat).
type DistribusiKelulusan struct {
	SangatBaik   int `json:"sangat_baik"`
	Baik         int `json:"baik"`
	Cukup        int `json:"cukup"`
	Kurang       int `json:"kurang"`
	SangatKurang int `json:"sangat_kurang"`
}

// RataPerAspek is the average score per one of the 5 penilaian aspects
// (unweighted — the raw 1-100 score submitted for that aspect).
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

type aspekRaw struct {
	kualitas, ketepatan, kepatuhanSOP, komunikasi, sikap float64
}

// GetDashboardPenilaianStats aggregates evaluasi_petugas into the summary
// statistics shown on the Dashboard Penilaian Kinerja Mitra screen. Only
// finalized assessments count — PPL and PML Mitra scores alike go through
// the same Subject Matter confirmation: "disetujui" finalizes the initial
// score as-is, "ditolak" finalizes as the average with Subject Matter's own
// re-score once submitted. Anything still "pending" (or rejected without a
// re-score yet) is excluded until finalized.
func GetDashboardPenilaianStats(f DaftarPenilaianFilter) (*DashboardPenilaianStats, error) {
	// periode_id is nullable only to preserve pre-migration rows (see
	// migrations/add_periode_ke_penilaian_kegiatan.sql) — those rows have no
	// periode to report against, so they're excluded here.
	clauses := []string{"e.periode_id IS NOT NULL"}
	var args []interface{}
	if f.KegiatanID > 0 {
		clauses = append(clauses, "e.kegiatan_id = ?")
		args = append(args, f.KegiatanID)
	}
	if f.PeriodeID > 0 {
		clauses = append(clauses, "e.periode_id = ?")
		args = append(args, f.PeriodeID)
	}
	if f.Tahun != "" {
		clauses = append(clauses, "e.tahun = ?")
		args = append(args, f.Tahun)
	}
	where := strings.Join(clauses, " AND ")

	rows, err := database.DB.Query(`
		SELECT e.periode_id, e.yang_dinilai_idsobat, e.peran_yang_dinilai, COALESCE(e.kecamatan, ''),
			e.tahap, e.status_konfirmasi,
			e.skor_kualitas, e.skor_ketepatan_waktu, e.skor_kepatuhan_sop, e.skor_komunikasi, e.skor_sikap
		FROM evaluasi_petugas e
		WHERE `+where, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	type key struct {
		periodeID int
		idsobat   string
	}
	type group struct {
		peran, kecamatan, status string
		tahap1, tahap2           *aspekRaw
	}
	groups := map[key]*group{}

	for rows.Next() {
		var pID, tahap int
		var idsobat, peran, kecamatan, status string
		var a aspekRaw
		if err := rows.Scan(&pID, &idsobat, &peran, &kecamatan, &tahap, &status,
			&a.kualitas, &a.ketepatan, &a.kepatuhanSOP, &a.komunikasi, &a.sikap); err != nil {
			return nil, err
		}
		k := key{pID, idsobat}
		g, ok := groups[k]
		if !ok {
			g = &group{peran: peran, kecamatan: kecamatan, status: status}
			groups[k] = g
		}
		aCopy := a
		if tahap == 1 {
			g.tahap1 = &aCopy
		} else if tahap == 2 {
			g.tahap2 = &aCopy
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	peranFilter := strings.ToLower(f.Peran)
	kecamatanFilter := strings.ToLower(f.Kecamatan)

	stats := &DashboardPenilaianStats{}
	var sumSkor, sumKualitas, sumKetepatan, sumKepatuhan, sumKomunikasi, sumSikap float64

	for _, g := range groups {
		if peranFilter != "" && !strings.EqualFold(g.peran, peranFilter) {
			continue
		}
		if kecamatanFilter != "" && !strings.Contains(strings.ToLower(g.kecamatan), kecamatanFilter) {
			continue
		}

		// PPL dan PML Mitra sama-sama lewat alur konfirmasi Subject Matter —
		// lihat catatan di GetDaftarPenilaian. Pending (atau ditolak tanpa
		// skor ulang) belum dihitung "final", jadi dilewati dari statistik.
		var final *aspekRaw
		switch g.status {
		case "disetujui":
			final = g.tahap1
		case "ditolak":
			if g.tahap1 != nil && g.tahap2 != nil {
				final = &aspekRaw{
					kualitas:     (g.tahap1.kualitas + g.tahap2.kualitas) / 2,
					ketepatan:    (g.tahap1.ketepatan + g.tahap2.ketepatan) / 2,
					kepatuhanSOP: (g.tahap1.kepatuhanSOP + g.tahap2.kepatuhanSOP) / 2,
					komunikasi:   (g.tahap1.komunikasi + g.tahap2.komunikasi) / 2,
					sikap:        (g.tahap1.sikap + g.tahap2.sikap) / 2,
				}
			}
		}
		if final == nil {
			continue
		}

		skor := WeightedSkor(final.kualitas, final.ketepatan, final.kepatuhanSOP, final.komunikasi, final.sikap)
		stats.TotalDinilai++
		sumSkor += skor
		sumKualitas += final.kualitas
		sumKetepatan += final.ketepatan
		sumKepatuhan += final.kepatuhanSOP
		sumKomunikasi += final.komunikasi
		sumSikap += final.sikap

		switch PredikatFromSkor(skor) {
		case "Sangat Baik":
			stats.DistribusiKelulusan.SangatBaik++
		case "Baik":
			stats.DistribusiKelulusan.Baik++
		case "Cukup":
			stats.DistribusiKelulusan.Cukup++
		case "Kurang":
			stats.DistribusiKelulusan.Kurang++
		default:
			stats.DistribusiKelulusan.SangatKurang++
		}
	}

	if stats.TotalDinilai > 0 {
		n := float64(stats.TotalDinilai)
		stats.RataRataSkor = sumSkor / n
		stats.RataPerAspek = RataPerAspek{
			Kualitas:       sumKualitas / n,
			KetepatanWaktu: sumKetepatan / n,
			KepatuhanSOP:   sumKepatuhan / n,
			Komunikasi:     sumKomunikasi / n,
			Sikap:          sumSikap / n,
		}
	}

	return stats, nil
}

// GetRekapDashboard is a thin alias over GetDaftarPenilaian — the Dashboard &
// Rekap table and the Daftar Penilaian screen show the same merged rows, just
// with a different filter set from their respective query params.
func GetRekapDashboard(f DaftarPenilaianFilter) ([]DaftarPenilaianItem, error) {
	return GetDaftarPenilaian(f)
}
