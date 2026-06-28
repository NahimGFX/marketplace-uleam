package routes

import (
	"marketplace-api/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func AuthRoutes(r chi.Router, s *handlers.Server) {
	r.Post("/auth/register", s.Registrar)
	r.Post("/auth/login", s.Login)
}
