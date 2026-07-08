package service

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"marketplace-api/internal/models"
	"marketplace-api/internal/storage"
)

// =======================
// MOCK REPOSITORY TESTIFY
// =======================

type comunidadRepoMock struct {
	mock.Mock
}

// -------- MESSAGE --------

func (m *comunidadRepoMock) ListarMessages() []models.Message {

	args := m.Called()

	return args.Get(0).([]models.Message)
}

func (m *comunidadRepoMock) BuscarMessagePorID(id int) (models.Message, bool) {

	args := m.Called(id)

	return args.Get(0).(models.Message), args.Bool(1)
}

func (m *comunidadRepoMock) CrearMessage(msg models.Message) models.Message {

	args := m.Called(msg)

	return args.Get(0).(models.Message)
}

func (m *comunidadRepoMock) ActualizarMessage(
	id int,
	msg models.Message,
) (models.Message, bool) {

	args := m.Called(id, msg)

	return args.Get(0).(models.Message), args.Bool(1)
}

func (m *comunidadRepoMock) BorrarMessage(id int) bool {

	args := m.Called(id)

	return args.Bool(0)
}

// -------- MISSION --------

func (m *comunidadRepoMock) ListarMissions() []models.Mission {

	args := m.Called()

	return args.Get(0).([]models.Mission)
}

func (m *comunidadRepoMock) BuscarMisionPorID(id int) (models.Mission, bool) {

	args := m.Called(id)

	return args.Get(0).(models.Mission), args.Bool(1)
}

func (m *comunidadRepoMock) CrearMision(
	mission models.Mission,
) models.Mission {

	args := m.Called(mission)

	return args.Get(0).(models.Mission)
}

func (m *comunidadRepoMock) ActualizarMision(
	id int,
	mission models.Mission,
) (models.Mission, bool) {

	args := m.Called(id, mission)

	return args.Get(0).(models.Mission), args.Bool(1)
}

func (m *comunidadRepoMock) BorrarMision(id int) bool {

	args := m.Called(id)

	return args.Bool(0)
}

// -------- USER MISSION --------

func (m *comunidadRepoMock) ListarUserMissions() []models.UserMission {

	args := m.Called()

	return args.Get(0).([]models.UserMission)
}

func (m *comunidadRepoMock) BuscarUserMissionPorID(
	id int,
) (models.UserMission, bool) {

	args := m.Called(id)

	return args.Get(0).(models.UserMission), args.Bool(1)
}

func (m *comunidadRepoMock) CrearUserMission(
	um models.UserMission,
) models.UserMission {

	args := m.Called(um)

	return args.Get(0).(models.UserMission)
}

func (m *comunidadRepoMock) ActualizarUserMission(
	id int,
	um models.UserMission,
) (models.UserMission, bool) {

	args := m.Called(id, um)

	return args.Get(0).(models.UserMission), args.Bool(1)
}

func (m *comunidadRepoMock) BorrarUserMission(id int) bool {

	args := m.Called(id)

	return args.Bool(0)
}

// verificar interfaz
var _ storage.ComunidadRepository = (*comunidadRepoMock)(nil)

// =======================
// TEST REGLA NEGOCIO
// =======================

func TestMessageService_ContentVacio_NoLlegaRepositorio(t *testing.T) {

	repo := new(comunidadRepoMock)

	service := NuevoMessageService(repo)

	_, err := service.Crear(
		models.Message{
			Content: "",
		},
	)

	require.ErrorIs(
		t,
		err,
		ErrContentVacio,
	)

	// IMPORTANTE:
	// No ponemos repo.On("CrearMessage")
	// porque el repositorio NO debe ejecutarse.

	repo.AssertNotCalled(
		t,
		"CrearMessage",
		mock.Anything,
	)

}
