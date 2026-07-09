package service

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"marketplace-api/internal/models"
)

func TestMissionService_DescriptionVacia_NoLlegaRepositorio(t *testing.T) {
	repo := new(comunidadRepoMock)
	svc := NuevoMissionService(repo)

	_, err := svc.Crear(models.Mission{Description: "   "})

	require.ErrorIs(t, err, ErrContentVacio)
	repo.AssertNotCalled(t, "CrearMision", mock.Anything)
	repo.AssertExpectations(t)
}

func TestMissionService_CrearValida_LlegaRepositorio(t *testing.T) {
	repo := new(comunidadRepoMock)
	svc := NuevoMissionService(repo)
	mission := models.Mission{
		Title:         "Primera venta",
		Description:   "Publicar un producto",
		RequiredLevel: 1,
		RewardPoints:  20,
	}
	guardada := mission
	guardada.ID = 7

	repo.On("CrearMision", mission).Return(guardada).Once()

	resultado, err := svc.Crear(mission)

	require.NoError(t, err)
	require.Equal(t, guardada, resultado)
	repo.AssertExpectations(t)
}

func TestMissionService_ActualizarDescriptionVacia_NoLlegaRepositorio(t *testing.T) {
	repo := new(comunidadRepoMock)
	svc := NuevoMissionService(repo)

	_, err := svc.Actualizar(1, models.Mission{Description: ""})

	require.ErrorIs(t, err, ErrContentVacio)
	repo.AssertNotCalled(t, "ActualizarMision", mock.Anything, mock.Anything)
	repo.AssertExpectations(t)
}

func TestMissionService_ActualizarNoEncontrado_RetornaError(t *testing.T) {
	repo := new(comunidadRepoMock)
	svc := NuevoMissionService(repo)
	mission := models.Mission{Description: "Actualizar perfil"}

	repo.On("ActualizarMision", 99, mission).Return(models.Mission{}, false).Once()

	_, err := svc.Actualizar(99, mission)

	require.ErrorIs(t, err, ErrNoEncontrado)
	repo.AssertExpectations(t)
}

func TestMissionService_ActualizarValida_LlegaRepositorio(t *testing.T) {
	repo := new(comunidadRepoMock)
	svc := NuevoMissionService(repo)
	mission := models.Mission{Description: "Actualizar perfil"}
	actualizada := mission
	actualizada.ID = 3

	repo.On("ActualizarMision", 3, mission).Return(actualizada, true).Once()

	resultado, err := svc.Actualizar(3, mission)

	require.NoError(t, err)
	require.Equal(t, actualizada, resultado)
	repo.AssertExpectations(t)
}

func TestMissionService_BorrarNoEncontrado_RetornaError(t *testing.T) {
	repo := new(comunidadRepoMock)
	svc := NuevoMissionService(repo)

	repo.On("BorrarMision", 50).Return(false).Once()

	err := svc.Borrar(50)

	require.ErrorIs(t, err, ErrNoEncontrado)
	repo.AssertExpectations(t)
}

func TestMissionService_BorrarExistente_NoRetornaError(t *testing.T) {
	repo := new(comunidadRepoMock)
	svc := NuevoMissionService(repo)

	repo.On("BorrarMision", 5).Return(true).Once()

	err := svc.Borrar(5)

	require.NoError(t, err)
	repo.AssertExpectations(t)
}
