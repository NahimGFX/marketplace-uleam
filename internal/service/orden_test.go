package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"marketplace-api/internal/models"
	"marketplace-api/internal/service"
	"marketplace-api/internal/storage"
)

// ordenRepoMock es un doble de storage.OrdenRepository (15 metodos:
// Categoria + Producto + Orden, el modulo completo). Se define aqui y
// se reutiliza en categoria_test.go y producto_test.go (mismo paquete).
type ordenRepoMock struct {
	mock.Mock
}

func (m *ordenRepoMock) ListarCategorias() []models.Categoria {
	return m.Called().Get(0).([]models.Categoria)
}
func (m *ordenRepoMock) BuscarCategoriaPorID(id int) (models.Categoria, bool) {
	a := m.Called(id)
	return a.Get(0).(models.Categoria), a.Bool(1)
}
func (m *ordenRepoMock) CrearCategoria(c models.Categoria) models.Categoria {
	return m.Called(c).Get(0).(models.Categoria)
}
func (m *ordenRepoMock) ActualizarCategoria(id int, datos models.Categoria) (models.Categoria, bool) {
	a := m.Called(id, datos)
	return a.Get(0).(models.Categoria), a.Bool(1)
}
func (m *ordenRepoMock) BorrarCategoria(id int) bool {
	return m.Called(id).Bool(0)
}

func (m *ordenRepoMock) ListarProductos() []models.Producto {
	return m.Called().Get(0).([]models.Producto)
}
func (m *ordenRepoMock) BuscarProductoPorID(id int) (models.Producto, bool) {
	a := m.Called(id)
	return a.Get(0).(models.Producto), a.Bool(1)
}
func (m *ordenRepoMock) CrearProducto(p models.Producto) models.Producto {
	return m.Called(p).Get(0).(models.Producto)
}
func (m *ordenRepoMock) ActualizarProducto(id int, datos models.Producto) (models.Producto, bool) {
	a := m.Called(id, datos)
	return a.Get(0).(models.Producto), a.Bool(1)
}
func (m *ordenRepoMock) BorrarProducto(id int) bool {
	return m.Called(id).Bool(0)
}

func (m *ordenRepoMock) ListarOrdenes() []models.Orden {
	return m.Called().Get(0).([]models.Orden)
}
func (m *ordenRepoMock) BuscarOrdenPorID(id int) (models.Orden, bool) {
	a := m.Called(id)
	return a.Get(0).(models.Orden), a.Bool(1)
}
func (m *ordenRepoMock) CrearOrden(o models.Orden) models.Orden {
	return m.Called(o).Get(0).(models.Orden)
}
func (m *ordenRepoMock) ActualizarOrden(id int, datos models.Orden) (models.Orden, bool) {
	a := m.Called(id, datos)
	return a.Get(0).(models.Orden), a.Bool(1)
}
func (m *ordenRepoMock) BorrarOrden(id int) bool {
	return m.Called(id).Bool(0)
}

// Red de seguridad en tiempo de compilacion: el mock DEBE cumplir el contrato.
var _ storage.OrdenRepository = (*ordenRepoMock)(nil)

// --- Crear: 3 validaciones distintas en validarOrden ---

