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
// Surat Pernyataan Penyelesaian Lapangan Termin I (PPL & PML)
// ---------------------------------------------------------------------------------

// usahaKeluargaLampiranRow adalah satu baris tabel lampiran: satu petugas (PPL, dipakai
// di lampiran surat pernyataan PML & Kepala) atau agregat sendiri (PML/PPL individu).
type usahaKeluargaLampiranRow struct {
	Nama          string
	Jabatan       string // "Petugas Lapangan" / "Pemeriksa Lapangan" - dipakai lampiran Kepala saja
	TargetPrelist int
	Realisasi     int
}

func persentase(realisasi, target int) float64 {
	if target == 0 {
		return 0
	}
	return float64(realisasi) / float64(target) * 100
}

// computeUsahaKeluargaSE2026 menjumlahkan target prelist (se2026.sls.target) dan realisasi
// hasil pendataan usaha+keluarga (se2026.progress.fasih_total, metode "FASIH total submit"
// sesuai arahan user) untuk satu user se2026 (PPL: by ppl_id, PML: by pml_id).
func computeUsahaKeluargaSE2026(uid int, isPML bool) (target int, realisasi int, err error) {
	col := "s.ppl_id"
	if isPML {
		col = "s.pml_id"
	}
	q := fmt.Sprintf(`
		SELECT COALESCE(SUM(s.target),0), COALESCE(SUM(p.fasih_total),0)
		FROM se2026.sls s LEFT JOIN se2026.progress p ON p.sls_id = s.id
		WHERE %s = ?`, col)
	err = database.DB.QueryRow(q, uid).Scan(&target, &realisasi)
	return
}

// buildSuratPernyataanSE2026 mengambil data petugas (PPL atau PML, dari rekap) dan
// menyusun map placeholder + (khusus PML) daftar PPL binaan untuk lampiran tabel.
func buildSuratPernyataanSE2026(rekapID int, tanggal string) (bappSE2026Data, string, []usahaKeluargaLampiranRow, error) {
	var d bappSE2026Data
	var idSpk, tahun, kegiatan string

	row := database.DB.QueryRow(`
		SELECT r.namamitra, COALESCE(r.id_spk,''), r.tahun, r.kegiatan
		FROM rekap r WHERE r.id = ?`, rekapID)
	if err := row.Scan(&d.NamaPetugas, &idSpk, &tahun, &kegiatan); err != nil {
		return d, "", nil, fmt.Errorf("rekap id %d tidak ditemukan", rekapID)
	}
	d.NamaPetugas = properCase(d.NamaPetugas)
	if idSpk == "" {
		return d, "", nil, fmt.Errorf("petugas ini belum memiliki nomor SPK")
	}
	jenis := jenisPetugasSE2026(kegiatan)

	tgl, err := time.Parse("2006-01-02", tanggal)
	if err != nil {
		tgl, _ = time.Parse("2006-01-02", "2026-08-15")
	}

	kind := "pernyataan_ppl"
	seRole := "ppl"
	if jenis == "pml" {
		kind = "pernyataan_pml"
		seRole = "pml"
	}
	nomor, err := suratNomorSE2026(idSpk, kind, 1, tahun, tgl)
	if err != nil {
		return d, jenis, nil, err
	}
	d.Nomor = nomor
	d.NomorSPK = idSpk

	database.DB.QueryRow(`
		SELECT nip FROM pegawai
		WHERE UPPER(CONVERT(nama USING utf8mb4)) COLLATE utf8mb4_unicode_ci = UPPER(?) COLLATE utf8mb4_unicode_ci
		LIMIT 1`, d.NamaPetugas).Scan(&d.NIKPetugas)

	d.Hari = hariIndonesia[tgl.Weekday().String()]
	d.TglTeks = dayToTeks(tgl.Day())
	d.BlnTeks = bulanTeksSE[tgl.Month()]
	d.TglAngka = fmt.Sprintf("%02d-%02d", tgl.Day(), int(tgl.Month()))

	uid, cnt := findSE2026UserID(d.NamaPetugas, seRole)
	if cnt != 1 {
		return d, jenis, nil, fmt.Errorf("nama '%s' tidak ketemu tunggal di se2026.users (%d kandidat) - jalankan sinkron realisasi dulu", d.NamaPetugas, cnt)
	}

	if jenis == "pcl" {
		// PPL: tidak ada lampiran tabel, cukup angka target/realisasi/persentase dirinya sendiri.
		target, realisasi, err := computeUsahaKeluargaSE2026(uid, false)
		if err != nil {
			return d, jenis, nil, fmt.Errorf("gagal menghitung realisasi: %v", err)
		}
		return d, jenis, []usahaKeluargaLampiranRow{{
			Nama: d.NamaPetugas, TargetPrelist: target, Realisasi: realisasi,
		}}, nil
	}

	// PML: daftar PPL binaan (dari se2026.sls, dikelompokkan per PPL).
	rows, err := database.DB.Query(`
		SELECT u.name, u.id
		FROM se2026.sls s
		JOIN se2026.users u ON u.id = s.ppl_id
		WHERE s.pml_id = ?
		GROUP BY u.id, u.name
		ORDER BY u.name`, uid)
	if err != nil {
		return d, jenis, nil, fmt.Errorf("gagal mengambil daftar PPL binaan: %v", err)
	}
	defer rows.Close()

	var lampiran []usahaKeluargaLampiranRow
	for rows.Next() {
		var nama string
		var pplUID int
		if err := rows.Scan(&nama, &pplUID); err != nil {
			continue
		}
		target, realisasi, err := computeUsahaKeluargaSE2026(pplUID, false)
		if err != nil {
			continue
		}
		lampiran = append(lampiran, usahaKeluargaLampiranRow{
			Nama: properCase(nama), TargetPrelist: target, Realisasi: realisasi,
		})
	}
	return d, jenis, lampiran, nil
}

