package models

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"sirekap/internal/database"
)

type UserTransportReportItem struct {
	ID           int    `json:"id"`
	PerihalFP    string `json:"perihalFp"`
	Tujuan       string `json:"tujuan"`
	TanggalMulai string `json:"tanggalMulai"`
	TanggalAkhir string `json:"tanggalAkhir"`
	LaporanText  string `json:"laporanText,omitempty"`
	LaporanFoto  string `json:"laporanFoto,omitempty"`
	CreatedAt    string `json:"createdAt,omitempty"`
}

type UserSPPDReportItem struct {
	ID                      int    `json:"id"`
	PerihalFP               string `json:"perihalFp"`
	Tujuan                  string `json:"tujuan"`
	TanggalMulai            string `json:"tanggalMulai"`
	TanggalAkhir            string `json:"tanggalAkhir"`
	WaktuBerangkatMulai     string `json:"waktuBerangkatMulai,omitempty"`
	WaktuBerangkatSelesai   string `json:"waktuBerangkatSelesai,omitempty"`
	IsianBerangkat          string `json:"isianBerangkat,omitempty"`
	WaktuPengawasan1Mulai   string `json:"waktuPengawasan1Mulai,omitempty"`
	WaktuPengawasan1Selesai string `json:"waktuPengawasan1Selesai,omitempty"`
	IsianPengawasan1        string `json:"isianPengawasan1,omitempty"`
	WaktuIsomaMulai         string `json:"waktuIsomaMulai,omitempty"`
	WaktuIsomaSelesai       string `json:"waktuIsomaSelesai,omitempty"`
	IsianIsoma              string `json:"isianIsoma,omitempty"`
	WaktuPengawasan2Mulai   string `json:"waktuPengawasan2Mulai,omitempty"`
	WaktuPengawasan2Selesai string `json:"waktuPengawasan2Selesai,omitempty"`
	IsianPengawasan2        string `json:"isianPengawasan2,omitempty"`
	WaktuPulangMulai        string `json:"waktuPulangMulai,omitempty"`
	WaktuPulangSelesai      string `json:"waktuPulangSelesai,omitempty"`
	IsianPulang             string `json:"isianPulang,omitempty"`
	LaporanFoto             string `json:"laporanFoto,omitempty"`
	CreatedAt               string `json:"createdAt,omitempty"`
}

type SaveSPPDReportPayload struct {
	WaktuBerangkatMulai     string
	WaktuBerangkatSelesai   string
	IsianBerangkat          string
	WaktuPengawasan1Mulai   string
	WaktuPengawasan1Selesai string
	IsianPengawasan1        string
	WaktuIsomaMulai         string
	WaktuIsomaSelesai       string
	IsianIsoma              string
	WaktuPengawasan2Mulai   string
	WaktuPengawasan2Selesai string
	IsianPengawasan2        string
	WaktuPulangMulai        string
	WaktuPulangSelesai      string
	IsianPulang             string
	LaporanFotoPath         string
}

func ListPendingTransportReports(userNIP, userName string) ([]UserTransportReportItem, error) {
	return listUserTransportReports(userNIP, userName, false)
}

func ListSubmittedTransportReports(userNIP, userName string) ([]UserTransportReportItem, error) {
	return listUserTransportReports(userNIP, userName, true)
}

func ListPendingSPPDReports(userNIP, userName string) ([]UserSPPDReportItem, error) {
	return listUserSPPDReports(userNIP, userName, false)
}

func ListSubmittedSPPDReports(userNIP, userName string) ([]UserSPPDReportItem, error) {
	return listUserSPPDReports(userNIP, userName, true)
}

