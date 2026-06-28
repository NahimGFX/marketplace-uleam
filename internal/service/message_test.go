package service

import (
	"testing"

	"github.com/stretchr/testify/require"

	"marketplace-api/internal/models"
	"marketplace-api/internal/storage"
)

// =======================
// FAKE REPOSITORY
// =======================

type comunidadRepoFake struct {
	llamado bool
}

func (f *comunidadRepoFake) ListarMessages() []models.Message {
	return nil
}

func (f *comunidadRepoFake) BuscarMessagePorID(id int) (models.Message, bool) {
	return models.Message{}, false
}

func (f *comunidadRepoFake) CrearMessage(m models.Message) models.Message {
	f.llamado = true
	return m
}

func (f *comunidadRepoFake) ActualizarMessage(id int, datos models.Message) (models.Message, bool) {
	return models.Message{}, true
}

func (f *comunidadRepoFake) BorrarMessage(id int) bool {
	return true
}

func (f *comunidadRepoFake) ListarMissions() []models.Mission {
	return nil
}

func (f *comunidadRepoFake) BuscarMisionPorID(id int) (models.Mission, bool) {
	return models.Mission{}, false
}

func (f *comunidadRepoFake) CrearMision(m models.Mission) models.Mission {
	return m
}

func (f *comunidadRepoFake) ActualizarMision(id int, datos models.Mission) (models.Mission, bool) {
	return models.Mission{}, true
}

func (f *comunidadRepoFake) BorrarMision(id int) bool {
	return true
}

func (f *comunidadRepoFake) ListarUserMissions() []models.UserMission {
	return nil
}

func (f *comunidadRepoFake) BuscarUserMissionPorID(id int) (models.UserMission, bool) {
	return models.UserMission{}, false
}

func (f *comunidadRepoFake) CrearUserMission(m models.UserMission) models.UserMission {
	return m
}

func (f *comunidadRepoFake) ActualizarUserMission(id int, datos models.UserMission) (models.UserMission, bool) {
	return models.UserMission{}, true
}

func (f *comunidadRepoFake) BorrarUserMission(id int) bool {
	return true
}

// Verificación de interfaz
var _ storage.ComunidadRepository = (*comunidadRepoFake)(nil)

// =======================
// TEST
// =======================

func TestMessageService_ContentVacio(t *testing.T) {
	repo := &comunidadRepoFake{}
	svc := NuevoMessageService(repo)

	_, err := svc.Crear(models.Message{
		Content: "",
	})

	require.ErrorIs(t, err, ErrContentVacio)
	require.False(t, repo.llamado, "el repositorio no debía ser llamado")
}
