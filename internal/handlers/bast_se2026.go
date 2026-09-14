package handlers

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"sirekap/internal/database"
	"sirekap/internal/models"
)

// BAST untuk SE2026 sengaja dipisah dari alur BAST generik (spk_bast.go) karena mitra
// SE2026 nomor SPK-nya di-import langsung ke rekap.id_spk (bukan lewat models.CreateSurat/
// tabel `surat`), jadi CreateBulkBASTHandler generik gagal (butuh baris `surat` stujuan=SPK
// yang tidak pernah ada). BAST tetap satu dokumen per orang untuk SELURUH kegiatan (bukan
// per-termin, "gabungan total" - arahan user), jadi nomornya tetap lewat models.CreateSurat
// biasa (tabel `surat`, stujuan=BAST) supaya konsisten & unik lintas sistem, cuma proses
// pembuatannya di-scope ke daftar payable SE2026 & dijalankan satu-satu (bukan lewat jalur
// generik yang butuh SPK-surat).

type createBulkBASTSE2026Item struct {
	RekapID int `json:"rekap_id"`
}

type createBulkBASTSE2026Payload struct {
	Tanggal string                     `json:"tanggal"`
	Items   []createBulkBASTSE2026Item `json:"items"`
}

// CreateBulkBASTSE2026Handler POST /api/rekap/spk/se2026/bast/create-bulk
// Urutan Items MENENTUKAN urutan nomor surat (counter naik per tanggal) - dipakai utk
// aturan "5 petugas pengganti dapat nomor buntut" (kirim mereka TERAKHIR dalam array).
func CreateBulkBASTSE2026Handler(w http.ResponseWriter, r *http.Request) {
	var payload createBulkBASTSE2026Payload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Payload tidak valid", http.StatusBadRequest)
		return
	}
	payload.Tanggal = strings.TrimSpace(payload.Tanggal)
	if payload.Tanggal == "" {
		http.Error(w, "Tanggal surat wajib diisi", http.StatusBadRequest)
		return
	}
	if len(payload.Items) == 0 {
		http.Error(w, "Pilih minimal satu petugas", http.StatusBadRequest)
		return
	}

	successCount := 0
	var createdNumbers, failedItems []string

	for _, item := range payload.Items {
		rekap, err := models.GetRekapRecordByID(item.RekapID)
		if err != nil {
			failedItems = append(failedItems, fmt.Sprintf("rekap_id %d: tidak ditemukan", item.RekapID))
			continue
		}
		mitra, err := models.GetMitraByIDSobat(rekap.IDSobat)
		if err != nil {
			failedItems = append(failedItems, fmt.Sprintf("%s: data mitra tidak ditemukan", rekap.NamaMitra))
			continue
		}

		existing, err := models.GetBASTSurat(rekap.IDSobat, rekap.Bulan, rekap.Tahun, rekap.Kegiatan)
		if err == nil && existing != nil {
			createdNumbers = append(createdNumbers, existing.NSurat)
			successCount++
			continue // idempotent, sudah pernah dibuat
		}

		nsurat, err := models.CreateSurat(rekap.IDSobat, mitra.NamaMitra, rekap.Bulan, rekap.Tahun, "BAST", payload.Tanggal, rekap.Kegiatan)
		if err != nil {
			failedItems = append(failedItems, fmt.Sprintf("%s: %v", mitra.NamaMitra, err))
			continue
		}

		docBytes, filename, err := generateBASTDocument(item.RekapID)
		if err != nil {
			failedItems = append(failedItems, fmt.Sprintf("%s: gagal generate dokumen (%v)", mitra.NamaMitra, err))
			continue
		}
		if err := saveGeneratedDocument(filename, docBytes); err != nil {
			failedItems = append(failedItems, fmt.Sprintf("%s: gagal simpan dokumen (%v)", mitra.NamaMitra, err))
			continue
		}

		successCount++
		createdNumbers = append(createdNumbers, nsurat)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": successCount > 0,
		"created": successCount,
		"failed":  len(failedItems),
		"numbers": createdNumbers,
		"errors":  failedItems,
	})
}

// DownloadAllBASTSE2026Handler GET /api/rekap/spk/se2026/bast/download-all?jenis=pcl
// Download ZIP semua BAST yang SUDAH punya nomor (dibuat lewat create-bulk/create di atas),
// discope ke kegiatan SE2026 & daftar payable (excluded_termin2 TIDAK relevan di sini krn
// BAST gabungan total, bukan per-termin - tapi tetap discope ke idsobat yg terdaftar payable
// biar tidak ikut mitra kegiatan lain).
func DownloadAllBASTSE2026Handler(w http.ResponseWriter, r *http.Request) {
	jenisFlt := strings.TrimSpace(r.URL.Query().Get("jenis"))
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
		SELECT r.id FROM rekap r
		WHERE r.kegiatan LIKE ?
		  AND (r.tanggaran = 2026 OR r.tahun = '2026')
		  AND EXISTS (
		      SELECT 1 FROM surat s
		      WHERE s.idsobat = r.idsobat AND s.sbulan = r.bulan AND s.stahun = r.tahun
		        AND s.stujuan = 'BAST' AND s.keterangan = r.kegiatan
		  )`, kegiatanFilter)
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
		http.Error(w, "Tidak ada BAST dengan nomor surat", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="BAST_SE2026_%s.zip"`, jenisFlt))

	zw := zip.NewWriter(w)
	defer zw.Close()

	for _, id := range ids {
		docBytes, filename, err := generateBASTDocument(id)
		if err != nil {
			continue
		}
		f, err := zw.Create(fmt.Sprintf("%03d_%s", id, filename))
		if err != nil {
			continue
		}
		f.Write(docBytes)
	}
}

// ListBASTStatusSE2026Handler GET /api/rekap/spk/se2026/bast/status?jenis=pcl
// Flag sudah_bast per rekap_id (dari tabel `surat`, bukan kolom rekap) - dipakai UI utk
// nentuin tombol Buat vs Download di kolom BAST.
func ListBASTStatusSE2026Handler(w http.ResponseWriter, r *http.Request) {
	jenis := strings.TrimSpace(r.URL.Query().Get("jenis"))
	var kegiatanFilter string
	if jenis == "pml" {
		kegiatanFilter = "%Pemeriksa Lapangan Sensus Ekonomi%"
	} else {
		kegiatanFilter = "%Pendataan Sensus Ekonomi%"
	}

	rows, err := database.DB.Query(`
		SELECT r.id,
		       EXISTS (
		           SELECT 1 FROM surat s
		           WHERE s.idsobat = r.idsobat AND s.sbulan = r.bulan AND s.stahun = r.tahun
		             AND s.stujuan = 'BAST' AND s.keterangan = r.kegiatan
		       )
		FROM rekap r
		WHERE r.kegiatan LIKE ? AND (r.tanggaran = 2026 OR r.tahun = '2026')`, kegiatanFilter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	result := map[string]bool{}
	for rows.Next() {
		var id int
		var sudah bool
		if err := rows.Scan(&id, &sudah); err == nil {
			result[strconv.Itoa(id)] = sudah
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