func GetPendingTransportReportByID(id int, userNIP, userName string) (*UserTransportReportItem, error) {
	userClause, userArgs := buildUserOwnershipClause(userNIP, userName)
	if userClause == "" {
		return nil, ValidationError{Message: "Identitas user tidak ditemukan"}
	}

	query := `
		SELECT
			m.id_matriks,
			COALESCE(NULLIF(m.perihal_fp, ''), ''),
			COALESCE(NULLIF(m.tujuan, ''), ''),
			m.tglmulai,
			m.tglselesai
		FROM matriks m
		LEFT JOIN laporan_perjalanan lp ON m.id_matriks = lp.id_matriks_fk
		WHERE m.id_matriks = ?
		  AND m.jenis_perjalanan = 'Transport Lokal'
		  AND lp.id_laporan IS NULL
		  AND (` + userClause + `)
		LIMIT 1`

	args := make([]interface{}, 0, len(userArgs)+1)
	args = append(args, id)
	args = append(args, userArgs...)

	row := database.DB.QueryRow(query, args...)

	var item UserTransportReportItem
	var perihal, tujuan sql.NullString
	var tanggalMulai, tanggalAkhir sql.NullTime
	if err := row.Scan(&item.ID, &perihal, &tujuan, &tanggalMulai, &tanggalAkhir); err != nil {
		return nil, err
	}

	if perihal.Valid {
		item.PerihalFP = perihal.String
	}
	if tujuan.Valid {
		item.Tujuan = tujuan.String
	}
	if tanggalMulai.Valid {
		item.TanggalMulai = tanggalMulai.Time.Format("2006-01-02")
	}
	if tanggalAkhir.Valid {
		item.TanggalAkhir = tanggalAkhir.Time.Format("2006-01-02")
	}

	return &item, nil
}

func GetPendingSPPDReportByID(id int, userNIP, userName string) (*UserSPPDReportItem, error) {
	userClause, userArgs := buildUserOwnershipClause(userNIP, userName)
	if userClause == "" {
		return nil, ValidationError{Message: "Identitas user tidak ditemukan"}
	}

	query := `
		SELECT
			m.id_matriks,
			COALESCE(NULLIF(m.perihal_fp, ''), ''),
			COALESCE(NULLIF(m.tujuan, ''), ''),
			m.tglmulai,
			m.tglselesai
		FROM matriks m
		LEFT JOIN laporan_sppd_detail lsd ON m.id_matriks = lsd.id_matriks_fk
		WHERE m.id_matriks = ?
		  AND m.jenis_perjalanan = 'Perjalanan Dinas'
		  AND lsd.id_laporan_sppd IS NULL
		  AND (` + userClause + `)
		LIMIT 1`

	args := make([]interface{}, 0, len(userArgs)+1)
	args = append(args, id)
	args = append(args, userArgs...)

	row := database.DB.QueryRow(query, args...)

	var item UserSPPDReportItem
	var perihal, tujuan sql.NullString
	var tanggalMulai, tanggalAkhir sql.NullTime
	if err := row.Scan(&item.ID, &perihal, &tujuan, &tanggalMulai, &tanggalAkhir); err != nil {
		return nil, err
	}

	if perihal.Valid {
		item.PerihalFP = perihal.String
	}
	if tujuan.Valid {
		item.Tujuan = tujuan.String
	}
	if tanggalMulai.Valid {
		item.TanggalMulai = tanggalMulai.Time.Format("2006-01-02")
	}
	if tanggalAkhir.Valid {
		item.TanggalAkhir = tanggalAkhir.Time.Format("2006-01-02")
	}

	return &item, nil
}

func SaveTransportReport(id int, laporanText, laporanFotoPath string) error {
	laporanText = strings.TrimSpace(laporanText)
	laporanFotoPath = strings.TrimSpace(laporanFotoPath)

	if laporanText == "" {
		return ValidationError{Message: "Isi laporan wajib diisi"}
	}
	if laporanFotoPath == "" {
		return ValidationError{Message: "Foto laporan wajib diunggah"}
	}

	hasCreatedAt, err := tableColumnExists("laporan_perjalanan", "created_at")
	if err != nil {
		return err
	}

	if hasCreatedAt {
		_, err = database.DB.Exec(`
			INSERT INTO laporan_perjalanan (id_matriks_fk, laporan_text, laporan_foto, created_at)
			VALUES (?, ?, ?, NOW())`,
			id, laporanText, laporanFotoPath,
		)
		return err
	}

	_, err = database.DB.Exec(`
		INSERT INTO laporan_perjalanan (id_matriks_fk, laporan_text, laporan_foto)
		VALUES (?, ?, ?)`,
		id, laporanText, laporanFotoPath,
	)
	return err
}

