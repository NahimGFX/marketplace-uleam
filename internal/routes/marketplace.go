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

	r.Group(func(r chi.Router) {
		r.Use(middleware.Auth(authSvc))

		// Categorías
		r.Get("/categorias", s.ListarCategorias)
		r.Get("/categorias/{id}", s.ObtenerCategoria)
		r.With(middleware.RolPermitido("admin")).Post("/categorias", s.CrearCategoria)
		r.With(middleware.RolPermitido("admin")).Put("/categorias/{id}", s.ActualizarCategoria)
		r.With(middleware.RolPermitido("admin")).Delete("/categorias/{id}", s.BorrarCategoria)

		// Productos
		r.Get("/productos", s.ListarProductos)
		r.Get("/productos/{id}", s.ObtenerProducto)
		r.Post("/productos", s.CrearProducto)
		r.Put("/productos/{id}", s.ActualizarProducto)
		r.Delete("/productos/{id}", s.BorrarProducto)

		// Órdenes
		r.Get("/ordenes", s.ListarOrdenes)
		r.Get("/ordenes/{id}", s.ObtenerOrden)
		r.Post("/ordenes", s.CrearOrden)
		r.Put("/ordenes/{id}", s.ActualizarOrden)
		r.Delete("/ordenes/{id}", s.BorrarOrden)
	})
}
