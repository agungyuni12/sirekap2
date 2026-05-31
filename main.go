package main

import (
	"encoding/base64"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"sirekap/internal/config"
	"sirekap/internal/database"
	"sirekap/internal/handlers"
	"sirekap/internal/middleware"
	"sirekap/internal/models"
	"sirekap/internal/storage"
)

// Template helper functions
var funcMap = template.FuncMap{
	// add adds two integers (for row numbering)
	"add": func(a, b int) int {
		return a + b
	},
	// sub subtracts two floats (for SBML excess calculation)
	"sub": func(a, b float64) float64 {
		return a - b
	},
	// gtFloat compares float64 with float64 (for SBML limit check)
	"gtFloat": func(a, b float64) bool {
		return a > b
	},
	// encodeLink encodes idsobat|bulan|tahun to base64 for URL obfuscation
	"encodeLink": func(idsobat, bulan, tahun string) string {
		data := fmt.Sprintf("%s|%s|%s", idsobat, bulan, tahun)
		return base64.URLEncoding.EncodeToString([]byte(data))
	},
	// initials extracts initials from a name (max 2 characters)
	"initials": func(name string) string {
		if name == "" {
			return "?"
		}
		parts := strings.Fields(name)
		if len(parts) >= 2 {
			return strings.ToUpper(string(parts[0][0])) + strings.ToUpper(string(parts[1][0]))
		}
		if len(name) >= 2 {
			return strings.ToUpper(name[:2])
		}
		return strings.ToUpper(string(name[0]))
	},
	// formatNumber formats a number with thousand separators (Indonesian style)
	"formatNumber": func(n interface{}) string {
		switch v := n.(type) {
		case int:
			return formatNumberInt(v)
		case float64:
			return formatNumberFloat(v)
		case int64:
			return formatNumberInt(int(v))
		default:
			return fmt.Sprintf("%v", n)
		}
	},
	// formatCurrency formats a number as Indonesian Rupiah
	"formatCurrency": func(n interface{}) string {
		switch v := n.(type) {
		case int:
			return formatNumberInt(v)
		case float64:
			return formatNumberFloat(v)
		case int64:
			return formatNumberInt(int(v))
		default:
			return fmt.Sprintf("%v", n)
		}
	},
}

// formatNumberInt formats an integer with thousand separators
func formatNumberInt(n int) string {
	str := fmt.Sprintf("%d", n)
	return addThousandSeparators(str)
}

// formatNumberFloat formats a float with thousand separators
func formatNumberFloat(n float64) string {
	// For whole numbers, don't show decimals
	if n == float64(int64(n)) {
		return formatNumberInt(int(n))
	}
	// Otherwise format with 2 decimal places
	str := fmt.Sprintf("%.2f", n)
	parts := strings.Split(str, ".")
	parts[0] = addThousandSeparators(parts[0])
	return strings.Join(parts, ",")
}

// addThousandSeparators adds dots as thousand separators (Indonesian format)
func addThousandSeparators(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	var result strings.Builder
	start := n % 3
	if start == 0 {
		start = 3
	}
	result.WriteString(s[:start])
	for i := start; i < n; i += 3 {
		result.WriteString(".")
		result.WriteString(s[i : i+3])
	}
	return result.String()
}

