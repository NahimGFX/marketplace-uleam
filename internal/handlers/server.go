// Package handlers contiene los handlers HTTP de la API de cafeteria.
package handlers

import "marketplace-api/internal/service"

// Server agrupa los servicios de los que dependen los handlers.
//
// Antes guardaba un storage.Almacen directo; ahora guarda la capa de servicio,
// que es la que tiene la logica de negocio. Los handlers quedan delgados:
// decodifican el request, llaman al servicio y traducen el resultado a HTTP.
type Server struct {
	//MODULO 1
	Users   *service.UserService
	Reviews *service.ReviewService
	Badges  *service.BadgeService
	// MÓDULO 2
	Categorias *service.CategoriaService
	Productos  *service.ProductoService
	Ordenes    *service.OrdenService
	//MODULO 3
	Messages     *service.MessageService
	Missions     *service.MissionService
	UserMissions *service.UserMissionService
	Auth         *service.AuthService
}

// NewServer construye un Server con sus servicios ya inyectados.
func NewServer(
	users *service.UserService,
	reviews *service.ReviewService,
	badges *service.BadgeService,
	messages *service.MessageService,
	missions *service.MissionService,
	usermissions *service.UserMissionService,
	auth *service.AuthService,
	categories *service.CategoriaService,
	products *service.ProductoService,
	orders *service.OrdenService,

) *Server {

	return &Server{
		Users:   users,
		Reviews: reviews,
		Badges:  badges,

		Categorias: categories,
		Productos:  products,
		Ordenes:    orders,

		Messages:     messages,
		Missions:     missions,
		UserMissions: usermissions,
		Auth:         auth,
	}
}
