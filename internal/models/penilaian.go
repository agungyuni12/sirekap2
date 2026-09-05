package models

import (
	"database/sql"
	"fmt"
	"strings"

	"sirekap/internal/database"
)

// Bobot aspek penilaian, sesuai "SIREKAP Penilaian Mitra.docx":
// Kualitas hasil kerja 30%, empat aspek lainnya masing-masing 17,5%.
const (
	BobotKualitas   = 0.30
	BobotKetepatan  = 0.175
	BobotKepatuhan  = 0.175
	BobotKomunikasi = 0.175
	BobotSikap      = 0.175
)

// weightedSkorSQL is the shared SQL expression for the weighted final score of
// one evaluasi_petugas row — keep in sync with WeightedSkor below.
const weightedSkorSQL = `(e.skor_kualitas * 0.30 + e.skor_ketepatan_waktu * 0.175 + e.skor_kepatuhan_sop * 0.175 + e.skor_komunikasi * 0.175 + e.skor_sikap * 0.175)`

// WeightedSkor computes the weighted score for one evaluasi_petugas row —
// the Go-side equivalent of weightedSkorSQL.
func WeightedSkor(kualitas, ketepatanWaktu, kepatuhanSOP, komunikasi, sikap float64) float64 {
	return kualitas*BobotKualitas + ketepatanWaktu*BobotKetepatan + kepatuhanSOP*BobotKepatuhan +
		komunikasi*BobotKomunikasi + sikap*BobotSikap
}

// PenilaianKegiatan is one entry of the kegiatan master list used specifically
// by Penilaian Mitra (seeded from Daftar_Nama_Kegiatan_Rapi.xlsx via
// migrations/rework_penilaian_mitra_workflow.sql). Deliberately kept separate
// from the `kegiatan` (budget) table used by the REKAPITULASI module — the two
// serve different purposes and are not meant to be kept in sync.
type PenilaianKegiatan struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
}

// ListPenilaianKegiatan returns the full Penilaian Mitra kegiatan list.
func ListPenilaianKegiatan() ([]PenilaianKegiatan, error) {
	rows, err := database.DB.Query("SELECT id, nama FROM penilaian_kegiatan ORDER BY urutan, nama")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []PenilaianKegiatan
	for rows.Next() {
		var k PenilaianKegiatan
		if err := rows.Scan(&k.ID, &k.Nama); err != nil {
			return nil, err
		}
		items = append(items, k)
	}
	return items, rows.Err()
}

// GetPenilaianKegiatanNama looks up a Penilaian Mitra kegiatan's name by id.
func GetPenilaianKegiatanNama(kegiatanID int) (string, error) {
	var nama string
	err := database.DB.QueryRow("SELECT nama FROM penilaian_kegiatan WHERE id = ?", kegiatanID).Scan(&nama)
	return nama, err
}

// PenilaianKegiatanPeriode is one admin-defined instance of a kegiatan (e.g.
// "Sakernas — Triwulan I, 2026") — roster and skor are scoped to this, not to
// the kegiatan as a whole, since the same kegiatan can repeat every period
// with a different roster each time.
type PenilaianKegiatanPeriode struct {
	ID         int    `json:"id"`
	KegiatanID int    `json:"kegiatan_id"`
	Kegiatan   string `json:"kegiatan,omitempty"`
	Periode    string `json:"periode"`
	Tahun      string `json:"tahun"`
}

