package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"sirekap/internal/database"
	"sirekap/internal/models"
)

// SPKDataSE2026 holds all data needed for generating SE2026 SPK documents.
type SPKDataSE2026 struct {
	Nomor string

	// Tanggal tanda tangan
	Hari     string
	TglTeks  string
	BlnTeks  string
	TglAngka string // format: DD-MM

	// Pihak Kedua
	NamaPetugas   string
	AlamatPetugas string

	// Honor
	HonorAngka    string
	HonorTerbilang string

	// Sanksi Pasal 11
	SanksiAngka    string
	SanksiTerbilang string

	// Lampiran - Target Pekerjaan
	JumlahSLS     int
	MuatanUsaha   int
	MuatanKeluarga int

	// Jenis: "pcl" atau "pml"
	Jenis string

	// Untuk PML: daftar PCL yang diawasi
	DaftarPCL []PCLBawahanPML
}

type PCLBawahanPML struct {
	Nama       string
	Kecamatan  string
	Desa       string
	JumlahSLS  int
}

type spkSE2026Request struct {
	RekapID        int    `json:"rekap_id"`
	Jenis          string `json:"jenis"`       // "pcl" atau "pml"
	NomorUrut      int    `json:"nomor_urut"`
	TanggalSPK     string `json:"tanggal_spk"` // YYYY-MM-DD
	AlamatPetugas  string `json:"alamat_petugas"`
	JumlahSLS      int    `json:"jumlah_sls"`
	MuatanUsaha    int    `json:"muatan_usaha"`
	MuatanKeluarga int    `json:"muatan_keluarga"`
	SanksiAngka    int64  `json:"sanksi_angka"`
	DaftarPCL      []struct {
		Nama      string `json:"nama"`
		Kecamatan string `json:"kecamatan"`
		Desa      string `json:"desa"`
		JumlahSLS int    `json:"jumlah_sls"`
	} `json:"daftar_pcl"`
}

var bulanTeks = map[time.Month]string{
	time.January: "Januari", time.February: "Februari", time.March: "Maret",
	time.April: "April", time.May: "Mei", time.June: "Juni",
	time.July: "Juli", time.August: "Agustus", time.September: "September",
	time.October: "Oktober", time.November: "November", time.December: "Desember",
}

var angkaTeks = []string{
	"", "satu", "dua", "tiga", "empat", "lima", "enam", "tujuh", "delapan", "sembilan",
	"sepuluh", "sebelas", "dua belas", "tiga belas", "empat belas", "lima belas",
	"enam belas", "tujuh belas", "delapan belas", "sembilan belas", "dua puluh",
	"dua puluh satu", "dua puluh dua", "dua puluh tiga", "dua puluh empat",
	"dua puluh lima", "dua puluh enam", "dua puluh tujuh", "dua puluh delapan",
	"dua puluh sembilan", "tiga puluh", "tiga puluh satu",
}

func tglTeks(day int) string {
	if day >= 1 && day <= len(angkaTeks)-1 {
		return angkaTeks[day]
	}
	return strconv.Itoa(day)
}

// nomorSPKSE2026 generates SPK number for SE2026.
// Format: B-XXX/SPK-SE2026/3201/PL.200/2026
func nomorSPKSE2026(urut int, jenis string) string {
	prefix := "B"
	if jenis == "pml" {
		prefix = "C"
	}
	return fmt.Sprintf("%s-%03d/SPK-SE2026/3201/PL.200/2026", prefix, urut)
}

// generateLampiranPCL generates the appendix table XML for PCL SPK.
func generateLampiranPCL(data SPKDataSE2026) string {
	honor, _ := strconv.ParseFloat(strings.ReplaceAll(data.HonorAngka, ".", ""), 64)
	termin1Nilai := honor * 0.4
	termin2Nilai := honor * 0.6

	sls1 := int(float64(data.JumlahSLS) * 0.4)
	if sls1 < 1 && data.JumlahSLS > 0 {
		sls1 = 1
	}
	sls2 := data.JumlahSLS - sls1

	usaha1 := int(float64(data.MuatanUsaha) * 0.4)
	usaha2 := data.MuatanUsaha - usaha1
	kel1 := int(float64(data.MuatanKeluarga) * 0.4)
	kel2 := data.MuatanKeluarga - kel1

	vol1 := buildVolumePCL(sls1, usaha1, kel1)
	vol2 := buildVolumePCL(sls2, usaha2, kel2)
	volTotal := buildVolumePCLTotal(data.JumlahSLS, data.MuatanUsaha, data.MuatanKeluarga)

	terbilangTotal := strings.Title(Terbilang(int64(honor))) + " rupiah"

	return fmt.Sprintf(`
DAFTAR URAIAN PEKERJAAN, WAKTU PENYELESAIAN, TARGET PEKERJAAN DAN NILAI PERJANJIAN

Termin I  | Minimal 1 bulan | 40%% | %s | %s
Termin II | 31 Agustus 2026  | 60%% | %s | %s
Total     | 15 Juni–31 Agustus 2026 | 100%% | %s | %s
Terbilang: %s`,
		vol1, FormatRupiah(termin1Nilai),
		vol2, FormatRupiah(termin2Nilai),
		volTotal, FormatRupiah(honor),
		terbilangTotal,
	)
}

