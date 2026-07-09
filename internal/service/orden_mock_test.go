package service

import (
	"github.com/stretchr/testify/mock"

	"marketplace-api/internal/models"
	"marketplace-api/internal/storage"
)

type ordenRepoMock struct {
	mock.Mock
}

func (m *ordenRepoMock) ListarCategorias() []models.Categoria {
	args := m.Called()
	return args.Get(0).([]models.Categoria)
}

func (m *ordenRepoMock) BuscarCategoriaPorID(id int) (models.Categoria, bool) {
	args := m.Called(id)
	return args.Get(0).(models.Categoria), args.Bool(1)
}

func (m *ordenRepoMock) CrearCategoria(c models.Categoria) models.Categoria {
	args := m.Called(c)
	return args.Get(0).(models.Categoria)
}

func (m *ordenRepoMock) ActualizarCategoria(id int, datos models.Categoria) (models.Categoria, bool) {
	args := m.Called(id, datos)
	return args.Get(0).(models.Categoria), args.Bool(1)
}

func (m *ordenRepoMock) BorrarCategoria(id int) bool {
	args := m.Called(id)
	return args.Bool(0)
}

func (m *ordenRepoMock) ListarProductos() []models.Producto {
	args := m.Called()
	return args.Get(0).([]models.Producto)
}

func (m *ordenRepoMock) BuscarProductoPorID(id int) (models.Producto, bool) {
	args := m.Called(id)
	return args.Get(0).(models.Producto), args.Bool(1)
}

func (m *ordenRepoMock) CrearProducto(p models.Producto) models.Producto {
	args := m.Called(p)
	return args.Get(0).(models.Producto)
}

func (m *ordenRepoMock) ActualizarProducto(id int, datos models.Producto) (models.Producto, bool) {
	args := m.Called(id, datos)
	return args.Get(0).(models.Producto), args.Bool(1)
}

func (m *ordenRepoMock) BorrarProducto(id int) bool {
	args := m.Called(id)
	return args.Bool(0)
}

func (m *ordenRepoMock) ListarOrdenes() []models.Orden {
	args := m.Called()
	return args.Get(0).([]models.Orden)
}

func (m *ordenRepoMock) BuscarOrdenPorID(id int) (models.Orden, bool) {
	args := m.Called(id)
	return args.Get(0).(models.Orden), args.Bool(1)
}

func (m *ordenRepoMock) CrearOrden(o models.Orden) models.Orden {
	args := m.Called(o)
	return args.Get(0).(models.Orden)
}

func (m *ordenRepoMock) ActualizarOrden(id int, datos models.Orden) (models.Orden, bool) {
	args := m.Called(id, datos)
	return args.Get(0).(models.Orden), args.Bool(1)
}

func (m *ordenRepoMock) BorrarOrden(id int) bool {
	args := m.Called(id)
	return args.Bool(0)
}

var _ storage.OrdenRepository = (*ordenRepoMock)(nil)
