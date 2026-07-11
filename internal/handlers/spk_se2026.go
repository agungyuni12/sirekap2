package handlers

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"sirekap/internal/database"
)

// spkSE2026Data holds data to fill placeholders in the SE2026 SPK template.
type spkSE2026Data struct {
	Nomor           string
	Hari            string
	TglTeks         string
	BlnTeks         string
	TglAngka        string
	NamaPetugas     string
	AlamatPetugas   string
	HonorAngka      string
	HonorTerbilang  string
	SanksiAngka     string
	SanksiTerbilang string
}

var bulanTeksSE = map[time.Month]string{
	time.January: "Januari", time.February: "Februari", time.March: "Maret",
	time.April: "April", time.May: "Mei", time.June: "Juni",
	time.July: "Juli", time.August: "Agustus", time.September: "September",
	time.October: "Oktober", time.November: "November", time.December: "Desember",
}

var tglTeksSE = []string{
	"", "satu", "dua", "tiga", "empat", "lima", "enam", "tujuh", "delapan", "sembilan",
	"sepuluh", "sebelas", "dua belas", "tiga belas", "empat belas", "lima belas",
	"enam belas", "tujuh belas", "delapan belas", "sembilan belas", "dua puluh",
	"dua puluh satu", "dua puluh dua", "dua puluh tiga", "dua puluh empat",
	"dua puluh lima", "dua puluh enam", "dua puluh tujuh", "dua puluh delapan",
	"dua puluh sembilan", "tiga puluh", "tiga puluh satu",
}

func dayToTeks(day int) string {
	if day >= 1 && day < len(tglTeksSE) {
		return tglTeksSE[day]
	}
	return strconv.Itoa(day)
}

func formatAngkaSE(n int64) string {
	s := strconv.FormatInt(n, 10)
	var result strings.Builder
	for i, c := range s {
		if i > 0 && (len(s)-i)%3 == 0 {
			result.WriteByte('.')
		}
		result.WriteRune(c)
	}
	return result.String()
}

// properCase mengubah nama kapital semua menjadi Proper Case (tiap kata kapital awal).
func properCase(s string) string {
	words := strings.Fields(strings.ToLower(s))
	for i, w := range words {
		if len(w) > 0 {
			words[i] = strings.ToUpper(string([]rune(w)[:1])) + string([]rune(w)[1:])
		}
	}
	return strings.Join(words, " ")
}

// TGR default per jenis (total biaya pelatihan = sanksi Pasal 11)
// PCL: 2.332.433 | PML: 2.335.840
const (
	tgrDefaultPCL int64 = 2332433
	tgrDefaultPML int64 = 2335840
)

// buildSPKSE2026 fetches rekap+mitra data and builds the template replacement map.
func buildSPKSE2026(rekapID int, tanggalSPK string) (spkSE2026Data, string, error) {
	var d spkSE2026Data

	row := database.DB.QueryRow(`
		SELECT r.namamitra, r.honor, r.id_spk, r.kegiatan,
		       COALESCE((SELECT m.alamat    FROM mitra m WHERE m.idsobat = r.idsobat LIMIT 1), ''),
		       COALESCE((SELECT m.kecamatan FROM mitra m WHERE m.idsobat = r.idsobat LIMIT 1), '')
		FROM rekap r
		WHERE r.id = ?`, rekapID)

	var honorStr, idSpk, kegiatan, alamat, kecamatan string
	if err := row.Scan(&d.NamaPetugas, &honorStr, &idSpk, &kegiatan, &alamat, &kecamatan); err != nil {
		return d, "", fmt.Errorf("rekap id %d tidak ditemukan", rekapID)
	}

	// Proper case nama
	d.NamaPetugas = properCase(d.NamaPetugas)
	d.Nomor = idSpk

	// Tanggal
	tgl, err := time.Parse("2006-01-02", tanggalSPK)
	if err != nil {
		tgl, _ = time.Parse("2006-01-02", "2026-05-31")
	}
	d.Hari = hariIndonesia[tgl.Weekday().String()]
	d.TglTeks = dayToTeks(tgl.Day())
	d.BlnTeks = bulanTeksSE[tgl.Month()]
	d.TglAngka = fmt.Sprintf("%02d-%02d", tgl.Day(), int(tgl.Month()))

	// Alamat: gabung alamat + kecamatan
	bagian := []string{}
	if alamat != "" {
		bagian = append(bagian, alamat)
	}
	if kecamatan != "" {
		bagian = append(bagian, "Kecamatan "+kecamatan)
	}
	if len(bagian) == 0 {
		bagian = append(bagian, "Kabupaten Dompu")
	}
	d.AlamatPetugas = strings.Join(bagian, ", ")

	// Honor
	honorVal, _ := strconv.ParseFloat(honorStr, 64)
	honorInt := int64(honorVal)
	d.HonorAngka = formatAngkaSE(honorInt)
	d.HonorTerbilang = strings.Title(Terbilang(honorInt))

	// Jenis template
	jenis := "pcl"
	sanksi := tgrDefaultPCL
	if strings.Contains(strings.ToLower(kegiatan), "pemeriksa") {
		jenis = "pml"
		sanksi = tgrDefaultPML
	}

	// Sanksi Pasal 11 = TGR (biaya pelatihan per orang)
	d.SanksiAngka = formatAngkaSE(sanksi)
	d.SanksiTerbilang = strings.Title(Terbilang(sanksi))

	return d, jenis, nil
}

