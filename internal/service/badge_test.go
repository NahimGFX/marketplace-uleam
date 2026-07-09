package service

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"marketplace-api/internal/models"
)

func TestBadgeService_NameVacio_NoLlegaRepositorio(t *testing.T) {
	repo := new(perfilRepoMock)
	svc := NuevoBadgeService(repo)

	_, err := svc.Crear(models.Badge{
		Name:        "",
		Description: "Badge de prueba",
		RequiredRep: 100,
	})

	require.ErrorIs(t, err, ErrNombreVacio)
	repo.AssertNotCalled(t, "CrearBadge", mock.Anything)
	repo.AssertExpectations(t)
}

func TestBadgeService_DescriptionVacia_NoLlegaRepositorio(t *testing.T) {
	repo := new(perfilRepoMock)
	svc := NuevoBadgeService(repo)

	_, err := svc.Crear(models.Badge{
		Name:        "Colaborador",
		Description: " ",
		RequiredRep: 100,
	})

	require.ErrorIs(t, err, ErrContentVacio)
	repo.AssertNotCalled(t, "CrearBadge", mock.Anything)
	repo.AssertExpectations(t)
}

func TestBadgeService_CrearValido_LlegaRepositorio(t *testing.T) {
	repo := new(perfilRepoMock)
	svc := NuevoBadgeService(repo)
	badge := models.Badge{
		Name:        "Colaborador",
		Description: "Ayuda a otros usuarios",
		RequiredRep: 100,
	}
	guardado := badge
	guardado.ID = 3

	repo.On("CrearBadge", badge).Return(guardado).Once()

	resultado, err := svc.Crear(badge)

	require.NoError(t, err)
	require.Equal(t, guardado, resultado)
	repo.AssertExpectations(t)
}

func TestBadgeService_ActualizarNoEncontrado_RetornaError(t *testing.T) {
	repo := new(perfilRepoMock)
	svc := NuevoBadgeService(repo)
	badge := models.Badge{
		Name:        "Colaborador",
		Description: "Ayuda a otros usuarios",
	}

	repo.On("ActualizarBadge", 99, badge).Return(models.Badge{}, false).Once()

	_, err := svc.Actualizar(99, badge)

	require.ErrorIs(t, err, ErrNoEncontrado)
	repo.AssertExpectations(t)
}

func TestBadgeService_BorrarNoEncontrado_RetornaError(t *testing.T) {
	repo := new(perfilRepoMock)
	svc := NuevoBadgeService(repo)

	repo.On("BorrarBadge", 99).Return(false).Once()

	err := svc.Borrar(99)

	require.ErrorIs(t, err, ErrNoEncontrado)
	repo.AssertExpectations(t)
}