func TestOrdenService_Crear(t *testing.T) {
	casos := []struct {
		nombre        string
		entrada       models.Orden
		errEsperado   error
		debePersistir bool
	}{
		{"producto_id invalido rechazado", models.Orden{ProductoID: 0, IDComprador: 1, Estado: "pendiente"}, service.ErrProductoInvalido, false},
		{"comprador invalido rechazado", models.Orden{ProductoID: 1, IDComprador: 0, Estado: "pendiente"}, service.ErrCompradorInvalido, false},
		{"estado vacio rechazado", models.Orden{ProductoID: 1, IDComprador: 1, Estado: ""}, service.ErrEstadoVacio, false},
		{"orden valida se persiste", models.Orden{ProductoID: 1, IDComprador: 1, Estado: "pendiente"}, nil, true},
	}
	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			repo := new(ordenRepoMock)
			if c.debePersistir {
				guardado := c.entrada
				guardado.ID = 1
				repo.On("CrearOrden", c.entrada).Return(guardado)
			}
			svc := service.NewOrdenService(repo)

			creado, err := svc.Crear(c.entrada)

			if c.errEsperado != nil {
				require.ErrorIs(t, err, c.errEsperado)
				repo.AssertNotCalled(t, "CrearOrden")
			} else {
				require.NoError(t, err)
				assert.Equal(t, 1, creado.ID)
				repo.AssertCalled(t, "CrearOrden", c.entrada)
			}
		})
	}
}

func TestOrdenService_Obtener(t *testing.T) {
	t.Run("existe", func(t *testing.T) {
		repo := new(ordenRepoMock)
		repo.On("BuscarOrdenPorID", 1).Return(models.Orden{ID: 1, Estado: "pendiente"}, true)
		o, ok := service.NewOrdenService(repo).Obtener(1)
		assert.True(t, ok)
		assert.Equal(t, "pendiente", o.Estado)
	})
	t.Run("no existe", func(t *testing.T) {
		repo := new(ordenRepoMock)
		repo.On("BuscarOrdenPorID", 999).Return(models.Orden{}, false)
		_, ok := service.NewOrdenService(repo).Obtener(999)
		assert.False(t, ok)
	})
}

func TestOrdenService_Actualizar(t *testing.T) {
	datos := models.Orden{ProductoID: 1, IDComprador: 1, Estado: "enviado"}

	t.Run("valido", func(t *testing.T) {
		repo := new(ordenRepoMock)
		actualizado := datos
		actualizado.ID = 1
		repo.On("ActualizarOrden", 1, datos).Return(actualizado, true)
		o, err := service.NewOrdenService(repo).Actualizar(1, datos)
		require.NoError(t, err)
		assert.Equal(t, 1, o.ID)
	})
	t.Run("no existe -> ErrNoEncontrado", func(t *testing.T) {
		repo := new(ordenRepoMock)
		repo.On("ActualizarOrden", 999, datos).Return(models.Orden{}, false)
		_, err := service.NewOrdenService(repo).Actualizar(999, datos)
		require.ErrorIs(t, err, service.ErrNoEncontrado)
	})
	t.Run("invalido no toca el repo", func(t *testing.T) {
		repo := new(ordenRepoMock)
		_, err := service.NewOrdenService(repo).Actualizar(1, models.Orden{Estado: ""})
		require.ErrorIs(t, err, service.ErrProductoInvalido) // ProductoID=0 se valida primero
		repo.AssertNotCalled(t, "ActualizarOrden")
	})
}

// Ojo: Borrar usa ErrOrdenNoEncontrada (no ErrNoEncontrado) - asi esta en tu service.go
func TestOrdenService_Borrar(t *testing.T) {
	t.Run("existe", func(t *testing.T) {
		repo := new(ordenRepoMock)
		repo.On("BorrarOrden", 1).Return(true)
		require.NoError(t, service.NewOrdenService(repo).Borrar(1))
	})
	t.Run("no existe -> ErrOrdenNoEncontrada", func(t *testing.T) {
		repo := new(ordenRepoMock)
		repo.On("BorrarOrden", 999).Return(false)
		require.ErrorIs(t, service.NewOrdenService(repo).Borrar(999), service.ErrOrdenNoEncontrada)
	})
}

func TestOrdenService_Listar(t *testing.T) {
	repo := new(ordenRepoMock)
	repo.On("ListarOrdenes").Return([]models.Orden{{ID: 1}, {ID: 2}})
	lista := service.NewOrdenService(repo).Listar()
	assert.Len(t, lista, 2)
	repo.AssertExpectations(t)
}