func main() {
	// Load configuration
	cfg := config.Load()

	// Connect to database
	if err := database.Connect(&cfg.DB); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()
	log.Println("Database connected successfully")

	// Initialize session store
	handlers.InitSession(cfg.Server.SessionKey, cfg.Server.SessionSecure)
	if err := storage.Init(cfg.Storage); err != nil {
		log.Fatalf("Failed to initialize object storage: %v", err)
	}

	// Initialize Router
	r := mux.NewRouter()
	r.Use(middleware.SecurityHeaders)

	// Static Assets
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Template Helper with FuncMap
	tmpl := func(w http.ResponseWriter, templateName string, data interface{}, files ...string) {
		t, err := template.New("").Funcs(funcMap).ParseFiles(files...)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Execute the base template (first file's name)
		err = t.ExecuteTemplate(w, "base.html", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	// Helper to merge user session data with page data
	mergeUserData := func(r *http.Request, pageData map[string]interface{}) map[string]interface{} {
		userData := handlers.GetUserFromSession(r)
		for k, v := range userData {
			pageData[k] = v
		}
		return pageData
	}

	// Public Routes
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	})

	r.HandleFunc("/login", handlers.LoginHandler).Methods("GET", "POST")
	r.HandleFunc("/logout", handlers.LogoutHandler)

	// Protected Routes - Dashboard (Admin)
	r.HandleFunc("/dashboard", middleware.RequireAdmin(func(w http.ResponseWriter, r *http.Request) {
		stats, err := models.GetDashboardStats()
		if err != nil {
			log.Printf("Error loading dashboard stats: %v", err)
			http.Error(w, "Failed to load dashboard", http.StatusInternalServerError)
			return
		}

		data := mergeUserData(r, map[string]interface{}{
			"Title":       "Dashboard",
			"PageTitle":   "Dashboard",
			"ShowSidebar": true,
			"ActivePage":  "dashboard",
			"Stats":       stats,
		})

		tmpl(w, "layouts/base.html", data,
			"templates/layouts/base.html",
			"templates/partials/sidebar.html",
			"templates/dashboard/index.html",
		)
	}))

	r.HandleFunc("/dashboard/input", middleware.RequireAdmin(func(w http.ResponseWriter, r *http.Request) {
		data := mergeUserData(r, map[string]interface{}{
			"Title":       "Input Kegiatan",
			"PageTitle":   "Input Kegiatan",
			"ShowSidebar": true,
			"ActivePage":  "input",
		})

		tmpl(w, "layouts/base.html", data,
			"templates/layouts/base.html",
			"templates/partials/sidebar.html",
			"templates/dashboard/input.html",
		)
	}))

	r.HandleFunc("/dashboard/rekap", middleware.RequireAdmin(func(w http.ResponseWriter, r *http.Request) {
		data := mergeUserData(r, map[string]interface{}{
			"Title":       "Rekapitulasi",
			"PageTitle":   "Rekapitulasi",
			"ShowSidebar": true,
			"ActivePage":  "rekap",
		})

		tmpl(w, "layouts/base.html", data,
			"templates/layouts/base.html",
			"templates/partials/sidebar.html",
			"templates/dashboard/rekap.html",
		)
	}))

	r.HandleFunc("/dashboard/kegiatan", middleware.RequireAdmin(func(w http.ResponseWriter, r *http.Request) {
		data := mergeUserData(r, map[string]interface{}{
			"Title":       "Master Kegiatan",
			"PageTitle":   "Master Kegiatan",
			"ShowSidebar": true,
			"ActivePage":  "kegiatan",
		})

		tmpl(w, "layouts/base.html", data,
			"templates/layouts/base.html",
			"templates/partials/sidebar.html",
			"templates/dashboard/kegiatan.html",
		)
	}))

	r.HandleFunc("/dashboard/surat", middleware.RequireAdmin(func(w http.ResponseWriter, r *http.Request) {
		data := mergeUserData(r, map[string]interface{}{
			"Title":       "Nomor Surat",
			"PageTitle":   "Nomor Surat",
			"ShowSidebar": true,
			"ActivePage":  "surat",
		})

		tmpl(w, "layouts/base.html", data,
			"templates/layouts/base.html",
			"templates/partials/sidebar.html",
			"templates/dashboard/surat.html",
		)
	}))

	r.HandleFunc("/dashboard/matriks", middleware.RequireAdmin(func(w http.ResponseWriter, r *http.Request) {
		data := mergeUserData(r, map[string]interface{}{
			"Title":       "Input Matriks",
			"PageTitle":   "Input Matriks",
			"ShowSidebar": true,
			"ActivePage":  "matriks",
		})

		tmpl(w, "layouts/base.html", data,
			"templates/layouts/base.html",
			"templates/partials/sidebar.html",
			"templates/dashboard/matriks.html",
		)
	}))

	r.HandleFunc("/dashboard/monitoring", middleware.RequireAdmin(func(w http.ResponseWriter, r *http.Request) {
		data := mergeUserData(r, map[string]interface{}{
			"Title":       "Monitoring Matriks",
			"PageTitle":   "Monitoring Matriks",
			"ShowSidebar": true,
			"ActivePage":  "monitoring",
		})

		tmpl(w, "layouts/base.html", data,
			"templates/layouts/base.html",
			"templates/partials/sidebar.html",
			"templates/dashboard/monitoring.html",
		)
	}))

	r.HandleFunc("/dashboard/kalender", middleware.RequireAdmin(func(w http.ResponseWriter, r *http.Request) {
		data := mergeUserData(r, map[string]interface{}{
			"Title":       "Kalender Matriks",
			"PageTitle":   "Kalender Matriks",
			"ShowSidebar": true,
			"ActivePage":  "kalender",
		})

		tmpl(w, "layouts/base.html", data,
			"templates/layouts/base.html",
			"templates/partials/sidebar.html",
			"templates/dashboard/kalender.html",
		)
	}))

	r.HandleFunc("/dashboard/administrasi", middleware.RequireAdmin(func(w http.ResponseWriter, r *http.Request) {
		data := mergeUserData(r, map[string]interface{}{
			"Title":       "Administrasi Matriks",
			"PageTitle":   "Administrasi Matriks",
			"ShowSidebar": true,
			"ActivePage":  "administrasi",
		})

		tmpl(w, "layouts/base.html", data,
			"templates/layouts/base.html",
			"templates/partials/sidebar.html",
			"templates/dashboard/administrasi.html",
		)
	}))

	r.HandleFunc("/dashboard/detail", middleware.RequireAdmin(func(w http.ResponseWriter, r *http.Request) {
		// Decode base64 parameter
		encodedData := r.URL.Query().Get("d")
		var idsobat, bulan, tahun string

		if encodedData != "" {
			decoded, err := base64.URLEncoding.DecodeString(encodedData)
			if err == nil {
				parts := strings.Split(string(decoded), "|")
				if len(parts) == 3 {
					idsobat = parts[0]
					bulan = parts[1]
					tahun = parts[2]
				}
			}
		}

		// Fallback to direct params (for backward compatibility)
		if idsobat == "" {
			idsobat = r.URL.Query().Get("idsobat")
			bulan = r.URL.Query().Get("bulan")
			tahun = r.URL.Query().Get("tahun")
		}

		// Get kegiatan list for this mitra
		kegiatan, err := models.GetMitraKegiatan(idsobat, bulan, tahun)
		if err != nil {
			log.Printf("Error getting kegiatan: %v", err)
		}

		// Get mitra name from first kegiatan or database
		var namaMitra string
		var totalHonor float64
		if len(kegiatan) > 0 {
			// Get mitra info from database
			mitraInfo, _ := models.GetMitraByIDSobat(idsobat)
			if mitraInfo != nil {
				namaMitra = mitraInfo.NamaMitra
			}
			// Calculate total honor
			for _, k := range kegiatan {
				totalHonor += k.Honor
			}
		}

		data := mergeUserData(r, map[string]interface{}{
			"Title":       "Detail Kegiatan",
			"PageTitle":   "Detail Kegiatan",
			"ShowSidebar": true,
			"ActivePage":  "rekap",
			"IDSobat":     idsobat,
			"NamaMitra":   namaMitra,
			"Bulan":       bulan,
			"Tahun":       tahun,
			"Kegiatan":    kegiatan,
			"TotalHonor":  totalHonor,
		})

		tmpl(w, "layouts/base.html", data,
			"templates/layouts/base.html",
			"templates/partials/sidebar.html",
			"templates/dashboard/detail.html",
		)
	}))

	// Protected Routes - User (Pengguna)
	r.HandleFunc("/lapor", middleware.RequireAuth(func(w http.ResponseWriter, r *http.Request) {
		data := mergeUserData(r, map[string]interface{}{
			"Title":       "Lapor Translok",
			"PageTitle":   "Lapor Translok",
			"ShowSidebar": true,
			"ActivePage":  "lapor",
		})

		tmpl(w, "layouts/base.html", data,
			"templates/layouts/base.html",
			"templates/partials/sidebar.html",
			"templates/user/lapor.html",
		)
	}))

	r.HandleFunc("/lapor/history", middleware.RequireAuth(func(w http.ResponseWriter, r *http.Request) {
		data := mergeUserData(r, map[string]interface{}{
			"Title":       "Riwayat Lapor Translok",
			"PageTitle":   "Riwayat Lapor Translok",
			"ShowSidebar": true,
			"ActivePage":  "lapor-history",
		})

		tmpl(w, "layouts/base.html", data,
			"templates/layouts/base.html",
			"templates/partials/sidebar.html",
			"templates/user/lapor_history.html",
		)
	}))

	r.HandleFunc("/lapor/sppd", middleware.RequireAuth(func(w http.ResponseWriter, r *http.Request) {
		data := mergeUserData(r, map[string]interface{}{
			"Title":       "Lapor SPPD",
			"PageTitle":   "Lapor SPPD",
			"ShowSidebar": true,
			"ActivePage":  "lapor-sppd",
		})

		tmpl(w, "layouts/base.html", data,
			"templates/layouts/base.html",
			"templates/partials/sidebar.html",
			"templates/user/lapor_sppd.html",
		)
	}))

	r.HandleFunc("/lapor/sppd/history", middleware.RequireAuth(func(w http.ResponseWriter, r *http.Request) {
		data := mergeUserData(r, map[string]interface{}{
			"Title":       "Riwayat Lapor SPPD",
			"PageTitle":   "Riwayat Lapor SPPD",
			"ShowSidebar": true,
			"ActivePage":  "lapor-sppd-history",
		})

		tmpl(w, "layouts/base.html", data,
			"templates/layouts/base.html",
			"templates/partials/sidebar.html",
			"templates/user/lapor_sppd_history.html",
		)
	}))

	// API Routes (protected - admin only)
	r.HandleFunc("/api/kegiatan/search", middleware.RequireAdmin(handlers.SearchKegiatanHandler)).Methods("GET")
	r.HandleFunc("/api/kegiatan/detail", middleware.RequireAdmin(handlers.GetKegiatanDetailHandler)).Methods("GET")
	r.HandleFunc("/api/master-kegiatan", middleware.RequireAdmin(handlers.ListKegiatanMasterHandler)).Methods("GET")
	r.HandleFunc("/api/master-kegiatan", middleware.RequireAdmin(handlers.CreateKegiatanMasterHandler)).Methods("POST")
	r.HandleFunc("/api/master-kegiatan/{id}", middleware.RequireAdmin(handlers.GetMasterKegiatanHandler)).Methods("GET")
	r.HandleFunc("/api/master-kegiatan/{id}", middleware.RequireAdmin(handlers.UpdateMasterKegiatanHandler)).Methods("PUT")
	r.HandleFunc("/api/master-kegiatan/{id}", middleware.RequireAdmin(handlers.DeleteMasterKegiatanHandler)).Methods("DELETE")
	r.HandleFunc("/api/surat", middleware.RequireAdmin(handlers.ListSuratHandler)).Methods("GET")
	r.HandleFunc("/api/surat", middleware.RequireAdmin(handlers.CreateSuratManualHandler)).Methods("POST")
	r.HandleFunc("/api/matriks/administrasi", middleware.RequireAdmin(handlers.ListMatriksAdministrasiHandler)).Methods("GET")
	r.HandleFunc("/api/matriks/administrasi/peserta", middleware.RequireAdmin(handlers.ListMatriksAdministrasiPesertaHandler)).Methods("GET")
	r.HandleFunc("/api/matriks/print/translok", middleware.RequireAdmin(handlers.DownloadMatriksTransportReportHandler)).Methods("GET")
	r.HandleFunc("/api/matriks/print/sppd", middleware.RequireAdmin(handlers.DownloadMatriksSPPDReportHandler)).Methods("GET")
	r.HandleFunc("/api/matriks/print/pernyataan", middleware.RequireAdmin(handlers.DownloadMatriksPernyataanKendaraanHandler)).Methods("GET")
	r.HandleFunc("/api/matriks/print/kwitansi/translok", middleware.RequireAdmin(handlers.DownloadMatriksKwitansiTranslokHandler)).Methods("GET")
	r.HandleFunc("/api/matriks/print/kwitansi/sppd", middleware.RequireAdmin(handlers.DownloadMatriksKwitansiSPPDHandler)).Methods("GET")
	r.HandleFunc("/api/matriks/calendar", middleware.RequireAdmin(handlers.ListMatriksCalendarHandler)).Methods("GET")
	r.HandleFunc("/api/matriks", middleware.RequireAdmin(handlers.ListMatriksHandler)).Methods("GET")
	r.HandleFunc("/api/matriks/import/template", middleware.RequireAdmin(handlers.DownloadMatriksImportTemplateHandler)).Methods("GET")
	r.HandleFunc("/api/matriks/kegiatan/search", middleware.RequireAdmin(handlers.SearchDaftarTranslokHandler)).Methods("GET")
	r.HandleFunc("/api/pegawai/search", middleware.RequireAdmin(handlers.SearchPegawaiHandler)).Methods("GET")
	r.HandleFunc("/api/matriks/check-conflict", middleware.RequireAdmin(handlers.CheckMatriksConflictHandler)).Methods("POST")
	r.HandleFunc("/api/matriks", middleware.RequireAdmin(handlers.CreateMatriksHandler)).Methods("POST")
	r.HandleFunc("/api/matriks/import", middleware.RequireAdmin(handlers.ImportMatriksHandler)).Methods("POST")
	r.HandleFunc("/api/matriks/{id}", middleware.RequireAdmin(handlers.GetMatriksHandler)).Methods("GET")
	r.HandleFunc("/api/matriks/{id}", middleware.RequireAdmin(handlers.UpdateMatriksHandler)).Methods("PUT")
	r.HandleFunc("/api/matriks/{id}", middleware.RequireAdmin(handlers.DeleteMatriksHandler)).Methods("DELETE")
	r.HandleFunc("/api/mitra/search", middleware.RequireAdmin(handlers.SearchMitraHandler)).Methods("GET")
	r.HandleFunc("/api/mitra/submit", middleware.RequireAdmin(handlers.SubmitMitraHandler)).Methods("POST")
	r.HandleFunc("/api/rekap/filter", middleware.RequireAdmin(handlers.FilterRekapHandler)).Methods("GET")
	r.HandleFunc("/api/rekap/kegiatan", middleware.RequireAdmin(handlers.GetMitraKegiatanHandler)).Methods("GET")

	// API Routes (protected - user)
	r.HandleFunc("/api/lapor/translok/pending", middleware.RequireAuth(handlers.ListPendingTransportReportsHandler)).Methods("GET")
	r.HandleFunc("/api/lapor/translok/history", middleware.RequireAuth(handlers.ListSubmittedTransportReportsHandler)).Methods("GET")
	r.HandleFunc("/api/lapor/translok/submit", middleware.RequireAuth(handlers.SubmitTransportReportHandler)).Methods("POST")
	r.HandleFunc("/api/lapor/sppd/pending", middleware.RequireAuth(handlers.ListPendingSPPDReportsHandler)).Methods("GET")
	r.HandleFunc("/api/lapor/sppd/history", middleware.RequireAuth(handlers.ListSubmittedSPPDReportsHandler)).Methods("GET")
	r.HandleFunc("/api/lapor/sppd/submit", middleware.RequireAuth(handlers.SubmitSPPDReportHandler)).Methods("POST")
	r.HandleFunc("/lapor/print/translok", middleware.RequireAuth(handlers.DownloadOwnTransportReportHandler)).Methods("GET")
	r.HandleFunc("/lapor/print/sppd", middleware.RequireAuth(handlers.DownloadOwnSPPDReportHandler)).Methods("GET")
	r.HandleFunc("/media/report-photo", middleware.RequireAuth(handlers.ServeReportPhotoHandler)).Methods("GET")

	// Kegiatan CRUD
	r.HandleFunc("/api/rekap/{id}", middleware.RequireAdmin(handlers.GetKegiatanHandler)).Methods("GET")
	r.HandleFunc("/api/rekap/{id}", middleware.RequireAdmin(handlers.UpdateKegiatanHandler)).Methods("PUT")
	r.HandleFunc("/api/rekap/{id}", middleware.RequireAdmin(handlers.DeleteKegiatanHandler)).Methods("DELETE")

	// SPK
	r.HandleFunc("/api/rekap/spk/bulk-options", middleware.RequireAdmin(handlers.ListBulkSPKOptionsHandler)).Methods("GET")
	r.HandleFunc("/api/rekap/spk/create", middleware.RequireAdmin(handlers.CreateSPKHandler)).Methods("POST")
	r.HandleFunc("/api/rekap/spk/create-bulk", middleware.RequireAdmin(handlers.CreateBulkSPKHandler)).Methods("POST")
	r.HandleFunc("/api/rekap/spk/download", middleware.RequireAdmin(handlers.DownloadSPKHandler)).Methods("GET")
	r.HandleFunc("/api/rekap/spk/download-all", middleware.RequireAdmin(handlers.DownloadFilteredSPKZipHandler)).Methods("GET")

	// Halaman SPK SE2026 - Cetak
	r.HandleFunc("/dashboard/se2026", middleware.RequireAdmin(func(w http.ResponseWriter, r *http.Request) {
		data := mergeUserData(r, map[string]interface{}{
			"Title":       "Cetak SPK SE2026",
			"PageTitle":   "SPK Sensus Ekonomi 2026",
			"ShowSidebar": true,
			"ActivePage":  "se2026",
		})
		tmpl(w, "layouts/base.html", data,
			"templates/layouts/base.html",
			"templates/partials/sidebar.html",
			"templates/dashboard/se2026.html",
		)
	}))

	// Halaman Import Petugas SE2026
	r.HandleFunc("/dashboard/se2026/import", middleware.RequireAdmin(func(w http.ResponseWriter, r *http.Request) {
		data := mergeUserData(r, map[string]interface{}{
			"Title":       "Import Petugas SE2026",
			"PageTitle":   "Import Petugas SE2026",
			"ShowSidebar": true,
			"ActivePage":  "se2026-import",
		})
		tmpl(w, "layouts/base.html", data,
			"templates/layouts/base.html",
			"templates/partials/sidebar.html",
			"templates/dashboard/se2026_import.html",
		)
	}))

	// API SPK + Import Sensus Ekonomi 2026
	r.HandleFunc("/api/rekap/spk/se2026/list", middleware.RequireAdmin(handlers.ListRekapSE2026Handler)).Methods("GET")
	r.HandleFunc("/api/rekap/spk/se2026/import/template", middleware.RequireAdmin(handlers.DownloadTemplateSE2026Handler)).Methods("GET")
	r.HandleFunc("/api/rekap/spk/se2026/import", middleware.RequireAdmin(handlers.ImportPetugasSE2026Handler)).Methods("POST")
	r.HandleFunc("/api/rekap/spk/se2026/download", middleware.RequireAdmin(handlers.DownloadSPKSE2026Handler)).Methods("GET")
	r.HandleFunc("/api/rekap/spk/se2026/download-all", middleware.RequireAdmin(handlers.DownloadAllSPKSE2026Handler)).Methods("GET")
	r.HandleFunc("/api/rekap/spk/se2026/update-sls", middleware.RequireAdmin(handlers.UpdateJumlahSLSHandler)).Methods("POST")

	// BAST
	r.HandleFunc("/api/rekap/bast/bulk-options", middleware.RequireAdmin(handlers.ListBulkBASTOptionsHandler)).Methods("GET")
	r.HandleFunc("/api/rekap/bast/create", middleware.RequireAdmin(handlers.CreateBASTHandler)).Methods("POST")
	r.HandleFunc("/api/rekap/bast/create-bulk", middleware.RequireAdmin(handlers.CreateBulkBASTHandler)).Methods("POST")
	r.HandleFunc("/api/rekap/bast/download", middleware.RequireAdmin(handlers.DownloadBASTHandler)).Methods("GET")
	r.HandleFunc("/api/rekap/bast/download-all", middleware.RequireAdmin(handlers.DownloadFilteredBASTZipHandler)).Methods("GET")

	// Server Configuration
	srv := &http.Server{
		Handler:      r,
		Addr:         cfg.Server.Addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Server starting on http://%s", cfg.Server.Addr)
	log.Fatal(srv.ListenAndServe())
}
