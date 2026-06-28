package routes

import (
	"marketplace-api/internal/handlers"
	"marketplace-api/internal/middleware"
	"marketplace-api/internal/service"

	"github.com/go-chi/chi/v5"
)

func MarketplaceRoutes(
	r chi.Router,
	s *handlers.Server,
	authSvc *service.AuthService,
) {
	// Rutas públicas (solo lectura)
	r.Get("/categorias", s.ListarCategorias)
	r.Get("/categorias/{id}", s.ObtenerCategoria)
	r.Get("/productos", s.ListarProductos)
	r.Get("/productos/{id}", s.ObtenerProducto)
	r.Get("/ordenes", s.ListarOrdenes)
	r.Get("/ordenes/{id}", s.ObtenerOrden)

	// Rutas protegidas (requieren token)
	r.Group(func(r chi.Router) {
		r.Use(middleware.Auth(authSvc))

		r.Post("/categorias", s.CrearCategoria)
		r.Put("/categorias/{id}", s.ActualizarCategoria)
		r.Delete("/categorias/{id}", s.BorrarCategoria)

		r.Post("/productos", s.CrearProducto)
		r.Put("/productos/{id}", s.ActualizarProducto)
		r.Delete("/productos/{id}", s.BorrarProducto)

		r.Post("/ordenes", s.CrearOrden)
		r.Put("/ordenes/{id}", s.ActualizarOrden)
		r.Delete("/ordenes/{id}", s.BorrarOrden)
	})
}
