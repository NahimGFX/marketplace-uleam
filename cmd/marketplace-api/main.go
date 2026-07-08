package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/glebarez/go-sqlite"
	"github.com/glebarez/sqlite"
	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"marketplace-api/internal/config"
	"marketplace-api/internal/handlers"
	"marketplace-api/internal/middleware"
	"marketplace-api/internal/models"
	"marketplace-api/internal/routes"
	"marketplace-api/internal/service"
	"marketplace-api/internal/storage"
)

func main() {
	cfg, err := config.Cargar()
	if err != nil {
		log.Fatal("configuracion invalida: ", err)
	}

	// =========================
	// GORM (usuarios + migraciones)
	// =========================
	gdb, err := abrirGORM(cfg)
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
	); err != nil {
		log.Fatal("fallo AutoMigrate: ", err)
	}

	almacenGorm := storage.NuevoAlmacenSQLite(gdb)
	almacenGorm.SembrarSiVacio()

	// =========================
	// BACKEND DINAMICO (GORM o SQLC)
	// =========================
	var almacen storage.Almacen

	switch cfg.Storage {
	case "sqlc":
		sdb, err := abrirSQLC(cfg)
		if err != nil {
			log.Fatal("no se pudo abrir sql.DB para sqlc: ", err)
		}

		almacen = storage.NuevoAlmacenSQLC(sdb)
		log.Println("Backend productos/categorias: SQLC")

	default:
		almacen = almacenGorm
		log.Println("Backend productos/categorias: GORM")
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
	productoService := service.NewPorductoService(almacen)
	ordenService := service.NewOrdenService(almacen)
	messageService := service.NuevoMessageService(almacen)
	missionService := service.NuevoMissionService(almacen)
	userMissionService := service.NuevoUserMissionService(almacen)
	authService := service.NuevoAuthServiceConJWT(usuarioRepo, cfg.JWTSecreto, cfg.JWTDuracion)

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

	httpServer := &http.Server{
		Addr:         cfg.Puerto,
		Handler:      r,
		ReadTimeout:  cfg.HTTPReadTimeout,
		WriteTimeout: cfg.HTTPWriteTimeout,
	}

	log.Printf("Servidor escuchando en http://localhost%s", cfg.Puerto)
	log.Fatal(httpServer.ListenAndServe())
}

func abrirGORM(cfg config.Config) (*gorm.DB, error) {
	switch cfg.DBDriver {
	case "postgres":
		return gorm.Open(postgres.Open(cfg.DBDSN), &gorm.Config{})
	case "sqlite":
		return gorm.Open(sqlite.Open(cfg.RutaDB), &gorm.Config{})
	default:
		return nil, fmt.Errorf("DB_DRIVER no soportado: %s", cfg.DBDriver)
	}
}

func abrirSQLC(cfg config.Config) (*sql.DB, error) {
	if cfg.DBDriver != "sqlite" {
		return nil, fmt.Errorf("STORAGE=sqlc solo esta disponible con DB_DRIVER=sqlite")
	}
	return sql.Open("sqlite", cfg.RutaDB)
}
