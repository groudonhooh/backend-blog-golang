package middleware

import "net/http"

var allowedOrigins = []string{
	"http://localhost:5173",
	"https://frontend-blog-react.vercel.app",
	"https://frontend-blog-react.vercel.app/",
	"http://103.174.114.55",
	"http://103.174.114.55/",
}

// EnableCORS menambahkan header agar frontend (mis. Vite/React) bisa mengakses API
func EnableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		origin := r.Header.Get("Origin")

		// Cek apakah origin ada dalam daftar allowedOrigins
		for _, allowed := range allowedOrigins {
			if origin == allowed {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				break
			}
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