// docFontRPr/docFontRPrBold adalah run-properties yang menyamakan font tabel lampiran
// dengan sisa dokumen (Bookman Old Style 12pt/sz=24), supaya isian tabel tidak "lompat"
// gaya-nya dibanding teks contoh aslinya.
const (
	docFontRPr     = `<w:rFonts w:ascii="Bookman Old Style" w:hAnsi="Bookman Old Style" w:eastAsia="Bookman Old Style" w:cs="Bookman Old Style"/><w:sz w:val="24"/><w:szCs w:val="24"/>`
	docFontRPrBold = `<w:rFonts w:ascii="Bookman Old Style" w:hAnsi="Bookman Old Style" w:eastAsia="Bookman Old Style" w:cs="Bookman Old Style"/><w:b/><w:sz w:val="24"/><w:szCs w:val="24"/>`
	// tcBordersXML: garis tunggal di semua sisi sel, sama seperti tabel contoh aslinya.
	tcBordersXML = `<w:tcBorders><w:top w:val="single" w:sz="4" w:space="0" w:color="000000"/><w:start w:val="single" w:sz="4" w:space="0" w:color="000000"/><w:bottom w:val="single" w:sz="4" w:space="0" w:color="000000"/><w:end w:val="single" w:sz="4" w:space="0" w:color="000000"/></w:tcBorders>`
)

// tcOpen membuka sebuah <w:tc> dengan lebar tetap (dxa) sesuai kolom tabel contoh asli,
// supaya tabel dinamis yang di-generate persis selebar & serapi tabel contoh di template.
func tcOpen(widthDxa int) string {
	return fmt.Sprintf(`<w:tc><w:tcPr><w:tcW w:w="%d" w:type="dxa"/>`+tcBordersXML+`</w:tcPr>`, widthDxa)
}

// pmlColW adalah lebar kolom (twips/dxa) persis seperti tabel contoh di
// surat_pernyataan_pml_se2026.docx asli (930/3630/1634/2326/2100, total 10620 dxa).
var pmlColW = [5]int{930, 3630, 1634, 2326, 2100}

