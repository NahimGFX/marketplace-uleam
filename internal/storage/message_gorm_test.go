package storage

import (
	"testing"

	"marketplace-api/internal/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func setupDBTest(t *testing.T) *gorm.DB {

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	if err != nil {
		t.Fatalf("no se pudo abrir la base de datos: %v", err)
	}

	err = db.AutoMigrate(&models.Message{})

	if err != nil {
		t.Fatalf("fallo migracion: %v", err)
	}

	return db
}

// =============================
// CREATE + READ + LIST
// =============================

func TestMessage_Crear_Buscar_Listar(t *testing.T) {

	db := setupDBTest(t)

	message := models.Message{
		SenderID:   1,
		ReceiverID: 2,
		Content:    "hola mundo",
	}

	err := db.Create(&message).Error

	if err != nil {
		t.Fatalf("error creando mensaje: %v", err)
	}

	if message.ID == 0 {
		t.Fatalf("esperaba ID generado")
	}

	var encontrado models.Message

	err = db.First(&encontrado, message.ID).Error

	if err != nil {
		t.Fatalf("no se encontro mensaje: %v", err)
	}

	if encontrado.SenderID != 1 {
		t.Fatalf("SenderID incorrecto")
	}

	if encontrado.ReceiverID != 2 {
		t.Fatalf("ReceiverID incorrecto")
	}

	if encontrado.Content != "hola mundo" {
		t.Fatalf("contenido incorrecto")
	}

	var mensajes []models.Message

	err = db.Find(&mensajes).Error

	if err != nil {
		t.Fatalf("error listando mensajes: %v", err)
	}

	if len(mensajes) != 1 {
		t.Fatalf("esperaba 1 mensaje")
	}
}

// =============================
// BUSCAR ID INEXISTENTE
// =============================

func TestMessage_BuscarNoExiste(t *testing.T) {

	db := setupDBTest(t)

	var message models.Message

	err := db.First(&message, 999).Error

	if err == nil {
		t.Fatalf("esperaba error buscando mensaje inexistente")
	}

	if err != gorm.ErrRecordNotFound {
		t.Fatalf("error diferente al esperado")
	}
}

// =============================
// ACTUALIZAR MENSAJE
// =============================

func TestMessage_Actualizar(t *testing.T) {

	db := setupDBTest(t)

	message := models.Message{
		SenderID:   1,
		ReceiverID: 2,
		Content:    "mensaje original",
	}

	err := db.Create(&message).Error

	if err != nil {
		t.Fatalf("error creando mensaje: %v", err)
	}

	message.Content = "mensaje actualizado"

	err = db.Save(&message).Error

	if err != nil {
		t.Fatalf("error actualizando: %v", err)
	}

	var actualizado models.Message

	err = db.First(&actualizado, message.ID).Error

	if err != nil {
		t.Fatalf("error buscando mensaje actualizado: %v", err)
	}

	if actualizado.Content != "mensaje actualizado" {
		t.Fatalf("no se actualizo correctamente")
	}
}

// =============================
// ELIMINAR MENSAJE
// =============================

func TestMessage_Eliminar(t *testing.T) {

	db := setupDBTest(t)

	message := models.Message{
		SenderID:   1,
		ReceiverID: 2,
		Content:    "mensaje borrar",
	}

	err := db.Create(&message).Error

	if err != nil {
		t.Fatalf("error creando mensaje: %v", err)
	}

	err = db.Delete(&message).Error

	if err != nil {
		t.Fatalf("error eliminando mensaje: %v", err)
	}

	var encontrado models.Message

	err = db.First(&encontrado, message.ID).Error

	if err == nil {
		t.Fatalf("el mensaje todavía existe")
	}

	if err != gorm.ErrRecordNotFound {
		t.Fatalf("error diferente al esperado")
	}
}

// =============================
// LISTAR VARIOS MENSAJES
// =============================

func TestMessage_ListarVarios(t *testing.T) {

	db := setupDBTest(t)

	mensajes := []models.Message{

		{
			SenderID:   1,
			ReceiverID: 2,
			Content:    "mensaje 1",
		},

		{
			SenderID:   2,
			ReceiverID: 3,
			Content:    "mensaje 2",
		},
	}

	err := db.Create(&mensajes).Error

	if err != nil {
		t.Fatalf("error creando mensajes: %v", err)
	}

	var resultado []models.Message

	err = db.Find(&resultado).Error

	if err != nil {
		t.Fatalf("error listando mensajes: %v", err)
	}

	if len(resultado) != 2 {
		t.Fatalf("esperaba 2 mensajes")
	}
}
