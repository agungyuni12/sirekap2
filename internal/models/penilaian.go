package models

import (
	"database/sql"
	"fmt"
	"strings"

	"sirekap/internal/database"
)

// PetugasPenilaianItem represents one petugas (PPL/PML) that can be scored
// for a given kegiatan + tahap, along with whether that tahap is already filled in.
type PetugasPenilaianItem struct {
	MitraID         string `json:"mitra_id"`
	Nama            string `json:"nama"`
	Peran           string `json:"peran"`
	StatusPenilaian string `json:"status_penilaian"` // "belum" | "selesai"
}

// RekapPenilaianItem represents the combined Tahap 1 + Tahap 2 score for one petugas.
type RekapPenilaianItem struct {
	MitraID         string   `json:"mitra_id"`
	MitraName       string   `json:"mitra_name"`
	Peran           string   `json:"peran"`
	NilaiTahap1     *float64 `json:"nilai_tahap1"`
	NilaiTahap2     *float64 `json:"nilai_tahap2"`
	NilaiAkhir      *float64 `json:"nilai_akhir"`
	StatusKelulusan string   `json:"status_kelulusan"` // "belum" | "lulus" | "perlu_perhatian"
}

// EvaluasiInput is the payload for saving one Tahap 1 or Tahap 2 score.
type EvaluasiInput struct {
	KegiatanID         int
	YangDinilaiIDSobat string
	Tahap              int
	SkorKualitas       float64
	SkorKetepatanWaktu float64
	SkorEtika          float64
	Catatan            string
}

// classifyPeranFromKegiatan infers PPL vs PML from the kegiatan name, following
// the same convention already used for SE2026 import (see se2026_import.go),
// since this app doesn't track PPL/PML as a separate column — a kegiatan's
// name itself designates which position it covers (e.g. "... (PML)").
func classifyPeranFromKegiatan(kegiatanNama string) string {
	upper := strings.ToUpper(kegiatanNama)
	if strings.Contains(upper, "PML") || strings.Contains(upper, "PEMERIKSA") {
		return "pml"
	}
	return "ppl"
}

// GetKegiatanNama looks up a kegiatan's name by id.
func GetKegiatanNama(kegiatanID int) (string, error) {
	var nama string
	err := database.DB.QueryRow("SELECT nama FROM kegiatan WHERE id = ?", kegiatanID).Scan(&nama)
	return nama, err
}

// ListKegiatanIDName returns a lightweight {id, nama} list, safe to expose to
// non-admin roles (unlike /api/master-kegiatan, it carries no budget figures).
func ListKegiatanIDName() ([]Kegiatan, error) {
	rows, err := database.DB.Query("SELECT id, nama FROM kegiatan ORDER BY id DESC LIMIT 500")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Kegiatan
	for rows.Next() {
		var k Kegiatan
		if err := rows.Scan(&k.ID, &k.Nama); err != nil {
			return nil, err
		}
		items = append(items, k)
	}
	return items, rows.Err()
}

