package models

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"sirekap/internal/database"
)

type ValidationError struct {
	Message string
}

func (e ValidationError) Error() string {
	return e.Message
}

type PegawaiSearchResult struct {
	NIP      string `json:"nip"`
	Nama     string `json:"nama"`
	Golongan string `json:"golongan"`
	Jabatan  string `json:"jabatan"`
}

type PegawaiDetail struct {
	NIP      string
	Nama     string
	Golongan string
	Jabatan  string
}

type DaftarTranslokSearchResult struct {
	Nama string `json:"nama"`
	Mak  string `json:"mak"`
}

type MatriksRecord struct {
	ID              int    `json:"id"`
	NIP             string `json:"nip"`
	Nama            string `json:"nama"`
	Golongan        string `json:"golongan"`
	JenisPerjalanan string `json:"jenisPerjalanan"`
	Kegiatan        string `json:"kegiatan"`
	Asal            string `json:"asal"`
	Tujuan          string `json:"tujuan"`
	TanggalMulai    string `json:"tanggalMulai"`
	TanggalSelesai  string `json:"tanggalSelesai"`
	NoSurtug        string `json:"noSurtug"`
	NoSpd           string `json:"noSpd"`
	Jumlah          int    `json:"jumlah"`
	Mak             string `json:"mak"`
	NoFP            string `json:"noFp"`
	PerihalFP       string `json:"perihalFp"`
	TglSurtugUtama  string `json:"tglSurtugUtama"`
	TglSpdUtama     string `json:"tglSpdUtama"`
}

type MatriksFilter struct {
	Search string
	Jenis  string
	Tahun  string
}

type MatriksCalendarFilter struct {
	Search string
	Jenis  string
	Year   int
	Month  int
}

type MatriksAdministrasiSummary struct {
	NoFP          string `json:"noFp"`
	Mak           string `json:"mak"`
	PerihalFP     string `json:"perihalFp"`
	JumlahPeserta int    `json:"jumlahPeserta"`
	JumlahLaporan int    `json:"jumlahLaporan"`
	SampleID      int    `json:"sampleId"`
}

type MatriksAdministrasiFilter struct {
	Search string
}

type MatriksAdministrasiPeserta struct {
	ID              int    `json:"id"`
	Nama            string `json:"nama"`
	NIP             string `json:"nip"`
	JenisPerjalanan string `json:"jenisPerjalanan"`
	Tujuan          string `json:"tujuan"`
	TanggalMulai    string `json:"tanggalMulai"`
	TanggalSelesai  string `json:"tanggalSelesai"`
	NoSurtug        string `json:"noSurtug"`
	NoSpd           string `json:"noSpd"`
	ReportType      string `json:"reportType"`
}

type MatriksPrintBase struct {
	ID              int
	Nama            string
	NamaGelar       string
	NIP             string
	Golongan        string
	Jabatan         string
	JenisPerjalanan string
	Kegiatan        string
	Asal            string
	Tujuan          string
	Mak             string
	PerihalFP       string
	NoSurtug        string
	NoSpd           string
	TanggalMulai    time.Time
	TanggalSelesai  time.Time
	TglSurtugUtama  time.Time
	TglSpdUtama     time.Time
}

type MatriksTransportPrintData struct {
	Base        MatriksPrintBase
	LaporanText string
	LaporanFoto string
}

type MatriksSPPDPrintData struct {
	Base                    MatriksPrintBase
	WaktuBerangkatMulai     string
	WaktuBerangkatSelesai   string
	IsianBerangkat          string
	WaktuPengawasan1Mulai   string
	WaktuPengawasan1Selesai string
	IsianPengawasan1        string
	WaktuIsomaMulai         string
	WaktuIsomaSelesai       string
	WaktuPengawasan2Mulai   string
	WaktuPengawasan2Selesai string
	IsianPengawasan2        string
	WaktuPulangMulai        string
	WaktuPulangSelesai      string
	IsianPulang             string
	LaporanFoto             string
}

type MatriksEntryPayload struct {
	NIP             string `json:"nip"`
	Nama            string `json:"nama"`
	Golongan        string `json:"golongan"`
	JenisPerjalanan string `json:"jenisPerjalanan"`
	Kegiatan        string `json:"kegiatan"`
	Asal            string `json:"asal"`
	Tujuan          string `json:"tujuan"`
	TanggalMulai    string `json:"tanggalMulai"`
	TanggalSelesai  string `json:"tanggalSelesai"`
	NoSurtug        string `json:"noSurtug"`
	NoSpd           string `json:"noSpd"`
	Jumlah          int    `json:"jumlah"`
	Mak             string `json:"mak"`
	NoFP            string `json:"noFp"`
	PerihalFP       string `json:"perihalFp"`
	TglSurtugUtama  string `json:"tglSurtugUtama"`
	TglSpdUtama     string `json:"tglSpdUtama"`
}

type MatriksParticipantPayload struct {
	NIP            string `json:"nip"`
	Nama           string `json:"nama"`
	Golongan       string `json:"golongan"`
	Asal           string `json:"asal"`
	Tujuan         string `json:"tujuan"`
	TanggalMulai   string `json:"tanggalMulai"`
	TanggalSelesai string `json:"tanggalSelesai"`
	NoSurtug       string `json:"noSurtug"`
	NoSpd          string `json:"noSpd"`
	Jumlah         int    `json:"jumlah"`
}

type MatriksCreatePayload struct {
	JenisPerjalanan string                      `json:"jenisPerjalanan"`
	Kegiatan        string                      `json:"kegiatan"`
	Mak             string                      `json:"mak"`
	NoFP            string                      `json:"noFp"`
	PerihalFP       string                      `json:"perihalFp"`
	TglSurtugUtama  string                      `json:"tglSurtugUtama"`
	TglSpdUtama     string                      `json:"tglSpdUtama"`
	Peserta         []MatriksParticipantPayload `json:"peserta"`
}

type MatriksConflictResult struct {
	Bentrok bool   `json:"bentrok"`
	Pesan   string `json:"pesan"`
}

type MatriksImportRowError struct {
	Row     int    `json:"row"`
	NIP     string `json:"nip,omitempty"`
	Nama    string `json:"nama,omitempty"`
	Message string `json:"message"`
}

