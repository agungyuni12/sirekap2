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

// errNotAssigned is returned by resolvePenilaiPeran when an organik account
// has no assignment for the periode they're trying to score.
var errNotAssigned = errors.New("penilai belum ditugaskan untuk periode ini")

// resolvePenilaiPeran decides the peran ("ppl" | "pml") a logged-in user is
// allowed to score in Tahap 1, for one specific periode instance. PML Mitra
// (level "pml_mitra") always scores PPL, unchanged. Everyone else — organik
// accounts, and admin if they ever try — must have been explicitly assigned
// a peran ("pml" scores PPL, "korwil" scores PML Mitra) for that periode via
// Kelola Petugas; without an assignment, access is denied (errNotAssigned).
// This is decided server-side from the session, never trusted from the
// request body.
func resolvePenilaiPeran(userID int, level string, periodeID int) (string, error) {
	if level == "pml_mitra" {
		return "ppl", nil
	}
	assigned, err := models.GetPenilaiPeranForUser(periodeID, userID)
	if errors.Is(err, sql.ErrNoRows) {
		return "", errNotAssigned
	}
	if err != nil {
		return "", err
	}
	if assigned == "korwil" {
		return "pml", nil
	}
	return "ppl", nil
}

// ListKegiatanForPenilaianHandler GET /api/penilaian/kegiatan[?mine=1]
// Kegiatan list dedicated to Penilaian Mitra (seeded from
// Daftar_Nama_Kegiatan_Rapi.xlsx), separate from the budget `kegiatan` table.
// With ?mine=1, returns only the kegiatan the logged-in organik account has
// at least one periode assignment in — used to drive step 1 (pilih Kegiatan)
// of Input Penilaian; step 2 (pilih Periode) is ListPeriodeForKegiatanHandler.
func ListKegiatanForPenilaianHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("mine") == "1" {
		userData := GetUserFromSession(r)
		userID, _ := userData["UserID"].(int)
		level, _ := userData["UserRole"].(string)
		if level != "pml_mitra" {
			periodes, err := models.ListPeriodeSayaForOrganik(userID)
			if err != nil {
				log.Printf("Error listing kegiatan saya: %v", err)
				apiError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Gagal memuat daftar kegiatan")
				return
			}
			seen := map[int]bool{}
			items := make([]models.PenilaianKegiatan, 0, len(periodes))
			for _, p := range periodes {
				if seen[p.KegiatanID] {
					continue
				}
				seen[p.KegiatanID] = true
				items = append(items, models.PenilaianKegiatan{ID: p.KegiatanID, Nama: p.Kegiatan})
			}
			writeJSON(w, http.StatusOK, map[string]interface{}{"data": items})
			return
		}
		// PML Mitra always scores PPL for every kegiatan — unchanged, unrestricted.
	}

	items, err := models.ListPenilaianKegiatan()
	if err != nil {
		log.Printf("Error listing kegiatan for penilaian: %v", err)
		apiError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Gagal memuat daftar kegiatan")
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{"data": items})
}

// ListPeriodeForKegiatanHandler GET /api/penilaian/kegiatan/periode?kegiatan_id=[&mine=1]
// Step 2 (pilih Periode) of Input Penilaian, and the periode picker on Kelola
// Petugas. With ?mine=1 (organik/admin, not pml_mitra), returns only periode
// instances the caller has been assigned to, each tagged with their peran.
func ListPeriodeForKegiatanHandler(w http.ResponseWriter, r *http.Request) {
	kegiatanID, err := strconv.Atoi(r.URL.Query().Get("kegiatan_id"))
	if err != nil || kegiatanID <= 0 {
		apiError(w, http.StatusBadRequest, "INVALID_KEGIATAN", "kegiatan_id tidak valid")
		return
	}

	if r.URL.Query().Get("mine") == "1" {
		userData := GetUserFromSession(r)
		userID, _ := userData["UserID"].(int)
		level, _ := userData["UserRole"].(string)
		if level != "pml_mitra" {
			all, err := models.ListPeriodeSayaForOrganik(userID)
			if err != nil {
				log.Printf("Error listing periode saya: %v", err)
				apiError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Gagal memuat daftar periode")
				return
			}
			items := make([]models.PeriodeSayaItem, 0, len(all))
			for _, it := range all {
				if it.KegiatanID == kegiatanID {
					items = append(items, it)
				}
			}
			writeJSON(w, http.StatusOK, map[string]interface{}{"data": items})
			return
		}
		// PML Mitra: unrestricted, falls through to the full periode list below.
	}

	items, err := models.ListPeriodeForKegiatan(kegiatanID)
	if err != nil {
		log.Printf("Error listing periode for kegiatan: %v", err)
		apiError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Gagal memuat daftar periode")
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{"data": items})
}

