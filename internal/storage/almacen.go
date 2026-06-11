package storage

import (
	"marketplace-api/internal/models"
)

type Almacen interface {
	// MODULO 1
	// Users
	ListarUsers() []models.User
	BuscarUserPorID(id int) (models.User, bool)
	CrearUser(u models.User) models.User
	ActualizarUser(id int, datos models.User) (models.User, bool)
	BorrarUser(id int) bool

	//

	// MODULO 2

	// MODULO 3
	ListarMessages() []models.Message
}
