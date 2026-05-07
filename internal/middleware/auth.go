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