// generateLampiranTableXML membangun <w:tbl> "No | Nama Petugas Lapangan | Target Prelist |
// Realisasi Hasil Pendataan (Usaha+Keluarga) | Presentase (%)" + baris "Jumlah" di akhir -
// lebar kolom & tblLayout=fixed disamakan dengan tabel contoh di template aslinya, dan
// tabel ditengahkan (jc=center) di halaman.
func generateLampiranTableXML(rows []usahaKeluargaLampiranRow) string {
	c := pmlColW
	var rowXML strings.Builder
	totalTarget, totalRealisasi := 0, 0
	for i, r := range rows {
		totalTarget += r.TargetPrelist
		totalRealisasi += r.Realisasi
		rowXML.WriteString(fmt.Sprintf(`
<w:tr>
%s<w:p><w:pPr><w:jc w:val="center"/></w:pPr><w:r><w:rPr>`+docFontRPr+`</w:rPr><w:t>%d</w:t></w:r></w:p></w:tc>
%s<w:p><w:r><w:rPr>`+docFontRPr+`</w:rPr><w:t>%s</w:t></w:r></w:p></w:tc>
%s<w:p><w:pPr><w:jc w:val="center"/></w:pPr><w:r><w:rPr>`+docFontRPr+`</w:rPr><w:t>%d</w:t></w:r></w:p></w:tc>
%s<w:p><w:pPr><w:jc w:val="center"/></w:pPr><w:r><w:rPr>`+docFontRPr+`</w:rPr><w:t>%d</w:t></w:r></w:p></w:tc>
%s<w:p><w:pPr><w:jc w:val="center"/></w:pPr><w:r><w:rPr>`+docFontRPr+`</w:rPr><w:t>%.1f</w:t></w:r></w:p></w:tc>
</w:tr>`,
			tcOpen(c[0]), i+1, tcOpen(c[1]), escapeXML(r.Nama), tcOpen(c[2]), r.TargetPrelist,
			tcOpen(c[3]), r.Realisasi, tcOpen(c[4]), persentase(r.Realisasi, r.TargetPrelist)))
	}
	rowXML.WriteString(fmt.Sprintf(`
<w:tr>
<w:tc><w:tcPr><w:tcW w:w="%d" w:type="dxa"/><w:gridSpan w:val="2"/>`+tcBordersXML+`</w:tcPr><w:p><w:r><w:rPr>`+docFontRPrBold+`</w:rPr><w:t>Jumlah</w:t></w:r></w:p></w:tc>
%s<w:p><w:pPr><w:jc w:val="center"/></w:pPr><w:r><w:rPr>`+docFontRPrBold+`</w:rPr><w:t>%d</w:t></w:r></w:p></w:tc>
%s<w:p><w:pPr><w:jc w:val="center"/></w:pPr><w:r><w:rPr>`+docFontRPrBold+`</w:rPr><w:t>%d</w:t></w:r></w:p></w:tc>
%s<w:p><w:pPr><w:jc w:val="center"/></w:pPr><w:r><w:rPr>`+docFontRPrBold+`</w:rPr><w:t>%.1f</w:t></w:r></w:p></w:tc>
</w:tr>`,
		c[0]+c[1], tcOpen(c[2]), totalTarget, tcOpen(c[3]), totalRealisasi, tcOpen(c[4]), persentase(totalRealisasi, totalTarget)))

	return fmt.Sprintf(`<w:tbl>
<w:tblPr>
<w:tblStyle w:val="TableGrid"/>
<w:tblW w:w="%d" w:type="dxa"/>
<w:jc w:val="center"/>
<w:tblLayout w:type="fixed"/>
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
<w:gridCol w:w="%d"/>
<w:gridCol w:w="%d"/>
<w:gridCol w:w="%d"/>
<w:gridCol w:w="%d"/>
<w:gridCol w:w="%d"/>
</w:tblGrid>
<w:tr>
%s<w:p><w:pPr><w:jc w:val="center"/></w:pPr><w:r><w:rPr>%s</w:rPr><w:t>No</w:t></w:r></w:p></w:tc>
%s<w:p><w:pPr><w:jc w:val="center"/></w:pPr><w:r><w:rPr>%s</w:rPr><w:t>Nama Petugas Lapangan</w:t></w:r></w:p></w:tc>
%s<w:p><w:pPr><w:jc w:val="center"/></w:pPr><w:r><w:rPr>%s</w:rPr><w:t>Target Prelist</w:t></w:r></w:p></w:tc>
%s<w:p><w:pPr><w:jc w:val="center"/></w:pPr><w:r><w:rPr>%s</w:rPr><w:t>Realisasi Hasil Pendataan (Usaha+Keluarga)</w:t></w:r></w:p></w:tc>
%s<w:p><w:pPr><w:jc w:val="center"/></w:pPr><w:r><w:rPr>%s</w:rPr><w:t>Presentase (%%)</w:t></w:r></w:p></w:tc>
</w:tr>`,
		c[0]+c[1]+c[2]+c[3]+c[4], c[0], c[1], c[2], c[3], c[4],
		tcOpen(c[0]), docFontRPrBold, tcOpen(c[1]), docFontRPrBold, tcOpen(c[2]), docFontRPrBold,
		tcOpen(c[3]), docFontRPrBold, tcOpen(c[4]), docFontRPrBold,
	) + rowXML.String() + `</w:tbl>`
}

