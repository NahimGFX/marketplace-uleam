package routes

import (
	"marketplace-api/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func MarketplaceRoutes(r chi.Router, s *handlers.Server) {

	r.Get("/categorias", s.ListarCategorias)
	r.Get("/categorias/{id}", s.ObtenerCategoria)
	r.Post("/categorias", s.CrearCategoria)
	r.Put("/categorias/{id}", s.ActualizarCategoria)
	r.Delete("/categorias/{id}", s.BorrarCategoria)

	r.Get("/productos", s.ListarProductos)
	r.Get("/productos/{id}", s.ObtenerProducto)
	r.Post("/productos", s.CrearProducto)
	r.Put("/productos/{id}", s.ActualizarProducto)
	r.Delete("/productos/{id}", s.BorrarProducto)

	r.Get("/ordenes", s.ListarOrdenes)
	r.Get("/ordenes/{id}", s.ObtenerOrden)
	r.Post("/ordenes", s.CrearOrden)
	r.Put("/ordenes/{id}", s.ActualizarOrden)
	r.Delete("/ordenes/{id}", s.BorrarOrden)
}
