package middleware

import "net/http"

// EnableCORS menambahkan header agar frontend (mis. Vite/React) bisa mengakses API
func EnableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Izinkan origin React-mu (ubah jika perlu)
		w.Header().Set("Access-Control-Allow-Origin", "https://frontend-blog-react.vercel.app/")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
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
