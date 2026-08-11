package handlers

import (
	"encoding/json"
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

// ListKegiatanForPenilaianHandler GET /api/penilaian/kegiatan?peran=ppl|pml
// A lightweight {id, nama} kegiatan list — unlike /api/master-kegiatan it
// carries no budget figures, so it's safe for non-admin roles (PML/PML Mitra).
// peran filters to only kegiatan designated for that position.
func ListKegiatanForPenilaianHandler(w http.ResponseWriter, r *http.Request) {
	items, err := models.ListKegiatanIDName(r.URL.Query().Get("peran"))
	if err != nil {
		log.Printf("Error listing kegiatan for penilaian: %v", err)
		apiError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Gagal memuat daftar kegiatan")
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{"data": items})
}

// ListPetugasPenilaianHandler GET /api/penilaian/petugas?kegiatan_id=&tahap=
func ListPetugasPenilaianHandler(w http.ResponseWriter, r *http.Request) {
	kegiatanID, err := strconv.Atoi(r.URL.Query().Get("kegiatan_id"))
	if err != nil || kegiatanID <= 0 {
		apiError(w, http.StatusBadRequest, "INVALID_KEGIATAN", "kegiatan_id tidak valid")
		return
	}
	tahap, err := strconv.Atoi(r.URL.Query().Get("tahap"))
	if err != nil || (tahap != 1 && tahap != 2) {
		apiError(w, http.StatusBadRequest, "INVALID_TAHAP", "tahap harus 1 atau 2")
		return
	}

	items, err := models.GetPetugasForPenilaian(kegiatanID, tahap)
	if err != nil {
		log.Printf("Error listing petugas penilaian: %v", err)
		apiError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Gagal memuat daftar petugas")
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{"data": items})
}

type penilaianPayload struct {
	KegiatanID         int     `json:"kegiatan_id"`
	YangDinilaiID      string  `json:"yang_dinilai_id"`
	Tahap              int     `json:"tahap"`
	SkorKualitas       float64 `json:"skor_kualitas"`
	SkorKetepatanWaktu float64 `json:"skor_ketepatan_waktu"`
	SkorKepatuhanSOP   float64 `json:"skor_kepatuhan_sop"`
	SkorKomunikasi     float64 `json:"skor_komunikasi"`
	SkorSikap          float64 `json:"skor_sikap"`
	Catatan            string  `json:"catatan"`
}

// SubmitPenilaianHandler POST /api/penilaian
// Tahap 1: PML (organik atau mitra) menilai PPL yang diawasi — nilai ini FINAL
// begitu disimpan, tidak ada review/finalisasi lanjutan oleh admin.
// Tahap 2: Subject Matter menilai PML Mitra secara langsung (independen dari Tahap 1).
func SubmitPenilaianHandler(w http.ResponseWriter, r *http.Request) {
	var payload penilaianPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		apiError(w, http.StatusBadRequest, "INVALID_BODY", "Payload tidak valid")
		return
	}

	if payload.KegiatanID <= 0 || payload.YangDinilaiID == "" {
		apiError(w, http.StatusUnprocessableEntity, "MISSING_FIELDS", "kegiatan_id dan yang_dinilai_id wajib diisi")
		return
	}
	if payload.Tahap != 1 && payload.Tahap != 2 {
		apiError(w, http.StatusUnprocessableEntity, "INVALID_TAHAP", "tahap harus 1 atau 2")
		return
	}
	if msg := models.ValidateSkor(payload.SkorKualitas, payload.SkorKetepatanWaktu, payload.SkorKepatuhanSOP, payload.SkorKomunikasi, payload.SkorSikap); msg != "" {
		apiError(w, http.StatusUnprocessableEntity, "INVALID_SKOR_RANGE", msg)
		return
	}

	peran, err := models.PeranForPetugas(payload.KegiatanID)
	if err != nil {
		apiError(w, http.StatusNotFound, "KEGIATAN_NOT_FOUND", "Kegiatan tidak ditemukan")
		return
	}

	if payload.Tahap == 1 && peran != "ppl" {
		apiError(w, http.StatusUnprocessableEntity, "INVALID_PERAN_TAHAP1", "Penilaian Tahap 1 hanya berlaku untuk PPL")
		return
	}

	if payload.Tahap == 2 && peran != "pml" {
		apiError(w, http.StatusUnprocessableEntity, "INVALID_PERAN_TAHAP2",
			"PPL hanya dinilai oleh PML (Tahap 1); tidak ada tahap penilaian lanjutan untuk PPL")
		return
	}

	userData := GetUserFromSession(r)
	penilaiID, _ := userData["UserID"].(int)
	if penilaiID <= 0 {
		apiError(w, http.StatusUnauthorized, "UNAUTHENTICATED", "Sesi tidak valid")
		return
	}

	err = models.UpsertEvaluasi(penilaiID, peran, models.EvaluasiInput{
		KegiatanID:         payload.KegiatanID,
		YangDinilaiIDSobat: payload.YangDinilaiID,
		Tahap:              payload.Tahap,
		SkorKualitas:       payload.SkorKualitas,
		SkorKetepatanWaktu: payload.SkorKetepatanWaktu,
		SkorKepatuhanSOP:   payload.SkorKepatuhanSOP,
		SkorKomunikasi:     payload.SkorKomunikasi,
		SkorSikap:          payload.SkorSikap,
		Catatan:            payload.Catatan,
	})
	if err != nil {
		log.Printf("Error saving penilaian: %v", err)
		apiError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Gagal menyimpan penilaian")
		return
	}

	writeJSON(w, http.StatusCreated, map[string]interface{}{
		"status":  "created",
		"message": "Penilaian berhasil disimpan",
	})
}

// RekapPenilaianHandler GET /api/penilaian/rekap?kegiatan_id=
func RekapPenilaianHandler(w http.ResponseWriter, r *http.Request) {
	kegiatanID, err := strconv.Atoi(r.URL.Query().Get("kegiatan_id"))
	if err != nil || kegiatanID <= 0 {
		apiError(w, http.StatusBadRequest, "INVALID_KEGIATAN", "kegiatan_id tidak valid")
		return
	}

	items, err := models.GetRekapPenilaian(kegiatanID)
	if err != nil {
		log.Printf("Error loading rekap penilaian: %v", err)
		apiError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Gagal memuat rekap penilaian")
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{"data": items})
}
