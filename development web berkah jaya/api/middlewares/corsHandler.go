package middlewares

import (
	"net/http"
)

func CorsMiddlewares(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		
		allowedOrigins := []string{
			"http://localhost:3000",
		}

		for _, origns := range allowedOrigins {
			if origin == origns {
				w.Header().Set("Control-Access-Allow-Origin", origns)
				w.Header().Set("Control-Access-Allow-Methods", "POST, GET, PUT, UPDATE, DELETE, OPTIONS")
				w.Header().Set("Control-Access-Allow-Header", "X-Requested-With, Content-Type, Authorization")
				w.Header().Set("Control-Access-Allow-Credentials", "true")
			}
		}

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}