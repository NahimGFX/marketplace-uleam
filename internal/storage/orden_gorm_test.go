package storage_test

import (
	"testing"

	"marketplace-api/internal/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func setupOrdenDBTest(t *testing.T) *gorm.DB {
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

func TestOrden_Crear_Buscar_Listar(t *testing.T) {
	db := setupOrdenDBTest(t)

	// CREAR categoria primero por foreign key
	categoria := models.Categoria{Name: "Libros"}
	db.Create(&categoria)

	if categoria.ID == 0 {
		t.Fatalf("esperaba ID generado para categoria")
	}

	// CREAR producto por foreign key
	producto := models.Producto{
		Nombre:      "Cálculo",
		Descripcion: "Libro de cálculo",
		Precio:      15.00,
		CategoriaID: uint(categoria.ID),
	}
	db.Create(&producto)

	if producto.ID == 0 {
		t.Fatalf("esperaba ID generado para producto")
	}

	// CREAR orden
	orden := models.Orden{
		ProductoID:  producto.ID,
		IDComprador: 1,
		Estado:      "pendiente",
	}
	db.Create(&orden)

	if orden.ID == 0 {
		t.Fatalf("esperaba ID generado para orden")
	}

	// BUSCAR
	var encontrada models.Orden
	err := db.First(&encontrada, orden.ID).Error
	if err != nil {
		t.Fatalf("no se pudo encontrar la orden")
	}
	if encontrada.Estado != "pendiente" {
		t.Fatalf("estado incorrecto")
	}

	// LISTAR
	var ordenes []models.Orden
	db.Find(&ordenes)

	if len(ordenes) != 1 {
		t.Fatalf("esperado 1 orden, obtenido %d", len(ordenes))
	}
	if ordenes[0].Estado != "pendiente" {
		t.Fatalf("dato incorrecto en lista")
	}
}
