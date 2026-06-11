package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"

	"marketplace-api/internal/handlers"
	"marketplace-api/internal/storage"
)

func main() {

	almacen := storage.NuevaMemoria()
	almacen.SeedUsers()
	almacen.SeedReviews()
	almacen.SeedBadges()
	almacen.SeedMessages()
	almacen.SeedMissions()

	servidor := handlers.NewServer(almacen)
	// 4. Router + middleware.
	r := chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)

	// 5. Rutas versionadas /api/v1/.
	r.Route("/api/v1", func(r chi.Router) {
		// MODULO 1
		// Users
		r.Get("/users", servidor.ListarUsers)
		r.Post("/users", servidor.CrearUser)
		r.Get("/users/{id}", servidor.ObtenerUser)
		r.Put("/users/{id}", servidor.ActualizarUser)
		r.Delete("/users/{id}", servidor.BorrarUser)

		// Reviews
		r.Get("/reviews", servidor.ListarReviews)
		r.Get("/reviews/{id}", servidor.ObteneReview)

		// Badges
		r.Get("/badges", servidor.ListarBadges)
		r.Get("/badges/{id}", servidor.ObteneBadge)

		// MODULO 2
		r.Get("/categorias", servidor.ListarCategorias)
		r.Get("/categorias/{id}", servidor.ObtenerCategoria)
		r.Post("/categorias", servidor.CrearCategoria)
		r.Put("/categorias/{id}", servidor.ActualizarCategoria)
		r.Delete("/categorias/{id}", servidor.BorrarCategoria)

		r.Get("/productos", servidor.ListarProductos)
		r.Get("/productos/{id}", servidor.ObtenerProducto)
		r.Post("/productos", servidor.CrearProducto)
		r.Put("/productos/{id}", servidor.ActualizarProducto)
		r.Delete("/productos/{id}", servidor.BorrarProducto)

		r.Get("/ordenes", servidor.ListarOrdenes)
		r.Get("/ordenes/{id}", servidor.ObtenerOrden)
		r.Post("/ordenes", servidor.CrearOrden)
		r.Put("/ordenes/{id}", servidor.ActualizarOrden)
		r.Delete("/ordenes/{id}", servidor.BorrarOrden)

		// MODULO 3
		///messages
		r.Get("/messages", servidor.ListarMessages)
		r.Get("/messages/{id}", servidor.ObtenerMessage)
		r.Post("/messages", servidor.CrearMessage)
		r.Put("/messages/{id}", servidor.ActualizarMessage)
		r.Delete("/messages/{id}", servidor.BorrarMessage)
		///missions
		r.Get("/missions", servidor.ListarMissions)
		r.Get("/missions/{id}", servidor.ObtenerMision)
		r.Post("/missions", servidor.CrearMision)
		r.Put("/missions/{id}", servidor.ActualizarMision)

	})

	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
