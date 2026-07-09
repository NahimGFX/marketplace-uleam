package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"marketplace-api/internal/models"
	"marketplace-api/internal/service"
)

// Usa el mismo comunidadRepoMock definido en message_test.go.

func TestMissionService_Crear(t *testing.T) {
	casos := []struct {
		nombre        string
		entrada       models.Mission
		errEsperado   error
		debePersistir bool
	}{
		{"descripcion vacia rechazada", models.Mission{Description: "   "}, service.ErrContentVacio, false},
		{"mision valida se persiste", models.Mission{Description: "Vende 5 productos esta semana"}, nil, true},
	}
	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			repo := new(comunidadRepoMock)
			if c.debePersistir {
				guardado := c.entrada
				guardado.ID = 1
				repo.On("CrearMision", c.entrada).Return(guardado)
			}
			svc := service.NuevoMissionService(repo)

			creado, err := svc.Crear(c.entrada)

			if c.errEsperado != nil {
				require.ErrorIs(t, err, c.errEsperado)
				repo.AssertNotCalled(t, "CrearMision")
			} else {
				require.NoError(t, err)
				assert.Equal(t, 1, creado.ID)
				repo.AssertCalled(t, "CrearMision", c.entrada)
			}
		})
	}
}

func TestMissionService_Obtener(t *testing.T) {
	t.Run("existe", func(t *testing.T) {
		repo := new(comunidadRepoMock)
		repo.On("BuscarMisionPorID", 1).Return(models.Mission{ID: 1, Description: "Vende 5"}, true)
		mi, err := service.NuevoMissionService(repo).Obtener(1)
		require.NoError(t, err)
		assert.Equal(t, "Vende 5", mi.Description)
	})
	t.Run("no existe -> ErrNoEncontrado", func(t *testing.T) {
		repo := new(comunidadRepoMock)
		repo.On("BuscarMisionPorID", 999).Return(models.Mission{}, false)
		_, err := service.NuevoMissionService(repo).Obtener(999)
		require.ErrorIs(t, err, service.ErrNoEncontrado)
	})
}

func TestMissionService_Actualizar(t *testing.T) {
	datos := models.Mission{Description: "Editada"}

	t.Run("valido", func(t *testing.T) {
		repo := new(comunidadRepoMock)
		actualizado := datos
		actualizado.ID = 1
		repo.On("ActualizarMision", 1, datos).Return(actualizado, true)
		mi, err := service.NuevoMissionService(repo).Actualizar(1, datos)
		require.NoError(t, err)
		assert.Equal(t, 1, mi.ID)
	})
	t.Run("no existe -> ErrNoEncontrado", func(t *testing.T) {
		repo := new(comunidadRepoMock)
		repo.On("ActualizarMision", 999, datos).Return(models.Mission{}, false)
		_, err := service.NuevoMissionService(repo).Actualizar(999, datos)
		require.ErrorIs(t, err, service.ErrNoEncontrado)
	})
	t.Run("invalido no toca el repo", func(t *testing.T) {
		repo := new(comunidadRepoMock)
		_, err := service.NuevoMissionService(repo).Actualizar(1, models.Mission{Description: ""})
		require.ErrorIs(t, err, service.ErrContentVacio)
		repo.AssertNotCalled(t, "ActualizarMision")
	})
}

func TestMissionService_Borrar(t *testing.T) {
	t.Run("existe", func(t *testing.T) {
		repo := new(comunidadRepoMock)
		repo.On("BorrarMision", 1).Return(true)
		require.NoError(t, service.NuevoMissionService(repo).Borrar(1))
	})
	t.Run("no existe -> ErrNoEncontrado", func(t *testing.T) {
		repo := new(comunidadRepoMock)
		repo.On("BorrarMision", 999).Return(false)
		require.ErrorIs(t, service.NuevoMissionService(repo).Borrar(999), service.ErrNoEncontrado)
	})
}

func TestMissionService_Listar(t *testing.T) {
	repo := new(comunidadRepoMock)
	repo.On("ListarMissions").Return([]models.Mission{{ID: 1}, {ID: 2}})
	lista := service.NuevoMissionService(repo).Listar()
	assert.Len(t, lista, 2)
	repo.AssertExpectations(t)
}