type periodePayload struct {
	KegiatanID int    `json:"kegiatan_id"`
	Periode    string `json:"periode"`
	Tahun      string `json:"tahun"`
}

// AddPeriodeKegiatanHandler POST /api/penilaian/kegiatan/periode — admin-only.
func AddPeriodeKegiatanHandler(w http.ResponseWriter, r *http.Request) {
	var payload periodePayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		apiError(w, http.StatusBadRequest, "INVALID_BODY", "Payload tidak valid")
		return
	}
	if payload.KegiatanID <= 0 || payload.Periode == "" || payload.Tahun == "" {
		apiError(w, http.StatusUnprocessableEntity, "MISSING_FIELDS", "kegiatan_id, periode, dan tahun wajib diisi")
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

	id, err := models.AddPeriodeKegiatan(payload.KegiatanID, payload.Periode, payload.Tahun, adminID)
	if err != nil {
		log.Printf("Error adding periode kegiatan: %v", err)
		apiError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Gagal menambahkan periode")
		return
	}
	writeJSON(w, http.StatusCreated, map[string]interface{}{"status": "created", "id": id, "message": "Periode ditambahkan"})
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
	PeriodeID          int     `json:"periode_id"`
	YangDinilaiID      string  `json:"yang_dinilai_id"`
	Kecamatan          string  `json:"kecamatan"`
	TanggalPenilaian   string  `json:"tanggal_penilaian"`
	SkorKualitas       float64 `json:"skor_kualitas"`
	SkorKetepatanWaktu float64 `json:"skor_ketepatan_waktu"`
	SkorKepatuhanSOP   float64 `json:"skor_kepatuhan_sop"`
	SkorKomunikasi     float64 `json:"skor_komunikasi"`
	SkorSikap          float64 `json:"skor_sikap"`
	Catatan            string  `json:"catatan"`
}

