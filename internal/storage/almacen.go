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

	// Reviews
	ListarReviews() []models.Review
	BuscarReviewPorID(id int) (models.Review, bool)
	BuscarBadgePorID(id int) (models.Badge, bool)

	// Badges
	ListarBadges() []models.Badge

	// MODULO 2
	// Categorias
	ListarCategorias() []models.Categoria
	BuscarCategoriaPorID(id int) (models.Categoria, bool)
	CrearCategoria(c models.Categoria) models.Categoria
	ActualizarCategoria(id int, datos models.Categoria) (models.Categoria, bool)
	BorrarCategoria(id int) bool

	// Productos
	ListarProductos() []models.Producto
	BuscarProductoPorID(id int) (models.Producto, bool)
	CrearProducto(p models.Producto) models.Producto
	ActualizarProducto(id int, datos models.Producto) (models.Producto, bool)
	BorrarProducto(id int) bool

	// Ordenes
	ListarOrdenes() []models.Orden
	BuscarOrdenPorID(id int) (models.Orden, bool)
	CrearOrden(o models.Orden) models.Orden
	ActualizarOrden(id int, datos models.Orden) (models.Orden, bool)
	BorrarOrden(id int) bool

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
	//// UserMissions
	ListarUsermissions() []models.UserMission
	BuscarUserMissionPorID(id int) (models.UserMission, bool)
	CrearUserMission(userMission models.UserMission) models.UserMission
	ActualizarUserMission(id int, datos models.UserMission) (models.UserMission, bool)
	BorrarUserMission(id int) bool
}
