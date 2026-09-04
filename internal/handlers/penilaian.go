package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"sirekap/internal/models"
)

// apiError writes the app-wide structured error payload:
// {"error": {"code": "...", "message": "..."}}
func apiError(w http.ResponseWriter, status int, code, message string) {
	writeJSON(w, status, map[string]interface{}{
		"error": map[string]string{
			"code":    code,
			"message": message,
		},
	})
}

// peranForRole maps the logged-in user's role to the peran they're allowed to
// score in Tahap 1 — "korwil" (atasan langsung organik / PJK-Korwil) scores
// PML Mitra, everyone else with penilaian access (organik PML, PML Mitra
// scoring the PPL they supervise) scores PPL. This is decided server-side
// from the session, never trusted from the request body.
func peranForRole(level string) string {
	if level == "korwil" {
		return "pml"
	}
	return "ppl"
}

// ListKegiatanForPenilaianHandler GET /api/penilaian/kegiatan
// Kegiatan list dedicated to Penilaian Mitra (seeded from
// Daftar_Nama_Kegiatan_Rapi.xlsx), separate from the budget `kegiatan` table.
func ListKegiatanForPenilaianHandler(w http.ResponseWriter, r *http.Request) {
	items, err := models.ListPenilaianKegiatan()
	if err != nil {
		log.Printf("Error listing kegiatan for penilaian: %v", err)
		apiError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Gagal memuat daftar kegiatan")
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{"data": items})
}

// SearchMitraPenilaianHandler GET /api/penilaian/mitra?q=
// Autocomplete for the "Nama petugas" field on Input Penilaian — searches all
// mitra directly (kegiatan for Penilaian Mitra no longer derives assignment
// from `rekap`, per the "daftar kegiatan terpisah" requirement).
func SearchMitraPenilaianHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	results, err := models.SearchMitra(query, "")
	if err != nil {
		log.Printf("Error searching mitra for penilaian: %v", err)
		apiError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Gagal mencari mitra")
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{"data": results})
}

type penilaianPayload struct {
	KegiatanID         int     `json:"kegiatan_id"`
	YangDinilaiID      string  `json:"yang_dinilai_id"`
	Kecamatan          string  `json:"kecamatan"`
	Periode            string  `json:"periode"`
	Tahun              string  `json:"tahun"`
	TanggalPenilaian   string  `json:"tanggal_penilaian"`
	SkorKualitas       float64 `json:"skor_kualitas"`
	SkorKetepatanWaktu float64 `json:"skor_ketepatan_waktu"`
	SkorKepatuhanSOP   float64 `json:"skor_kepatuhan_sop"`
	SkorKomunikasi     float64 `json:"skor_komunikasi"`
	SkorSikap          float64 `json:"skor_sikap"`
	Catatan            string  `json:"catatan"`
}

func (p penilaianPayload) toInput(kegiatanID int, peran string) models.EvaluasiInput {
	return models.EvaluasiInput{
		KegiatanID:         kegiatanID,
		YangDinilaiIDSobat: p.YangDinilaiID,
		Peran:              peran,
		Kecamatan:          p.Kecamatan,
		Periode:            p.Periode,
		Tahun:              p.Tahun,
		TanggalPenilaian:   p.TanggalPenilaian,
		SkorKualitas:       p.SkorKualitas,
		SkorKetepatanWaktu: p.SkorKetepatanWaktu,
		SkorKepatuhanSOP:   p.SkorKepatuhanSOP,
		SkorKomunikasi:     p.SkorKomunikasi,
		SkorSikap:          p.SkorSikap,
		Catatan:            p.Catatan,
	}
}

func decodePenilaianPayload(w http.ResponseWriter, r *http.Request) (penilaianPayload, bool) {
	var payload penilaianPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		apiError(w, http.StatusBadRequest, "INVALID_BODY", "Payload tidak valid")
		return payload, false
	}
	if payload.KegiatanID <= 0 || payload.YangDinilaiID == "" {
		apiError(w, http.StatusUnprocessableEntity, "MISSING_FIELDS", "kegiatan_id dan yang_dinilai_id wajib diisi")
		return payload, false
	}
	if msg := models.ValidateSkor(payload.SkorKualitas, payload.SkorKetepatanWaktu, payload.SkorKepatuhanSOP, payload.SkorKomunikasi, payload.SkorSikap); msg != "" {
		apiError(w, http.StatusUnprocessableEntity, "INVALID_SKOR_RANGE", msg)
		return payload, false
	}
	return payload, true
}

