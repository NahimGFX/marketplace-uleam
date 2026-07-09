package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"marketplace-api/internal/models"
	"marketplace-api/internal/service"
)

// Nota: usa el mismo perfilRepoMock definido en user_test.go
// (mismo paquete service_test, no hay que redeclararlo).

func TestReviewService_Crear(t *testing.T) {
	casos := []struct {
		nombre        string
		entrada       models.Review
		errEsperado   error
		debePersistir bool
	}{
		{"reviewer_id requerido", models.Review{ReviewerID: 0, ReviewedID: 2, Comment: "ok"}, service.ErrReviewerIDRequerido, false},
		{"reviewed_id requerido", models.Review{ReviewerID: 1, ReviewedID: 0, Comment: "ok"}, service.ErrReviewedIDRequerido, false},
		{"comentario vacio rechazado", models.Review{ReviewerID: 1, ReviewedID: 2, Comment: "  "}, service.ErrContentVacio, false},
		{"review valida se persiste", models.Review{ReviewerID: 1, ReviewedID: 2, Rating: 5, Comment: "Buen vendedor"}, nil, true},
	}
	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			repo := new(perfilRepoMock)
			if c.debePersistir {
				guardado := c.entrada
				guardado.ID = 1
				repo.On("CrearReview", c.entrada).Return(guardado)
			}
			svc := service.NuevoReviewService(repo)

			creado, err := svc.Crear(c.entrada)

			if c.errEsperado != nil {
				require.ErrorIs(t, err, c.errEsperado)
				repo.AssertNotCalled(t, "CrearReview")
			} else {
				require.NoError(t, err)
				assert.Equal(t, 1, creado.ID)
				repo.AssertCalled(t, "CrearReview", c.entrada)
			}
		})
	}
}

func TestReviewService_Obtener(t *testing.T) {
	t.Run("existe", func(t *testing.T) {
		repo := new(perfilRepoMock)
		repo.On("BuscarReviewPorID", 1).Return(models.Review{ID: 1, Comment: "ok"}, true)
		r, err := service.NuevoReviewService(repo).Obtener(1)
		require.NoError(t, err)
		assert.Equal(t, "ok", r.Comment)
	})
	t.Run("no existe -> ErrNoEncontrado", func(t *testing.T) {
		repo := new(perfilRepoMock)
		repo.On("BuscarReviewPorID", 999).Return(models.Review{}, false)
		_, err := service.NuevoReviewService(repo).Obtener(999)
		require.ErrorIs(t, err, service.ErrNoEncontrado)
	})
}

func TestReviewService_Actualizar(t *testing.T) {
	datos := models.Review{ReviewerID: 1, ReviewedID: 2, Rating: 4, Comment: "Editado"}

	t.Run("valido", func(t *testing.T) {
		repo := new(perfilRepoMock)
		actualizado := datos
		actualizado.ID = 1
		repo.On("ActualizarReview", 1, datos).Return(actualizado, true)
		r, err := service.NuevoReviewService(repo).Actualizar(1, datos)
		require.NoError(t, err)
		assert.Equal(t, 1, r.ID)
	})
	t.Run("no existe -> ErrNoEncontrado", func(t *testing.T) {
		repo := new(perfilRepoMock)
		repo.On("ActualizarReview", 999, datos).Return(models.Review{}, false)
		_, err := service.NuevoReviewService(repo).Actualizar(999, datos)
		require.ErrorIs(t, err, service.ErrNoEncontrado)
	})
	t.Run("invalido no toca el repo", func(t *testing.T) {
		repo := new(perfilRepoMock)
		_, err := service.NuevoReviewService(repo).Actualizar(1, models.Review{ReviewerID: 0})
		require.ErrorIs(t, err, service.ErrReviewerIDRequerido)
		repo.AssertNotCalled(t, "ActualizarReview")
	})
}

func TestReviewService_Borrar(t *testing.T) {
	t.Run("existe", func(t *testing.T) {
		repo := new(perfilRepoMock)
		repo.On("BorrarReview", 1).Return(true)
		require.NoError(t, service.NuevoReviewService(repo).Borrar(1))
	})
	t.Run("no existe -> ErrNoEncontrado", func(t *testing.T) {
		repo := new(perfilRepoMock)
		repo.On("BorrarReview", 999).Return(false)
		require.ErrorIs(t, service.NuevoReviewService(repo).Borrar(999), service.ErrNoEncontrado)
	})
}

func TestReviewService_Listar(t *testing.T) {
	repo := new(perfilRepoMock)
	repo.On("ListarReviews").Return([]models.Review{{ID: 1}, {ID: 2}})
	lista := service.NuevoReviewService(repo).Listar()
	assert.Len(t, lista, 2)
	repo.AssertExpectations(t)
}
