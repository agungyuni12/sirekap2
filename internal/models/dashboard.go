package models

import "sirekap/internal/database"

type DashboardStats struct {
	TotalKegiatan      int
	TotalPagu          float64
	TotalRealisasi     float64
	RealisasiPercent   float64
	TotalMitraAktif    int
	TotalSPKTerbit     int
	TotalMatriks       int
	TotalPerjalananDin int
}

func GetDashboardStats() (*DashboardStats, error) {
	stats := &DashboardStats{}

	if err := database.DB.QueryRow(`SELECT COUNT(*) FROM kegiatan`).Scan(&stats.TotalKegiatan); err != nil {
		return nil, err
	}

	if err := database.DB.QueryRow(`SELECT COALESCE(SUM(total), 0) FROM kegiatan`).Scan(&stats.TotalPagu); err != nil {
		return nil, err
	}

	if err := database.DB.QueryRow(`SELECT COALESCE(SUM(honor), 0) FROM rekap`).Scan(&stats.TotalRealisasi); err != nil {
		return nil, err
	}

	if stats.TotalPagu > 0 {
		stats.RealisasiPercent = (stats.TotalRealisasi / stats.TotalPagu) * 100
		if stats.RealisasiPercent > 100 {
			stats.RealisasiPercent = 100
		}
	}

	if err := database.DB.QueryRow(`SELECT COUNT(DISTINCT idsobat) FROM rekap`).Scan(&stats.TotalMitraAktif); err != nil {
		return nil, err
	}

	if err := database.DB.QueryRow(`SELECT COUNT(*) FROM surat WHERE stujuan = 'SPK'`).Scan(&stats.TotalSPKTerbit); err != nil {
		return nil, err
	}

	if err := database.DB.QueryRow(`SELECT COUNT(*) FROM matriks`).Scan(&stats.TotalMatriks); err != nil {
		return nil, err
	}

	if err := database.DB.QueryRow(`SELECT COUNT(*) FROM matriks WHERE jenis_perjalanan = 'Perjalanan Dinas'`).Scan(&stats.TotalPerjalananDin); err != nil {
		return nil, err
	}

	return stats, nil
}
