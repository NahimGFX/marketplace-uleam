// Package handlers contiene los handlers HTTP de la API de cafeteria.
package handlers

import "marketplace-api/internal/service"

// Server agrupa los servicios de los que dependen los handlers.
type Server struct {
	// Modulo 3
	Messages     *service.MessageService
	Missions     *service.MissionService
	UserMissions *service.UserMissionService
	Auth         *service.AuthService

	// Modulo 2
	Categorias *service.CategoriaService
	Productos  *service.ProductoService
	Ordenes    *service.OrdenService
}

func NewServer(
	messages *service.MessageService,
	missions *service.MissionService,
	usermissions *service.UserMissionService,
	auth *service.AuthService,
	categorias *service.CategoriaService,
	productos *service.ProductoService,
	ordenes *service.OrdenService,
) *Server {
	return &Server{
		Messages:     messages,
		Missions:     missions,
		UserMissions: usermissions,
		Auth:         auth,
		Categorias:   categorias,
		Productos:    productos,
		Ordenes:      ordenes,
	}
}
