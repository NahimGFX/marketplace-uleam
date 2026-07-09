package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"marketplace-api/internal/models"
	"marketplace-api/internal/service"
)

// Usa el mismo comunidadRepoMock definido en message_test.go.

func TestUserMissionService_Crear(t *testing.T) {
	casos := []struct {
		nombre        string
		entrada       models.UserMission
		errEsperado   error
		debePersistir bool
	}{
		{"user_id requerido", models.UserMission{UserID: 0, MissionID: 1}, service.ErrUserIDRequerido, false},
		{"mission_id requerido", models.UserMission{UserID: 1, MissionID: 0}, service.ErrMissionIDRequerido, false},
		{"usermission valida se persiste", models.UserMission{UserID: 1, MissionID: 2}, nil, true},
	}
	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			repo := new(comunidadRepoMock)
			if c.debePersistir {
				guardado := c.entrada
				guardado.ID = 1
				repo.On("CrearUserMission", c.entrada).Return(guardado)
			}
			svc := service.NuevoUserMissionService(repo)

			creado, err := svc.Crear(c.entrada)

			if c.errEsperado != nil {
				require.ErrorIs(t, err, c.errEsperado)
				repo.AssertNotCalled(t, "CrearUserMission")
			} else {
				require.NoError(t, err)
				assert.Equal(t, 1, creado.ID)
				repo.AssertCalled(t, "CrearUserMission", c.entrada)
			}
		})
	}
}

func TestUserMissionService_Obtener(t *testing.T) {
	t.Run("existe", func(t *testing.T) {
		repo := new(comunidadRepoMock)
		repo.On("BuscarUserMissionPorID", 1).Return(models.UserMission{ID: 1, UserID: 1, MissionID: 2}, true)
		um, err := service.NuevoUserMissionService(repo).Obtener(1)
		require.NoError(t, err)
		assert.Equal(t, 2, um.MissionID)
	})
	t.Run("no existe -> ErrNoEncontrado", func(t *testing.T) {
		repo := new(comunidadRepoMock)
		repo.On("BuscarUserMissionPorID", 999).Return(models.UserMission{}, false)
		_, err := service.NuevoUserMissionService(repo).Obtener(999)
		require.ErrorIs(t, err, service.ErrNoEncontrado)
	})
}

func TestUserMissionService_Actualizar(t *testing.T) {
	datos := models.UserMission{UserID: 1, MissionID: 3}

	t.Run("valido", func(t *testing.T) {
		repo := new(comunidadRepoMock)
		actualizado := datos
		actualizado.ID = 1
		repo.On("ActualizarUserMission", 1, datos).Return(actualizado, true)
		um, err := service.NuevoUserMissionService(repo).Actualizar(1, datos)
		require.NoError(t, err)
		assert.Equal(t, 1, um.ID)
	})
	t.Run("no existe -> ErrNoEncontrado", func(t *testing.T) {
		repo := new(comunidadRepoMock)
		repo.On("ActualizarUserMission", 999, datos).Return(models.UserMission{}, false)
		_, err := service.NuevoUserMissionService(repo).Actualizar(999, datos)
		require.ErrorIs(t, err, service.ErrNoEncontrado)
	})
	t.Run("invalido no toca el repo", func(t *testing.T) {
		repo := new(comunidadRepoMock)
		_, err := service.NuevoUserMissionService(repo).Actualizar(1, models.UserMission{UserID: 0})
		require.ErrorIs(t, err, service.ErrUserIDRequerido)
		repo.AssertNotCalled(t, "ActualizarUserMission")
	})
}

func TestUserMissionService_Borrar(t *testing.T) {
	t.Run("existe", func(t *testing.T) {
		repo := new(comunidadRepoMock)
		repo.On("BorrarUserMission", 1).Return(true)
		require.NoError(t, service.NuevoUserMissionService(repo).Borrar(1))
	})
	t.Run("no existe -> ErrNoEncontrado", func(t *testing.T) {
		repo := new(comunidadRepoMock)
		repo.On("BorrarUserMission", 999).Return(false)
		require.ErrorIs(t, service.NuevoUserMissionService(repo).Borrar(999), service.ErrNoEncontrado)
	})
}

func TestUserMissionService_Listar(t *testing.T) {
	repo := new(comunidadRepoMock)
	repo.On("ListarUserMissions").Return([]models.UserMission{{ID: 1}, {ID: 2}})
	lista := service.NuevoUserMissionService(repo).Listar()
	assert.Len(t, lista, 2)
	repo.AssertExpectations(t)
}
