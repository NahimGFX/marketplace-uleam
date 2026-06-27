package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"

	"marketplace-api/internal/handlers"
	"marketplace-api/internal/routes"
	"marketplace-api/internal/storage"
)

func main() {

	almacen := storage.NuevaMemoria()
	almacen.SeedUsers()
	almacen.SeedReviews()
	almacen.SeedBadges()
	almacen.SeedMessages()
	almacen.SeedMissions()
	almacen.SeedUsermissions()

	servidor := handlers.NewServer(almacen)

	r := chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)

	r.Route("/api/v1", func(r chi.Router) {

		routes.PerfiñRoutes(r, servidor)
		routes.MarketplaceRoutes(r, servidor)
		routes.ComunidadRoutes(r, servidor)

	})

	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