func buildVolumePCL(sls, usaha, keluarga int) string {
	parts := []string{}
	if sls > 0 {
		parts = append(parts, fmt.Sprintf("%d SLS/sub-SLS", sls))
	}
	if usaha > 0 && keluarga > 0 {
		parts = append(parts, fmt.Sprintf("%d usaha/%d keluarga", usaha, keluarga))
	} else if usaha > 0 {
		parts = append(parts, fmt.Sprintf("%d usaha besar", usaha))
	} else if keluarga > 0 {
		parts = append(parts, fmt.Sprintf("%d keluarga", keluarga))
	}
	if len(parts) == 0 {
		return "..."
	}
	return strings.Join(parts, " dan/atau ")
}

func buildVolumePCLTotal(sls, usaha, keluarga int) string {
	if sls > 0 {
		return fmt.Sprintf("Seluruh Muatan %d SLS/sub-SLS", sls)
	}
	if usaha > 0 {
		return fmt.Sprintf("Seluruh %d usaha besar", usaha)
	}
	return "..."
}

// generateLampiranPMLXML generates the appendix table XML for PML SPK.
func generateLampiranPMLXML(data SPKDataSE2026) string {
	honor, _ := strconv.ParseFloat(strings.ReplaceAll(data.HonorAngka, ".", ""), 64)
	termin1Nilai := honor * 0.4
	termin2Nilai := honor * 0.6

	slsTotal := 0
	for _, pcl := range data.DaftarPCL {
		slsTotal += pcl.JumlahSLS
	}
	sls1 := int(float64(slsTotal) * 0.4)
	sls2 := slsTotal - sls1

	terbilangTotal := strings.Title(Terbilang(int64(honor))) + " rupiah"

	lampiranText := fmt.Sprintf(`
DAFTAR URAIAN PEKERJAAN, WAKTU PENYELESAIAN, TARGET PEKERJAAN DAN NILAI PERJANJIAN

Termin I  | Minimal 1 bulan | 40%% | %d SLS/Sub-SLS | %s
Termin II | 31 Agustus 2026  | 60%% | %d SLS/Sub-SLS | %s
Total     | 15 Juni–31 Agustus 2026 | 100%% | %d SLS/Sub-SLS | %s
Terbilang: %s

DAFTAR PETUGAS DAN WILAYAH KERJA:`,
		sls1, FormatRupiah(termin1Nilai),
		sls2, FormatRupiah(termin2Nilai),
		slsTotal, FormatRupiah(honor),
		terbilangTotal,
	)

	for i, pcl := range data.DaftarPCL {
		lampiranText += fmt.Sprintf("\n%d. %s | %s | %s | %d SLS",
			i+1, pcl.Nama, pcl.Kecamatan, pcl.Desa, pcl.JumlahSLS)
	}
	return lampiranText
}

