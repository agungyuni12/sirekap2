package handlers

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/sessions"
	"sirekap/internal/models"
	"sirekap/internal/ratelimit"
)

var store *sessions.CookieStore

// InitSession initializes the session store with the given key and secure flag
func InitSession(key string, secure bool) {
	store = sessions.NewCookieStore([]byte(key))
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400, // 1 day
		HttpOnly: true,
		Secure:   secure,
		SameSite: http.SameSiteStrictMode,
	}
}

// GetStore returns the session store
func GetStore() *sessions.CookieStore {
	return store
}

// LoginPageData holds data for the login template
type LoginPageData struct {
	Error   string
	Success string
}

// LoginHandler handles GET and POST for /login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "sirekap-session")

	// If already logged in, redirect
	if userID, ok := session.Values["user_id"].(int); ok && userID > 0 {
		level, _ := session.Values["level"].(string)
		redirectByLevel(w, r, level)
		return
	}

	if r.Method == http.MethodGet {
		renderLogin(w, LoginPageData{})
		return
	}

	// POST - Process login
	clientIP := strings.Split(r.RemoteAddr, ":")[0]
	if fwd := r.Header.Get("X-Forwarded-For"); fwd != "" {
		clientIP = strings.TrimSpace(strings.SplitN(fwd, ",", 2)[0])
	}

	if ratelimit.IsBlocked(clientIP) {
		renderLogin(w, LoginPageData{Error: "Terlalu banyak percobaan login. Coba lagi dalam 15 menit."})
		return
	}

	if err := r.ParseForm(); err != nil {
		renderLogin(w, LoginPageData{Error: "Terjadi kesalahan pada sistem."})
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		renderLogin(w, LoginPageData{Error: "Username dan password harus diisi."})
		return
	}

	// Find user
	user, err := models.GetUserByUsernameOrEmail(username)
	if err != nil {
		if err == sql.ErrNoRows {
			ratelimit.RecordFailure(clientIP)
			renderLogin(w, LoginPageData{Error: "Login Gagal. Username atau password yang Anda masukkan salah."})
			return
		}
		log.Printf("Database error: %v", err)
		renderLogin(w, LoginPageData{Error: "Terjadi kesalahan pada sistem. Silakan coba lagi."})
		return
	}

	// Verify password
	if !user.CheckPassword(password) {
		ratelimit.RecordFailure(clientIP)
		renderLogin(w, LoginPageData{Error: "Login Gagal. Username atau password yang Anda masukkan salah."})
		return
	}

	ratelimit.Reset(clientIP)

	// Login successful - save to session
	session.Values["user_id"] = user.ID
	session.Values["username"] = user.Username
	session.Values["nama"] = user.Nama
	session.Values["nip"] = user.NIP
	session.Values["email"] = user.Email
	session.Values["level"] = user.Level
	session.Values["initials"] = user.GetInitials()

	if err := session.Save(r, w); err != nil {
		log.Printf("Session save error: %v", err)
		renderLogin(w, LoginPageData{Error: "Terjadi kesalahan pada sistem."})
		return
	}

	// Redirect based on role
	redirectByLevel(w, r, user.Level)
}

// LogoutHandler handles logout
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "sirekap-session")
	
	// Clear session
	session.Values = make(map[interface{}]interface{})
	session.Options.MaxAge = -1
	session.Save(r, w)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// GetUserFromSession retrieves user data from session
func GetUserFromSession(r *http.Request) map[string]interface{} {
	session, _ := store.Get(r, "sirekap-session")
	
	data := make(map[string]interface{})
	
	if userID, ok := session.Values["user_id"].(int); ok {
		data["UserID"] = userID
		data["UserName"] = session.Values["nama"]
		data["Username"] = session.Values["username"]
		data["UserNIP"] = session.Values["nip"]
		data["UserEmail"] = session.Values["email"]
		data["UserRole"] = session.Values["level"]
		data["UserInitials"] = session.Values["initials"]
		data["IsLoggedIn"] = true
		// See IsExternalAccount below — level "pengguna" covers both organik
		// (BPS staff) and external mitra, split only by email domain.
		// Presentation-only (e.g. hide Daftar Penilaian); never used to grant
		// PML-Mitra-style capabilities.
		data["IsExternalAccount"] = IsExternalAccount(r)
	} else {
		data["IsLoggedIn"] = false
	}
	
	return data
}

// IsAuthenticated checks if user is authenticated
func IsAuthenticated(r *http.Request) bool {
	session, _ := store.Get(r, "sirekap-session")
	userID, ok := session.Values["user_id"].(int)
	return ok && userID > 0
}

// GetUserLevel returns the user's level from session
func GetUserLevel(r *http.Request) string {
	session, _ := store.Get(r, "sirekap-session")
	if level, ok := session.Values["level"].(string); ok {
		return level
	}
	return ""
}

// IsExternalAccount reports whether the current session belongs to a
// "pengguna"-level account whose email isn't @bps.go.id — see
// models.User.IsExternalAccount for why this exists and its limits
// (presentation-only, never a stand-in for PML-Mitra capabilities).
func IsExternalAccount(r *http.Request) bool {
	session, _ := store.Get(r, "sirekap-session")
	level, _ := session.Values["level"].(string)
	email, _ := session.Values["email"].(string)
	return level == "pengguna" && !strings.HasSuffix(strings.ToLower(email), "@bps.go.id")
}

func renderLogin(w http.ResponseWriter, data LoginPageData) {
	tmpl, err := template.ParseFiles("templates/login.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}

func redirectByLevel(w http.ResponseWriter, r *http.Request, level string) {
	switch level {
	case "admin":
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	case "pml_mitra":
		http.Redirect(w, r, "/penilaian", http.StatusSeeOther)
	default:
		http.Redirect(w, r, "/lapor", http.StatusSeeOther)
	}
}
