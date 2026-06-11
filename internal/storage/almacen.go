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
	ListarMessages() []models.Message
}
