package service

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"marketplace-api/internal/models"
)

func TestMessageService_ContentVacio_NoLlegaRepositorio(t *testing.T) {
	repo := new(comunidadRepoMock)
	svc := NuevoMessageService(repo)

	_, err := svc.Crear(models.Message{Content: "   "})

	require.ErrorIs(t, err, ErrContentVacio)
	repo.AssertNotCalled(t, "CrearMessage", mock.Anything)
	repo.AssertExpectations(t)
}

func TestMessageService_CrearValido_LlegaRepositorio(t *testing.T) {
	repo := new(comunidadRepoMock)
	svc := NuevoMessageService(repo)
	msg := models.Message{
		SenderID:   1,
		ReceiverID: 2,
		Content:    "Hola",
	}
	guardado := msg
	guardado.ID = 10

	repo.On("CrearMessage", msg).Return(guardado).Once()

	resultado, err := svc.Crear(msg)

	require.NoError(t, err)
	require.Equal(t, guardado, resultado)
	repo.AssertExpectations(t)
}

func TestMessageService_ActualizarContentVacio_NoLlegaRepositorio(t *testing.T) {
	repo := new(comunidadRepoMock)
	svc := NuevoMessageService(repo)

	_, err := svc.Actualizar(1, models.Message{Content: ""})

	require.ErrorIs(t, err, ErrContentVacio)
	repo.AssertNotCalled(t, "ActualizarMessage", mock.Anything, mock.Anything)
	repo.AssertExpectations(t)
}

func TestMessageService_ActualizarNoEncontrado_RetornaError(t *testing.T) {
	repo := new(comunidadRepoMock)
	svc := NuevoMessageService(repo)
	msg := models.Message{Content: "Mensaje editado"}

	repo.On("ActualizarMessage", 99, msg).Return(models.Message{}, false).Once()

	_, err := svc.Actualizar(99, msg)

	require.ErrorIs(t, err, ErrNoEncontrado)
	repo.AssertExpectations(t)
}

func TestMessageService_ActualizarValido_LlegaRepositorio(t *testing.T) {
	repo := new(comunidadRepoMock)
	svc := NuevoMessageService(repo)
	msg := models.Message{Content: "Mensaje editado"}
	actualizado := msg
	actualizado.ID = 1

	repo.On("ActualizarMessage", 1, msg).Return(actualizado, true).Once()

	resultado, err := svc.Actualizar(1, msg)

	require.NoError(t, err)
	require.Equal(t, actualizado, resultado)
	repo.AssertExpectations(t)
}

func TestMessageService_BorrarNoEncontrado_RetornaError(t *testing.T) {
	repo := new(comunidadRepoMock)
	svc := NuevoMessageService(repo)

	repo.On("BorrarMessage", 99).Return(false).Once()

	err := svc.Borrar(99)

	require.ErrorIs(t, err, ErrNoEncontrado)
	repo.AssertExpectations(t)
}

func TestMessageService_BorrarExistente_NoRetornaError(t *testing.T) {
	repo := new(comunidadRepoMock)
	svc := NuevoMessageService(repo)

	repo.On("BorrarMessage", 1).Return(true).Once()

	err := svc.Borrar(1)

	require.NoError(t, err)
	repo.AssertExpectations(t)
}
