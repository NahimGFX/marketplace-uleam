package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"marketplace-api/internal/models"
	"marketplace-api/internal/service"
	"marketplace-api/internal/storage"
)

// comunidadRepoMock es un doble de storage.ComunidadRepository (15 metodos:
// Message + Mission + UserMission, el modulo completo).
type comunidadRepoMock struct {
	mock.Mock
}

func (m *comunidadRepoMock) ListarMessages() []models.Message {
	return m.Called().Get(0).([]models.Message)
}
func (m *comunidadRepoMock) BuscarMessagePorID(id int) (models.Message, bool) {
	a := m.Called(id)
	return a.Get(0).(models.Message), a.Bool(1)
}
func (m *comunidadRepoMock) CrearMessage(msg models.Message) models.Message {
	return m.Called(msg).Get(0).(models.Message)
}
func (m *comunidadRepoMock) ActualizarMessage(id int, datos models.Message) (models.Message, bool) {
	a := m.Called(id, datos)
	return a.Get(0).(models.Message), a.Bool(1)
}
func (m *comunidadRepoMock) BorrarMessage(id int) bool {
	return m.Called(id).Bool(0)
}

func (m *comunidadRepoMock) ListarMissions() []models.Mission {
	return m.Called().Get(0).([]models.Mission)
}
func (m *comunidadRepoMock) BuscarMisionPorID(id int) (models.Mission, bool) {
	a := m.Called(id)
	return a.Get(0).(models.Mission), a.Bool(1)
}
func (m *comunidadRepoMock) CrearMision(mi models.Mission) models.Mission {
	return m.Called(mi).Get(0).(models.Mission)
}
func (m *comunidadRepoMock) ActualizarMision(id int, datos models.Mission) (models.Mission, bool) {
	a := m.Called(id, datos)
	return a.Get(0).(models.Mission), a.Bool(1)
}
func (m *comunidadRepoMock) BorrarMision(id int) bool {
	return m.Called(id).Bool(0)
}

func (m *comunidadRepoMock) ListarUserMissions() []models.UserMission {
	return m.Called().Get(0).([]models.UserMission)
}
func (m *comunidadRepoMock) BuscarUserMissionPorID(id int) (models.UserMission, bool) {
	a := m.Called(id)
	return a.Get(0).(models.UserMission), a.Bool(1)
}
func (m *comunidadRepoMock) CrearUserMission(um models.UserMission) models.UserMission {
	return m.Called(um).Get(0).(models.UserMission)
}
func (m *comunidadRepoMock) ActualizarUserMission(id int, datos models.UserMission) (models.UserMission, bool) {
	a := m.Called(id, datos)
	return a.Get(0).(models.UserMission), a.Bool(1)
}
func (m *comunidadRepoMock) BorrarUserMission(id int) bool {
	return m.Called(id).Bool(0)
}

// Red de seguridad en tiempo de compilacion: el mock DEBE cumplir el contrato.
var _ storage.ComunidadRepository = (*comunidadRepoMock)(nil)

// --- Crear: unica validacion, Content no vacio ---

func TestMessageService_Crear(t *testing.T) {
	casos := []struct {
		nombre        string
		entrada       models.Message
		errEsperado   error
		debePersistir bool
	}{
		{"content vacio rechazado", models.Message{Content: "   "}, service.ErrContentVacio, false},
		{"message valido se persiste", models.Message{Content: "Hola, tienes el producto disponible?"}, nil, true},
	}
	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			repo := new(comunidadRepoMock)
			if c.debePersistir {
				guardado := c.entrada
				guardado.ID = 1
				repo.On("CrearMessage", c.entrada).Return(guardado)
			}
			svc := service.NuevoMessageService(repo)

			creado, err := svc.Crear(c.entrada)

			if c.errEsperado != nil {
				require.ErrorIs(t, err, c.errEsperado)
				repo.AssertNotCalled(t, "CrearMessage")
			} else {
				require.NoError(t, err)
				assert.Equal(t, 1, creado.ID)
				repo.AssertCalled(t, "CrearMessage", c.entrada)
			}
		})
	}
}

func TestMessageService_Obtener(t *testing.T) {
	t.Run("existe", func(t *testing.T) {
		repo := new(comunidadRepoMock)
		repo.On("BuscarMessagePorID", 1).Return(models.Message{ID: 1, Content: "hola"}, true)
		m, err := service.NuevoMessageService(repo).Obtener(1)
		require.NoError(t, err)
		assert.Equal(t, "hola", m.Content)
	})
	t.Run("no existe -> ErrNoEncontrado", func(t *testing.T) {
		repo := new(comunidadRepoMock)
		repo.On("BuscarMessagePorID", 999).Return(models.Message{}, false)
		_, err := service.NuevoMessageService(repo).Obtener(999)
		require.ErrorIs(t, err, service.ErrNoEncontrado)
	})
}

func TestMessageService_Actualizar(t *testing.T) {
	datos := models.Message{Content: "Editado"}

	t.Run("valido", func(t *testing.T) {
		repo := new(comunidadRepoMock)
		actualizado := datos
		actualizado.ID = 1
		repo.On("ActualizarMessage", 1, datos).Return(actualizado, true)
		m, err := service.NuevoMessageService(repo).Actualizar(1, datos)
		require.NoError(t, err)
		assert.Equal(t, 1, m.ID)
	})
	t.Run("no existe -> ErrNoEncontrado", func(t *testing.T) {
		repo := new(comunidadRepoMock)
		repo.On("ActualizarMessage", 999, datos).Return(models.Message{}, false)
		_, err := service.NuevoMessageService(repo).Actualizar(999, datos)
		require.ErrorIs(t, err, service.ErrNoEncontrado)
	})
	t.Run("invalido no toca el repo", func(t *testing.T) {
		repo := new(comunidadRepoMock)
		_, err := service.NuevoMessageService(repo).Actualizar(1, models.Message{Content: ""})
		require.ErrorIs(t, err, service.ErrContentVacio)
		repo.AssertNotCalled(t, "ActualizarMessage")
	})
}

func TestMessageService_Borrar(t *testing.T) {
	t.Run("existe", func(t *testing.T) {
		repo := new(comunidadRepoMock)
		repo.On("BorrarMessage", 1).Return(true)
		require.NoError(t, service.NuevoMessageService(repo).Borrar(1))
	})
	t.Run("no existe -> ErrNoEncontrado", func(t *testing.T) {
		repo := new(comunidadRepoMock)
		repo.On("BorrarMessage", 999).Return(false)
		require.ErrorIs(t, service.NuevoMessageService(repo).Borrar(999), service.ErrNoEncontrado)
	})
}

func TestMessageService_Listar(t *testing.T) {
	repo := new(comunidadRepoMock)
	repo.On("ListarMessages").Return([]models.Message{{ID: 1}, {ID: 2}})
	lista := service.NuevoMessageService(repo).Listar()
	assert.Len(t, lista, 2)
	repo.AssertExpectations(t)
}
