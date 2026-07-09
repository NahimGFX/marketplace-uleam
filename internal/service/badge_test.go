package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"marketplace-api/internal/models"
	"marketplace-api/internal/service"
)

func TestBadgeService_Crear(t *testing.T) {
	casos := []struct {
		nombre        string
		entrada       models.Badge
		errEsperado   error
		debePersistir bool
	}{
		{"nombre vacio rechazado", models.Badge{Name: "  ", Description: "algo", RequiredRep: 10}, service.ErrNombreVacio, false},
		{"descripcion vacia rechazada", models.Badge{Name: "Novato", Description: "  ", RequiredRep: 10}, service.ErrContentVacio, false},
		{"badge valido se persiste", models.Badge{Name: "Novato", Description: "10 puntos", RequiredRep: 10}, nil, true},
	}
	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			repo := new(perfilRepoMock)
			if c.debePersistir {
				guardado := c.entrada
				guardado.ID = 1
				repo.On("CrearBadge", c.entrada).Return(guardado)
			}
			svc := service.NuevoBadgeService(repo)

			creado, err := svc.Crear(c.entrada)

			if c.errEsperado != nil {
				require.ErrorIs(t, err, c.errEsperado)
				repo.AssertNotCalled(t, "CrearBadge")
			} else {
				require.NoError(t, err)
				assert.Equal(t, 1, creado.ID)
				repo.AssertCalled(t, "CrearBadge", c.entrada)
			}
		})
	}
}

func TestBadgeService_Obtener(t *testing.T) {
	t.Run("existe", func(t *testing.T) {
		repo := new(perfilRepoMock)
		repo.On("BuscarBadgePorID", 1).Return(models.Badge{ID: 1, Name: "Novato"}, true)
		b, err := service.NuevoBadgeService(repo).Obtener(1)
		require.NoError(t, err)
		assert.Equal(t, "Novato", b.Name)
	})
	t.Run("no existe -> ErrNoEncontrado", func(t *testing.T) {
		repo := new(perfilRepoMock)
		repo.On("BuscarBadgePorID", 999).Return(models.Badge{}, false)
		_, err := service.NuevoBadgeService(repo).Obtener(999)
		require.ErrorIs(t, err, service.ErrNoEncontrado)
	})
}

func TestBadgeService_Actualizar(t *testing.T) {
	datos := models.Badge{Name: "Experto", Description: "100 puntos", RequiredRep: 100}

	t.Run("valido", func(t *testing.T) {
		repo := new(perfilRepoMock)
		actualizado := datos
		actualizado.ID = 1
		repo.On("ActualizarBadge", 1, datos).Return(actualizado, true)
		b, err := service.NuevoBadgeService(repo).Actualizar(1, datos)
		require.NoError(t, err)
		assert.Equal(t, 1, b.ID)
	})
	t.Run("no existe -> ErrNoEncontrado", func(t *testing.T) {
		repo := new(perfilRepoMock)
		repo.On("ActualizarBadge", 999, datos).Return(models.Badge{}, false)
		_, err := service.NuevoBadgeService(repo).Actualizar(999, datos)
		require.ErrorIs(t, err, service.ErrNoEncontrado)
	})
	t.Run("invalido no toca el repo", func(t *testing.T) {
		repo := new(perfilRepoMock)
		_, err := service.NuevoBadgeService(repo).Actualizar(1, models.Badge{Name: ""})
		require.ErrorIs(t, err, service.ErrNombreVacio)
		repo.AssertNotCalled(t, "ActualizarBadge")
	})
}

func TestBadgeService_Borrar(t *testing.T) {
	t.Run("existe", func(t *testing.T) {
		repo := new(perfilRepoMock)
		repo.On("BorrarBadge", 1).Return(true)
		require.NoError(t, service.NuevoBadgeService(repo).Borrar(1))
	})
	t.Run("no existe -> ErrNoEncontrado", func(t *testing.T) {
		repo := new(perfilRepoMock)
		repo.On("BorrarBadge", 999).Return(false)
		require.ErrorIs(t, service.NuevoBadgeService(repo).Borrar(999), service.ErrNoEncontrado)
	})
}

func TestBadgeService_Listar(t *testing.T) {
	repo := new(perfilRepoMock)
	repo.On("ListarBadges").Return([]models.Badge{{ID: 1}, {ID: 2}})
	lista := service.NuevoBadgeService(repo).Listar()
	assert.Len(t, lista, 2)
	repo.AssertExpectations(t)
}