// buildSPKData builds SPKDataSE2026 from request and DB query.
func buildSPKData(req spkSE2026Request) (SPKDataSE2026, error) {
	var data SPKDataSE2026

	// Query rekap for mitra info
	row := database.DB.QueryRow(`
		SELECT r.namamitra, r.honor, m.alamat, m.kecamatan
		FROM rekap r
		LEFT JOIN mitra m ON m.idsobat = r.idsobat
		WHERE r.id = ?`, req.RekapID)

	var honorStr string
	var alamat, kecamatan *string
	err := row.Scan(&data.NamaPetugas, &honorStr, &alamat, &kecamatan)
	if err != nil {
		return data, fmt.Errorf("rekap tidak ditemukan: %w", err)
	}

	honorVal, _ := strconv.ParseFloat(honorStr, 64)

	// Tanggal
	tgl, err := time.Parse("2006-01-02", req.TanggalSPK)
	if err != nil {
		tgl = time.Now()
	}
	hariEn := tgl.Weekday().String()
	data.Hari = hariIndonesia[hariEn]
	data.TglTeks = tglTeks(tgl.Day())
	data.BlnTeks = bulanTeks[tgl.Month()]
	data.TglAngka = fmt.Sprintf("%02d-%02d", tgl.Day(), int(tgl.Month()))

	data.Nomor = nomorSPKSE2026(req.NomorUrut, req.Jenis)
	data.Jenis = req.Jenis

	if req.AlamatPetugas != "" {
		data.AlamatPetugas = req.AlamatPetugas
	} else if alamat != nil && *alamat != "" {
		data.AlamatPetugas = *alamat
	} else if kecamatan != nil {
		data.AlamatPetugas = *kecamatan + ", Kab. Dompu"
	} else {
		data.AlamatPetugas = "Kab. Dompu"
	}

	// Honor formatting
	honorInt := int64(honorVal)
	data.HonorAngka = formatAngka(honorInt)
	data.HonorTerbilang = strings.Title(Terbilang(honorInt))

	// Sanksi default: sama dengan 1 OB SBKS
	sanksi := req.SanksiAngka
	if sanksi == 0 {
		sanksi = int64(honorVal * 0.4) // 40% jika tidak diisi
	}
	data.SanksiAngka = formatAngka(sanksi)
	data.SanksiTerbilang = strings.Title(Terbilang(sanksi))

	data.JumlahSLS = req.JumlahSLS
	data.MuatanUsaha = req.MuatanUsaha
	data.MuatanKeluarga = req.MuatanKeluarga

	for _, p := range req.DaftarPCL {
		data.DaftarPCL = append(data.DaftarPCL, PCLBawahanPML{
			Nama:      p.Nama,
			Kecamatan: p.Kecamatan,
			Desa:      p.Desa,
			JumlahSLS: p.JumlahSLS,
		})
	}

	return data, nil
}

func formatAngka(n int64) string {
	s := strconv.FormatInt(n, 10)
	result := ""
	for i, c := range s {
		if i > 0 && (len(s)-i)%3 == 0 {
			result += "."
		}
		result += string(c)
	}
	return result
}

