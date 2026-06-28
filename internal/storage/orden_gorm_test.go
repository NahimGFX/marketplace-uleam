package storage_test

import (
	"testing"

	"marketplace-api/internal/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func setupDBTest(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("no se pudo abrir la base de datos en memoria: %v", err)
	}
	db.AutoMigrate(
		&models.Categoria{},
		&models.Producto{},
		&models.Orden{},
	)
	return db
}

func TestOrdenRepo_Crear_Listar(t *testing.T) {
	db := setupDBTest(t)

	// Crear categoría y producto primero por las foreign keys
	categoria := models.Categoria{Name: "Libros"}
	db.Create(&categoria)

	producto := models.Producto{
		Nombre:      "Cálculo",
		Descripcion: "Libro de cálculo",
		Precio:      15.00,
		CategoriaID: uint(categoria.ID),
	}
	db.Create(&producto)

	// Crear orden
	orden := models.Orden{
		ProductoID:  producto.ID,
		IDComprador: 1,
		Estado:      "pendiente",
	}
	db.Create(&orden)

	// Listar y verificar
	var ordenes []models.Orden
	db.Find(&ordenes)

	if len(ordenes) != 1 {
		t.Fatalf("esperado 1 orden, obtenido %d", len(ordenes))
	}
	if ordenes[0].Estado != "pendiente" {
		t.Fatalf("estado esperado 'pendiente', obtenido '%s'", ordenes[0].Estado)
	}
}