// toInput builds the EvaluasiInput to save — kegiatanID/periode/tahun are
// resolved server-side from periodeID (via models.GetPeriodeInfo), never
// trusted from the client, so callers must pass those in explicitly rather
// than relying on payload fields.
func (p penilaianPayload) toInput(kegiatanID, periodeID int, periode, tahun, peran string) models.EvaluasiInput {
	return models.EvaluasiInput{
		KegiatanID:         kegiatanID,
		PeriodeID:          periodeID,
		YangDinilaiIDSobat: p.YangDinilaiID,
		Peran:              peran,
		Kecamatan:          p.Kecamatan,
		Periode:            periode,
		Tahun:              tahun,
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
	if payload.PeriodeID <= 0 || payload.YangDinilaiID == "" {
		apiError(w, http.StatusUnprocessableEntity, "MISSING_FIELDS", "periode_id dan yang_dinilai_id wajib diisi")
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

	kegiatanID, _, periode, tahun, err := models.GetPeriodeInfo(payload.PeriodeID)
	if errors.Is(err, sql.ErrNoRows) {
		apiError(w, http.StatusNotFound, "PERIODE_NOT_FOUND", "Periode tidak ditemukan")
		return
	}
	if err != nil {
		log.Printf("Error loading periode info: %v", err)
		apiError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Gagal memeriksa periode")
		return
	}

	userData := GetUserFromSession(r)
	penilaiID, _ := userData["UserID"].(int)
	if penilaiID <= 0 {
		apiError(w, http.StatusUnauthorized, "UNAUTHENTICATED", "Sesi tidak valid")
		return
	}
	level, _ := userData["UserRole"].(string)
	peran, err := resolvePenilaiPeran(penilaiID, level, payload.PeriodeID)
	if errors.Is(err, errNotAssigned) {
		apiError(w, http.StatusForbidden, "NOT_ASSIGNED", "Anda belum ditugaskan untuk periode ini. Hubungi admin.")
		return
	}
	if err != nil {
		log.Printf("Error resolving penilai peran: %v", err)
		apiError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Gagal memeriksa penugasan")
		return
	}

	// Petugas yang sudah punya skor Tahap 1 untuk periode ini tidak boleh
	// dinilai ulang lewat form ini — mencegah nilai lama tertimpa tanpa
	// sadar (roster picker juga sudah menyembunyikan mereka, ini jaga-jaga
	// kalau ada yang coba lewat API langsung).
	if _, _, err := models.GetTahap1Status(payload.PeriodeID, payload.YangDinilaiID); err == nil {
		apiError(w, http.StatusConflict, "ALREADY_SCORED", "Petugas ini sudah dinilai untuk periode ini.")
		return
	} else if !errors.Is(err, sql.ErrNoRows) {
		log.Printf("Error checking existing skor: %v", err)
		apiError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Gagal memeriksa penilaian sebelumnya")
		return
	}

	in := payload.toInput(kegiatanID, payload.PeriodeID, periode, tahun, peran)
	if err := models.UpsertEvaluasiTahap1(penilaiID, in); err != nil {
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
	PeriodeID int    `json:"periode_id"`
	IDSobat   string `json:"idsobat"`
	Setuju    bool   `json:"setuju"`
	Catatan   string `json:"catatan"`
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
	if payload.PeriodeID <= 0 || payload.IDSobat == "" {
		apiError(w, http.StatusUnprocessableEntity, "MISSING_FIELDS", "periode_id dan idsobat wajib diisi")
		return
	}

	userData := GetUserFromSession(r)
	subjectMatterID, _ := userData["UserID"].(int)
	if subjectMatterID <= 0 {
		apiError(w, http.StatusUnauthorized, "UNAUTHENTICATED", "Sesi tidak valid")
		return
	}

	err := models.KonfirmasiPenilaian(payload.PeriodeID, payload.IDSobat, payload.Setuju, subjectMatterID, payload.Catatan)
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

	_, status, err := models.GetTahap1Status(payload.PeriodeID, payload.YangDinilaiID)
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

	kegiatanID, _, periode, tahun, err := models.GetPeriodeInfo(payload.PeriodeID)
	if err != nil {
		log.Printf("Error loading periode info: %v", err)
		apiError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Gagal memeriksa periode")
		return
	}

	userData := GetUserFromSession(r)
	subjectMatterID, _ := userData["UserID"].(int)
	if subjectMatterID <= 0 {
		apiError(w, http.StatusUnauthorized, "UNAUTHENTICATED", "Sesi tidak valid")
		return
	}

	in := payload.toInput(kegiatanID, payload.PeriodeID, periode, tahun, "pml")
	if err := models.SubmitPenilaianUlang(subjectMatterID, in); err != nil {
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
	periodeID, _ := strconv.Atoi(r.URL.Query().Get("periode_id"))
	return models.DaftarPenilaianFilter{
		KegiatanID:       kegiatanID,
		PeriodeID:        periodeID,
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

// DetailPenilaianHandler GET /api/penilaian/detail?periode_id=&idsobat=
// Riwayat penilaian per aspek — the drill-down behind one Daftar Penilaian row.
func DetailPenilaianHandler(w http.ResponseWriter, r *http.Request) {
	periodeID, err := strconv.Atoi(r.URL.Query().Get("periode_id"))
	if err != nil || periodeID <= 0 {
		apiError(w, http.StatusBadRequest, "INVALID_PERIODE", "periode_id tidak valid")
		return
	}
	idsobat := r.URL.Query().Get("idsobat")
	if idsobat == "" {
		apiError(w, http.StatusBadRequest, "INVALID_IDSOBAT", "idsobat tidak valid")
		return
	}

	detail, err := models.GetDetailPenilaian(periodeID, idsobat)
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

// ListRosterPetugasHandler GET /api/penilaian/kegiatan/petugas?periode_id=&peran=
// Roster petugas (PPL/PML Mitra/organik) terdaftar untuk satu periode
// instance — dipakai Input Penilaian untuk memilih siapa yang dinilai.
// Terbuka untuk semua role yang punya akses Penilaian (read-only; hanya
// admin yang bisa menambah/menghapus).
func ListRosterPetugasHandler(w http.ResponseWriter, r *http.Request) {
	periodeID, err := strconv.Atoi(r.URL.Query().Get("periode_id"))
	if err != nil || periodeID <= 0 {
		apiError(w, http.StatusBadRequest, "INVALID_PERIODE", "periode_id tidak valid")
		return
	}
	excludeScored := r.URL.Query().Get("exclude_scored") == "1"
	items, err := models.ListPetugasKegiatan(periodeID, r.URL.Query().Get("peran"), excludeScored)
	if err != nil {
		log.Printf("Error listing roster petugas: %v", err)
		apiError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Gagal memuat daftar petugas")
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{"data": items})
}

type rosterPayload struct {
	PeriodeID int    `json:"periode_id"`
	IDSobat   string `json:"idsobat"`
	UserID    int    `json:"user_id"`
	Peran     string `json:"peran"`
}

// AddRosterPetugasHandler POST /api/penilaian/kegiatan/petugas — admin-only.
// Menambahkan satu entri ke roster periode: mitra (idsobat, peran ppl/pml —
// yang dinilai) ATAU akun organik (user_id, peran pml/korwil — penugasan
// penilai). Korwil dan PML tidak boleh menambah sendiri karena "gk nentu
// siapa aja" yang sah bertugas di periode itu.
func AddRosterPetugasHandler(w http.ResponseWriter, r *http.Request) {
	var payload rosterPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		apiError(w, http.StatusBadRequest, "INVALID_BODY", "Payload tidak valid")
		return
	}
	if payload.PeriodeID <= 0 {
		apiError(w, http.StatusUnprocessableEntity, "MISSING_FIELDS", "periode_id wajib diisi")
		return
	}
	hasIDSobat := payload.IDSobat != ""
	hasUserID := payload.UserID > 0
	if hasIDSobat == hasUserID {
		apiError(w, http.StatusUnprocessableEntity, "MISSING_FIELDS", "isi salah satu: idsobat (mitra) atau user_id (organik)")
		return
	}
	if hasIDSobat && payload.Peran != "ppl" && payload.Peran != "pml" {
		apiError(w, http.StatusUnprocessableEntity, "INVALID_PERAN", "Peran mitra harus ppl atau pml")
		return
	}
	if hasUserID && payload.Peran != "pml" && payload.Peran != "korwil" {
		apiError(w, http.StatusUnprocessableEntity, "INVALID_PERAN", "Peran organik harus pml atau korwil")
		return
	}
	if _, _, _, _, err := models.GetPeriodeInfo(payload.PeriodeID); err != nil {
		apiError(w, http.StatusNotFound, "PERIODE_NOT_FOUND", "Periode tidak ditemukan")
		return
	}

	userData := GetUserFromSession(r)
	adminID, _ := userData["UserID"].(int)
	if adminID <= 0 {
		apiError(w, http.StatusUnauthorized, "UNAUTHENTICATED", "Sesi tidak valid")
		return
	}

	if err := models.AddPetugasKeKegiatan(payload.PeriodeID, payload.IDSobat, payload.UserID, payload.Peran, adminID); err != nil {
		log.Printf("Error adding roster petugas: %v", err)
		apiError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Gagal menambahkan petugas")
		return
	}
	writeJSON(w, http.StatusCreated, map[string]interface{}{"status": "created", "message": "Petugas ditambahkan ke periode"})
}

// RemoveRosterPetugasHandler DELETE /api/penilaian/kegiatan/petugas?periode_id=&idsobat=|user_id= — admin-only.
func RemoveRosterPetugasHandler(w http.ResponseWriter, r *http.Request) {
	periodeID, err := strconv.Atoi(r.URL.Query().Get("periode_id"))
	if err != nil || periodeID <= 0 {
		apiError(w, http.StatusBadRequest, "INVALID_PERIODE", "periode_id tidak valid")
		return
	}
	idsobat := r.URL.Query().Get("idsobat")
	userID, _ := strconv.Atoi(r.URL.Query().Get("user_id"))
	if idsobat == "" && userID <= 0 {
		apiError(w, http.StatusBadRequest, "INVALID_TARGET", "idsobat atau user_id wajib diisi")
		return
	}

	err = models.RemovePetugasDariKegiatan(periodeID, idsobat, userID)
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

// SearchPenilaiOrganikHandler GET /api/penilaian/penilai/search?q= — admin-only.
// Autocomplete akun organik untuk tab PML/Korwil di Kelola Petugas.
func SearchPenilaiOrganikHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	// Tab Korwil ("atasan langsung organik") cuma boleh diisi organik asli
	// (email @bps.go.id) — PML Mitra eksternal gak bisa jadi Korwil.
	bpsOnly := r.URL.Query().Get("peran") == "korwil"
	results, err := models.SearchPenilaiOrganik(query, bpsOnly)
	if err != nil {
		log.Printf("Error searching penilai organik: %v", err)
		apiError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Gagal mencari akun organik")
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{"data": results})
}