// GetPetugasForPenilaian lists the distinct petugas assigned to a kegiatan
// (sourced from `rekap`, the app's actual assignment record) with their
// Tahap-`tahap` scoring status.
func GetPetugasForPenilaian(kegiatanID, tahap int) ([]PetugasPenilaianItem, error) {
	kegiatanNama, err := GetKegiatanNama(kegiatanID)
	if err != nil {
		return nil, err
	}
	peran := classifyPeranFromKegiatan(kegiatanNama)

	rows, err := database.DB.Query(
		`SELECT DISTINCT idsobat, namamitra FROM rekap WHERE kegiatan = ? AND idsobat != '' ORDER BY namamitra`,
		kegiatanNama,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var mitraList []struct{ IDSobat, Nama string }
	for rows.Next() {
		var idsobat, nama string
		if err := rows.Scan(&idsobat, &nama); err != nil {
			return nil, err
		}
		mitraList = append(mitraList, struct{ IDSobat, Nama string }{idsobat, nama})
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	items := make([]PetugasPenilaianItem, 0, len(mitraList))
	for _, m := range mitraList {
		var count int
		err := database.DB.QueryRow(
			`SELECT COUNT(*) FROM evaluasi_petugas WHERE kegiatan_id = ? AND yang_dinilai_idsobat = ? AND tahap = ?`,
			kegiatanID, m.IDSobat, tahap,
		).Scan(&count)
		if err != nil {
			return nil, err
		}
		status := "belum"
		if count > 0 {
			status = "selesai"
		}
		items = append(items, PetugasPenilaianItem{MitraID: m.IDSobat, Nama: m.Nama, Peran: peran, StatusPenilaian: status})
	}
	return items, nil
}

// PeranForPetugas returns the ppl/pml classification for a given kegiatan.
func PeranForPetugas(kegiatanID int) (string, error) {
	nama, err := GetKegiatanNama(kegiatanID)
	if err != nil {
		return "", err
	}
	return classifyPeranFromKegiatan(nama), nil
}

// HasTahap1 reports whether a Tahap 1 score already exists for this petugas.
func HasTahap1(kegiatanID int, idsobat string) (bool, error) {
	var count int
	err := database.DB.QueryRow(
		`SELECT COUNT(*) FROM evaluasi_petugas WHERE kegiatan_id = ? AND yang_dinilai_idsobat = ? AND tahap = 1`,
		kegiatanID, idsobat,
	).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// UpsertEvaluasi saves (or overwrites) one petugas' score for one tahap.
func UpsertEvaluasi(penilaiID int, peran string, in EvaluasiInput) error {
	_, err := database.DB.Exec(`
		INSERT INTO evaluasi_petugas
			(kegiatan_id, penilai_id, yang_dinilai_idsobat, peran_yang_dinilai, tahap, skor_kualitas, skor_ketepatan_waktu, skor_etika, catatan)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE
			penilai_id = VALUES(penilai_id),
			skor_kualitas = VALUES(skor_kualitas),
			skor_ketepatan_waktu = VALUES(skor_ketepatan_waktu),
			skor_etika = VALUES(skor_etika),
			catatan = VALUES(catatan)
	`, in.KegiatanID, penilaiID, in.YangDinilaiIDSobat, peran, in.Tahap,
		in.SkorKualitas, in.SkorKetepatanWaktu, in.SkorEtika, in.Catatan)
	return err
}

// GetRekapPenilaian returns the combined Tahap 1 / Tahap 2 view for every
// petugas on a kegiatan, used by the Subject Matter finalization screen.
func GetRekapPenilaian(kegiatanID int) ([]RekapPenilaianItem, error) {
	kegiatanNama, err := GetKegiatanNama(kegiatanID)
	if err != nil {
		return nil, err
	}
	peran := classifyPeranFromKegiatan(kegiatanNama)

	query := `
		SELECT r.idsobat, r.namamitra,
			e1.avg_skor AS nilai_tahap1,
			e2.avg_skor AS nilai_tahap2
		FROM (SELECT DISTINCT idsobat, namamitra FROM rekap WHERE kegiatan = ? AND idsobat != '') r
		LEFT JOIN (
			SELECT yang_dinilai_idsobat, (skor_kualitas + skor_ketepatan_waktu + skor_etika) / 3 AS avg_skor
			FROM evaluasi_petugas WHERE kegiatan_id = ? AND tahap = 1
		) e1 ON e1.yang_dinilai_idsobat = r.idsobat
		LEFT JOIN (
			SELECT yang_dinilai_idsobat, (skor_kualitas + skor_ketepatan_waktu + skor_etika) / 3 AS avg_skor
			FROM evaluasi_petugas WHERE kegiatan_id = ? AND tahap = 2
		) e2 ON e2.yang_dinilai_idsobat = r.idsobat
		ORDER BY r.namamitra
	`
	rows, err := database.DB.Query(query, kegiatanNama, kegiatanID, kegiatanID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []RekapPenilaianItem
	for rows.Next() {
		var idsobat, nama string
		var nilai1, nilai2 sql.NullFloat64
		if err := rows.Scan(&idsobat, &nama, &nilai1, &nilai2); err != nil {
			return nil, err
		}

		item := RekapPenilaianItem{MitraID: idsobat, MitraName: nama, Peran: peran, StatusKelulusan: "belum"}
		if nilai1.Valid {
			v := nilai1.Float64
			item.NilaiTahap1 = &v
		}
		if nilai2.Valid {
			v := nilai2.Float64
			item.NilaiTahap2 = &v
		}

		// Tahap 2 is the authoritative final review for PPL; PML Mitra is
		// scored directly in Tahap 2 with no Tahap 1 step.
		if item.NilaiTahap2 != nil {
			item.NilaiAkhir = item.NilaiTahap2
		} else if peran == "pml" {
			item.NilaiAkhir = nil
		} else if item.NilaiTahap1 != nil {
			item.NilaiAkhir = item.NilaiTahap1
		}

		if item.NilaiAkhir != nil {
			if *item.NilaiAkhir >= 70 {
				item.StatusKelulusan = "lulus"
			} else {
				item.StatusKelulusan = "perlu_perhatian"
			}
		}

		items = append(items, item)
	}
	return items, rows.Err()
}

// ValidateSkor returns an error message if any score is outside 1-100.
func ValidateSkor(kualitas, ketepatan, etika float64) string {
	for _, s := range []float64{kualitas, ketepatan, etika} {
		if s < 1 || s > 100 {
			return fmt.Sprintf("Skor harus berada pada rentang 1-100 (nilai diterima: %.2f)", s)
		}
	}
	return ""
}