// ListPeriodeForKegiatan returns every periode instance admin has created for
// one kegiatan, newest first.
func ListPeriodeForKegiatan(kegiatanID int) ([]PenilaianKegiatanPeriode, error) {
	rows, err := database.DB.Query(
		`SELECT id, kegiatan_id, periode, tahun FROM penilaian_kegiatan_periode
		 WHERE kegiatan_id = ? ORDER BY tahun DESC, id DESC`,
		kegiatanID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []PenilaianKegiatanPeriode
	for rows.Next() {
		var it PenilaianKegiatanPeriode
		if err := rows.Scan(&it.ID, &it.KegiatanID, &it.Periode, &it.Tahun); err != nil {
			return nil, err
		}
		items = append(items, it)
	}
	return items, rows.Err()
}

// AddPeriodeKegiatan creates (or returns the id of an already-existing) periode
// instance for one kegiatan. Admin-only, enforced at the handler.
func AddPeriodeKegiatan(kegiatanID int, periode, tahun string, addedBy int) (int, error) {
	res, err := database.DB.Exec(`
		INSERT INTO penilaian_kegiatan_periode (kegiatan_id, periode, tahun, added_by)
		VALUES (?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE id = LAST_INSERT_ID(id)
	`, kegiatanID, periode, tahun, addedBy)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	return int(id), err
}

// GetPeriodeInfo resolves one periode instance's kegiatan_id, kegiatan nama,
// periode, and tahun — used server-side to fill in evaluasi_petugas instead
// of trusting periode/tahun from the request body.
func GetPeriodeInfo(periodeID int) (kegiatanID int, kegiatanNama, periode, tahun string, err error) {
	err = database.DB.QueryRow(`
		SELECT p.kegiatan_id, k.nama, p.periode, p.tahun
		FROM penilaian_kegiatan_periode p
		JOIN penilaian_kegiatan k ON k.id = p.kegiatan_id
		WHERE p.id = ?
	`, periodeID).Scan(&kegiatanID, &kegiatanNama, &periode, &tahun)
	return kegiatanID, kegiatanNama, periode, tahun, err
}

// EvaluasiInput is the payload for saving one penilaian (Tahap 1 initial score,
// or the Subject Matter's re-score on rejection).
type EvaluasiInput struct {
	KegiatanID         int
	PeriodeID          int
	YangDinilaiIDSobat string
	Peran              string // "ppl" | "pml" — who is being scored
	Kecamatan          string // wilayah tugas for this specific assignment
	Periode            string
	Tahun              string
	TanggalPenilaian   string // "YYYY-MM-DD"
	SkorKualitas       float64
	SkorKetepatanWaktu float64
	SkorKepatuhanSOP   float64
	SkorKomunikasi     float64
	SkorSikap          float64
	Catatan            string
}

// ValidateSkor returns an error message if any score is outside 1-100.
func ValidateSkor(kualitas, ketepatan, kepatuhanSOP, komunikasi, sikap float64) string {
	for _, s := range []float64{kualitas, ketepatan, kepatuhanSOP, komunikasi, sikap} {
		if s < 1 || s > 100 {
			return fmt.Sprintf("Skor harus berada pada rentang 1-100 (nilai diterima: %.2f)", s)
		}
	}
	return ""
}

// UpsertEvaluasiTahap1 saves the initial score for one yang-dinilai (PPL by
// PML organik/mitra, or PML Mitra by atasan langsung organik/Korwil). Every
// Tahap 1 score — regardless of peran — goes to status_konfirmasi "pending",
// awaiting the Subject Matter's approve/reject decision; that's the "Dari
// nilai tersebut, subject matter mengkonfirmasi nilai" step in the PRD, which
// applies to both PPL and PML Mitra scores, not just PML Mitra.
//
// Menyimpan ulang (upsert) skor Tahap 1 mereset status konfirmasi ke "pending"
// dan membuang skor ulang Subject Matter sebelumnya (jika ada), karena
// keduanya merujuk ke nilai awal yang sudah tidak berlaku lagi.
func UpsertEvaluasiTahap1(penilaiID int, in EvaluasiInput) error {
	tx, err := database.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
		INSERT INTO evaluasi_petugas
			(kegiatan_id, periode_id, penilai_id, yang_dinilai_idsobat, peran_yang_dinilai, kecamatan, periode, tahun, tanggal_penilaian, tahap,
			 skor_kualitas, skor_ketepatan_waktu, skor_kepatuhan_sop, skor_komunikasi, skor_sikap, catatan, status_konfirmasi)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, 1, ?, ?, ?, ?, ?, ?, 'pending')
		ON DUPLICATE KEY UPDATE
			penilai_id = VALUES(penilai_id),
			kecamatan = VALUES(kecamatan),
			tanggal_penilaian = VALUES(tanggal_penilaian),
			skor_kualitas = VALUES(skor_kualitas),
			skor_ketepatan_waktu = VALUES(skor_ketepatan_waktu),
			skor_kepatuhan_sop = VALUES(skor_kepatuhan_sop),
			skor_komunikasi = VALUES(skor_komunikasi),
			skor_sikap = VALUES(skor_sikap),
			catatan = VALUES(catatan),
			status_konfirmasi = 'pending',
			confirmed_by = NULL,
			confirmed_at = NULL,
			catatan_konfirmasi = NULL
	`, in.KegiatanID, in.PeriodeID, penilaiID, in.YangDinilaiIDSobat, in.Peran, in.Kecamatan, in.Periode, in.Tahun, nullableDate(in.TanggalPenilaian),
		in.SkorKualitas, in.SkorKetepatanWaktu, in.SkorKepatuhanSOP, in.SkorKomunikasi, in.SkorSikap, in.Catatan)
	if err != nil {
		return err
	}

	if _, err := tx.Exec(
		`DELETE FROM evaluasi_petugas WHERE periode_id = ? AND yang_dinilai_idsobat = ? AND tahap = 2`,
		in.PeriodeID, in.YangDinilaiIDSobat,
	); err != nil {
		return err
	}

	return tx.Commit()
}

func nullableDate(s string) interface{} {
	if s == "" {
		return nil
	}
	return s
}

// GetTahap1Status returns the peran and status_konfirmasi of one Tahap 1 score
// ("pending" | "disetujui" | "ditolak") — peran is needed by
// SubmitPenilaianUlang, since the Subject Matter's re-score must carry the
// same peran as what it's re-scoring (PPL or PML Mitra).
func GetTahap1Status(periodeID int, idsobat string) (peran, status string, err error) {
	err = database.DB.QueryRow(
		`SELECT peran_yang_dinilai, status_konfirmasi FROM evaluasi_petugas
		 WHERE periode_id = ? AND yang_dinilai_idsobat = ? AND tahap = 1`,
		periodeID, idsobat,
	).Scan(&peran, &status)
	return peran, status, err
}

// KonfirmasiPenilaian records the Subject Matter's approve/reject decision on
// one Tahap 1 score (PPL or PML Mitra).
func KonfirmasiPenilaian(periodeID int, idsobat string, setuju bool, subjectMatterID int, catatan string) error {
	status := "ditolak"
	if setuju {
		status = "disetujui"
	}
	res, err := database.DB.Exec(`
		UPDATE evaluasi_petugas
		SET status_konfirmasi = ?, confirmed_by = ?, confirmed_at = CURRENT_TIMESTAMP, catatan_konfirmasi = ?
		WHERE periode_id = ? AND yang_dinilai_idsobat = ? AND tahap = 1
	`, status, subjectMatterID, catatan, periodeID, idsobat)
	if err != nil {
		return err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return sql.ErrNoRows
	}
	return nil
}

// SubmitPenilaianUlang saves the Subject Matter's own score after rejecting a
// Tahap 1 score (PPL or PML Mitra). Final skor akhir becomes the average of
// the two — computed on read in GetDaftarPenilaian, not stored.
func SubmitPenilaianUlang(subjectMatterID int, in EvaluasiInput) error {
	_, err := database.DB.Exec(`
		INSERT INTO evaluasi_petugas
			(kegiatan_id, periode_id, penilai_id, yang_dinilai_idsobat, peran_yang_dinilai, kecamatan, periode, tahun, tanggal_penilaian, tahap,
			 skor_kualitas, skor_ketepatan_waktu, skor_kepatuhan_sop, skor_komunikasi, skor_sikap, catatan, status_konfirmasi)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, 2, ?, ?, ?, ?, ?, ?, 'disetujui')
		ON DUPLICATE KEY UPDATE
			penilai_id = VALUES(penilai_id),
			kecamatan = VALUES(kecamatan),
			tanggal_penilaian = VALUES(tanggal_penilaian),
			skor_kualitas = VALUES(skor_kualitas),
			skor_ketepatan_waktu = VALUES(skor_ketepatan_waktu),
			skor_kepatuhan_sop = VALUES(skor_kepatuhan_sop),
			skor_komunikasi = VALUES(skor_komunikasi),
			skor_sikap = VALUES(skor_sikap),
			catatan = VALUES(catatan)
	`, in.KegiatanID, in.PeriodeID, subjectMatterID, in.YangDinilaiIDSobat, in.Peran, in.Kecamatan, in.Periode, in.Tahun, nullableDate(in.TanggalPenilaian),
		in.SkorKualitas, in.SkorKetepatanWaktu, in.SkorKepatuhanSOP, in.SkorKomunikasi, in.SkorSikap, in.Catatan)
	return err
}

// PendingKonfirmasiItem is one Tahap 1 score (PPL or PML Mitra) awaiting the
// Subject Matter's approve/reject decision.
type PendingKonfirmasiItem struct {
	KegiatanID       int     `json:"kegiatan_id"`
	PeriodeID        int     `json:"periode_id"`
	Kegiatan         string  `json:"kegiatan"`
	Periode          string  `json:"periode"`
	Tahun            string  `json:"tahun"`
	IDSobat          string  `json:"idsobat"`
	NamaMitra        string  `json:"nama_mitra"`
	Peran            string  `json:"peran"`
	Kecamatan        string  `json:"kecamatan"`
	TanggalPenilaian string  `json:"tanggal_penilaian"`
	NamaPenilai      string  `json:"nama_penilai"`
	SkorAwal         float64 `json:"skor_awal"`
	Catatan          string  `json:"catatan"`
}

// ListPendingKonfirmasi returns every Tahap 1 score (PPL and PML Mitra) still
// waiting on the Subject Matter.
func ListPendingKonfirmasi() ([]PendingKonfirmasiItem, error) {
	rows, err := database.DB.Query(`
		SELECT e.kegiatan_id, e.periode_id, COALESCE(pk.nama, ''), COALESCE(e.periode, ''), COALESCE(e.tahun, ''),
			e.yang_dinilai_idsobat, COALESCE(m.nmitra, ''), e.peran_yang_dinilai,
			COALESCE(e.kecamatan, ''), COALESCE(e.tanggal_penilaian, ''), COALESCE(u.nama, ''),
			` + weightedSkorSQL + `, COALESCE(e.catatan, '')
		FROM evaluasi_petugas e
		LEFT JOIN penilaian_kegiatan pk ON pk.id = e.kegiatan_id
		LEFT JOIN mitra m ON m.idsobat = e.yang_dinilai_idsobat
		LEFT JOIN user u ON u.id = e.penilai_id
		WHERE e.tahap = 1 AND e.status_konfirmasi = 'pending' AND e.periode_id IS NOT NULL
		ORDER BY e.tanggal_penilaian DESC, m.nmitra
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []PendingKonfirmasiItem
	for rows.Next() {
		var it PendingKonfirmasiItem
		var tanggal sql.NullString
		if err := rows.Scan(&it.KegiatanID, &it.PeriodeID, &it.Kegiatan, &it.Periode, &it.Tahun, &it.IDSobat, &it.NamaMitra, &it.Peran, &it.Kecamatan, &tanggal, &it.NamaPenilai, &it.SkorAwal, &it.Catatan); err != nil {
			return nil, err
		}
		it.Peran = strings.ToUpper(it.Peran)
		it.TanggalPenilaian = tanggal.String
		items = append(items, it)
	}
	return items, rows.Err()
}

