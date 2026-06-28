package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/glebarez/go-sqlite" // driver sqlc
	"github.com/glebarez/sqlite"      // driver gorm
	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"

	"marketplace-api/internal/handlers"
	"marketplace-api/internal/middleware"
	"marketplace-api/internal/models"
	"marketplace-api/internal/routes"
	"marketplace-api/internal/service"
	"marketplace-api/internal/storage"
)

func main() {

	// 1. GORM es el DUENO DEL ESQUEMA: abre la DB, migra y siembra.
	//    Ahora tambien migra la tabla de usuarios.
	gdb, err := gorm.Open(sqlite.Open("marketplace.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("no se pudo abrir la base de datos: ", err)
	}

	if err := gdb.AutoMigrate(
		&models.Message{},
		&models.Mission{},
		&models.UserMission{},
		&models.Usuario{},
	); err != nil {
		log.Fatal("fallo AutoMigrate: ", err)
	}

	almacenGorm := storage.NuevoAlmacenSQLite(gdb)
	almacenGorm.SembrarSiVacio()

	// 2. Elegir el backend de productos/categorias segun STORAGE (igual que antes).
	var almacen storage.Almacen

	switch os.Getenv("STORAGE") {

	case "sqlc":
		sdb, err := sql.Open("sqlite", "marketplace.db")
		if err != nil {
			log.Fatal("no se pudo abrir sql.DB para sqlc: ", err)
		}

		almacen = storage.NuevoAlmacenSQLC(sdb)
		log.Println("Backend productos/categorías: SQLC")

	default:
		almacen = almacenGorm
		log.Println("Backend productos/categorías: GORM")
	}

	// 3. Los usuarios viven SIEMPRE en GORM (decision de la semana). Por eso NO
	//    cerramos gdb aunque el backend de productos sea sqlc: GORM mantiene su
	//    conexion para la tabla de usuarios.
	usuarioRepo := storage.NuevoUsuarioGORM(gdb)

	// 4. Capa de servicio. Cada servicio recibe SOLO la interfaz estrecha que
	//    necesita; almacen (Almacen) cumple ProductoRepository y CategoriaRepository
	//    por embedding, asi que es asignable a ambos parametros.
	message := service.NuevoMessageService(almacen)
	missions := service.NuevoMissionService(almacen)
	userMissions := service.NuevoUserMissionService(almacen)
	auth := service.NuevoAuthService(usuarioRepo)
	// 5. Server con los servicios inyectados.
	servidor := handlers.NewServer(message, missions, userMissions, auth)

	r := chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(middleware.CORS)

	r.Route("/api/v1", func(r chi.Router) {

		routes.PerfiñRoutes(r, servidor)
		routes.MarketplaceRoutes(r, servidor)

		// 🔐 ahora sí pasas authSvc
		routes.AuthRoutes(r, servidor)
		routes.ComunidadRoutes(r, servidor, auth)
	})

	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