// CreateSPKSE2026Handler handles POST /api/rekap/spk/se2026/create
func CreateSPKSE2026Handler(w http.ResponseWriter, r *http.Request) {
	var req spkSE2026Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if req.RekapID == 0 {
		http.Error(w, "rekap_id wajib diisi", http.StatusBadRequest)
		return
	}
	if req.Jenis != "pcl" && req.Jenis != "pml" {
		http.Error(w, "jenis harus 'pcl' atau 'pml'", http.StatusBadRequest)
		return
	}
	if req.NomorUrut == 0 {
		http.Error(w, "nomor_urut wajib diisi", http.StatusBadRequest)
		return
	}

	data, err := buildSPKData(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var templatePath string
	if req.Jenis == "pcl" {
		templatePath = "static/templates-se2026/spk_pcl_se2026.docx"
	} else {
		templatePath = "static/templates-se2026/spk_pml_se2026.docx"
	}

	var lampiranInfo string
	if req.Jenis == "pcl" {
		lampiranInfo = generateLampiranPCL(data)
	} else {
		lampiranInfo = generateLampiranPMLXML(data)
	}

	honorVal, _ := strconv.ParseFloat(strings.ReplaceAll(data.HonorAngka, ".", ""), 64)
	replacements := map[string]string{
		"${nomor}":            data.Nomor,
		"${hari}":             data.Hari,
		"${tgl_teks}":         data.TglTeks,
		"${bln_teks}":         data.BlnTeks,
		"${tgl_angka}":        data.TglAngka,
		"${nama_petugas}":     data.NamaPetugas,
		"${alamat_petugas}":   data.AlamatPetugas,
		"${honor_angka}":      data.HonorAngka,
		"${honor_terbilang}":  data.HonorTerbilang,
		"${sanksi_angka}":     data.SanksiAngka,
		"${sanksi_terbilang}": data.SanksiTerbilang,
		"${terbilang_total}":  strings.Title(Terbilang(int64(honorVal))) + " rupiah",
		"${lampiran_info}":    lampiranInfo,
	}

	docxBytes, err := replaceInDocx(templatePath, replacements)
	if err != nil {
		http.Error(w, "Gagal generate SPK: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Save nomor SPK ke rekap
	_, _ = database.DB.Exec(
		`UPDATE rekap SET id_spk = ? WHERE id = ?`,
		data.Nomor, req.RekapID,
	)

	jenisFname := "PCL"
	if req.Jenis == "pml" {
		jenisFname = "PML"
	}
	namaSafe := strings.Map(func(r rune) rune {
		if r == ' ' {
			return '_'
		}
		if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '_' || r == '-' {
			return r
		}
		return -1
	}, data.NamaPetugas)

	filename := fmt.Sprintf("SPK_SE2026_%s_%s.docx", jenisFname, namaSafe)
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, filename))
	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.wordprocessingml.document")
	w.Header().Set("Content-Length", strconv.Itoa(len(docxBytes)))
	w.Write(docxBytes)
}

// ListRekapSE2026Handler returns rekap entries for SE2026 kegiatan.
func ListRekapSE2026Handler(w http.ResponseWriter, r *http.Request) {
	jenis := strings.TrimSpace(r.URL.Query().Get("jenis")) // "pcl" atau "pml"

	var kegiatanFilter string
	if jenis == "pml" {
		kegiatanFilter = "%Pemeriksa Lapangan Sensus Ekonomi%"
	} else {
		kegiatanFilter = "%Pendataan Sensus Ekonomi%"
	}

	rows, err := database.DB.Query(`
		SELECT r.id, r.idsobat, r.namamitra, r.kegiatan, r.honor, r.jumlah_sls, r.id_spk,
		       m.alamat, m.kecamatan
		FROM rekap r
		LEFT JOIN mitra m ON m.idsobat = r.idsobat
		WHERE r.kegiatan LIKE ?
		  AND (r.tanggaran = 2026 OR r.tahun = '2026')
		ORDER BY r.namamitra`, kegiatanFilter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type RekapSE2026 struct {
		ID        int     `json:"id"`
		IDSobat   string  `json:"idsobat"`
		Nama      string  `json:"nama"`
		Kegiatan  string  `json:"kegiatan"`
		Honor     float64 `json:"honor"`
		JumlahSLS int     `json:"jumlah_sls"`
		IDSpk     string  `json:"id_spk"`
		Alamat    string  `json:"alamat"`
		Kecamatan string  `json:"kecamatan"`
	}

	var items []RekapSE2026
	for rows.Next() {
		var item RekapSE2026
		var idSpk, alamat, kecamatan *string
		var honorStr string
		if err := rows.Scan(&item.ID, &item.IDSobat, &item.Nama, &item.Kegiatan,
			&honorStr, &item.JumlahSLS, &idSpk, &alamat, &kecamatan); err != nil {
			continue
		}
		item.Honor, _ = strconv.ParseFloat(honorStr, 64)
		if idSpk != nil {
			item.IDSpk = *idSpk
		}
		if alamat != nil {
			item.Alamat = *alamat
		}
		if kecamatan != nil {
			item.Kecamatan = *kecamatan
		}
		items = append(items, item)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// GetNextNomorSPKSE2026Handler returns the next available nomor urut.
func GetNextNomorSPKSE2026Handler(w http.ResponseWriter, r *http.Request) {
	jenis := strings.TrimSpace(r.URL.Query().Get("jenis"))
	prefix := "B-%"
	if jenis == "pml" {
		prefix = "C-%"
	}

	var maxNomor int
	err := database.DB.QueryRow(`
		SELECT COALESCE(MAX(
			CAST(SUBSTRING_INDEX(SUBSTRING_INDEX(id_spk, '-', 2), '-', -1) AS UNSIGNED)
		), 0)
		FROM rekap
		WHERE id_spk LIKE ? AND id_spk LIKE '%SPK-SE2026%'`, prefix).Scan(&maxNomor)
	if err != nil {
		maxNomor = 0
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"next_urut": maxNomor + 1,
		"preview":   nomorSPKSE2026(maxNomor+1, jenis),
	})
}

// UpdateJumlahSLSHandler updates the jumlah_sls field in rekap for SE2026.
func UpdateJumlahSLSHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		RekapID   int `json:"rekap_id"`
		JumlahSLS int `json:"jumlah_sls"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	_, err := database.DB.Exec(`UPDATE rekap SET jumlah_sls = ? WHERE id = ?`, req.JumlahSLS, req.RekapID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

// verifikasi model import
var _ = models.GetDashboardStats