// SubmitPenilaianHandler POST /api/penilaian
// Tahap 1 initial score: PML (organik atau mitra) menilai PPL — final
// seketika. Korwil (atasan langsung organik / PJK-Korwil) menilai PML Mitra —
// menunggu konfirmasi Subject Matter.
func SubmitPenilaianHandler(w http.ResponseWriter, r *http.Request) {
	payload, ok := decodePenilaianPayload(w, r)
	if !ok {
		return
	}

	if _, err := models.GetPenilaianKegiatanNama(payload.KegiatanID); err != nil {
		apiError(w, http.StatusNotFound, "KEGIATAN_NOT_FOUND", "Kegiatan tidak ditemukan")
		return
	}

	userData := GetUserFromSession(r)
	penilaiID, _ := userData["UserID"].(int)
	if penilaiID <= 0 {
		apiError(w, http.StatusUnauthorized, "UNAUTHENTICATED", "Sesi tidak valid")
		return
	}
	level, _ := userData["UserRole"].(string)
	peran := peranForRole(level)

	if err := models.UpsertEvaluasiTahap1(penilaiID, payload.toInput(payload.KegiatanID, peran)); err != nil {
		log.Printf("Error saving penilaian: %v", err)
		apiError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Gagal menyimpan penilaian")
		return
	}

	writeJSON(w, http.StatusCreated, map[string]interface{}{
		"status":  "created",
		"message": "Penilaian berhasil disimpan",
	})
}

// ListPendingKonfirmasiHandler GET /api/penilaian/konfirmasi/pending
// Subject Matter's queue of PML Mitra scores awaiting approve/reject.
func ListPendingKonfirmasiHandler(w http.ResponseWriter, r *http.Request) {
	items, err := models.ListPendingKonfirmasi()
	if err != nil {
		log.Printf("Error listing pending konfirmasi: %v", err)
		apiError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Gagal memuat antrean konfirmasi")
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{"data": items})
}

type konfirmasiPayload struct {
	KegiatanID int    `json:"kegiatan_id"`
	IDSobat    string `json:"idsobat"`
	Setuju     bool   `json:"setuju"`
	Catatan    string `json:"catatan"`
}

// KonfirmasiPenilaianHandler POST /api/penilaian/konfirmasi
// Subject Matter approves or rejects one PML Mitra's Tahap 1 score. Approve ->
// that score is final. Reject -> Subject Matter must submit their own score
// next (POST /api/penilaian/ulang); final = average of the two.
func KonfirmasiPenilaianHandler(w http.ResponseWriter, r *http.Request) {
	var payload konfirmasiPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		apiError(w, http.StatusBadRequest, "INVALID_BODY", "Payload tidak valid")
		return
	}
	if payload.KegiatanID <= 0 || payload.IDSobat == "" {
		apiError(w, http.StatusUnprocessableEntity, "MISSING_FIELDS", "kegiatan_id dan idsobat wajib diisi")
		return
	}

	userData := GetUserFromSession(r)
	subjectMatterID, _ := userData["UserID"].(int)
	if subjectMatterID <= 0 {
		apiError(w, http.StatusUnauthorized, "UNAUTHENTICATED", "Sesi tidak valid")
		return
	}

	err := models.KonfirmasiPenilaianPMLMitra(payload.KegiatanID, payload.IDSobat, payload.Setuju, subjectMatterID, payload.Catatan)
	if errors.Is(err, sql.ErrNoRows) {
		apiError(w, http.StatusNotFound, "NOT_FOUND", "Penilaian PML Mitra tidak ditemukan")
		return
	}
	if err != nil {
		log.Printf("Error confirming penilaian: %v", err)
		apiError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Gagal menyimpan konfirmasi")
		return
	}

	message := "Penilaian disetujui."
	if !payload.Setuju {
		message = "Penilaian ditolak. Silakan berikan penilaian ulang."
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{"status": "ok", "message": message})
}

