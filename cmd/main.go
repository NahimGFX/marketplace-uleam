// Command cafeteria-api arranca el servidor HTTP de la Cafetería Universitaria.
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
		r.Get("/ruses", servidor.ListarUsers)

		// MODULO 2

		// MODULO 3
		///messages
		r.Get("/rmesa", servidor.ListarMessages)
		r.Get("/lmesa/{id}", servidor.ObtenerMessage)
		r.Post("/cremesa", servidor.CrearMessage)
		r.Put("/amesa/{id}", servidor.ActualizarMessage)
		r.Delete("/dmesa/{id}", servidor.BorrarMessage)
		///missions
		r.Get("/rmission", servidor.ListarMissions)
		r.Get("/lmission/{id}", servidor.ObtenerMision)
		r.Post("/cremission", servidor.CrearMision)
		r.Put("/amission/{id}", servidor.ActualizarMision)
	})

	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
