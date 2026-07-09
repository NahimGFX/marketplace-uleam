package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/glebarez/go-sqlite"
	"github.com/glebarez/sqlite"
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

	// =========================
	// GORM (usuarios + migraciones)
	// =========================
	gdb, err := gorm.Open(sqlite.Open("marketplace.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("no se pudo abrir la base de datos: ", err)
	}

	if err := gdb.AutoMigrate(
		&models.User{},
		&models.Review{},
		&models.Badge{},
		&models.Categoria{},
		&models.Producto{},
		&models.Orden{},
		&models.Message{},
		&models.Mission{},
		&models.UserMission{},
		&models.Usuario{},
	); err != nil {
		log.Fatal("fallo AutoMigrate: ", err)
	}

	almacenGorm := storage.NuevoAlmacenSQLite(gdb)
	almacenGorm.SembrarSiVacio()

	// =========================
	// BACKEND DINÁMICO (GORM o SQLC)
	// =========================
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

	// =========================
	// REPOSITORIO DE USUARIOS (SIEMPRE GORM)
	// =========================
	usuarioRepo := storage.NuevoUsuarioGORM(gdb)

	// =========================
	// SERVICIOS
	// =========================
	userService := service.NuevoUserService(almacen)
	reviewService := service.NuevoReviewService(almacen)
	badgeService := service.NuevoBadgeService(almacen)
	categoriaService := service.NewCategoriaService(almacen)
	productoService := service.NewProductoService(almacen)
	ordenService := service.NewOrdenService(almacen)
	messageService := service.NuevoMessageService(almacen)
	missionService := service.NuevoMissionService(almacen)
	userMissionService := service.NuevoUserMissionService(almacen)
	authService := service.NuevoAuthService(usuarioRepo)

	// =========================
	// SERVER
	// =========================
	servidor := handlers.NewServer(
		userService,        // UserService
		reviewService,      // ReviewService
		badgeService,       // BadgeService
		messageService,     // MessageService
		missionService,     // MissionService
		userMissionService, // UserMissionService
		authService,        // AuthService
		categoriaService,   // CategoriaService
		productoService,    // ProductoService
		ordenService,       // OrdenService
	)

	// =========================
	// ROUTER
	// =========================
	r := chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(middleware.CORS)

	r.Route("/api/v1", func(r chi.Router) {
		routes.PerfilRoutes(r, servidor, authService)
		routes.MarketplaceRoutes(r, servidor, authService)
		routes.AuthRoutes(r, servidor)
		routes.ComunidadRoutes(r, servidor, authService)
	})

	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
