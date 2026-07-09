package storage_test

import (
	"testing"

	"marketplace-api/internal/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// =========================
// SETUP DB EN MEMORIA
// =========================
func setupBadgeDBTest(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("no se pudo abrir la base de datos en memoria: %v", err)
	}

	// IMPORTANTE: crear la tabla en memoria
	err = db.AutoMigrate(&models.Badge{})
	if err != nil {
		t.Fatalf("error en AutoMigrate: %v", err)
	}

	return db
}

// =========================
// TEST CRUD BÁSICO BADGE
// =========================
func TestBadge_Crear_Buscar_Listar(t *testing.T) {
	db := setupBadgeDBTest(t)

	// CREAR
	badge := models.Badge{
		Name:        "Experto",
		Description: "Alcanza 100 puntos de reputación",
		RequiredRep: 100,
	}

	result := db.Create(&badge)
	if result.Error != nil {
		t.Fatalf("error al crear badge: %v", result.Error)
	}

	if badge.ID == 0 {
		t.Fatalf("esperaba ID generado")
	}

	// BUSCAR
	var encontrado models.Badge
	err := db.First(&encontrado, badge.ID).Error
	if err != nil {
		t.Fatalf("no se pudo encontrar el badge: %v", err)
	}

	if encontrado.Name != "Experto" {
		t.Fatalf("nombre incorrecto")
	}

	if encontrado.Description != "Alcanza 100 puntos de reputación" {
		t.Fatalf("descripción incorrecta")
	}

	if encontrado.RequiredRep != 100 {
		t.Fatalf("RequiredRep incorrecto")
	}

	// LISTAR
	var badges []models.Badge
	db.Find(&badges)

	if len(badges) != 1 {
		t.Fatalf("esperado 1 badge, obtenido %d", len(badges))
	}

	if badges[0].Name != "Experto" {
		t.Fatalf("dato incorrecto en lista")
	}
}