// SubmitPenilaianUlangHandler POST /api/penilaian/ulang
// Subject Matter's own score for a PML Mitra whose Tahap 1 score they
// rejected. Only allowed once that Tahap 1 score's status is "ditolak".
func SubmitPenilaianUlangHandler(w http.ResponseWriter, r *http.Request) {
	payload, ok := decodePenilaianPayload(w, r)
	if !ok {
		return
	}

	status, err := models.GetStatusKonfirmasi(payload.KegiatanID, payload.YangDinilaiID)
	if errors.Is(err, sql.ErrNoRows) {
		apiError(w, http.StatusNotFound, "NOT_FOUND", "Penilaian PML Mitra tidak ditemukan")
		return
	}
	if err != nil {
		log.Printf("Error checking status konfirmasi: %v", err)
		apiError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Gagal memeriksa status penilaian")
		return
	}
	if status != "ditolak" {
		apiError(w, http.StatusUnprocessableEntity, "NOT_REJECTED", "Penilaian ulang hanya berlaku setelah nilai awal ditolak")
		return
	}

	userData := GetUserFromSession(r)
	subjectMatterID, _ := userData["UserID"].(int)
	if subjectMatterID <= 0 {
		apiError(w, http.StatusUnauthorized, "UNAUTHENTICATED", "Sesi tidak valid")
		return
	}

	if err := models.SubmitPenilaianUlang(subjectMatterID, payload.toInput(payload.KegiatanID, "pml")); err != nil {
		log.Printf("Error saving penilaian ulang: %v", err)
		apiError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Gagal menyimpan penilaian ulang")
		return
	}

	writeJSON(w, http.StatusCreated, map[string]interface{}{
		"status":  "created",
		"message": "Penilaian ulang tersimpan. Skor akhir adalah rata-rata kedua penilaian.",
	})
}

// daftarPenilaianFilters reads the shared query params used by the Daftar
// Penilaian listing, the Dashboard stats, and the Excel export.
func daftarPenilaianFilters(r *http.Request) models.DaftarPenilaianFilter {
	kegiatanID, _ := strconv.Atoi(r.URL.Query().Get("kegiatan_id"))
	return models.DaftarPenilaianFilter{
		KegiatanID:       kegiatanID,
		Tahun:            r.URL.Query().Get("tahun"),
		Peran:            r.URL.Query().Get("peran"),
		Kecamatan:        r.URL.Query().Get("kecamatan"),
		Predikat:         r.URL.Query().Get("predikat"),
		StatusKonfirmasi: r.URL.Query().Get("status_konfirmasi"),
		Search:           r.URL.Query().Get("search"),
	}
}

// DaftarPenilaianHandler GET /api/penilaian/daftar
func DaftarPenilaianHandler(w http.ResponseWriter, r *http.Request) {
	items, err := models.GetDaftarPenilaian(daftarPenilaianFilters(r))
	if err != nil {
		log.Printf("Error loading daftar penilaian: %v", err)
		apiError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Gagal memuat daftar penilaian")
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{
		"data":       items,
		"total_data": len(items),
	})
}

// DetailPenilaianHandler GET /api/penilaian/detail?kegiatan_id=&idsobat=
// Riwayat penilaian per aspek — the drill-down behind one Daftar Penilaian row.
func DetailPenilaianHandler(w http.ResponseWriter, r *http.Request) {
	kegiatanID, err := strconv.Atoi(r.URL.Query().Get("kegiatan_id"))
	if err != nil || kegiatanID <= 0 {
		apiError(w, http.StatusBadRequest, "INVALID_KEGIATAN", "kegiatan_id tidak valid")
		return
	}
	idsobat := r.URL.Query().Get("idsobat")
	if idsobat == "" {
		apiError(w, http.StatusBadRequest, "INVALID_IDSOBAT", "idsobat tidak valid")
		return
	}

	detail, err := models.GetDetailPenilaian(kegiatanID, idsobat)
	if errors.Is(err, sql.ErrNoRows) {
		apiError(w, http.StatusNotFound, "NOT_FOUND", "Penilaian tidak ditemukan")
		return
	}
	if err != nil {
		log.Printf("Error loading detail penilaian: %v", err)
		apiError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Gagal memuat detail penilaian")
		return
	}
	writeJSON(w, http.StatusOK, detail)
}