func generateSPKDocx(d spkSE2026Data, jenis string) ([]byte, error) {
	tpl := "static/templates-se2026/spk_pcl_se2026.docx"
	if jenis == "pml" {
		tpl = "static/templates-se2026/spk_pml_se2026.docx"
	}
	return replaceInDocx(tpl, map[string]string{
		"nomor":            d.Nomor,
		"hari":             d.Hari,
		"tgl_teks":         d.TglTeks,
		"bln_teks":         d.BlnTeks,
		"tgl_angka":        d.TglAngka,
		"nama_petugas":     d.NamaPetugas,
		"alamat_petugas":   d.AlamatPetugas,
		"honor_angka":      d.HonorAngka,
		"honor_terbilang":  d.HonorTerbilang,
		"sanksi_angka":     d.SanksiAngka,
		"sanksi_terbilang": d.SanksiTerbilang,
	})
}

func safeName(s string) string {
	return strings.Map(func(r rune) rune {
		if r == ' ' {
			return '_'
		}
		if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '_' || r == '-' {
			return r
		}
		return -1
	}, s)
}

// DownloadSPKSE2026Handler GET /api/rekap/spk/se2026/download?rekap_id=X&tanggal=2026-05-31
// Download SPK SE2026 untuk satu petugas yang sudah punya nomor surat.
func DownloadSPKSE2026Handler(w http.ResponseWriter, r *http.Request) {
	rekapID, _ := strconv.Atoi(r.URL.Query().Get("rekap_id"))
	if rekapID == 0 {
		http.Error(w, "rekap_id wajib", http.StatusBadRequest)
		return
	}
	tanggal := r.URL.Query().Get("tanggal")
	if tanggal == "" {
		tanggal = "2026-05-31"
	}

	d, jenis, err := buildSPKSE2026(rekapID, tanggal)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if d.Nomor == "" {
		http.Error(w, "Petugas ini belum memiliki nomor surat SPK", http.StatusBadRequest)
		return
	}

	docx, err := generateSPKDocx(d, jenis)
	if err != nil {
		http.Error(w, "Gagal generate SPK: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fname := fmt.Sprintf("SPK_SE2026_%s.docx", safeName(d.NamaPetugas))
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, fname))
	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.wordprocessingml.document")
	w.Header().Set("Content-Length", strconv.Itoa(len(docx)))
	w.Write(docx)
}

