package middleware

import (
	"net/http"
	"strings"
)

func RequiereToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" || !strings.HasPrefix(header, "Bearer ") {
			http.Error(w, `{"error":"no autorizado"}`, http.StatusUnauthorized)
			return
		}
		token := strings.TrimPrefix(header, "Bearer ")
		if strings.TrimSpace(token) == "" {
			http.Error(w, `{"error":"no autorizado"}`, http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

//es para el test 401
