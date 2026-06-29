package routes

import (
	"marketplace-api/internal/handlers"
	"marketplace-api/internal/middleware"
	"marketplace-api/internal/service"

	"github.com/go-chi/chi/v5"
)

func PerfilRoutes(r chi.Router, s *handlers.Server, authSvc *service.AuthService) {

	r.Group(func(r chi.Router) {
		r.Use(middleware.Auth(authSvc))
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
	})
}
