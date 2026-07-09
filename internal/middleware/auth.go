package middleware

import (
	"context"
	"marketplace-api/internal/service"
	"net/http"
	"strings"
)

type claveContexto string

const ClaveUsuarioID claveContexto = "usuarioID"
const ClaveRol claveContexto = "rol"

func Auth(auth *service.AuthService) func(http.Handler) http.Handler {
	return func(siguiente http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			encabezado := r.Header.Get("Authorization")
			partes := strings.SplitN(encabezado, " ", 2)
			if len(partes) != 2 || !strings.EqualFold(partes[0], "Bearer") {
				responderNoAutorizado(w)
				return
			}
			usuarioID, err := auth.ValidarToken(partes[1])
			if err != nil {
				responderNoAutorizado(w)
				return
			}
			rol, err := auth.RolDesdeToken(partes[1])
			if err != nil {
				responderNoAutorizado(w)
				return
			}
			ctx := context.WithValue(r.Context(), ClaveUsuarioID, usuarioID)
			ctx = context.WithValue(ctx, ClaveRol, rol)
			siguiente.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func RolPermitido(roles ...string) func(http.Handler) http.Handler {
	permitidos := map[string]bool{}
	for _, role := range roles {
		permitidos[strings.ToLower(strings.TrimSpace(role))] = true
	}
	return func(siguiente http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rol, _ := r.Context().Value(ClaveRol).(string)
			if !permitidos[strings.ToLower(strings.TrimSpace(rol))] {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusForbidden)
				_, _ = w.Write([]byte(`{"error":"rol no autorizado"}`))
				return
			}
			siguiente.ServeHTTP(w, r)
		})
	}
}

func responderNoAutorizado(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	_, _ = w.Write([]byte(`{"error":"token ausente o invalido"}`))
}