// DaftarPenilaianItem is one merged assessment row — Tahap 1 (+ Tahap 2 when
// the confirmation was rejected) — used by both "Daftar Penilaian" and the
// Dashboard & Rekap screens.
type DaftarPenilaianItem struct {
	KegiatanID       int      `json:"kegiatan_id"`
	PeriodeID        int      `json:"periode_id"`
	Kegiatan         string   `json:"kegiatan"`
	IDSobat          string   `json:"idsobat"`
	NamaMitra        string   `json:"nama_mitra"`
	Peran            string   `json:"peran"`
	Kecamatan        string   `json:"kecamatan"`
	Periode          string   `json:"periode"`
	Tahun            string   `json:"tahun"`
	TanggalPenilaian string   `json:"tanggal_penilaian"`
	NamaPenilaiAwal  string   `json:"nama_penilai_awal"`
	SkorAwal         *float64 `json:"skor_awal"`
	StatusKonfirmasi string   `json:"status_konfirmasi"` // "" untuk PPL (tidak berlaku)
	NamaPenilaiUlang string   `json:"nama_penilai_ulang,omitempty"`
	SkorUlang        *float64 `json:"skor_ulang,omitempty"`
	SkorAkhir        *float64 `json:"skor_akhir"`
	Predikat         string   `json:"predikat"`
	Catatan          string   `json:"catatan"`
}

