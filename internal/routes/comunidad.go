package routes

import (
	"marketplace-api/internal/handlers"
	"marketplace-api/internal/middleware"
	"marketplace-api/internal/service"

	"github.com/go-chi/chi/v5"
)

func ComunidadRoutes(
	r chi.Router,
	s *handlers.Server,
	authSvc *service.AuthService,
) {

	r.Group(func(r chi.Router) {
		r.Use(middleware.Auth(authSvc))

		r.Get("/messages", s.ListarMessages)
		r.Get("/messages/{id}", s.ObtenerMessage)
		r.Post("/messages", s.CrearMessage)
		r.Put("/messages/{id}", s.ActualizarMessage)
		r.Delete("/messages/{id}", s.BorrarMessage)

		r.Get("/missions", s.ListarMissions)
		r.Get("/missions/{id}", s.ObtenerMision)
		r.With(middleware.RolPermitido("admin")).Post("/missions", s.CrearMision)
		r.With(middleware.RolPermitido("admin")).Put("/missions/{id}", s.ActualizarMision)
		r.With(middleware.RolPermitido("admin")).Delete("/missions/{id}", s.BorrarMision)

		r.Get("/usermissions", s.ListarUsermissions)
		r.Get("/usermissions/{id}", s.ObtenerUserMission)
		r.Post("/usermissions", s.CrearUserMission)
		r.Put("/usermissions/{id}", s.ActualizarUserMission)
		r.Delete("/usermissions/{id}", s.BorrarUserMission)
	})
}
