package routes

import (
	"marketplace-api/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func ComunidadRoutes(r chi.Router, s *handlers.Server) {

	r.Get("/messages", s.ListarMessages)
	r.Get("/messages/{id}", s.ObtenerMessage)
	r.Post("/messages", s.CrearMessage)
	r.Put("/messages/{id}", s.ActualizarMessage)
	r.Delete("/messages/{id}", s.BorrarMessage)

	r.Get("/missions", s.ListarMissions)
	r.Get("/missions/{id}", s.ObtenerMision)
	r.Post("/missions", s.CrearMision)
	r.Put("/missions/{id}", s.ActualizarMision)
	r.Delete("/missions/{id}", s.BorrarMision)

	r.Get("/usermissions", s.ListarUsermissions)
	r.Get("/usermissions/{id}", s.ObtenerUserMission)
	r.Post("/usermissions", s.CrearUserMission)
	r.Put("/usermissions/{id}", s.ActualizarUserMission)
	r.Delete("/usermissions/{id}", s.BorrarUserMission)
}
