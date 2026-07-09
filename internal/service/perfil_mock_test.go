package service

import (
	"github.com/stretchr/testify/mock"

	"marketplace-api/internal/models"
	"marketplace-api/internal/storage"
)

type perfilRepoMock struct {
	mock.Mock
}

func (m *perfilRepoMock) ListarUsers() []models.User {
	args := m.Called()
	return args.Get(0).([]models.User)
}

func (m *perfilRepoMock) BuscarUserPorID(id int) (models.User, bool) {
	args := m.Called(id)
	return args.Get(0).(models.User), args.Bool(1)
}

func (m *perfilRepoMock) CrearUser(u models.User) models.User {
	args := m.Called(u)
	return args.Get(0).(models.User)
}

func (m *perfilRepoMock) ActualizarUser(id int, datos models.User) (models.User, bool) {
	args := m.Called(id, datos)
	return args.Get(0).(models.User), args.Bool(1)
}

func (m *perfilRepoMock) BorrarUser(id int) bool {
	args := m.Called(id)
	return args.Bool(0)
}

func (m *perfilRepoMock) ListarReviews() []models.Review {
	args := m.Called()
	return args.Get(0).([]models.Review)
}

func (m *perfilRepoMock) BuscarReviewPorID(id int) (models.Review, bool) {
	args := m.Called(id)
	return args.Get(0).(models.Review), args.Bool(1)
}

func (m *perfilRepoMock) CrearReview(r models.Review) models.Review {
	args := m.Called(r)
	return args.Get(0).(models.Review)
}

func (m *perfilRepoMock) ActualizarReview(id int, datos models.Review) (models.Review, bool) {
	args := m.Called(id, datos)
	return args.Get(0).(models.Review), args.Bool(1)
}

func (m *perfilRepoMock) BorrarReview(id int) bool {
	args := m.Called(id)
	return args.Bool(0)
}

func (m *perfilRepoMock) ListarBadges() []models.Badge {
	args := m.Called()
	return args.Get(0).([]models.Badge)
}

func (m *perfilRepoMock) BuscarBadgePorID(id int) (models.Badge, bool) {
	args := m.Called(id)
	return args.Get(0).(models.Badge), args.Bool(1)
}

func (m *perfilRepoMock) CrearBadge(b models.Badge) models.Badge {
	args := m.Called(b)
	return args.Get(0).(models.Badge)
}

func (m *perfilRepoMock) ActualizarBadge(id int, datos models.Badge) (models.Badge, bool) {
	args := m.Called(id, datos)
	return args.Get(0).(models.Badge), args.Bool(1)
}

func (m *perfilRepoMock) BorrarBadge(id int) bool {
	args := m.Called(id)
	return args.Bool(0)
}

var _ storage.PerfilRepository = (*perfilRepoMock)(nil)
