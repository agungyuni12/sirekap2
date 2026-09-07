package middleware

import (
	"net/http"

	"sirekap/internal/handlers"
)

// RequireAuth middleware ensures user is authenticated
func RequireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !handlers.IsAuthenticated(r) {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next(w, r)
	}
}

// RequireAdmin middleware ensures user is admin
func RequireAdmin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !handlers.IsAuthenticated(r) {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		if handlers.GetUserLevel(r) != "admin" {
			http.Redirect(w, r, "/lapor", http.StatusSeeOther)
			return
		}

		next(w, r)
	}
}

// RequireNotPMLMitra middleware ensures the "pml_mitra" role (a mitra acting as
// PML, restricted to the Penilaian Kinerja PPL + Lapor Translok Mandiri menus)
// cannot reach routes outside that scope, such as SPPD reporting.
func RequireNotPMLMitra(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !handlers.IsAuthenticated(r) {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		if handlers.GetUserLevel(r) == "pml_mitra" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		next(w, r)
	}
}

// RequireNotExternalDaftarPenilaian blocks Daftar Penilaian (everyone's
// performance history) from "pml_mitra" and from "pengguna" accounts whose
// email isn't @bps.go.id (see handlers.IsExternalAccount) — matches the
// sidebar/page-button visibility, so direct URL access can't bypass it.
// Deliberately narrower than RequireNotPMLMitra: it does NOT touch Penilaian
// scoring/assignment access, which stays gated by the per-periode Kelola
// Petugas assignment regardless of email — an external account can still be
// PPL, not PML, so it must never be granted PML-Mitra-style capabilities.
func RequireNotExternalDaftarPenilaian(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !handlers.IsAuthenticated(r) {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		if handlers.GetUserLevel(r) == "pml_mitra" || handlers.IsExternalAccount(r) {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		next(w, r)
	}
}
