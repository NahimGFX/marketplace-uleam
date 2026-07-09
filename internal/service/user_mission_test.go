package service

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"marketplace-api/internal/models"
)

func TestUserMissionService_UserIDRequerido_NoLlegaRepositorio(t *testing.T) {
	repo := new(comunidadRepoMock)
	svc := NuevoUserMissionService(repo)

	_, err := svc.Crear(models.UserMission{
		UserID:    0,
		MissionID: 3,
	})

	require.ErrorIs(t, err, ErrUserIDRequerido)
	repo.AssertNotCalled(t, "CrearUserMission", mock.Anything)
	repo.AssertExpectations(t)
}

func TestUserMissionService_MissionIDRequerido_NoLlegaRepositorio(t *testing.T) {
	repo := new(comunidadRepoMock)
	svc := NuevoUserMissionService(repo)

	_, err := svc.Crear(models.UserMission{
		UserID:    2,
		MissionID: 0,
	})

	require.ErrorIs(t, err, ErrMissionIDRequerido)
	repo.AssertNotCalled(t, "CrearUserMission", mock.Anything)
	repo.AssertExpectations(t)
}

func TestUserMissionService_CrearValida_LlegaRepositorio(t *testing.T) {
	repo := new(comunidadRepoMock)
	svc := NuevoUserMissionService(repo)
	userMission := models.UserMission{
		UserID:    2,
		MissionID: 3,
		Completed: false,
	}
	guardada := userMission
	guardada.ID = 4

	repo.On("CrearUserMission", userMission).Return(guardada).Once()

	resultado, err := svc.Crear(userMission)

	require.NoError(t, err)
	require.Equal(t, guardada, resultado)
	repo.AssertExpectations(t)
}

func TestUserMissionService_ActualizarMissionIDRequerido_NoLlegaRepositorio(t *testing.T) {
	repo := new(comunidadRepoMock)
	svc := NuevoUserMissionService(repo)

	_, err := svc.Actualizar(1, models.UserMission{UserID: 2})

	require.ErrorIs(t, err, ErrMissionIDRequerido)
	repo.AssertNotCalled(t, "ActualizarUserMission", mock.Anything, mock.Anything)
	repo.AssertExpectations(t)
}

func TestUserMissionService_ActualizarNoEncontrado_RetornaError(t *testing.T) {
	repo := new(comunidadRepoMock)
	svc := NuevoUserMissionService(repo)
	userMission := models.UserMission{
		UserID:    2,
		MissionID: 3,
		Completed: true,
	}

	repo.On("ActualizarUserMission", 99, userMission).Return(models.UserMission{}, false).Once()

	_, err := svc.Actualizar(99, userMission)

	require.ErrorIs(t, err, ErrNoEncontrado)
	repo.AssertExpectations(t)
}

func TestUserMissionService_ActualizarValida_LlegaRepositorio(t *testing.T) {
	repo := new(comunidadRepoMock)
	svc := NuevoUserMissionService(repo)
	userMission := models.UserMission{
		UserID:    2,
		MissionID: 3,
		Completed: true,
	}
	actualizada := userMission
	actualizada.ID = 9

	repo.On("ActualizarUserMission", 9, userMission).Return(actualizada, true).Once()

	resultado, err := svc.Actualizar(9, userMission)

	require.NoError(t, err)
	require.Equal(t, actualizada, resultado)
	repo.AssertExpectations(t)
}

func TestUserMissionService_BorrarNoEncontrado_RetornaError(t *testing.T) {
	repo := new(comunidadRepoMock)
	svc := NuevoUserMissionService(repo)

	repo.On("BorrarUserMission", 100).Return(false).Once()

	err := svc.Borrar(100)

	require.ErrorIs(t, err, ErrNoEncontrado)
	repo.AssertExpectations(t)
}

func TestUserMissionService_BorrarExistente_NoRetornaError(t *testing.T) {
	repo := new(comunidadRepoMock)
	svc := NuevoUserMissionService(repo)

	repo.On("BorrarUserMission", 6).Return(true).Once()

	err := svc.Borrar(6)

	require.NoError(t, err)
	repo.AssertExpectations(t)
}
