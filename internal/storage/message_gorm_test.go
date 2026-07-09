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

	db.AutoMigrate(&models.Message{})
	return db
}

func TestMessage_Crear_Buscar_Listar(t *testing.T) {
	db := setupDBTest(t)

	// CREAR
	message := models.Message{
		SenderID:   1,
		ReceiverID: 1,
		Content:    "hola mundo"}
	db.Create(&message)

	if message.ID == 0 {
		t.Fatalf("esperaba ID generado")
	}

	// BUSCAR
	var encontrado models.Message
	err := db.First(&encontrado, message.ID).Error
	if err != nil {
		t.Fatalf("no se pudo encontrar el mensaje")
	}
	if encontrado.SenderID != 1 {
		t.Fatalf("SenderID incorrecto")
	}
	if encontrado.ReceiverID != 1 {
		t.Fatalf("ReceiverID incorrecto")
	}

	if encontrado.Content != "hola mundo" {
		t.Fatalf("contenido incorrecto")
	}

	// LISTAR
	var mensajes []models.Message
	db.Find(&mensajes)

	if len(mensajes) != 1 {
		t.Fatalf("esperado 1 mensaje, obtenido %d", len(mensajes))
	}

	if mensajes[0].Content != "hola mundo" {
		t.Fatalf("dato incorrecto en lista")
	}
}