// DownloadAllSPKSE2026Handler GET /api/rekap/spk/se2026/download-all?jenis=pcl&tanggal=2026-05-31
// Download semua SPK SE2026 yang sudah punya nomor surat sebagai ZIP.
func DownloadAllSPKSE2026Handler(w http.ResponseWriter, r *http.Request) {
	jenisFlt := strings.TrimSpace(r.URL.Query().Get("jenis")) // "pcl", "pml", atau kosong = semua
	tanggal := r.URL.Query().Get("tanggal")
	if tanggal == "" {
		tanggal = "2026-05-31"
	}

	var kegiatanFilter string
	switch jenisFlt {
	case "pml":
		kegiatanFilter = "%Pemeriksa Lapangan Sensus Ekonomi%"
	case "pcl":
		kegiatanFilter = "%Pendataan Sensus Ekonomi%"
	default:
		kegiatanFilter = "%Sensus Ekonomi 2026%"
	}

	rows, err := database.DB.Query(`
		SELECT r.id
		FROM rekap r
		WHERE r.kegiatan LIKE ?
		  AND r.id_spk LIKE '%SPK-SE2026%'
		  AND (r.tanggaran = 2026 OR r.tahun = '2026')
		ORDER BY CAST(SUBSTRING_INDEX(SUBSTRING_INDEX(r.id_spk, '-', 2), '-', -1) AS UNSIGNED)`,
		kegiatanFilter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var ids []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err == nil {
			ids = append(ids, id)
		}
	}

	if len(ids) == 0 {
		http.Error(w, "Tidak ada SPK SE2026 dengan nomor surat", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="SPK_SE2026_Gelombang1_%s.zip"`, jenisFlt))

	zw := zip.NewWriter(w)
	defer zw.Close()

	for _, id := range ids {
		d, jenis, err := buildSPKSE2026(id, tanggal)
		if err != nil {
			continue
		}
		docx, err := generateSPKDocx(d, jenis)
		if err != nil {
			continue
		}
		fname := fmt.Sprintf("SPK_SE2026_%s_%s.docx", safeName(d.Nomor), safeName(d.NamaPetugas))
		f, err := zw.Create(fname)
		if err != nil {
			continue
		}
		f.Write(docx)
	}
}

// ListRekapSE2026Handler GET /api/rekap/spk/se2026/list?jenis=pcl
// Mengembalikan semua rekap SE2026, dengan flag sudah_spk = punya nomor surat.
func ListRekapSE2026Handler(w http.ResponseWriter, r *http.Request) {
	jenis := strings.TrimSpace(r.URL.Query().Get("jenis"))
	var kegiatanFilter string
	if jenis == "pml" {
		kegiatanFilter = "%Pemeriksa Lapangan Sensus Ekonomi%"
	} else {
		kegiatanFilter = "%Pendataan Sensus Ekonomi%"
	}

	rows, err := database.DB.Query(`
		SELECT r.id, r.idsobat, r.namamitra, r.kegiatan, r.honor,
		       COALESCE(r.id_spk, ''),
		       COALESCE(r.id_bapp1, ''), COALESCE(r.id_bapp2, ''), COALESCE(r.id_pernyataan1, ''),
		       COALESCE(NULLIF(r.jumlah_sls,''),'0'), r.realisasi_sls,
		       COALESCE((SELECT m.alamat FROM mitra m WHERE m.idsobat = r.idsobat LIMIT 1), ''),
		       COALESCE((SELECT m.kecamatan FROM mitra m WHERE m.idsobat = r.idsobat LIMIT 1), '')
		FROM rekap r
		WHERE r.kegiatan LIKE ?
		  AND (r.tanggaran = 2026 OR r.tahun = '2026')
		ORDER BY CASE WHEN r.id_spk IS NOT NULL AND r.id_spk != '' THEN 0 ELSE 1 END,
		         CAST(SUBSTRING_INDEX(SUBSTRING_INDEX(COALESCE(r.id_spk,'B-999'), '-', 2), '-', -1) AS UNSIGNED)`,
		kegiatanFilter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type Item struct {
		ID              int     `json:"id"`
		IDSobat         string  `json:"idsobat"`
		Nama            string  `json:"nama"`
		Kegiatan        string  `json:"kegiatan"`
		Honor           float64 `json:"honor"`
		IDSpk           string  `json:"id_spk"`
		SudahSPK        bool    `json:"sudah_spk"`
		IDBapp1         string  `json:"id_bapp1"`
		SudahBapp1      bool    `json:"sudah_bapp1"`
		IDBapp2         string  `json:"id_bapp2"`
		SudahBapp2      bool    `json:"sudah_bapp2"`
		IDPernyataan1   string  `json:"id_pernyataan1"`
		SudahPernyataan bool    `json:"sudah_pernyataan"`
		TargetSLS       int     `json:"target_sls"`
		RealisasiSLS    int     `json:"realisasi_sls"`
		Alamat          string  `json:"alamat"`
		Kecamatan       string  `json:"kecamatan"`
	}

	var items []Item
	for rows.Next() {
		var item Item
		var honorStr, jumlahSLSStr string
		if err := rows.Scan(&item.ID, &item.IDSobat, &item.Nama, &item.Kegiatan,
			&honorStr, &item.IDSpk, &item.IDBapp1, &item.IDBapp2, &item.IDPernyataan1,
			&jumlahSLSStr, &item.RealisasiSLS, &item.Alamat, &item.Kecamatan); err != nil {
			continue
		}
		item.Honor, _ = strconv.ParseFloat(honorStr, 64)
		item.TargetSLS, _ = strconv.Atoi(strings.TrimSpace(jumlahSLSStr))
		item.SudahSPK = strings.Contains(item.IDSpk, "SPK-SE2026")
		item.SudahBapp1 = item.IDBapp1 != ""
		item.SudahBapp2 = item.IDBapp2 != ""
		item.SudahPernyataan = item.IDPernyataan1 != ""
		items = append(items, item)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// UpdateJumlahSLSHandler POST /api/rekap/spk/se2026/update-sls
func UpdateJumlahSLSHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		RekapID   int `json:"rekap_id"`
		JumlahSLS int `json:"jumlah_sls"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	if _, err := database.DB.Exec(`UPDATE rekap SET jumlah_sls = ? WHERE id = ?`, req.JumlahSLS, req.RekapID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