func UserOwnsTransportReport(id int, userNIP, userName string) (bool, error) {
	userClause, userArgs := buildUserOwnershipClause(userNIP, userName)
	if userClause == "" {
		return false, nil
	}

	query := `
		SELECT COUNT(*)
		FROM matriks m
		WHERE m.id_matriks = ?
		  AND m.jenis_perjalanan = 'Transport Lokal'
		  AND (` + userClause + `)`

	args := make([]interface{}, 0, len(userArgs)+1)
	args = append(args, id)
	args = append(args, userArgs...)

	var count int
	if err := database.DB.QueryRow(query, args...).Scan(&count); err != nil {
		return false, err
	}

	return count > 0, nil
}

func SaveSPPDReport(id int, payload SaveSPPDReportPayload) error {
	payload.WaktuBerangkatMulai = strings.TrimSpace(payload.WaktuBerangkatMulai)
	payload.WaktuBerangkatSelesai = strings.TrimSpace(payload.WaktuBerangkatSelesai)
	payload.IsianBerangkat = strings.TrimSpace(payload.IsianBerangkat)
	payload.WaktuPengawasan1Mulai = strings.TrimSpace(payload.WaktuPengawasan1Mulai)
	payload.WaktuPengawasan1Selesai = strings.TrimSpace(payload.WaktuPengawasan1Selesai)
	payload.IsianPengawasan1 = strings.TrimSpace(payload.IsianPengawasan1)
	payload.WaktuIsomaMulai = strings.TrimSpace(payload.WaktuIsomaMulai)
	payload.WaktuIsomaSelesai = strings.TrimSpace(payload.WaktuIsomaSelesai)
	payload.IsianIsoma = strings.TrimSpace(payload.IsianIsoma)
	payload.WaktuPengawasan2Mulai = strings.TrimSpace(payload.WaktuPengawasan2Mulai)
	payload.WaktuPengawasan2Selesai = strings.TrimSpace(payload.WaktuPengawasan2Selesai)
	payload.IsianPengawasan2 = strings.TrimSpace(payload.IsianPengawasan2)
	payload.WaktuPulangMulai = strings.TrimSpace(payload.WaktuPulangMulai)
	payload.WaktuPulangSelesai = strings.TrimSpace(payload.WaktuPulangSelesai)
	payload.IsianPulang = strings.TrimSpace(payload.IsianPulang)
	payload.LaporanFotoPath = strings.TrimSpace(payload.LaporanFotoPath)

	required := []string{
		payload.WaktuBerangkatMulai, payload.WaktuBerangkatSelesai, payload.IsianBerangkat,
		payload.WaktuPengawasan1Mulai, payload.WaktuPengawasan1Selesai, payload.IsianPengawasan1,
		payload.WaktuIsomaMulai, payload.WaktuIsomaSelesai, payload.IsianIsoma,
		payload.WaktuPengawasan2Mulai, payload.WaktuPengawasan2Selesai, payload.IsianPengawasan2,
		payload.WaktuPulangMulai, payload.WaktuPulangSelesai, payload.IsianPulang, payload.LaporanFotoPath,
	}
	for _, value := range required {
		if value == "" {
			return ValidationError{Message: "Semua kolom laporan SPPD wajib diisi"}
		}
	}

	hasCreatedAt, err := tableColumnExists("laporan_sppd_detail", "created_at")
	if err != nil {
		return err
	}

	if hasCreatedAt {
		_, err = database.DB.Exec(`
			INSERT INTO laporan_sppd_detail (
				id_matriks_fk,
				waktu_berangkat_mulai, waktu_berangkat_selesai, isian_berangkat,
				waktu_pengawasan1_mulai, waktu_pengawasan1_selesai, isian_pengawasan1,
				waktu_isoma_mulai, waktu_isoma_selesai, isian_isoma,
				waktu_pengawasan2_mulai, waktu_pengawasan2_selesai, isian_pengawasan2,
				waktu_pulang_mulai, waktu_pulang_selesai, isian_pulang,
				laporan_foto, created_at
			) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW())`,
			id,
			payload.WaktuBerangkatMulai, payload.WaktuBerangkatSelesai, payload.IsianBerangkat,
			payload.WaktuPengawasan1Mulai, payload.WaktuPengawasan1Selesai, payload.IsianPengawasan1,
			payload.WaktuIsomaMulai, payload.WaktuIsomaSelesai, payload.IsianIsoma,
			payload.WaktuPengawasan2Mulai, payload.WaktuPengawasan2Selesai, payload.IsianPengawasan2,
			payload.WaktuPulangMulai, payload.WaktuPulangSelesai, payload.IsianPulang,
			payload.LaporanFotoPath,
		)
		return err
	}

	_, err = database.DB.Exec(`
		INSERT INTO laporan_sppd_detail (
			id_matriks_fk,
			waktu_berangkat_mulai, waktu_berangkat_selesai, isian_berangkat,
			waktu_pengawasan1_mulai, waktu_pengawasan1_selesai, isian_pengawasan1,
			waktu_isoma_mulai, waktu_isoma_selesai, isian_isoma,
			waktu_pengawasan2_mulai, waktu_pengawasan2_selesai, isian_pengawasan2,
			waktu_pulang_mulai, waktu_pulang_selesai, isian_pulang,
			laporan_foto
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		id,
		payload.WaktuBerangkatMulai, payload.WaktuBerangkatSelesai, payload.IsianBerangkat,
		payload.WaktuPengawasan1Mulai, payload.WaktuPengawasan1Selesai, payload.IsianPengawasan1,
		payload.WaktuIsomaMulai, payload.WaktuIsomaSelesai, payload.IsianIsoma,
		payload.WaktuPengawasan2Mulai, payload.WaktuPengawasan2Selesai, payload.IsianPengawasan2,
		payload.WaktuPulangMulai, payload.WaktuPulangSelesai, payload.IsianPulang,
		payload.LaporanFotoPath,
	)
	return err
}

func UserOwnsSPPDReport(id int, userNIP, userName string) (bool, error) {
	userClause, userArgs := buildUserOwnershipClause(userNIP, userName)
	if userClause == "" {
		return false, nil
	}

	query := `
		SELECT COUNT(*)
		FROM matriks m
		WHERE m.id_matriks = ?
		  AND m.jenis_perjalanan = 'Perjalanan Dinas'
		  AND (` + userClause + `)`

	args := make([]interface{}, 0, len(userArgs)+1)
	args = append(args, id)
	args = append(args, userArgs...)

	var count int
	if err := database.DB.QueryRow(query, args...).Scan(&count); err != nil {
		return false, err
	}

	return count > 0, nil
}

func listUserTransportReports(userNIP, userName string, submitted bool) ([]UserTransportReportItem, error) {
	hasReportTable, err := tableExists("laporan_perjalanan")
	if err != nil {
		return nil, err
	}

	if !hasReportTable {
		return []UserTransportReportItem{}, nil
	}

	userClause, userArgs := buildUserOwnershipClause(userNIP, userName)
	if userClause == "" {
		return []UserTransportReportItem{}, nil
	}

	hasCreatedAt, err := tableColumnExists("laporan_perjalanan", "created_at")
	if err != nil {
		return nil, err
	}

	selectCreatedAt := "NULL AS created_at"
	orderBy := "m.id_matriks DESC"
	if hasCreatedAt {
		selectCreatedAt = "lp.created_at"
		if submitted {
			orderBy = "lp.created_at DESC, m.id_matriks DESC"
		}
	}

	reportCondition := "lp.id_laporan IS NULL"
	if submitted {
		reportCondition = "lp.id_laporan IS NOT NULL"
	}

	query := fmt.Sprintf(`
		SELECT
			m.id_matriks,
			COALESCE(NULLIF(m.perihal_fp, ''), ''),
			COALESCE(NULLIF(m.tujuan, ''), ''),
			m.tglmulai,
			m.tglselesai,
			COALESCE(lp.laporan_text, ''),
			COALESCE(lp.laporan_foto, ''),
			%s
		FROM matriks m
		LEFT JOIN laporan_perjalanan lp ON m.id_matriks = lp.id_matriks_fk
		WHERE m.jenis_perjalanan = 'Transport Lokal'
		  AND %s
		  AND (%s)
		ORDER BY %s`, selectCreatedAt, reportCondition, userClause, orderBy)

	rows, err := database.DB.Query(query, userArgs...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]UserTransportReportItem, 0)
	for rows.Next() {
		var item UserTransportReportItem
		var perihal, tujuan, laporanText, laporanFoto sql.NullString
		var tanggalMulai, tanggalAkhir, createdAt sql.NullTime
		if err := rows.Scan(
			&item.ID,
			&perihal,
			&tujuan,
			&tanggalMulai,
			&tanggalAkhir,
			&laporanText,
			&laporanFoto,
			&createdAt,
		); err != nil {
			return nil, err
		}

		if perihal.Valid {
			item.PerihalFP = perihal.String
		}
		if tujuan.Valid {
			item.Tujuan = tujuan.String
		}
		if tanggalMulai.Valid {
			item.TanggalMulai = tanggalMulai.Time.Format("2006-01-02")
		}
		if tanggalAkhir.Valid {
			item.TanggalAkhir = tanggalAkhir.Time.Format("2006-01-02")
		}
		if laporanText.Valid {
			item.LaporanText = laporanText.String
		}
		if laporanFoto.Valid {
			item.LaporanFoto = laporanFoto.String
		}
		if createdAt.Valid {
			item.CreatedAt = createdAt.Time.Format(time.RFC3339)
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func listUserSPPDReports(userNIP, userName string, submitted bool) ([]UserSPPDReportItem, error) {
	hasReportTable, err := tableExists("laporan_sppd_detail")
	if err != nil {
		return nil, err
	}

	if !hasReportTable {
		return []UserSPPDReportItem{}, nil
	}

	userClause, userArgs := buildUserOwnershipClause(userNIP, userName)
	if userClause == "" {
		return []UserSPPDReportItem{}, nil
	}

	hasCreatedAt, err := tableColumnExists("laporan_sppd_detail", "created_at")
	if err != nil {
		return nil, err
	}

	selectCreatedAt := "NULL AS created_at"
	orderBy := "m.id_matriks DESC"
	if hasCreatedAt {
		selectCreatedAt = "lsd.created_at"
		if submitted {
			orderBy = "lsd.created_at DESC, m.id_matriks DESC"
		}
	}

	reportCondition := "lsd.id_laporan_sppd IS NULL"
	if submitted {
		reportCondition = "lsd.id_laporan_sppd IS NOT NULL"
	}

	query := fmt.Sprintf(`
		SELECT
			m.id_matriks,
			COALESCE(NULLIF(m.perihal_fp, ''), ''),
			COALESCE(NULLIF(m.tujuan, ''), ''),
			m.tglmulai,
			m.tglselesai,
			COALESCE(TIME_FORMAT(lsd.waktu_berangkat_mulai, '%%H:%%i'), ''),
			COALESCE(TIME_FORMAT(lsd.waktu_berangkat_selesai, '%%H:%%i'), ''),
			COALESCE(lsd.isian_berangkat, ''),
			COALESCE(TIME_FORMAT(lsd.waktu_pengawasan1_mulai, '%%H:%%i'), ''),
			COALESCE(TIME_FORMAT(lsd.waktu_pengawasan1_selesai, '%%H:%%i'), ''),
			COALESCE(lsd.isian_pengawasan1, ''),
			COALESCE(TIME_FORMAT(lsd.waktu_isoma_mulai, '%%H:%%i'), ''),
			COALESCE(TIME_FORMAT(lsd.waktu_isoma_selesai, '%%H:%%i'), ''),
			COALESCE(lsd.isian_isoma, ''),
			COALESCE(TIME_FORMAT(lsd.waktu_pengawasan2_mulai, '%%H:%%i'), ''),
			COALESCE(TIME_FORMAT(lsd.waktu_pengawasan2_selesai, '%%H:%%i'), ''),
			COALESCE(lsd.isian_pengawasan2, ''),
			COALESCE(TIME_FORMAT(lsd.waktu_pulang_mulai, '%%H:%%i'), ''),
			COALESCE(TIME_FORMAT(lsd.waktu_pulang_selesai, '%%H:%%i'), ''),
			COALESCE(lsd.isian_pulang, ''),
			COALESCE(lsd.laporan_foto, ''),
			%s
		FROM matriks m
		LEFT JOIN laporan_sppd_detail lsd ON m.id_matriks = lsd.id_matriks_fk
		WHERE m.jenis_perjalanan = 'Perjalanan Dinas'
		  AND %s
		  AND (%s)
		ORDER BY %s`, selectCreatedAt, reportCondition, userClause, orderBy)

	rows, err := database.DB.Query(query, userArgs...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]UserSPPDReportItem, 0)
	for rows.Next() {
		var item UserSPPDReportItem
		var perihal, tujuan, foto sql.NullString
		var tanggalMulai, tanggalAkhir, createdAt sql.NullTime
		var wbm, wbs, ib, wp1m, wp1s, ip1, wim, wis, ii, wp2m, wp2s, ip2, wpm, wps, ip sql.NullString
		if err := rows.Scan(
			&item.ID,
			&perihal,
			&tujuan,
			&tanggalMulai,
			&tanggalAkhir,
			&wbm, &wbs, &ib,
			&wp1m, &wp1s, &ip1,
			&wim, &wis, &ii,
			&wp2m, &wp2s, &ip2,
			&wpm, &wps, &ip,
			&foto,
			&createdAt,
		); err != nil {
			return nil, err
		}

		if perihal.Valid {
			item.PerihalFP = perihal.String
		}
		if tujuan.Valid {
			item.Tujuan = tujuan.String
		}
		if tanggalMulai.Valid {
			item.TanggalMulai = tanggalMulai.Time.Format("2006-01-02")
		}
		if tanggalAkhir.Valid {
			item.TanggalAkhir = tanggalAkhir.Time.Format("2006-01-02")
		}
		if wbm.Valid {
			item.WaktuBerangkatMulai = wbm.String
		}
		if wbs.Valid {
			item.WaktuBerangkatSelesai = wbs.String
		}
		if ib.Valid {
			item.IsianBerangkat = ib.String
		}
		if wp1m.Valid {
			item.WaktuPengawasan1Mulai = wp1m.String
		}
		if wp1s.Valid {
			item.WaktuPengawasan1Selesai = wp1s.String
		}
		if ip1.Valid {
			item.IsianPengawasan1 = ip1.String
		}
		if wim.Valid {
			item.WaktuIsomaMulai = wim.String
		}
		if wis.Valid {
			item.WaktuIsomaSelesai = wis.String
		}
		if ii.Valid {
			item.IsianIsoma = ii.String
		}
		if wp2m.Valid {
			item.WaktuPengawasan2Mulai = wp2m.String
		}
		if wp2s.Valid {
			item.WaktuPengawasan2Selesai = wp2s.String
		}
		if ip2.Valid {
			item.IsianPengawasan2 = ip2.String
		}
		if wpm.Valid {
			item.WaktuPulangMulai = wpm.String
		}
		if wps.Valid {
			item.WaktuPulangSelesai = wps.String
		}
		if ip.Valid {
			item.IsianPulang = ip.String
		}
		if foto.Valid {
			item.LaporanFoto = foto.String
		}
		if createdAt.Valid {
			item.CreatedAt = createdAt.Time.Format(time.RFC3339)
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func buildUserOwnershipClause(userNIP, userName string) (string, []interface{}) {
	userNIP = strings.TrimSpace(userNIP)
	userName = strings.TrimSpace(userName)

	// NIP adalah identifier yang lebih andal (unik per pegawai), jadi jadikan
	// satu-satunya syarat saat tersedia. Mewajibkan nama ikut cocok persis
	// (AND) membuat laporan hilang dari listing begitu ada selisih spasi atau
	// kapitalisasi antara data login dan hasil import CSV matriks.
	if userNIP != "" {
		return "REPLACE(COALESCE(m.nip, ''), ' ', '') = ?", []interface{}{strings.ReplaceAll(userNIP, " ", "")}
	}
	if userName != "" {
		return "LOWER(REPLACE(COALESCE(m.nama, ''), ' ', '')) = LOWER(REPLACE(?, ' ', ''))", []interface{}{userName}
	}

	return "", nil
}
