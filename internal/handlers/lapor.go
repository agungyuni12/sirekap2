package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"sirekap/internal/models"
	"sirekap/internal/storage"
)

func ListPendingTransportReportsHandler(w http.ResponseWriter, r *http.Request) {
	userData := GetUserFromSession(r)
	userNIP, _ := userData["UserNIP"].(string)
	userName, _ := userData["UserName"].(string)

	items, err := models.ListPendingTransportReports(userNIP, userName)
	if err != nil {
		http.Error(w, "Gagal memuat daftar perjalanan", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"data":      items,
		"totalData": len(items),
	})
}

func ListSubmittedTransportReportsHandler(w http.ResponseWriter, r *http.Request) {
	userData := GetUserFromSession(r)
	userNIP, _ := userData["UserNIP"].(string)
	userName, _ := userData["UserName"].(string)

	items, err := models.ListSubmittedTransportReports(userNIP, userName)
	if err != nil {
		http.Error(w, "Gagal memuat riwayat laporan", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"data":      items,
		"totalData": len(items),
	})
}

func ListPendingSPPDReportsHandler(w http.ResponseWriter, r *http.Request) {
	userData := GetUserFromSession(r)
	userNIP, _ := userData["UserNIP"].(string)
	userName, _ := userData["UserName"].(string)

	items, err := models.ListPendingSPPDReports(userNIP, userName)
	if err != nil {
		http.Error(w, "Gagal memuat daftar perjalanan SPPD", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"data":      items,
		"totalData": len(items),
	})
}

func ListSubmittedSPPDReportsHandler(w http.ResponseWriter, r *http.Request) {
	userData := GetUserFromSession(r)
	userNIP, _ := userData["UserNIP"].(string)
	userName, _ := userData["UserName"].(string)

	items, err := models.ListSubmittedSPPDReports(userNIP, userName)
	if err != nil {
		http.Error(w, "Gagal memuat riwayat laporan SPPD", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"data":      items,
		"totalData": len(items),
	})
}

func SubmitTransportReportHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(12 << 20); err != nil {
		http.Error(w, "Form upload tidak valid", http.StatusBadRequest)
		return
	}

	userData := GetUserFromSession(r)
	userNIP, _ := userData["UserNIP"].(string)
	userName, _ := userData["UserName"].(string)

	id, err := parsePositiveInt(r.FormValue("id_matriks"))
	if err != nil {
		http.Error(w, "ID perjalanan tidak valid", http.StatusBadRequest)
		return
	}

	laporanText := strings.TrimSpace(r.FormValue("laporan_text"))
	if laporanText == "" {
		http.Error(w, "Isi laporan wajib diisi", http.StatusBadRequest)
		return
	}

	item, err := models.GetPendingTransportReportByID(id, userNIP, userName)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Perjalanan tidak ditemukan atau sudah pernah dilaporkan", http.StatusNotFound)
			return
		}
		http.Error(w, "Gagal memvalidasi perjalanan", http.StatusInternalServerError)
		return
	}
	if item == nil {
		http.Error(w, "Perjalanan tidak ditemukan", http.StatusNotFound)
		return
	}

	file, header, err := r.FormFile("foto_laporan")
	if err != nil {
		http.Error(w, "Foto laporan wajib diunggah", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fotoPath, err := storeReportPhoto(file, header.Filename)
	if err != nil {
		var validationErr models.ValidationError
		if errors.As(err, &validationErr) {
			http.Error(w, validationErr.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, "Gagal menyimpan foto laporan", http.StatusInternalServerError)
		return
	}

	if err := models.SaveTransportReport(id, laporanText, fotoPath); err != nil {
		_ = storage.DeleteReportPhoto(fotoPath)

		var validationErr models.ValidationError
		if errors.As(err, &validationErr) {
			http.Error(w, validationErr.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, "Gagal menyimpan laporan", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusCreated, map[string]interface{}{
		"success": true,
		"message": "Laporan translok berhasil dikirim",
	})
}

func DownloadOwnTransportReportHandler(w http.ResponseWriter, r *http.Request) {
	userData := GetUserFromSession(r)
	userNIP, _ := userData["UserNIP"].(string)
	userName, _ := userData["UserName"].(string)

	id, ok := parseIDQuery(w, r)
	if !ok {
		return
	}

	owned, err := models.UserOwnsTransportReport(id, userNIP, userName)
	if err != nil {
		http.Error(w, "Gagal memvalidasi akses dokumen", http.StatusInternalServerError)
		return
	}
	if !owned {
		http.Error(w, "Dokumen tidak ditemukan", http.StatusNotFound)
		return
	}

	DownloadMatriksTransportReportHandler(w, r)
}

func SubmitSPPDReportHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(12 << 20); err != nil {
		http.Error(w, "Form upload tidak valid", http.StatusBadRequest)
		return
	}

	userData := GetUserFromSession(r)
	userNIP, _ := userData["UserNIP"].(string)
	userName, _ := userData["UserName"].(string)

	id, err := parsePositiveInt(r.FormValue("id_matriks"))
	if err != nil {
		http.Error(w, "ID perjalanan tidak valid", http.StatusBadRequest)
		return
	}

	item, err := models.GetPendingSPPDReportByID(id, userNIP, userName)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Perjalanan SPPD tidak ditemukan atau sudah pernah dilaporkan", http.StatusNotFound)
			return
		}
		http.Error(w, "Gagal memvalidasi perjalanan SPPD", http.StatusInternalServerError)
		return
	}
	if item == nil {
		http.Error(w, "Perjalanan SPPD tidak ditemukan", http.StatusNotFound)
		return
	}

	file, _, err := r.FormFile("foto_laporan")
	if err != nil {
		http.Error(w, "Foto laporan wajib diunggah", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fotoPath, err := storeReportPhoto(file, "lapor_sppd")
	if err != nil {
		var validationErr models.ValidationError
		if errors.As(err, &validationErr) {
			http.Error(w, validationErr.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, "Gagal menyimpan foto laporan", http.StatusInternalServerError)
		return
	}

	payload := models.SaveSPPDReportPayload{
		WaktuBerangkatMulai:     r.FormValue("waktu_berangkat_mulai"),
		WaktuBerangkatSelesai:   r.FormValue("waktu_berangkat_selesai"),
		IsianBerangkat:          r.FormValue("isian_berangkat"),
		WaktuPengawasan1Mulai:   r.FormValue("waktu_pengawasan1_mulai"),
		WaktuPengawasan1Selesai: r.FormValue("waktu_pengawasan1_selesai"),
		IsianPengawasan1:        r.FormValue("isian_pengawasan1"),
		WaktuIsomaMulai:         r.FormValue("waktu_isoma_mulai"),
		WaktuIsomaSelesai:       r.FormValue("waktu_isoma_selesai"),
		IsianIsoma:              firstNonEmpty(strings.TrimSpace(r.FormValue("isian_isoma")), "Istirahat Sholat dan Makan"),
		WaktuPengawasan2Mulai:   r.FormValue("waktu_pengawasan2_mulai"),
		WaktuPengawasan2Selesai: r.FormValue("waktu_pengawasan2_selesai"),
		IsianPengawasan2:        r.FormValue("isian_pengawasan2"),
		WaktuPulangMulai:        r.FormValue("waktu_pulang_mulai"),
		WaktuPulangSelesai:      r.FormValue("waktu_pulang_selesai"),
		IsianPulang:             r.FormValue("isian_pulang"),
		LaporanFotoPath:         fotoPath,
	}

	if err := models.SaveSPPDReport(id, payload); err != nil {
		_ = storage.DeleteReportPhoto(fotoPath)

		var validationErr models.ValidationError
		if errors.As(err, &validationErr) {
			http.Error(w, validationErr.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, "Gagal menyimpan laporan SPPD", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusCreated, map[string]interface{}{
		"success": true,
		"message": "Laporan SPPD berhasil dikirim",
	})
}

func DownloadOwnSPPDReportHandler(w http.ResponseWriter, r *http.Request) {
	userData := GetUserFromSession(r)
	userNIP, _ := userData["UserNIP"].(string)
	userName, _ := userData["UserName"].(string)

	id, ok := parseIDQuery(w, r)
	if !ok {
		return
	}

	owned, err := models.UserOwnsSPPDReport(id, userNIP, userName)
	if err != nil {
		http.Error(w, "Gagal memvalidasi akses dokumen", http.StatusInternalServerError)
		return
	}
	if !owned {
		http.Error(w, "Dokumen tidak ditemukan", http.StatusNotFound)
		return
	}

	DownloadMatriksSPPDReportHandler(w, r)
}

func storeReportPhoto(file multipartFile, originalName string) (string, error) {
	const maxSize = 10 * 1024 * 1024

	limitedReader := io.LimitReader(file, maxSize+1)
	content, err := io.ReadAll(limitedReader)
	if err != nil {
		return "", err
	}
	if int64(len(content)) > maxSize {
		return "", models.ValidationError{Message: "Ukuran foto maksimal 10MB"}
	}
	if len(content) == 0 {
		return "", models.ValidationError{Message: "File foto tidak boleh kosong"}
	}

	contentType := http.DetectContentType(content)
	ext := ""
	switch contentType {
	case "image/jpeg":
		ext = ".jpg"
	case "image/png":
		ext = ".png"
	case "image/gif":
		ext = ".gif"
	default:
		return "", models.ValidationError{Message: "Foto harus berupa JPG, PNG, atau GIF"}
	}

	return storage.SaveReportPhoto(content, ext)
}

type multipartFile interface {
	io.Reader
}

func ServeReportPhotoHandler(w http.ResponseWriter, r *http.Request) {
	ref := strings.TrimSpace(r.URL.Query().Get("ref"))
	if ref == "" {
		http.NotFound(w, r)
		return
	}

	content, contentType, err := storage.LoadReportPhoto(ref)
	if err != nil {
		log.Printf("Error loading report photo %q: %v", ref, err)
		http.NotFound(w, r)
		return
	}

	if strings.TrimSpace(contentType) != "" {
		w.Header().Set("Content-Type", contentType)
	}
	w.Header().Set("Cache-Control", "private, max-age=300")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(content)
}

func parsePositiveInt(value string) (int, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return 0, errors.New("empty")
	}

	var id int
	_, err := fmt.Sscanf(value, "%d", &id)
	if err != nil || id <= 0 {
		return 0, errors.New("invalid")
	}
	return id, nil
}

func writeJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}