type matriksQueryer interface {
	QueryRow(query string, args ...interface{}) *sql.Row
}

type matriksPreparedEntry struct {
	Payload        *MatriksEntryPayload
	TanggalMulai   time.Time
	TanggalSelesai time.Time
	TglSurtugUtama time.Time
	TglSpdUtama    time.Time
}

func SearchPegawai(query string) ([]PegawaiSearchResult, error) {
	sqlQuery := `
		SELECT
			COALESCE(NULLIF(nip, ''), '') AS nip,
			COALESCE(NULLIF(namagelar, ''), NULLIF(nama, ''), '') AS nama,
			COALESCE(NULLIF(golongan, ''), '') AS golongan,
			COALESCE(NULLIF(jabatan, ''), '') AS jabatan
		FROM pegawai
		WHERE 1 = 1`

	searchClause, searchArgs := buildFlexibleSearchClause(
		[]string{
			"nip",
			"nama",
			"namagelar",
			"CONCAT(COALESCE(nama, ''), ' ', COALESCE(namagelar, ''))",
		},
		query,
	)
	sqlQuery += searchClause + `
		ORDER BY COALESCE(NULLIF(namagelar, ''), NULLIF(nama, ''), '') ASC
		LIMIT 20`

	rows, err := database.DB.Query(sqlQuery, searchArgs...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := make([]PegawaiSearchResult, 0)
	for rows.Next() {
		var item PegawaiSearchResult
		if err := rows.Scan(&item.NIP, &item.Nama, &item.Golongan, &item.Jabatan); err != nil {
			return nil, err
		}
		results = append(results, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func GetPegawaiByNIP(nip string) (*PegawaiDetail, error) {
	nip = strings.TrimSpace(nip)
	if nip == "" {
		return nil, sql.ErrNoRows
	}

	row := database.DB.QueryRow(`
		SELECT
			COALESCE(NULLIF(nip, ''), '') AS nip,
			COALESCE(NULLIF(namagelar, ''), NULLIF(nama, ''), '') AS nama,
			COALESCE(NULLIF(golongan, ''), '') AS golongan,
			COALESCE(NULLIF(jabatan, ''), '') AS jabatan
		FROM pegawai
		WHERE nip = ?
		LIMIT 1`, nip)

	var item PegawaiDetail
	if err := row.Scan(&item.NIP, &item.Nama, &item.Golongan, &item.Jabatan); err != nil {
		return nil, err
	}

	return &item, nil
}

func SearchDaftarTranslok(query string) ([]DaftarTranslokSearchResult, error) {
	sqlQuery := `
		SELECT
			COALESCE(NULLIF(kegiatan, ''), '') AS kegiatan,
			COALESCE(NULLIF(mak, ''), '') AS mak
		FROM daftartranslok
		WHERE 1 = 1`

	searchClause, searchArgs := buildFlexibleSearchClause([]string{"kegiatan", "CONCAT(COALESCE(kegiatan, ''), ' ', COALESCE(mak, ''))"}, query)
	sqlQuery += searchClause + `
		ORDER BY kegiatan ASC
		LIMIT 20`

	rows, err := database.DB.Query(sqlQuery, searchArgs...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := make([]DaftarTranslokSearchResult, 0)
	for rows.Next() {
		var item DaftarTranslokSearchResult
		if err := rows.Scan(&item.Nama, &item.Mak); err != nil {
			return nil, err
		}
		results = append(results, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func ListMatriks(filter MatriksFilter) ([]MatriksRecord, error) {
	query := `
		SELECT
			id_matriks,
			nip,
			nama,
			golongan,
			jenis_perjalanan,
			kegiatan,
			asal,
			tujuan,
			tglmulai,
			tglselesai,
			no_surtug,
			no_spd,
			jumlah,
			mak,
			no_fp,
			perihal_fp,
			tgl_surtug_utama,
			tgl_spd_utama
		FROM matriks
		WHERE 1 = 1`

	args := make([]interface{}, 0, 12)

	if filter.Tahun != "" {
		query += " AND YEAR(tglmulai) = ?"
		args = append(args, filter.Tahun)
	}

	if strings.TrimSpace(filter.Jenis) != "" {
		query += " AND jenis_perjalanan = ?"
		args = append(args, strings.TrimSpace(filter.Jenis))
	}

	if strings.TrimSpace(filter.Search) != "" {
		searchClause, searchArgs := buildFlexibleSearchClause(
			[]string{"nip", "nama", "kegiatan", "no_surtug", "no_spd", "no_fp", "asal", "tujuan"},
			filter.Search,
		)
		query += searchClause
		args = append(args, searchArgs...)
	}

	query += " ORDER BY tglmulai DESC, id_matriks DESC"

	rows, err := database.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]MatriksRecord, 0)
	for rows.Next() {
		item, err := scanMatriksRecord(rows)
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

func GetMatriksByID(id int) (*MatriksRecord, error) {
	row := database.DB.QueryRow(`
		SELECT
			id_matriks,
			nip,
			nama,
			golongan,
			jenis_perjalanan,
			kegiatan,
			asal,
			tujuan,
			tglmulai,
			tglselesai,
			no_surtug,
			no_spd,
			jumlah,
			mak,
			no_fp,
			perihal_fp,
			tgl_surtug_utama,
			tgl_spd_utama
		FROM matriks
		WHERE id_matriks = ?`, id)

	item, err := scanMatriksRecord(row)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func UpdateMatriks(id int, payload *MatriksEntryPayload) error {
	prepared, err := prepareMatriksEntryPayload(payload, id)
	if err != nil {
		return err
	}

	_, err = database.DB.Exec(`
		UPDATE matriks
		SET
			nip = ?,
			nama = ?,
			golongan = ?,
			jenis_perjalanan = ?,
			kegiatan = ?,
			asal = ?,
			tujuan = ?,
			tglmulai = ?,
			tglselesai = ?,
			no_surtug = ?,
			no_spd = ?,
			jumlah = ?,
			mak = ?,
			no_fp = ?,
			perihal_fp = ?,
			tgl_surtug_utama = ?,
			tgl_spd_utama = ?
		WHERE id_matriks = ?`,
		prepared.Payload.NIP,
		prepared.Payload.Nama,
		prepared.Payload.Golongan,
		prepared.Payload.JenisPerjalanan,
		prepared.Payload.Kegiatan,
		prepared.Payload.Asal,
		prepared.Payload.Tujuan,
		prepared.TanggalMulai.Format("2006-01-02"),
		prepared.TanggalSelesai.Format("2006-01-02"),
		prepared.Payload.NoSurtug,
		prepared.Payload.NoSpd,
		prepared.Payload.Jumlah,
		prepared.Payload.Mak,
		prepared.Payload.NoFP,
		prepared.Payload.PerihalFP,
		prepared.TglSurtugUtama.Format("2006-01-02"),
		prepared.TglSpdUtama.Format("2006-01-02"),
		id,
	)

	return err
}

func DeleteMatriks(id int) error {
	_, err := database.DB.Exec("DELETE FROM matriks WHERE id_matriks = ?", id)
	return err
}

func GetMatriksTahunOptions() ([]string, error) {
	rows, err := database.DB.Query(`
		SELECT DISTINCT YEAR(tglmulai) AS tahun
		FROM matriks
		WHERE tglmulai IS NOT NULL
		ORDER BY tahun DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var options []string
	for rows.Next() {
		var tahun sql.NullInt64
		if err := rows.Scan(&tahun); err != nil {
			return nil, err
		}
		if tahun.Valid {
			options = append(options, fmt.Sprintf("%d", tahun.Int64))
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return options, nil
}

func ListMatriksCalendar(filter MatriksCalendarFilter) ([]MatriksRecord, error) {
	if filter.Year <= 0 || filter.Month <= 0 || filter.Month > 12 {
		return nil, ValidationError{Message: "Filter bulan kalender tidak valid"}
	}

	startOfMonth := time.Date(filter.Year, time.Month(filter.Month), 1, 0, 0, 0, 0, time.UTC)
	endOfMonth := startOfMonth.AddDate(0, 1, -1)

	query := `
		SELECT
			id_matriks,
			nip,
			nama,
			golongan,
			jenis_perjalanan,
			kegiatan,
			asal,
			tujuan,
			tglmulai,
			tglselesai,
			no_surtug,
			no_spd,
			jumlah,
			mak,
			no_fp,
			perihal_fp,
			tgl_surtug_utama,
			tgl_spd_utama
		FROM matriks
		WHERE tglmulai <= ?
		  AND tglselesai >= ?`

	args := []interface{}{
		endOfMonth.Format("2006-01-02"),
		startOfMonth.Format("2006-01-02"),
	}

	if strings.TrimSpace(filter.Jenis) != "" {
		query += " AND jenis_perjalanan = ?"
		args = append(args, strings.TrimSpace(filter.Jenis))
	}

	if strings.TrimSpace(filter.Search) != "" {
		searchClause, searchArgs := buildFlexibleSearchClause(
			[]string{"nip", "nama", "kegiatan", "no_surtug", "no_spd", "asal", "tujuan"},
			filter.Search,
		)
		query += searchClause
		args = append(args, searchArgs...)
	}

	query += " ORDER BY tglmulai ASC, nama ASC, id_matriks ASC"

	rows, err := database.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]MatriksRecord, 0)
	for rows.Next() {
		item, err := scanMatriksRecord(rows)
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

func ListMatriksAdministrasi(filter MatriksAdministrasiFilter) ([]MatriksAdministrasiSummary, error) {
	hasTransportReport, err := tableExists("laporan_perjalanan")
	if err != nil {
		return nil, err
	}
	hasSPPDReport, err := tableExists("laporan_sppd_detail")
	if err != nil {
		return nil, err
	}

	selectReports := "0"
	joinReports := ""
	if hasTransportReport || hasSPPDReport {
		reportParts := make([]string, 0, 2)
		if hasTransportReport {
			joinReports += " LEFT JOIN laporan_perjalanan lp ON m.id_matriks = lp.id_matriks_fk"
			reportParts = append(reportParts, "COUNT(DISTINCT lp.id_laporan)")
		}
		if hasSPPDReport {
			joinReports += " LEFT JOIN laporan_sppd_detail lsd ON m.id_matriks = lsd.id_matriks_fk"
			reportParts = append(reportParts, "COUNT(DISTINCT lsd.id_laporan_sppd)")
		}
		selectReports = strings.Join(reportParts, " + ")
	}

	query := fmt.Sprintf(`
		SELECT
			m.no_fp,
			MAX(COALESCE(NULLIF(m.mak, ''), '')) AS mak,
			MAX(COALESCE(NULLIF(m.perihal_fp, ''), '')) AS perihal_fp,
			COUNT(DISTINCT m.id_matriks) AS jumlah_peserta,
			%s AS jumlah_laporan,
			MIN(m.id_matriks) AS sample_id
		FROM matriks m
		%s
		WHERE m.no_fp IS NOT NULL
		  AND m.no_fp != ''`, selectReports, joinReports)

	args := make([]interface{}, 0, 3)
	if strings.TrimSpace(filter.Search) != "" {
		searchClause, searchArgs := buildFlexibleSearchClause(
			[]string{"m.perihal_fp", "m.no_fp", "m.mak"},
			filter.Search,
		)
		query += searchClause
		args = append(args, searchArgs...)
	}

	query += `
		GROUP BY m.no_fp
		ORDER BY MAX(m.tgl_surtug_utama) DESC, sample_id DESC`

	rows, err := database.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]MatriksAdministrasiSummary, 0)
	for rows.Next() {
		var item MatriksAdministrasiSummary
		var noFP, mak, perihal sql.NullString
		if err := rows.Scan(&noFP, &mak, &perihal, &item.JumlahPeserta, &item.JumlahLaporan, &item.SampleID); err != nil {
			return nil, err
		}
		if noFP.Valid {
			item.NoFP = noFP.String
		}
		if mak.Valid {
			item.Mak = mak.String
		}
		if perihal.Valid {
			item.PerihalFP = perihal.String
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func ListMatriksAdministrasiPeserta(noFP string) ([]MatriksAdministrasiPeserta, error) {
	noFP = strings.TrimSpace(noFP)
	if noFP == "" {
		return nil, ValidationError{Message: "Nomor FP wajib diisi"}
	}

	hasTransportReport, err := tableExists("laporan_perjalanan")
	if err != nil {
		return nil, err
	}
	hasSPPDReport, err := tableExists("laporan_sppd_detail")
	if err != nil {
		return nil, err
	}

	reportSelect := "'' AS report_type"
	joinReports := ""
	if hasTransportReport || hasSPPDReport {
		caseParts := make([]string, 0, 2)
		if hasTransportReport {
			joinReports += " LEFT JOIN laporan_perjalanan lp ON m.id_matriks = lp.id_matriks_fk"
			caseParts = append(caseParts, "WHEN lp.id_laporan IS NOT NULL THEN 'transport'")
		}
		if hasSPPDReport {
			joinReports += " LEFT JOIN laporan_sppd_detail lsd ON m.id_matriks = lsd.id_matriks_fk"
			caseParts = append(caseParts, "WHEN lsd.id_laporan_sppd IS NOT NULL THEN 'sppd'")
		}
		reportSelect = "CASE " + strings.Join(caseParts, " ") + " ELSE '' END AS report_type"
	}

	query := fmt.Sprintf(`
		SELECT
			m.id_matriks,
			COALESCE(NULLIF(m.nama, ''), ''),
			COALESCE(NULLIF(m.nip, ''), ''),
			COALESCE(NULLIF(m.jenis_perjalanan, ''), ''),
			COALESCE(NULLIF(m.tujuan, ''), ''),
			m.tglmulai,
			m.tglselesai,
			COALESCE(NULLIF(m.no_surtug, ''), ''),
			COALESCE(NULLIF(m.no_spd, ''), ''),
			%s
		FROM matriks m
		%s
		WHERE m.no_fp = ?
		ORDER BY m.tglmulai ASC, m.nama ASC, m.id_matriks ASC`, reportSelect, joinReports)

	rows, err := database.DB.Query(query, noFP)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]MatriksAdministrasiPeserta, 0)
	for rows.Next() {
		var item MatriksAdministrasiPeserta
		var jenis, tujuan, noSurtug, noSpd, reportType sql.NullString
		var tanggalMulai, tanggalSelesai sql.NullTime
		if err := rows.Scan(
			&item.ID,
			&item.Nama,
			&item.NIP,
			&jenis,
			&tujuan,
			&tanggalMulai,
			&tanggalSelesai,
			&noSurtug,
			&noSpd,
			&reportType,
		); err != nil {
			return nil, err
		}

		if jenis.Valid {
			item.JenisPerjalanan = jenis.String
		}
		if tujuan.Valid {
			item.Tujuan = tujuan.String
		}
		if tanggalMulai.Valid {
			item.TanggalMulai = tanggalMulai.Time.Format("2006-01-02")
		}
		if tanggalSelesai.Valid {
			item.TanggalSelesai = tanggalSelesai.Time.Format("2006-01-02")
		}
		if noSurtug.Valid {
			item.NoSurtug = noSurtug.String
		}
		if noSpd.Valid {
			item.NoSpd = noSpd.String
		}
		if reportType.Valid {
			item.ReportType = reportType.String
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func GetMatriksPrintBase(id int) (*MatriksPrintBase, error) {
	row := database.DB.QueryRow(`
		SELECT
			m.id_matriks,
			COALESCE(NULLIF(m.nama, ''), ''),
			COALESCE(NULLIF(p.namagelar, ''), COALESCE(NULLIF(m.nama, ''), '')),
			COALESCE(NULLIF(m.nip, ''), ''),
			COALESCE(NULLIF(m.golongan, ''), ''),
			COALESCE(NULLIF(p.jabatan, ''), ''),
			COALESCE(NULLIF(m.jenis_perjalanan, ''), ''),
			COALESCE(NULLIF(m.kegiatan, ''), ''),
			COALESCE(NULLIF(m.asal, ''), ''),
			COALESCE(NULLIF(m.tujuan, ''), ''),
			COALESCE(NULLIF(m.mak, ''), ''),
			COALESCE(NULLIF(m.perihal_fp, ''), ''),
			COALESCE(NULLIF(m.no_surtug, ''), ''),
			COALESCE(NULLIF(m.no_spd, ''), ''),
			m.tglmulai,
			m.tglselesai,
			m.tgl_surtug_utama,
			m.tgl_spd_utama
		FROM matriks m
		LEFT JOIN pegawai p ON m.nip = p.nip
		WHERE m.id_matriks = ?`, id)

	var item MatriksPrintBase
	if err := row.Scan(
		&item.ID,
		&item.Nama,
		&item.NamaGelar,
		&item.NIP,
		&item.Golongan,
		&item.Jabatan,
		&item.JenisPerjalanan,
		&item.Kegiatan,
		&item.Asal,
		&item.Tujuan,
		&item.Mak,
		&item.PerihalFP,
		&item.NoSurtug,
		&item.NoSpd,
		&item.TanggalMulai,
		&item.TanggalSelesai,
		&item.TglSurtugUtama,
		&item.TglSpdUtama,
	); err != nil {
		return nil, err
	}

	return &item, nil
}

func GetMatriksTransportPrintData(id int) (*MatriksTransportPrintData, error) {
	row := database.DB.QueryRow(`
		SELECT
			m.id_matriks,
			COALESCE(NULLIF(m.nama, ''), ''),
			COALESCE(NULLIF(p.namagelar, ''), COALESCE(NULLIF(m.nama, ''), '')),
			COALESCE(NULLIF(m.nip, ''), ''),
			COALESCE(NULLIF(m.golongan, ''), ''),
			COALESCE(NULLIF(p.jabatan, ''), ''),
			COALESCE(NULLIF(m.jenis_perjalanan, ''), ''),
			COALESCE(NULLIF(m.kegiatan, ''), ''),
			COALESCE(NULLIF(m.asal, ''), ''),
			COALESCE(NULLIF(m.tujuan, ''), ''),
			COALESCE(NULLIF(m.mak, ''), ''),
			COALESCE(NULLIF(m.perihal_fp, ''), ''),
			COALESCE(NULLIF(m.no_surtug, ''), ''),
			COALESCE(NULLIF(m.no_spd, ''), ''),
			m.tglmulai,
			m.tglselesai,
			m.tgl_surtug_utama,
			m.tgl_spd_utama,
			COALESCE(lp.laporan_text, ''),
			COALESCE(lp.laporan_foto, '')
		FROM matriks m
		LEFT JOIN pegawai p ON m.nip = p.nip
		INNER JOIN laporan_perjalanan lp ON m.id_matriks = lp.id_matriks_fk
		WHERE m.id_matriks = ?`, id)

	var item MatriksTransportPrintData
	if err := row.Scan(
		&item.Base.ID,
		&item.Base.Nama,
		&item.Base.NamaGelar,
		&item.Base.NIP,
		&item.Base.Golongan,
		&item.Base.Jabatan,
		&item.Base.JenisPerjalanan,
		&item.Base.Kegiatan,
		&item.Base.Asal,
		&item.Base.Tujuan,
		&item.Base.Mak,
		&item.Base.PerihalFP,
		&item.Base.NoSurtug,
		&item.Base.NoSpd,
		&item.Base.TanggalMulai,
		&item.Base.TanggalSelesai,
		&item.Base.TglSurtugUtama,
		&item.Base.TglSpdUtama,
		&item.LaporanText,
		&item.LaporanFoto,
	); err != nil {
		return nil, err
	}

	return &item, nil
}

func GetMatriksSPPDPrintData(id int) (*MatriksSPPDPrintData, error) {
	row := database.DB.QueryRow(`
		SELECT
			m.id_matriks,
			COALESCE(NULLIF(m.nama, ''), ''),
			COALESCE(NULLIF(p.namagelar, ''), COALESCE(NULLIF(m.nama, ''), '')),
			COALESCE(NULLIF(m.nip, ''), ''),
			COALESCE(NULLIF(m.golongan, ''), ''),
			COALESCE(NULLIF(p.jabatan, ''), ''),
			COALESCE(NULLIF(m.jenis_perjalanan, ''), ''),
			COALESCE(NULLIF(m.kegiatan, ''), ''),
			COALESCE(NULLIF(m.asal, ''), ''),
			COALESCE(NULLIF(m.tujuan, ''), ''),
			COALESCE(NULLIF(m.mak, ''), ''),
			COALESCE(NULLIF(m.perihal_fp, ''), ''),
			COALESCE(NULLIF(m.no_surtug, ''), ''),
			COALESCE(NULLIF(m.no_spd, ''), ''),
			m.tglmulai,
			m.tglselesai,
			m.tgl_surtug_utama,
			m.tgl_spd_utama,
			COALESCE(TIME_FORMAT(lsd.waktu_berangkat_mulai, '%H:%i'), ''),
			COALESCE(TIME_FORMAT(lsd.waktu_berangkat_selesai, '%H:%i'), ''),
			COALESCE(lsd.isian_berangkat, ''),
			COALESCE(TIME_FORMAT(lsd.waktu_pengawasan1_mulai, '%H:%i'), ''),
			COALESCE(TIME_FORMAT(lsd.waktu_pengawasan1_selesai, '%H:%i'), ''),
			COALESCE(lsd.isian_pengawasan1, ''),
			COALESCE(TIME_FORMAT(lsd.waktu_isoma_mulai, '%H:%i'), ''),
			COALESCE(TIME_FORMAT(lsd.waktu_isoma_selesai, '%H:%i'), ''),
			COALESCE(TIME_FORMAT(lsd.waktu_pengawasan2_mulai, '%H:%i'), ''),
			COALESCE(TIME_FORMAT(lsd.waktu_pengawasan2_selesai, '%H:%i'), ''),
			COALESCE(lsd.isian_pengawasan2, ''),
			COALESCE(TIME_FORMAT(lsd.waktu_pulang_mulai, '%H:%i'), ''),
			COALESCE(TIME_FORMAT(lsd.waktu_pulang_selesai, '%H:%i'), ''),
			COALESCE(lsd.isian_pulang, ''),
			COALESCE(lsd.laporan_foto, '')
		FROM matriks m
		LEFT JOIN pegawai p ON m.nip = p.nip
		INNER JOIN laporan_sppd_detail lsd ON m.id_matriks = lsd.id_matriks_fk
		WHERE m.id_matriks = ?`, id)

	var item MatriksSPPDPrintData
	if err := row.Scan(
		&item.Base.ID,
		&item.Base.Nama,
		&item.Base.NamaGelar,
		&item.Base.NIP,
		&item.Base.Golongan,
		&item.Base.Jabatan,
		&item.Base.JenisPerjalanan,
		&item.Base.Kegiatan,
		&item.Base.Asal,
		&item.Base.Tujuan,
		&item.Base.Mak,
		&item.Base.PerihalFP,
		&item.Base.NoSurtug,
		&item.Base.NoSpd,
		&item.Base.TanggalMulai,
		&item.Base.TanggalSelesai,
		&item.Base.TglSurtugUtama,
		&item.Base.TglSpdUtama,
		&item.WaktuBerangkatMulai,
		&item.WaktuBerangkatSelesai,
		&item.IsianBerangkat,
		&item.WaktuPengawasan1Mulai,
		&item.WaktuPengawasan1Selesai,
		&item.IsianPengawasan1,
		&item.WaktuIsomaMulai,
		&item.WaktuIsomaSelesai,
		&item.WaktuPengawasan2Mulai,
		&item.WaktuPengawasan2Selesai,
		&item.IsianPengawasan2,
		&item.WaktuPulangMulai,
		&item.WaktuPulangSelesai,
		&item.IsianPulang,
		&item.LaporanFoto,
	); err != nil {
		return nil, err
	}

	return &item, nil
}

func ParseMatriksDate(value string) (time.Time, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return time.Time{}, ValidationError{Message: "Tanggal wajib diisi"}
	}

	layouts := []string{
		"2006-01-02",
		"02/01/2006",
		"02-01-2006",
	}

	for _, layout := range layouts {
		parsed, err := time.Parse(layout, value)
		if err == nil {
			return parsed, nil
		}
	}

	return time.Time{}, ValidationError{Message: fmt.Sprintf("Format tanggal tidak valid: %s", value)}
}

func prepareMatriksEntryPayload(payload *MatriksEntryPayload, excludeID int) (*matriksPreparedEntry, error) {
	if payload == nil {
		return nil, ValidationError{Message: "Data matriks tidak tersedia"}
	}

	payload.NIP = strings.TrimSpace(payload.NIP)
	payload.Nama = strings.TrimSpace(payload.Nama)
	payload.Golongan = strings.TrimSpace(payload.Golongan)
	payload.JenisPerjalanan = strings.TrimSpace(payload.JenisPerjalanan)
	payload.Kegiatan = strings.TrimSpace(payload.Kegiatan)
	payload.Asal = strings.TrimSpace(payload.Asal)
	payload.Tujuan = strings.TrimSpace(payload.Tujuan)
	payload.TanggalMulai = strings.TrimSpace(payload.TanggalMulai)
	payload.TanggalSelesai = strings.TrimSpace(payload.TanggalSelesai)
	payload.NoSurtug = strings.TrimSpace(payload.NoSurtug)
	payload.NoSpd = strings.TrimSpace(payload.NoSpd)
	payload.Mak = strings.TrimSpace(payload.Mak)
	payload.NoFP = strings.TrimSpace(payload.NoFP)
	payload.PerihalFP = strings.TrimSpace(payload.PerihalFP)
	payload.TglSurtugUtama = strings.TrimSpace(payload.TglSurtugUtama)
	payload.TglSpdUtama = strings.TrimSpace(payload.TglSpdUtama)

	if payload.NIP == "" {
		return nil, ValidationError{Message: "NIP wajib diisi"}
	}
	if payload.Nama == "" {
		return nil, ValidationError{Message: "Nama wajib diisi"}
	}
	if payload.JenisPerjalanan == "" {
		return nil, ValidationError{Message: "Jenis perjalanan wajib dipilih"}
	}
	if payload.Kegiatan == "" {
		return nil, ValidationError{Message: "Kegiatan wajib diisi"}
	}
	if payload.Asal == "" || payload.Tujuan == "" {
		return nil, ValidationError{Message: "Asal dan tujuan wajib diisi"}
	}
	if payload.NoSurtug == "" || payload.NoSpd == "" {
		return nil, ValidationError{Message: "Nomor surtug dan SPD wajib diisi"}
	}
	if payload.Jumlah <= 0 {
		return nil, ValidationError{Message: "Jumlah perjalanan tidak valid"}
	}

	tanggalMulai, err := ParseMatriksDate(payload.TanggalMulai)
	if err != nil {
		return nil, ValidationError{Message: "Tanggal mulai tidak valid"}
	}
	tanggalSelesai, err := ParseMatriksDate(payload.TanggalSelesai)
	if err != nil {
		return nil, ValidationError{Message: "Tanggal selesai tidak valid"}
	}
	if tanggalSelesai.Before(tanggalMulai) {
		return nil, ValidationError{Message: "Tanggal selesai tidak boleh lebih awal dari tanggal mulai"}
	}

	tglSurtugUtama, err := ParseMatriksDate(payload.TglSurtugUtama)
	if err != nil {
		return nil, ValidationError{Message: "Tanggal surtug utama tidak valid"}
	}
	tglSpdUtama, err := ParseMatriksDate(payload.TglSpdUtama)
	if err != nil {
		return nil, ValidationError{Message: "Tanggal SPD utama tidak valid"}
	}

	conflict, err := CheckMatriksConflict(payload.NIP, tanggalMulai, tanggalSelesai, excludeID)
	if err != nil {
		return nil, err
	}
	if conflict != nil && conflict.Bentrok {
		return nil, ValidationError{Message: conflict.Pesan}
	}

	return &matriksPreparedEntry{
		Payload:        payload,
		TanggalMulai:   tanggalMulai,
		TanggalSelesai: tanggalSelesai,
		TglSurtugUtama: tglSurtugUtama,
		TglSpdUtama:    tglSpdUtama,
	}, nil
}

func CheckMatriksConflict(nip string, tanggalMulai, tanggalSelesai time.Time, excludeID int) (*MatriksConflictResult, error) {
	return checkMatriksConflict(database.DB, nip, tanggalMulai, tanggalSelesai, excludeID)
}

func checkMatriksConflict(queryer matriksQueryer, nip string, tanggalMulai, tanggalSelesai time.Time, excludeID int) (*MatriksConflictResult, error) {
	result := &MatriksConflictResult{Bentrok: false, Pesan: ""}

	nip = strings.TrimSpace(nip)
	if nip == "" || tanggalMulai.IsZero() || tanggalSelesai.IsZero() {
		return result, nil
	}

	query := `
		SELECT jenis_perjalanan, kegiatan, tglmulai, tglselesai
		FROM matriks
		WHERE nip = ?
		  AND tglmulai <= ?
		  AND tglselesai >= ?`

	args := []interface{}{nip, tanggalSelesai.Format("2006-01-02"), tanggalMulai.Format("2006-01-02")}
	if excludeID > 0 {
		query += " AND id_matriks != ?"
		args = append(args, excludeID)
	}
	query += " LIMIT 1"

	var jenis, kegiatan string
	var mulai, selesai time.Time
	err := queryer.QueryRow(query, args...).Scan(&jenis, &kegiatan, &mulai, &selesai)
	if err != nil {
		if err == sql.ErrNoRows {
			return result, nil
		}
		return nil, err
	}

	result.Bentrok = true
	result.Pesan = fmt.Sprintf(
		"Jadwal bentrok dengan %s - %s (%s - %s)",
		strings.TrimSpace(jenis),
		strings.TrimSpace(kegiatan),
		mulai.Format("02/01/2006"),
		selesai.Format("02/01/2006"),
	)

	return result, nil
}

func CreateMatriksBatch(payload *MatriksCreatePayload) (int, error) {
	if payload == nil {
		return 0, ValidationError{Message: "Payload matriks tidak tersedia"}
	}

	payload.JenisPerjalanan = strings.TrimSpace(payload.JenisPerjalanan)
	payload.Kegiatan = strings.TrimSpace(payload.Kegiatan)
	payload.Mak = strings.TrimSpace(payload.Mak)
	payload.NoFP = strings.TrimSpace(payload.NoFP)
	payload.PerihalFP = strings.TrimSpace(payload.PerihalFP)
	payload.TglSurtugUtama = strings.TrimSpace(payload.TglSurtugUtama)
	payload.TglSpdUtama = strings.TrimSpace(payload.TglSpdUtama)

	if payload.JenisPerjalanan == "" {
		return 0, ValidationError{Message: "Jenis perjalanan wajib dipilih"}
	}
	if payload.Kegiatan == "" {
		return 0, ValidationError{Message: "Kegiatan wajib diisi"}
	}
	if len(payload.Peserta) == 0 {
		return 0, ValidationError{Message: "Tambahkan minimal satu peserta"}
	}

	tglSurtugUtama, err := ParseMatriksDate(payload.TglSurtugUtama)
	if err != nil {
		return 0, err
	}

	tglSpdUtama, err := ParseMatriksDate(payload.TglSpdUtama)
	if err != nil {
		return 0, err
	}

	tx, err := database.DB.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`
		INSERT INTO matriks
			(nip, nama, golongan, jenis_perjalanan, kegiatan, asal, tujuan, tglmulai, tglselesai, no_surtug, no_spd, jumlah, mak, no_fp, perihal_fp, tgl_surtug_utama, tgl_spd_utama)
		VALUES
			(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	inserted := 0
	for index := range payload.Peserta {
		peserta := payload.Peserta[index]

		peserta.NIP = strings.TrimSpace(peserta.NIP)
		peserta.Nama = strings.TrimSpace(peserta.Nama)
		peserta.Golongan = strings.TrimSpace(peserta.Golongan)
		peserta.Asal = strings.TrimSpace(peserta.Asal)
		peserta.Tujuan = strings.TrimSpace(peserta.Tujuan)
		peserta.NoSurtug = strings.TrimSpace(peserta.NoSurtug)
		peserta.NoSpd = strings.TrimSpace(peserta.NoSpd)
		peserta.TanggalMulai = strings.TrimSpace(peserta.TanggalMulai)
		peserta.TanggalSelesai = strings.TrimSpace(peserta.TanggalSelesai)

		rowLabel := fmt.Sprintf("Peserta %d", index+1)
		if peserta.Nama != "" {
			rowLabel = peserta.Nama
		}

		if peserta.NIP == "" {
			return 0, ValidationError{Message: fmt.Sprintf("NIP wajib diisi untuk %s", rowLabel)}
		}
		if peserta.Nama == "" {
			return 0, ValidationError{Message: fmt.Sprintf("Nama wajib diisi untuk %s", rowLabel)}
		}
		if peserta.Asal == "" || peserta.Tujuan == "" {
			return 0, ValidationError{Message: fmt.Sprintf("Asal dan tujuan wajib diisi untuk %s", rowLabel)}
		}
		if peserta.NoSurtug == "" || peserta.NoSpd == "" {
			return 0, ValidationError{Message: fmt.Sprintf("Nomor surtug dan SPD wajib diisi untuk %s", rowLabel)}
		}
		if peserta.Jumlah <= 0 {
			return 0, ValidationError{Message: fmt.Sprintf("Jumlah perjalanan tidak valid untuk %s", rowLabel)}
		}

		tglMulai, err := ParseMatriksDate(peserta.TanggalMulai)
		if err != nil {
			return 0, ValidationError{Message: fmt.Sprintf("Tanggal mulai tidak valid untuk %s", rowLabel)}
		}

		tglSelesai, err := ParseMatriksDate(peserta.TanggalSelesai)
		if err != nil {
			return 0, ValidationError{Message: fmt.Sprintf("Tanggal selesai tidak valid untuk %s", rowLabel)}
		}

		if tglSelesai.Before(tglMulai) {
			return 0, ValidationError{Message: fmt.Sprintf("Tanggal selesai tidak boleh lebih awal dari tanggal mulai untuk %s", rowLabel)}
		}

		conflict, err := checkMatriksConflict(tx, peserta.NIP, tglMulai, tglSelesai, 0)
		if err != nil {
			return 0, err
		}
		if conflict != nil && conflict.Bentrok {
			return 0, ValidationError{Message: fmt.Sprintf("%s memiliki jadwal bentrok. %s", rowLabel, conflict.Pesan)}
		}

		if _, err := stmt.Exec(
			peserta.NIP,
			peserta.Nama,
			peserta.Golongan,
			payload.JenisPerjalanan,
			payload.Kegiatan,
			peserta.Asal,
			peserta.Tujuan,
			tglMulai.Format("2006-01-02"),
			tglSelesai.Format("2006-01-02"),
			peserta.NoSurtug,
			peserta.NoSpd,
			peserta.Jumlah,
			payload.Mak,
			payload.NoFP,
			payload.PerihalFP,
			tglSurtugUtama.Format("2006-01-02"),
			tglSpdUtama.Format("2006-01-02"),
		); err != nil {
			return 0, err
		}

		inserted++
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return inserted, nil
}

func ImportMatriksEntries(entries []MatriksEntryPayload) (int, []MatriksImportRowError, error) {
	if len(entries) == 0 {
		return 0, []MatriksImportRowError{{
			Row:     0,
			Message: "File import tidak berisi data",
		}}, nil
	}

	preparedEntries := make([]*matriksPreparedEntry, 0, len(entries))
	rowErrors := make([]MatriksImportRowError, 0)

	for index := range entries {
		rowNumber := index + 2
		entry := entries[index]
		entry.NIP = strings.TrimSpace(entry.NIP)
		entry.Nama = strings.TrimSpace(entry.Nama)
		entry.Golongan = strings.TrimSpace(entry.Golongan)

		if entry.NIP != "" {
			if pegawai, err := GetPegawaiByNIP(entry.NIP); err == nil {
				if entry.Nama == "" {
					entry.Nama = pegawai.Nama
				}
				if entry.Golongan == "" {
					entry.Golongan = pegawai.Golongan
				}
			}
		}

		prepared, err := prepareMatriksEntryPayload(&entry, 0)
		if err != nil {
			rowErrors = append(rowErrors, MatriksImportRowError{
				Row:     rowNumber,
				NIP:     entry.NIP,
				Nama:    entry.Nama,
				Message: err.Error(),
			})
			continue
		}

		if conflictMessage := checkPreparedEntriesConflict(preparedEntries, prepared); conflictMessage != "" {
			rowErrors = append(rowErrors, MatriksImportRowError{
				Row:     rowNumber,
				NIP:     entry.NIP,
				Nama:    entry.Nama,
				Message: conflictMessage,
			})
			continue
		}

		preparedEntries = append(preparedEntries, prepared)
	}

	if len(rowErrors) > 0 {
		return 0, rowErrors, nil
	}

	tx, err := database.DB.Begin()
	if err != nil {
		return 0, nil, err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`
		INSERT INTO matriks
			(nip, nama, golongan, jenis_perjalanan, kegiatan, asal, tujuan, tglmulai, tglselesai, no_surtug, no_spd, jumlah, mak, no_fp, perihal_fp, tgl_surtug_utama, tgl_spd_utama)
		VALUES
			(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return 0, nil, err
	}
	defer stmt.Close()

	inserted := 0
	for _, prepared := range preparedEntries {
		if _, err := stmt.Exec(
			prepared.Payload.NIP,
			prepared.Payload.Nama,
			prepared.Payload.Golongan,
			prepared.Payload.JenisPerjalanan,
			prepared.Payload.Kegiatan,
			prepared.Payload.Asal,
			prepared.Payload.Tujuan,
			prepared.TanggalMulai.Format("2006-01-02"),
			prepared.TanggalSelesai.Format("2006-01-02"),
			prepared.Payload.NoSurtug,
			prepared.Payload.NoSpd,
			prepared.Payload.Jumlah,
			prepared.Payload.Mak,
			prepared.Payload.NoFP,
			prepared.Payload.PerihalFP,
			prepared.TglSurtugUtama.Format("2006-01-02"),
			prepared.TglSpdUtama.Format("2006-01-02"),
		); err != nil {
			return 0, nil, err
		}
		inserted++
	}

	if err := tx.Commit(); err != nil {
		return 0, nil, err
	}

	return inserted, nil, nil
}

func checkPreparedEntriesConflict(existing []*matriksPreparedEntry, candidate *matriksPreparedEntry) string {
	if candidate == nil || candidate.Payload == nil {
		return ""
	}

	for _, item := range existing {
		if item == nil || item.Payload == nil {
			continue
		}
		if strings.TrimSpace(item.Payload.NIP) != strings.TrimSpace(candidate.Payload.NIP) {
			continue
		}
		if item.TanggalMulai.After(candidate.TanggalSelesai) || item.TanggalSelesai.Before(candidate.TanggalMulai) {
			continue
		}

		return fmt.Sprintf(
			"Jadwal bentrok dengan data import lain: %s (%s - %s)",
			strings.TrimSpace(item.Payload.Kegiatan),
			item.TanggalMulai.Format("02/01/2006"),
			item.TanggalSelesai.Format("02/01/2006"),
		)
	}

	return ""
}

type matriksScanner interface {
	Scan(dest ...interface{}) error
}

func scanMatriksRecord(scanner matriksScanner) (MatriksRecord, error) {
	var item MatriksRecord
	var nip, nama, golongan, jenis, kegiatan, asal, tujuan, noSurtug, noSpd, mak, noFP, perihal sql.NullString
	var jumlah sql.NullInt64
	var tanggalMulai, tanggalSelesai, tglSurtugUtama, tglSpdUtama sql.NullTime

	err := scanner.Scan(
		&item.ID,
		&nip,
		&nama,
		&golongan,
		&jenis,
		&kegiatan,
		&asal,
		&tujuan,
		&tanggalMulai,
		&tanggalSelesai,
		&noSurtug,
		&noSpd,
		&jumlah,
		&mak,
		&noFP,
		&perihal,
		&tglSurtugUtama,
		&tglSpdUtama,
	)
	if err != nil {
		return item, err
	}

	if nip.Valid {
		item.NIP = nip.String
	}
	if nama.Valid {
		item.Nama = nama.String
	}
	if golongan.Valid {
		item.Golongan = golongan.String
	}
	if jenis.Valid {
		item.JenisPerjalanan = jenis.String
	}
	if kegiatan.Valid {
		item.Kegiatan = kegiatan.String
	}
	if asal.Valid {
		item.Asal = asal.String
	}
	if tujuan.Valid {
		item.Tujuan = tujuan.String
	}
	if tanggalMulai.Valid {
		item.TanggalMulai = tanggalMulai.Time.Format("2006-01-02")
	}
	if tanggalSelesai.Valid {
		item.TanggalSelesai = tanggalSelesai.Time.Format("2006-01-02")
	}
	if noSurtug.Valid {
		item.NoSurtug = noSurtug.String
	}
	if noSpd.Valid {
		item.NoSpd = noSpd.String
	}
	if jumlah.Valid {
		item.Jumlah = int(jumlah.Int64)
	}
	if mak.Valid {
		item.Mak = mak.String
	}
	if noFP.Valid {
		item.NoFP = noFP.String
	}
	if perihal.Valid {
		item.PerihalFP = perihal.String
	}
	if tglSurtugUtama.Valid {
		item.TglSurtugUtama = tglSurtugUtama.Time.Format("2006-01-02")
	}
	if tglSpdUtama.Valid {
		item.TglSpdUtama = tglSpdUtama.Time.Format("2006-01-02")
	}

	return item, nil
}

func tableExists(tableName string) (bool, error) {
	var count int
	err := database.DB.QueryRow(`
		SELECT COUNT(*)
		FROM information_schema.tables
		WHERE table_schema = DATABASE()
		  AND table_name = ?`, tableName).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func tableColumnExists(tableName, columnName string) (bool, error) {
	var count int
	err := database.DB.QueryRow(`
		SELECT COUNT(*)
		FROM information_schema.columns
		WHERE table_schema = DATABASE()
		  AND table_name = ?
		  AND column_name = ?`, tableName, columnName).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
