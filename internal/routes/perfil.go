//go:build ignore

package routes

import (
	"marketplace-api/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func PerfilRoutes(r chi.Router, s *handlers.Server) {

	// Users
	r.Get("/users", s.ListarUsers)
	r.Post("/users", s.CrearUser)
	r.Get("/users/{id}", s.ObtenerUser)
	r.Put("/users/{id}", s.ActualizarUser)
	r.Delete("/users/{id}", s.BorrarUser)

	// Reviews
	r.Get("/reviews", s.ListarReviews)
	r.Get("/reviews/{id}", s.ObtenerReview)
	r.Post("/reviews", s.CrearReview)
	r.Put("/reviews/{id}", s.ActualizarReview)
	r.Delete("/reviews/{id}", s.BorrarReview)

	// Badges
	r.Get("/badges", s.ListarBadges)
	r.Get("/badges/{id}", s.ObtenerBadge)
	r.Post("/badges", s.CrearBadge)
	r.Put("/badges/{id}", s.ActualizarBadge)
	r.Delete("/badges/{id}", s.BorrarBadge)
}
