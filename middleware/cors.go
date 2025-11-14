package middleware

import (
	"net/http"
	"strings"
)

// EnableCORS menambahkan header agar frontend (mis. Vite/React) bisa mengakses API
func EnableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		origin := r.Header.Get("Origin")

		// Cek apakah origin ada dalam daftar allowedOrigins
		if strings.Contains(origin, "vercel.app") || strings.Contains(origin, "103.174.114.55") {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}

		// Izinkan origin React-mu (ubah jika perlu)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Untuk preflight request (OPTIONS)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
