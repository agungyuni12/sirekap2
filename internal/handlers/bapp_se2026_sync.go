package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"sirekap/internal/database"
)

// progresExprSQL mengembalikan fragmen SQL untuk persentase progres per-SLS, metode-aware,
// persis 3 rumus yang dipakai dashboard sigempar (monitoringse263/handlers/admin.go
// computePctProgres & progresSortExprGeneric). metode divalidasi terhadap set tetap
// (bukan diinterpolasi langsung) supaya aman dari SQL injection.
func progresExprSQL(metode string) string {
	switch metode {
	case "2": // Total/Prelist
		return `(CASE WHEN COALESCE(s.target_prelist_resmi,0)=0 THEN 0 ELSE LEAST(COALESCE(p.jumlah_submit,0)/s.target_prelist_resmi,1) END)`
	case "3": // Prelist/Prelist (residual)
		return `(CASE WHEN COALESCE(s.target_prelist_resmi,0)=0 THEN 0 ELSE GREATEST(LEAST((s.target_prelist_resmi-GREATEST(COALESCE(p.fasih_total,0)-COALESCE(p.jumlah_submit,0),0))/s.target_prelist_resmi,1),0) END)`
	default: // "1" atau kosong: Total/Total
		return `(CASE WHEN COALESCE(p.fasih_total,0)=0 THEN 0 ELSE LEAST(COALESCE(p.jumlah_submit,0)/p.fasih_total,1) END)`
	}
}

const slsSelesaiThresholdPPL = 0.95 // sama seperti slsSelesaiThreshold di monitoringse263/handlers/admin.go
const verifikasiThresholdPML = 0.40 // ambang verifikasi PML sesuai arahan user (belum ada aturan ter-kode di sigempar)

type syncRealisasiResult struct {
	UpdatedPPL int      `json:"updated_ppl"`
	UpdatedPML int      `json:"updated_pml"`
	Unmatched  []string `json:"unmatched"`
}

// findSE2026UserID mencari se2026.users by nama (case/collation-safe) untuk role tertentu.
// Mengembalikan (id, jumlahKandidat). Kalau jumlahKandidat != 1, id tidak boleh dipakai
// (nama tidak ketemu atau ambigu/duplikat).
func findSE2026UserID(nama, role string) (int, int) {
	var cnt int
	database.DB.QueryRow(`
		SELECT COUNT(*) FROM se2026.users
		WHERE role = ? AND UPPER(CONVERT(name USING utf8mb4)) COLLATE utf8mb4_unicode_ci = UPPER(?) COLLATE utf8mb4_unicode_ci`,
		role, nama).Scan(&cnt)
	if cnt != 1 {
		return 0, cnt
	}
	var id int
	database.DB.QueryRow(`
		SELECT id FROM se2026.users
		WHERE role = ? AND UPPER(CONVERT(name USING utf8mb4)) COLLATE utf8mb4_unicode_ci = UPPER(?) COLLATE utf8mb4_unicode_ci`,
		role, nama).Scan(&id)
	return id, cnt
}

