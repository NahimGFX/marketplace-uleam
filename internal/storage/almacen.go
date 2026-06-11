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

	// Messages
	ListarMessages() []models.Message
	BuscarMessagePorID(id int) (models.Message, bool)
	CrearMessage(message models.Message) models.Message
	ActualizarMessage(id int, datos models.Message) (models.Message, bool)
	BorrarMessage(id int) bool
	//// Misiones
	ListarMissions() []models.Mission
	BuscarMisionPorID(id int) (models.Mission, bool)
	CrearMision(mission models.Mission) models.Mission
	ActualizarMision(id int, datos models.Mission) (models.Mission, bool)
	BorrarMision(id int) bool
}
