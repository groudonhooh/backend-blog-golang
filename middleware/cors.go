package middleware

import (
	"net/http"
)

var allowedOrigins = []string{
	"http://localhost:5173", // Ganti dengan origin React-mu
	"https://frontend-blog-vercel.vercel.app",
	"https://api-bloghub.my.id",
}

func isAllowedOrigin(origin string) bool {
	for _, o := range allowedOrigins {
		if o == origin {
			return true
		}
	}
	return false
}

// EnableCORS menambahkan header agar frontend (mis. Vite/React) bisa mengakses API
func EnableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		origin := r.Header.Get("Origin")

		// Cek apakah origin ada dalam daftar allowedOrigins
		if isAllowedOrigin(origin) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
		}

		// Izinkan origin React-mu (ubah jika perlu)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Untuk preflight request (OPTIONS)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