// DaftarPenilaianFilter holds the shared filters for GetDaftarPenilaian.
type DaftarPenilaianFilter struct {
	KegiatanID       int
	PeriodeID        int
	Tahun            string
	Peran            string
	Kecamatan        string
	Predikat         string
	StatusKonfirmasi string
	Search           string
}

// GetDaftarPenilaian merges each (kegiatan, yang-dinilai) pair's Tahap 1 (+
// Tahap 2, when rejected) rows into one row with a computed skor_akhir, then
// applies filters — peran/predikat/kecamatan/search are derived or free-text,
// so they're evaluated in Go after rows are assembled, same pattern as the
// previous GetRekapDashboard.
func GetDaftarPenilaian(f DaftarPenilaianFilter) ([]DaftarPenilaianItem, error) {
	// periode_id is nullable only to preserve pre-migration rows (see
	// migrations/add_periode_ke_penilaian_kegiatan.sql) — those rows have no
	// periode to report against, so they're excluded rather than crashing the
	// scan below (periode_id is scanned as a plain int, not sql.NullInt64).
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
		SELECT e.kegiatan_id, e.periode_id, COALESCE(pk.nama, ''), e.yang_dinilai_idsobat, COALESCE(m.nmitra, ''),
			e.peran_yang_dinilai, COALESCE(e.kecamatan, ''), COALESCE(e.periode, ''), COALESCE(e.tahun, ''),
			COALESCE(e.tanggal_penilaian, ''), e.tahap, `+weightedSkorSQL+`,
			e.status_konfirmasi, COALESCE(u.nama, ''), COALESCE(e.catatan, '')
		FROM evaluasi_petugas e
		LEFT JOIN penilaian_kegiatan pk ON pk.id = e.kegiatan_id
		LEFT JOIN mitra m ON m.idsobat = e.yang_dinilai_idsobat
		LEFT JOIN user u ON u.id = e.penilai_id
		WHERE `+where+`
		ORDER BY e.periode_id, e.yang_dinilai_idsobat, e.tahap
	`, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	type key struct {
		periodeID int
		idsobat   string
	}
	var order []key
	byKey := map[key]*DaftarPenilaianItem{}

	for rows.Next() {
		var kID, periodeID int
		var kegiatanNama, idsobat, namaMitra, peran, kecamatan, periode, tahun, tanggal, statusKonfirmasi, namaPenilai, catatan string
		var tahap int
		var skor float64
		if err := rows.Scan(&kID, &periodeID, &kegiatanNama, &idsobat, &namaMitra, &peran, &kecamatan, &periode, &tahun,
			&tanggal, &tahap, &skor, &statusKonfirmasi, &namaPenilai, &catatan); err != nil {
			return nil, err
		}

		k := key{periodeID, idsobat}
		item, ok := byKey[k]
		if !ok {
			item = &DaftarPenilaianItem{
				KegiatanID: kID,
				PeriodeID:  periodeID,
				Kegiatan:   kegiatanNama,
				IDSobat:    idsobat,
				NamaMitra:  namaMitra,
				Peran:      strings.ToUpper(peran),
				Kecamatan:  kecamatan,
				Periode:    periode,
				Tahun:      tahun,
			}
			byKey[k] = item
			order = append(order, k)
		}

		skorCopy := skor
		if tahap == 1 {
			item.TanggalPenilaian = tanggal
			item.NamaPenilaiAwal = namaPenilai
			item.SkorAwal = &skorCopy
			item.StatusKonfirmasi = statusKonfirmasi
			if catatan != "" {
				item.Catatan = catatan
			}
		} else if tahap == 2 {
			item.NamaPenilaiUlang = namaPenilai
			item.SkorUlang = &skorCopy
			if catatan != "" {
				item.Catatan = catatan // tahap 2 (Subject Matter) note wins when both have one
			}
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	peranFilter := strings.ToLower(f.Peran)
	kecamatanFilter := strings.ToLower(f.Kecamatan)
	items := make([]DaftarPenilaianItem, 0, len(order))
	for _, k := range order {
		item := *byKey[k]

		switch strings.ToLower(item.Peran) {
		case "ppl":
			// Tidak ada tahap konfirmasi untuk PPL — final seketika PML submit.
			item.SkorAkhir = item.SkorAwal
		case "pml":
			switch item.StatusKonfirmasi {
			case "disetujui":
				item.SkorAkhir = item.SkorAwal
			case "ditolak":
				if item.SkorAwal != nil && item.SkorUlang != nil {
					avg := (*item.SkorAwal + *item.SkorUlang) / 2
					item.SkorAkhir = &avg
				}
			}
		}
		if item.SkorAkhir != nil {
			item.Predikat = PredikatFromSkor(*item.SkorAkhir)
		}

		if peranFilter != "" && !strings.EqualFold(item.Peran, peranFilter) {
			continue
		}
		if kecamatanFilter != "" && !strings.Contains(strings.ToLower(item.Kecamatan), kecamatanFilter) {
			continue
		}
		if f.Predikat != "" && !strings.EqualFold(item.Predikat, f.Predikat) {
			continue
		}
		if f.StatusKonfirmasi != "" && !strings.EqualFold(item.StatusKonfirmasi, f.StatusKonfirmasi) {
			continue
		}
		if f.Search != "" {
			q := strings.ToLower(f.Search)
			if !strings.Contains(strings.ToLower(item.NamaMitra), q) && !strings.Contains(strings.ToLower(item.IDSobat), q) {
				continue
			}
		}
		items = append(items, item)
	}

	return items, nil
}

// GetDetailPenilaian returns the full aspect-by-aspect breakdown for one
// (kegiatan, yang-dinilai) assessment — "riwayat penilaiannya bisa dicek lagi
// per aspek nya" from the PRD.
type AspekSkor struct {
	Label string  `json:"label"`
	Bobot float64 `json:"bobot"`
	Skor  float64 `json:"skor"`
}

type DetailPenilaian struct {
	DaftarPenilaianItem
	AspekAwal  []AspekSkor `json:"aspek_awal"`
	AspekUlang []AspekSkor `json:"aspek_ulang,omitempty"`
}

func aspekBreakdown(kualitas, ketepatan, kepatuhanSOP, komunikasi, sikap float64) []AspekSkor {
	return []AspekSkor{
		{Label: "Kualitas Hasil Kerja", Bobot: BobotKualitas, Skor: kualitas},
		{Label: "Ketepatan Waktu", Bobot: BobotKetepatan, Skor: ketepatan},
		{Label: "Kepatuhan terhadap SOP", Bobot: BobotKepatuhan, Skor: kepatuhanSOP},
		{Label: "Komunikasi & Koordinasi", Bobot: BobotKomunikasi, Skor: komunikasi},
		{Label: "Sikap/Perilaku", Bobot: BobotSikap, Skor: sikap},
	}
}

// GetDetailPenilaian fetches one assessment's full per-aspect breakdown.
func GetDetailPenilaian(periodeID int, idsobat string) (*DetailPenilaian, error) {
	rows, err := database.DB.Query(`
		SELECT e.kegiatan_id, e.periode_id, COALESCE(pk.nama, ''), e.yang_dinilai_idsobat, COALESCE(m.nmitra, ''),
			e.peran_yang_dinilai, COALESCE(e.kecamatan, ''), COALESCE(e.periode, ''), COALESCE(e.tahun, ''),
			COALESCE(e.tanggal_penilaian, ''), e.tahap,
			e.skor_kualitas, e.skor_ketepatan_waktu, e.skor_kepatuhan_sop, e.skor_komunikasi, e.skor_sikap,
			e.status_konfirmasi, COALESCE(u.nama, ''), COALESCE(e.catatan, '')
		FROM evaluasi_petugas e
		LEFT JOIN penilaian_kegiatan pk ON pk.id = e.kegiatan_id
		LEFT JOIN mitra m ON m.idsobat = e.yang_dinilai_idsobat
		LEFT JOIN user u ON u.id = e.penilai_id
		WHERE e.periode_id = ? AND e.yang_dinilai_idsobat = ?
		ORDER BY e.tahap
	`, periodeID, idsobat)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var d *DetailPenilaian
	found := false
	for rows.Next() {
		found = true
		var kID, pID, tahap int
		var kegiatanNama, idsb, namaMitra, peran, kecamatan, periode, tahun, tanggal, statusKonfirmasi, namaPenilai, catatan string
		var kualitas, ketepatan, kepatuhanSOP, komunikasi, sikap float64
		if err := rows.Scan(&kID, &pID, &kegiatanNama, &idsb, &namaMitra, &peran, &kecamatan, &periode, &tahun, &tanggal, &tahap,
			&kualitas, &ketepatan, &kepatuhanSOP, &komunikasi, &sikap, &statusKonfirmasi, &namaPenilai, &catatan); err != nil {
			return nil, err
		}

		if d == nil {
			d = &DetailPenilaian{DaftarPenilaianItem: DaftarPenilaianItem{
				KegiatanID: kID, PeriodeID: pID, Kegiatan: kegiatanNama, IDSobat: idsb, NamaMitra: namaMitra,
				Peran: strings.ToUpper(peran), Kecamatan: kecamatan, Periode: periode, Tahun: tahun,
			}}
		}

		skor := WeightedSkor(kualitas, ketepatan, kepatuhanSOP, komunikasi, sikap)
		skorCopy := skor
		if tahap == 1 {
			d.TanggalPenilaian = tanggal
			d.NamaPenilaiAwal = namaPenilai
			d.SkorAwal = &skorCopy
			d.AspekAwal = aspekBreakdown(kualitas, ketepatan, kepatuhanSOP, komunikasi, sikap)
			if peran == "pml" {
				d.StatusKonfirmasi = statusKonfirmasi
			}
			if catatan != "" {
				d.Catatan = catatan
			}
		} else if tahap == 2 {
			d.NamaPenilaiUlang = namaPenilai
			d.SkorUlang = &skorCopy
			d.AspekUlang = aspekBreakdown(kualitas, ketepatan, kepatuhanSOP, komunikasi, sikap)
			if catatan != "" {
				d.Catatan = catatan
			}
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	if !found {
		return nil, sql.ErrNoRows
	}

	switch strings.ToLower(d.Peran) {
	case "ppl":
		d.SkorAkhir = d.SkorAwal
	case "pml":
		switch d.StatusKonfirmasi {
		case "disetujui":
			d.SkorAkhir = d.SkorAwal
		case "ditolak":
			if d.SkorAwal != nil && d.SkorUlang != nil {
				avg := (*d.SkorAwal + *d.SkorUlang) / 2
				d.SkorAkhir = &avg
			}
		}
	}
	if d.SkorAkhir != nil {
		d.Predikat = PredikatFromSkor(*d.SkorAkhir)
	}

	return d, nil
}

// RosterPetugasItem is one entry on a periode's Penilaian Mitra roster.
// Two kinds share the same roster: mitra entries (peran "ppl"/"pml" — who is
// being scored, identified by IDSobat) and organik entries (peran
// "pml"/"korwil" — which organik account is authorized to score PPL or PML
// Mitra for this periode, identified by UserID). Only admin (Subject
// Matter) may add/remove roster entries — Korwil and PML (organik/mitra)
// can only pick from what's already there when scoring, per "PML sm korwil
// gk boleh tambah petugas".
type RosterPetugasItem struct {
	Tipe      string `json:"tipe"` // "mitra" | "organik"
	IDSobat   string `json:"idsobat,omitempty"`
	UserID    int    `json:"user_id,omitempty"`
	Nama      string `json:"nama"`
	Kecamatan string `json:"kecamatan,omitempty"`
	Email     string `json:"email,omitempty"`
	Peran     string `json:"peran"`
}

// ListPetugasKegiatan returns one periode's Penilaian Mitra roster —
// mitra entries and organik entries combined — optionally filtered to one
// peran ("ppl" | "pml" | "korwil").
func ListPetugasKegiatan(periodeID int, peran string) ([]RosterPetugasItem, error) {
	query := `
		SELECT 'mitra' AS tipe, p.idsobat, 0 AS user_id, COALESCE(m.nmitra, '') AS nama, COALESCE(m.kecamatan, '') AS kecamatan, '' AS email, p.peran
		FROM penilaian_kegiatan_petugas p
		LEFT JOIN mitra m ON m.idsobat = p.idsobat
		WHERE p.periode_id = ? AND p.user_id IS NULL`
	args := []interface{}{periodeID}
	if peran != "" {
		query += " AND p.peran = ?"
		args = append(args, peran)
	}
	query += `
		UNION ALL
		SELECT 'organik', '', p.user_id, COALESCE(u.nama, ''), '', COALESCE(u.email, ''), p.peran
		FROM penilaian_kegiatan_petugas p
		LEFT JOIN user u ON u.id = p.user_id
		WHERE p.periode_id = ? AND p.user_id IS NOT NULL`
	args = append(args, periodeID)
	if peran != "" {
		query += " AND p.peran = ?"
		args = append(args, peran)
	}
	query += " ORDER BY nama"

	rows, err := database.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []RosterPetugasItem
	for rows.Next() {
		var it RosterPetugasItem
		if err := rows.Scan(&it.Tipe, &it.IDSobat, &it.UserID, &it.Nama, &it.Kecamatan, &it.Email, &it.Peran); err != nil {
			return nil, err
		}
		items = append(items, it)
	}
	return items, rows.Err()
}

// AddPetugasKeKegiatan adds (or updates the peran of) one entry on a
// periode's Penilaian Mitra roster — pass idsobat for a mitra entry
// ("ppl"/"pml") or userID for an organik entry ("pml"/"korwil"), never both.
// Admin-only, enforced at the handler.
func AddPetugasKeKegiatan(periodeID int, idsobat string, userID int, peran string, addedBy int) error {
	var idsobatArg, userIDArg interface{}
	if idsobat != "" {
		idsobatArg = idsobat
	} else {
		userIDArg = userID
	}
	_, err := database.DB.Exec(`
		INSERT INTO penilaian_kegiatan_petugas (periode_id, idsobat, user_id, peran, added_by)
		VALUES (?, ?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE peran = VALUES(peran)
	`, periodeID, idsobatArg, userIDArg, peran, addedBy)
	return err
}

// RemovePetugasDariKegiatan removes one entry from a periode's roster —
// pass idsobat for a mitra entry or userID for an organik entry.
func RemovePetugasDariKegiatan(periodeID int, idsobat string, userID int) error {
	var res sql.Result
	var err error
	if idsobat != "" {
		res, err = database.DB.Exec(
			`DELETE FROM penilaian_kegiatan_petugas WHERE periode_id = ? AND idsobat = ?`,
			periodeID, idsobat,
		)
	} else {
		res, err = database.DB.Exec(
			`DELETE FROM penilaian_kegiatan_petugas WHERE periode_id = ? AND user_id = ?`,
			periodeID, userID,
		)
	}
	if err != nil {
		return err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return sql.ErrNoRows
	}
	return nil
}

// OrganikSearchResult is one akun organik matched by SearchPenilaiOrganik.
type OrganikSearchResult struct {
	UserID int    `json:"user_id"`
	Nama   string `json:"nama"`
	Email  string `json:"email"`
}

// SearchPenilaiOrganik searches akun organik (level "pengguna") by nama,
// email, or NIP — used to populate the PML/Korwil tabs on Kelola Petugas.
func SearchPenilaiOrganik(query string) ([]OrganikSearchResult, error) {
	var results []OrganikSearchResult
	sqlQuery := `SELECT id, nama, email FROM user WHERE level = 'pengguna'`
	searchClause, args := buildFlexibleSearchClause([]string{"nama", "email", "nip"}, query)
	sqlQuery += searchClause + " ORDER BY nama LIMIT 20"

	rows, err := database.DB.Query(sqlQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var r OrganikSearchResult
		if err := rows.Scan(&r.UserID, &r.Nama, &r.Email); err != nil {
			return nil, err
		}
		results = append(results, r)
	}
	return results, rows.Err()
}

// GetPenilaiPeranForUser returns the peran ("pml" | "korwil") an organik
// account has been assigned for one periode — sql.ErrNoRows if the admin
// hasn't assigned them yet, which callers must treat as "access denied".
func GetPenilaiPeranForUser(periodeID, userID int) (string, error) {
	var peran string
	err := database.DB.QueryRow(
		`SELECT peran FROM penilaian_kegiatan_petugas WHERE periode_id = ? AND user_id = ?`,
		periodeID, userID,
	).Scan(&peran)
	return peran, err
}

// PeriodeSayaItem is one periode instance an organik account has been
// assigned to, along with the peran ("pml" | "korwil") they hold for it.
type PeriodeSayaItem struct {
	PeriodeID  int    `json:"periode_id"`
	KegiatanID int    `json:"kegiatan_id"`
	Kegiatan   string `json:"kegiatan"`
	Periode    string `json:"periode"`
	Tahun      string `json:"tahun"`
	Peran      string `json:"peran"`
}

// ListPeriodeSayaForOrganik returns every periode instance an organik account
// is assigned to (any peran), used to build the kegiatan/periode dropdowns on
// the Input Penilaian form — periode without an assignment never appear.
func ListPeriodeSayaForOrganik(userID int) ([]PeriodeSayaItem, error) {
	rows, err := database.DB.Query(`
		SELECT per.id, pk.id, pk.nama, per.periode, per.tahun, p.peran
		FROM penilaian_kegiatan_petugas p
		JOIN penilaian_kegiatan_periode per ON per.id = p.periode_id
		JOIN penilaian_kegiatan pk ON pk.id = per.kegiatan_id
		WHERE p.user_id = ?
		ORDER BY pk.urutan, pk.nama, per.tahun DESC, per.id DESC
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []PeriodeSayaItem
	for rows.Next() {
		var it PeriodeSayaItem
		if err := rows.Scan(&it.PeriodeID, &it.KegiatanID, &it.Kegiatan, &it.Periode, &it.Tahun, &it.Peran); err != nil {
			return nil, err
		}
		items = append(items, it)
	}
	return items, rows.Err()
}
