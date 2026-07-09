package service

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"marketplace-api/internal/models"
)

func TestUserService_NombreVacio_NoLlegaRepositorio(t *testing.T) {
	repo := new(perfilRepoMock)
	svc := NuevoUserService(repo)

	_, err := svc.Crear(models.User{
		Name:     " ",
		Email:    "user@uleam.edu.ec",
		Password: "secret",
	})

	require.ErrorIs(t, err, ErrNombreVacio)
	repo.AssertNotCalled(t, "CrearUser", mock.Anything)
	repo.AssertExpectations(t)
}

func TestUserService_EmailVacio_NoLlegaRepositorio(t *testing.T) {
	repo := new(perfilRepoMock)
	svc := NuevoUserService(repo)

	_, err := svc.Crear(models.User{
		Name:     "Ana",
		Email:    " ",
		Password: "secret",
	})

	require.ErrorIs(t, err, ErrEmailVacio)
	repo.AssertNotCalled(t, "CrearUser", mock.Anything)
	repo.AssertExpectations(t)
}

func TestUserService_CrearValido_LlegaRepositorio(t *testing.T) {
	repo := new(perfilRepoMock)
	svc := NuevoUserService(repo)
	user := models.User{
		Name:     "Ana",
		Email:    "ana@uleam.edu.ec",
		Password: "secret",
	}
	guardado := user
	guardado.ID = 1

	repo.On("CrearUser", user).Return(guardado).Once()

	resultado, err := svc.Crear(user)

	require.NoError(t, err)
	require.Equal(t, guardado, resultado)
	repo.AssertExpectations(t)
}

func TestUserService_ActualizarPasswordVacia_NoLlegaRepositorio(t *testing.T) {
	repo := new(perfilRepoMock)
	svc := NuevoUserService(repo)

	_, err := svc.Actualizar(1, models.User{
		Name:     "Ana",
		Email:    "ana@uleam.edu.ec",
		Password: "",
	})

	require.ErrorIs(t, err, ErrPasswordVacia)
	repo.AssertNotCalled(t, "ActualizarUser", mock.Anything, mock.Anything)
	repo.AssertExpectations(t)
}

func TestUserService_ActualizarNoEncontrado_RetornaError(t *testing.T) {
	repo := new(perfilRepoMock)
	svc := NuevoUserService(repo)
	user := models.User{
		Name:     "Ana",
		Email:    "ana@uleam.edu.ec",
		Password: "secret",
	}

	repo.On("ActualizarUser", 99, user).Return(models.User{}, false).Once()

	_, err := svc.Actualizar(99, user)

	require.ErrorIs(t, err, ErrNoEncontrado)
	repo.AssertExpectations(t)
}

func TestUserService_BorrarNoEncontrado_RetornaError(t *testing.T) {
	repo := new(perfilRepoMock)
	svc := NuevoUserService(repo)

	repo.On("BorrarUser", 99).Return(false).Once()

	err := svc.Borrar(99)

	require.ErrorIs(t, err, ErrNoEncontrado)
	repo.AssertExpectations(t)
}