// ListRosterPetugasHandler GET /api/penilaian/kegiatan/petugas?kegiatan_id=&peran=
// Roster petugas (PPL/PML Mitra) terdaftar untuk satu kegiatan — dipakai Input
// Penilaian untuk memilih siapa yang dinilai. Terbuka untuk semua role yang
// punya akses Penilaian (read-only; hanya admin yang bisa menambah/menghapus).
func ListRosterPetugasHandler(w http.ResponseWriter, r *http.Request) {
	kegiatanID, err := strconv.Atoi(r.URL.Query().Get("kegiatan_id"))
	if err != nil || kegiatanID <= 0 {
		apiError(w, http.StatusBadRequest, "INVALID_KEGIATAN", "kegiatan_id tidak valid")
		return
	}
	items, err := models.ListPetugasKegiatan(kegiatanID, r.URL.Query().Get("peran"))
	if err != nil {
		log.Printf("Error listing roster petugas: %v", err)
		apiError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Gagal memuat daftar petugas")
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{"data": items})
}

type rosterPayload struct {
	KegiatanID int    `json:"kegiatan_id"`
	IDSobat    string `json:"idsobat"`
	Peran      string `json:"peran"`
}

// AddRosterPetugasHandler POST /api/penilaian/kegiatan/petugas — admin-only.
// Menambahkan mitra ke roster kegiatan; Korwil dan PML tidak boleh menambah
// sendiri karena "gk nentu siapa aja" yang sah bertugas di kegiatan itu.
func AddRosterPetugasHandler(w http.ResponseWriter, r *http.Request) {
	var payload rosterPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		apiError(w, http.StatusBadRequest, "INVALID_BODY", "Payload tidak valid")
		return
	}
	if payload.KegiatanID <= 0 || payload.IDSobat == "" || (payload.Peran != "ppl" && payload.Peran != "pml") {
		apiError(w, http.StatusUnprocessableEntity, "MISSING_FIELDS", "kegiatan_id, idsobat, dan peran (ppl/pml) wajib diisi")
		return
	}
	if _, err := models.GetPenilaianKegiatanNama(payload.KegiatanID); err != nil {
		apiError(w, http.StatusNotFound, "KEGIATAN_NOT_FOUND", "Kegiatan tidak ditemukan")
		return
	}

	userData := GetUserFromSession(r)
	adminID, _ := userData["UserID"].(int)
	if adminID <= 0 {
		apiError(w, http.StatusUnauthorized, "UNAUTHENTICATED", "Sesi tidak valid")
		return
	}

	if err := models.AddPetugasKeKegiatan(payload.KegiatanID, payload.IDSobat, payload.Peran, adminID); err != nil {
		log.Printf("Error adding roster petugas: %v", err)
		apiError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Gagal menambahkan petugas")
		return
	}
	writeJSON(w, http.StatusCreated, map[string]interface{}{"status": "created", "message": "Petugas ditambahkan ke kegiatan"})
}

// RemoveRosterPetugasHandler DELETE /api/penilaian/kegiatan/petugas?kegiatan_id=&idsobat= — admin-only.
func RemoveRosterPetugasHandler(w http.ResponseWriter, r *http.Request) {
	kegiatanID, err := strconv.Atoi(r.URL.Query().Get("kegiatan_id"))
	if err != nil || kegiatanID <= 0 {
		apiError(w, http.StatusBadRequest, "INVALID_KEGIATAN", "kegiatan_id tidak valid")
		return
	}
	idsobat := r.URL.Query().Get("idsobat")
	if idsobat == "" {
		apiError(w, http.StatusBadRequest, "INVALID_IDSOBAT", "idsobat tidak valid")
		return
	}

	err = models.RemovePetugasDariKegiatan(kegiatanID, idsobat)
	if errors.Is(err, sql.ErrNoRows) {
		apiError(w, http.StatusNotFound, "NOT_FOUND", "Petugas tidak ditemukan di kegiatan ini")
		return
	}
	if err != nil {
		log.Printf("Error removing roster petugas: %v", err)
		apiError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Gagal menghapus petugas")
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{"status": "ok", "message": "Petugas dihapus dari kegiatan"})
}
