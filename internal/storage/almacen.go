package storage

import (
	"marketplace-api/internal/models"
)

type Almacen interface {
	// MODULO 1
	// Users
	ListarUsers() []models.User

	// MODULO 2

	// MODULO 3
	ListarMessages() []models.Message
	BuscarMessagePorID(id int) (models.Message, bool)
	CrearMessage(message models.Message) models.Message
	ActualizarMessage(id int, datos models.Message) (models.Message, bool)
	BorrarMessage(id int) bool
}
