package service

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"marketplace-api/internal/models"
)

func TestOrdenService_ProductoIDInvalido_NoLlegaRepositorio(t *testing.T) {
	repo := new(ordenRepoMock)
	svc := NewOrdenService(repo)

	_, err := svc.Crear(models.Orden{
		ProductoID:  0,
		IDComprador: 1,
		Estado:      "pendiente",
	})

	require.ErrorIs(t, err, ErrProductoInvalido)
	repo.AssertNotCalled(t, "CrearOrden", mock.Anything)
	repo.AssertExpectations(t)
}

func TestOrdenService_CompradorInvalido_NoLlegaRepositorio(t *testing.T) {
	repo := new(ordenRepoMock)
	svc := NewOrdenService(repo)

	_, err := svc.Crear(models.Orden{
		ProductoID:  1,
		IDComprador: 0,
		Estado:      "pendiente",
	})

	require.ErrorIs(t, err, ErrCompradorInvalido)
	repo.AssertNotCalled(t, "CrearOrden", mock.Anything)
	repo.AssertExpectations(t)
}

func TestOrdenService_EstadoVacio_NoLlegaRepositorio(t *testing.T) {
	repo := new(ordenRepoMock)
	svc := NewOrdenService(repo)

	_, err := svc.Crear(models.Orden{
		ProductoID:  1,
		IDComprador: 1,
		Estado:      "",
	})

	require.ErrorIs(t, err, ErrEstadoVacio)
	repo.AssertNotCalled(t, "CrearOrden", mock.Anything)
	repo.AssertExpectations(t)
}

func TestOrdenService_CrearValida_LlegaRepositorio(t *testing.T) {
	repo := new(ordenRepoMock)
	svc := NewOrdenService(repo)
	orden := models.Orden{
		ProductoID:  1,
		IDComprador: 2,
		Estado:      "pendiente",
	}
	guardada := orden
	guardada.ID = 5

	repo.On("CrearOrden", orden).Return(guardada).Once()

	resultado, err := svc.Crear(orden)

	require.NoError(t, err)
	require.Equal(t, guardada, resultado)
	repo.AssertExpectations(t)
}

func TestOrdenService_ActualizarNoEncontrado_RetornaError(t *testing.T) {
	repo := new(ordenRepoMock)
	svc := NewOrdenService(repo)
	orden := models.Orden{
		ProductoID:  1,
		IDComprador: 2,
		Estado:      "pagado",
	}

	repo.On("ActualizarOrden", 99, orden).Return(models.Orden{}, false).Once()

	_, err := svc.Actualizar(99, orden)

	require.ErrorIs(t, err, ErrNoEncontrado)
	repo.AssertExpectations(t)
}

func TestOrdenService_BorrarNoEncontrado_RetornaError(t *testing.T) {
	repo := new(ordenRepoMock)
	svc := NewOrdenService(repo)

	repo.On("BorrarOrden", 99).Return(false).Once()

	err := svc.Borrar(99)

	require.ErrorIs(t, err, ErrOrdenNoEncontrada)
	repo.AssertExpectations(t)
}
