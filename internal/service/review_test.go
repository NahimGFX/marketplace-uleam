package service

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"marketplace-api/internal/models"
)

func TestReviewService_ReviewerIDRequerido_NoLlegaRepositorio(t *testing.T) {
	repo := new(perfilRepoMock)
	svc := NuevoReviewService(repo)

	_, err := svc.Crear(models.Review{
		ReviewedID: 2,
		Comment:    "Buen vendedor",
	})

	require.ErrorIs(t, err, ErrReviewerIDRequerido)
	repo.AssertNotCalled(t, "CrearReview", mock.Anything)
	repo.AssertExpectations(t)
}

func TestReviewService_ReviewedIDRequerido_NoLlegaRepositorio(t *testing.T) {
	repo := new(perfilRepoMock)
	svc := NuevoReviewService(repo)

	_, err := svc.Crear(models.Review{
		ReviewerID: 1,
		Comment:    "Buen vendedor",
	})

	require.ErrorIs(t, err, ErrReviewedIDRequerido)
	repo.AssertNotCalled(t, "CrearReview", mock.Anything)
	repo.AssertExpectations(t)
}

func TestReviewService_CommentVacio_NoLlegaRepositorio(t *testing.T) {
	repo := new(perfilRepoMock)
	svc := NuevoReviewService(repo)

	_, err := svc.Crear(models.Review{
		ReviewerID: 1,
		ReviewedID: 2,
		Comment:    " ",
	})

	require.ErrorIs(t, err, ErrContentVacio)
	repo.AssertNotCalled(t, "CrearReview", mock.Anything)
	repo.AssertExpectations(t)
}

func TestReviewService_CrearValida_LlegaRepositorio(t *testing.T) {
	repo := new(perfilRepoMock)
	svc := NuevoReviewService(repo)
	review := models.Review{
		ReviewerID: 1,
		ReviewedID: 2,
		Rating:     5,
		Comment:    "Buen vendedor",
	}
	guardada := review
	guardada.ID = 4

	repo.On("CrearReview", review).Return(guardada).Once()

	resultado, err := svc.Crear(review)

	require.NoError(t, err)
	require.Equal(t, guardada, resultado)
	repo.AssertExpectations(t)
}

func TestReviewService_ActualizarNoEncontrado_RetornaError(t *testing.T) {
	repo := new(perfilRepoMock)
	svc := NuevoReviewService(repo)
	review := models.Review{
		ReviewerID: 1,
		ReviewedID: 2,
		Comment:    "Actualizado",
	}

	repo.On("ActualizarReview", 99, review).Return(models.Review{}, false).Once()

	_, err := svc.Actualizar(99, review)

	require.ErrorIs(t, err, ErrNoEncontrado)
	repo.AssertExpectations(t)
}

func TestReviewService_BorrarNoEncontrado_RetornaError(t *testing.T) {
	repo := new(perfilRepoMock)
	svc := NuevoReviewService(repo)

	repo.On("BorrarReview", 99).Return(false).Once()

	err := svc.Borrar(99)

	require.ErrorIs(t, err, ErrNoEncontrado)
	repo.AssertExpectations(t)
}