// SyncRealisasiSE2026Handler POST /api/rekap/spk/se2026/sync-realisasi?metode=1|2|3
// Menghitung ulang jumlah_sls (target) & realisasi_sls untuk semua PPL & PML SE2026,
// dari data se2026.sls + se2026.progress (bukan input manual).
func SyncRealisasiSE2026Handler(w http.ResponseWriter, r *http.Request) {
	metode := r.URL.Query().Get("metode")
	if metode != "1" && metode != "2" && metode != "3" {
		metode = "1"
	}
	pctExpr := progresExprSQL(metode)

	result := syncRealisasiResult{}

	type rekapRow struct {
		id   int
		nama string
	}

	// ---- PPL: threshold 95% (metode-aware) ----
	pplRows, err := database.DB.Query(`
		SELECT r.id, r.namamitra FROM rekap r
		WHERE r.kegiatan = 'Pendataan Sensus Ekonomi 2026' AND (r.tanggaran = 2026 OR r.tahun = '2026')`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var pplList []rekapRow
	for pplRows.Next() {
		var rr rekapRow
		if err := pplRows.Scan(&rr.id, &rr.nama); err == nil {
			pplList = append(pplList, rr)
		}
	}
	pplRows.Close()

	for _, item := range pplList {
		uid, cnt := findSE2026UserID(item.nama, "ppl")
		if cnt != 1 {
			result.Unmatched = append(result.Unmatched, fmt.Sprintf("PPL: %s (%d kandidat di se2026.users)", item.nama, cnt))
			continue
		}
		var target int
		database.DB.QueryRow(`SELECT COUNT(*) FROM se2026.sls WHERE ppl_id = ?`, uid).Scan(&target)

		var realisasi int
		q := fmt.Sprintf(`
			SELECT COUNT(*) FROM se2026.sls s LEFT JOIN se2026.progress p ON p.sls_id = s.id
			WHERE s.ppl_id = ? AND %s >= %v`, pctExpr, slsSelesaiThresholdPPL)
		database.DB.QueryRow(q, uid).Scan(&realisasi)

		if _, err := database.DB.Exec(`UPDATE rekap SET jumlah_sls = ?, realisasi_sls = ? WHERE id = ?`,
			target, realisasi, item.id); err == nil {
			result.UpdatedPPL++
		}
	}

	// ---- PML: threshold 40% rasio verifikasi (fasih_approved_pengawas/jumlah_submit) ----
	pmlRows, err := database.DB.Query(`
		SELECT r.id, r.namamitra FROM rekap r
		WHERE r.kegiatan LIKE '%Pemeriksa Lapangan Sensus Ekonomi%' AND (r.tanggaran = 2026 OR r.tahun = '2026')`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var pmlList []rekapRow
	for pmlRows.Next() {
		var rr rekapRow
		if err := pmlRows.Scan(&rr.id, &rr.nama); err == nil {
			pmlList = append(pmlList, rr)
		}
	}
	pmlRows.Close()

	for _, item := range pmlList {
		uid, cnt := findSE2026UserID(item.nama, "pml")
		if cnt != 1 {
			result.Unmatched = append(result.Unmatched, fmt.Sprintf("PML: %s (%d kandidat di se2026.users)", item.nama, cnt))
			continue
		}
		var target int
		database.DB.QueryRow(`SELECT COUNT(*) FROM se2026.sls WHERE pml_id = ?`, uid).Scan(&target)

		var realisasi int
		database.DB.QueryRow(`
			SELECT COUNT(*) FROM se2026.sls s LEFT JOIN se2026.progress p ON p.sls_id = s.id
			WHERE s.pml_id = ? AND COALESCE(p.jumlah_submit,0) > 0
			  AND (COALESCE(p.fasih_approved_pengawas,0)/p.jumlah_submit) >= ?`,
			uid, verifikasiThresholdPML).Scan(&realisasi)

		if _, err := database.DB.Exec(`UPDATE rekap SET jumlah_sls = ?, realisasi_sls = ? WHERE id = ?`,
			target, realisasi, item.id); err == nil {
			result.UpdatedPML++
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// ---------------------------------------------------------------------------------
// Surat Pernyataan Penyelesaian Lapangan Termin I (dari PML)
// ---------------------------------------------------------------------------------

type pplLampiranRow struct {
	Nama      string
	Kecamatan string
	Desa      string
	JumlahSLS int
	Realisasi int
}

// buildSuratPernyataanPML mengambil data PML (dari rekap) + daftar PPL binaannya
// (dari se2026.sls, dikelompokkan per PPL) untuk lampiran tabel.
func buildSuratPernyataanPML(rekapID int, tanggal string) (bappSE2026Data, []pplLampiranRow, error) {
	var d bappSE2026Data
	var idSpk, tahun string

	row := database.DB.QueryRow(`
		SELECT r.namamitra, COALESCE(r.id_spk,''), r.tahun
		FROM rekap r WHERE r.id = ? AND r.kegiatan LIKE '%Pemeriksa Lapangan Sensus Ekonomi%'`, rekapID)
	if err := row.Scan(&d.NamaPetugas, &idSpk, &tahun); err != nil {
		return d, nil, fmt.Errorf("rekap PML id %d tidak ditemukan", rekapID)
	}
	d.NamaPetugas = properCase(d.NamaPetugas)
	if idSpk == "" {
		return d, nil, fmt.Errorf("petugas ini belum memiliki nomor SPK")
	}

	nomor, err := bappNomorFromSPK(idSpk, "pernyataan", 1, tahun)
	if err != nil {
		return d, nil, err
	}
	d.Nomor = nomor
	d.NomorSPK = idSpk

	database.DB.QueryRow(`
		SELECT nip FROM pegawai
		WHERE UPPER(CONVERT(nama USING utf8mb4)) COLLATE utf8mb4_unicode_ci = UPPER(?) COLLATE utf8mb4_unicode_ci
		LIMIT 1`, d.NamaPetugas).Scan(&d.NIKPetugas)

	tgl, err := time.Parse("2006-01-02", tanggal)
	if err != nil {
		tgl, _ = time.Parse("2006-01-02", "2026-08-15")
	}
	d.Hari = hariIndonesia[tgl.Weekday().String()]
	d.TglTeks = dayToTeks(tgl.Day())
	d.BlnTeks = bulanTeksSE[tgl.Month()]
	d.TglAngka = fmt.Sprintf("%02d-%02d", tgl.Day(), int(tgl.Month()))

	uid, cnt := findSE2026UserID(d.NamaPetugas, "pml")
	if cnt != 1 {
		return d, nil, fmt.Errorf("nama PML '%s' tidak ketemu tunggal di se2026.users (%d kandidat) - jalankan sinkron realisasi dulu", d.NamaPetugas, cnt)
	}

	rows, err := database.DB.Query(`
		SELECT u.name, COALESCE(s.nama_kec,''), COALESCE(s.nama_desa,''), COUNT(*) AS jml_sls,
		       SUM(CASE WHEN COALESCE(p.jumlah_submit,0) >= s.target THEN 1 ELSE 0 END) AS realisasi_sls
		FROM se2026.sls s
		JOIN se2026.users u ON u.id = s.ppl_id
		LEFT JOIN se2026.progress p ON p.sls_id = s.id
		WHERE s.pml_id = ?
		GROUP BY u.id, u.name, s.nama_kec, s.nama_desa
		ORDER BY u.name`, uid)
	if err != nil {
		return d, nil, fmt.Errorf("gagal mengambil daftar PPL binaan: %v", err)
	}
	defer rows.Close()

	var lampiran []pplLampiranRow
	for rows.Next() {
		var lr pplLampiranRow
		if err := rows.Scan(&lr.Nama, &lr.Kecamatan, &lr.Desa, &lr.JumlahSLS, &lr.Realisasi); err == nil {
			lr.Nama = properCase(lr.Nama)
			lampiran = append(lampiran, lr)
		}
	}
	return d, lampiran, nil
}

// generateLampiranTableXML membangun <w:tbl> untuk daftar PPL binaan PML.
func generateLampiranTableXML(rows []pplLampiranRow) string {
	var rowXML strings.Builder
	for i, r := range rows {
		rowXML.WriteString(fmt.Sprintf(`
<w:tr>
<w:tc><w:p><w:pPr><w:jc w:val="center"/></w:pPr><w:r><w:t>%d</w:t></w:r></w:p></w:tc>
<w:tc><w:p><w:r><w:t>%s</w:t></w:r></w:p></w:tc>
<w:tc><w:p><w:pPr><w:jc w:val="center"/></w:pPr><w:r><w:t>%s</w:t></w:r></w:p></w:tc>
<w:tc><w:p><w:pPr><w:jc w:val="center"/></w:pPr><w:r><w:t>%s</w:t></w:r></w:p></w:tc>
<w:tc><w:p><w:pPr><w:jc w:val="center"/></w:pPr><w:r><w:t>%d</w:t></w:r></w:p></w:tc>
<w:tc><w:p><w:pPr><w:jc w:val="center"/></w:pPr><w:r><w:t>%d</w:t></w:r></w:p></w:tc>
</w:tr>`, i+1, escapeXML(r.Nama), escapeXML(r.Kecamatan), escapeXML(r.Desa), r.JumlahSLS, r.Realisasi))
	}

	return `<w:tbl>
<w:tblPr>
<w:tblStyle w:val="TableGrid"/>
<w:tblW w:w="0" w:type="auto"/>
<w:tblBorders>
<w:top w:val="single" w:sz="4" w:space="0" w:color="000000"/>
<w:left w:val="single" w:sz="4" w:space="0" w:color="000000"/>
<w:bottom w:val="single" w:sz="4" w:space="0" w:color="000000"/>
<w:right w:val="single" w:sz="4" w:space="0" w:color="000000"/>
<w:insideH w:val="single" w:sz="4" w:space="0" w:color="000000"/>
<w:insideV w:val="single" w:sz="4" w:space="0" w:color="000000"/>
</w:tblBorders>
</w:tblPr>
<w:tblGrid>
<w:gridCol w:w="700"/>
<w:gridCol w:w="2400"/>
<w:gridCol w:w="1600"/>
<w:gridCol w:w="1600"/>
<w:gridCol w:w="1400"/>
<w:gridCol w:w="1400"/>
</w:tblGrid>
<w:tr>
<w:tc><w:p><w:r><w:rPr><w:b/></w:rPr><w:t>No</w:t></w:r></w:p></w:tc>
<w:tc><w:p><w:r><w:rPr><w:b/></w:rPr><w:t>Nama Petugas Lapangan Sensus</w:t></w:r></w:p></w:tc>
<w:tc><w:p><w:r><w:rPr><w:b/></w:rPr><w:t>Kecamatan</w:t></w:r></w:p></w:tc>
<w:tc><w:p><w:r><w:rPr><w:b/></w:rPr><w:t>Desa</w:t></w:r></w:p></w:tc>
<w:tc><w:p><w:r><w:rPr><w:b/></w:rPr><w:t>Jumlah SLS/Sub-SLS</w:t></w:r></w:p></w:tc>
<w:tc><w:p><w:r><w:rPr><w:b/></w:rPr><w:t>Realisasi</w:t></w:r></w:p></w:tc>
</w:tr>` + rowXML.String() + `</w:tbl>`
}

func generateSuratPernyataanDocx(d bappSE2026Data, lampiran []pplLampiranRow) ([]byte, error) {
	return replaceInDocx("static/templates-se2026/surat_pernyataan_pml_se2026.docx", map[string]string{
		"nomor":          d.Nomor,
		"nomor_spk":      d.NomorSPK,
		"nama_petugas":   d.NamaPetugas,
		"nik_petugas":    d.NIKPetugas,
		"nama_ketua_tim": ketuaTimNamaSE2026,
		"nip_ketua_tim":  ketuaTimNIPSE2026,
		"table":          generateLampiranTableXML(lampiran),
	})
}

// assignPernyataanNumber sama seperti assignBAPPNumber tapi untuk Surat Pernyataan PML.
func assignPernyataanNumber(rekapID int, tanggal string) error {
	var idSpk, tahun string
	var existing sql.NullString
	err := database.DB.QueryRow(`SELECT COALESCE(id_spk,''), tahun, id_pernyataan1 FROM rekap WHERE id=?`, rekapID).
		Scan(&idSpk, &tahun, &existing)
	if err != nil {
		return fmt.Errorf("rekap id %d tidak ditemukan", rekapID)
	}
	if existing.Valid && existing.String != "" {
		return nil
	}
	if idSpk == "" {
		return fmt.Errorf("petugas belum memiliki nomor SPK")
	}
	nomor, err := bappNomorFromSPK(idSpk, "pernyataan", 1, tahun)
	if err != nil {
		return err
	}
	tgl, err := time.Parse("2006-01-02", tanggal)
	if err != nil {
		return fmt.Errorf("format tanggal salah, gunakan YYYY-MM-DD")
	}
	_, err = database.DB.Exec(`UPDATE rekap SET id_pernyataan1=?, tgl_pernyataan1=? WHERE id=?`,
		nomor, tgl.Format("2006-01-02"), rekapID)
	return err
}

// CreatePernyataanHandler POST /api/rekap/spk/se2026/pernyataan/create
func CreatePernyataanHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		RekapID int    `json:"rekap_id"`
		Tanggal string `json:"tanggal"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	if err := assignPernyataanNumber(req.RekapID, req.Tanggal); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

// DownloadPernyataanSE2026Handler GET /api/rekap/spk/se2026/pernyataan/download?rekap_id=X&tanggal=2026-08-15
func DownloadPernyataanSE2026Handler(w http.ResponseWriter, r *http.Request) {
	rekapID, _ := strconv.Atoi(r.URL.Query().Get("rekap_id"))
	if rekapID == 0 {
		http.Error(w, "rekap_id wajib", http.StatusBadRequest)
		return
	}
	tanggal := r.URL.Query().Get("tanggal")
	if tanggal == "" {
		tanggal = "2026-08-15"
	}

	d, lampiran, err := buildSuratPernyataanPML(rekapID, tanggal)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if d.Nomor == "" {
		http.Error(w, "Petugas ini belum memiliki nomor Surat Pernyataan - buat nomornya dulu", http.StatusBadRequest)
		return
	}

	docx, err := generateSuratPernyataanDocx(d, lampiran)
	if err != nil {
		http.Error(w, "Gagal generate Surat Pernyataan: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fname := fmt.Sprintf("Surat_Pernyataan_SE2026_%s.docx", safeName(d.NamaPetugas))
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, fname))
	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.wordprocessingml.document")
	w.Header().Set("Content-Length", strconv.Itoa(len(docx)))
	w.Write(docx)
}
