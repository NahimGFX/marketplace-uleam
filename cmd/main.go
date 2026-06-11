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
	almacen.SeedUsermissions()

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
		r.Delete("/missions/{id}", servidor.BorrarMision)
		///usermissions
		r.Get("/usermissions", servidor.ListarUsermissions)
		r.Get("/usermissions/{id}", servidor.ObtenerUserMission)
		r.Post("/usermissions", servidor.CrearUserMission)
		r.Put("/usermissions/{id}", servidor.ActualizarUserMission)
		r.Delete("/usermissions/{id}", servidor.BorrarUserMission)
	})

	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