func generateSuratPernyataanDocx(d bappSE2026Data, jenis string, lampiran []usahaKeluargaLampiranRow) ([]byte, error) {
	if jenis == "pml" {
		return replaceInDocx("static/templates-se2026/surat_pernyataan_pml_se2026.docx", map[string]string{
			"nomor":          d.Nomor,
			"nomor_spk":      d.NomorSPK,
			"tgl_teks":       d.TglTeks,
			"bln_teks":       d.BlnTeks,
			"nama_petugas":   d.NamaPetugas,
			"nik_petugas":    d.NIKPetugas,
			"nama_ketua_tim": ketuaTimNamaSE2026,
			"nip_ketua_tim":  ketuaTimNIPSE2026,
			"table":          generateLampiranTableXML(lampiran),
		})
	}
	// PPL: tanpa tabel, angka langsung di badan surat (lampiran hanya 1 baris berisi diri sendiri).
	target, realisasi := 0, 0
	if len(lampiran) == 1 {
		target, realisasi = lampiran[0].TargetPrelist, lampiran[0].Realisasi
	}
	return replaceInDocx("static/templates-se2026/surat_pernyataan_ppl_se2026.docx", map[string]string{
		"nomor":               d.Nomor,
		"nomor_spk":           d.NomorSPK,
		"tgl_teks":            d.TglTeks,
		"bln_teks":            d.BlnTeks,
		"nama_petugas":        d.NamaPetugas,
		"nik_petugas":         d.NIKPetugas,
		"nama_ketua_tim":      ketuaTimNamaSE2026,
		"nip_ketua_tim":       ketuaTimNIPSE2026,
		"target_prelist":      strconv.Itoa(target),
		"realisasi_usaha_kel": strconv.Itoa(realisasi),
		"persentase":          fmt.Sprintf("%.1f", persentase(realisasi, target)),
	})
}

// assignPernyataanNumber sama seperti assignBAPPNumber tapi untuk Surat Pernyataan
// Penyelesaian Lapangan (PPL & PML sama-sama pakai kolom id_pernyataan1/tgl_pernyataan1 -
// aman karena baris rekap PPL dan PML untuk 1 orang tidak pernah sama, kegiatan-nya beda).
func assignPernyataanNumber(rekapID int, tanggal string) error {
	var idSpk, tahun, kegiatan string
	var existing sql.NullString
	err := database.DB.QueryRow(`SELECT COALESCE(id_spk,''), tahun, kegiatan, id_pernyataan1 FROM rekap WHERE id=?`, rekapID).
		Scan(&idSpk, &tahun, &kegiatan, &existing)
	if err != nil {
		return fmt.Errorf("rekap id %d tidak ditemukan", rekapID)
	}
	if existing.Valid && existing.String != "" {
		return nil
	}
	if idSpk == "" {
		return fmt.Errorf("petugas belum memiliki nomor SPK")
	}
	tgl, err := time.Parse("2006-01-02", tanggal)
	if err != nil {
		return fmt.Errorf("format tanggal salah, gunakan YYYY-MM-DD")
	}
	kind := "pernyataan_ppl"
	if jenisPetugasSE2026(kegiatan) == "pml" {
		kind = "pernyataan_pml"
	}
	nomor, err := suratNomorSE2026(idSpk, kind, 1, tahun, tgl)
	if err != nil {
		return err
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

	d, jenis, lampiran, err := buildSuratPernyataanSE2026(rekapID, tanggal)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if d.Nomor == "" {
		http.Error(w, "Petugas ini belum memiliki nomor Surat Pernyataan - buat nomornya dulu", http.StatusBadRequest)
		return
	}

	docx, err := generateSuratPernyataanDocx(d, jenis, lampiran)
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
