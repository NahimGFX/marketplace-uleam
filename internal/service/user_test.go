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

// perfilRepoMock es un doble de storage.PerfilRepository (15 metodos:
// User + Review + Badge, el modulo completo). Cada metodo solo registra
// la llamada y devuelve lo que el test configuro con On(...).
type perfilRepoMock struct {
	mock.Mock
}

func (m *perfilRepoMock) ListarUsers() []models.User {
	return m.Called().Get(0).([]models.User)
}
func (m *perfilRepoMock) BuscarUserPorID(id int) (models.User, bool) {
	a := m.Called(id)
	return a.Get(0).(models.User), a.Bool(1)
}
func (m *perfilRepoMock) CrearUser(u models.User) models.User {
	return m.Called(u).Get(0).(models.User)
}
func (m *perfilRepoMock) ActualizarUser(id int, datos models.User) (models.User, bool) {
	a := m.Called(id, datos)
	return a.Get(0).(models.User), a.Bool(1)
}
func (m *perfilRepoMock) BorrarUser(id int) bool {
	return m.Called(id).Bool(0)
}

func (m *perfilRepoMock) ListarReviews() []models.Review {
	return m.Called().Get(0).([]models.Review)
}
func (m *perfilRepoMock) BuscarReviewPorID(id int) (models.Review, bool) {
	a := m.Called(id)
	return a.Get(0).(models.Review), a.Bool(1)
}
func (m *perfilRepoMock) CrearReview(r models.Review) models.Review {
	return m.Called(r).Get(0).(models.Review)
}
func (m *perfilRepoMock) ActualizarReview(id int, datos models.Review) (models.Review, bool) {
	a := m.Called(id, datos)
	return a.Get(0).(models.Review), a.Bool(1)
}
func (m *perfilRepoMock) BorrarReview(id int) bool {
	return m.Called(id).Bool(0)
}

func (m *perfilRepoMock) ListarBadges() []models.Badge {
	return m.Called().Get(0).([]models.Badge)
}
func (m *perfilRepoMock) BuscarBadgePorID(id int) (models.Badge, bool) {
	a := m.Called(id)
	return a.Get(0).(models.Badge), a.Bool(1)
}
func (m *perfilRepoMock) CrearBadge(b models.Badge) models.Badge {
	return m.Called(b).Get(0).(models.Badge)
}
func (m *perfilRepoMock) ActualizarBadge(id int, datos models.Badge) (models.Badge, bool) {
	a := m.Called(id, datos)
	return a.Get(0).(models.Badge), a.Bool(1)
}
func (m *perfilRepoMock) BorrarBadge(id int) bool {
	return m.Called(id).Bool(0)
}

// Red de seguridad en tiempo de compilacion: el mock DEBE cumplir el contrato.
var _ storage.PerfilRepository = (*perfilRepoMock)(nil)

// --- Crear: 3 validaciones distintas en UserService.validarUser ---

func TestUserService_Crear(t *testing.T) {
	casos := []struct {
		nombre        string
		entrada       models.User
		errEsperado   error
		debePersistir bool
	}{
		{"nombre vacio rechazado", models.User{Name: "  ", Email: "a@uleam.edu.ec", Password: "123"}, service.ErrNombreVacio, false},
		{"email vacio rechazado", models.User{Name: "Juan", Email: "  ", Password: "123"}, service.ErrEmailVacio, false},
		{"password vacia rechazada", models.User{Name: "Juan", Email: "a@uleam.edu.ec", Password: "  "}, service.ErrPasswordVacia, false},
		{"usuario valido se persiste", models.User{Name: "Juan", Email: "a@uleam.edu.ec", Password: "123"}, nil, true},
	}
	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			repo := new(perfilRepoMock)
			if c.debePersistir {
				guardado := c.entrada
				guardado.ID = 1
				repo.On("CrearUser", c.entrada).Return(guardado)
			}
			svc := service.NuevoUserService(repo)

			creado, err := svc.Crear(c.entrada)

			if c.errEsperado != nil {
				require.ErrorIs(t, err, c.errEsperado)
				repo.AssertNotCalled(t, "CrearUser")
			} else {
				require.NoError(t, err)
				assert.Equal(t, 1, creado.ID)
				repo.AssertCalled(t, "CrearUser", c.entrada)
			}
		})
	}
}

// --- Obtener ---

func TestUserService_Obtener(t *testing.T) {
	t.Run("existe", func(t *testing.T) {
		repo := new(perfilRepoMock)
		repo.On("BuscarUserPorID", 1).Return(models.User{ID: 1, Name: "Juan"}, true)
		u, err := service.NuevoUserService(repo).Obtener(1)
		require.NoError(t, err)
		assert.Equal(t, "Juan", u.Name)
	})
	t.Run("no existe -> ErrNoEncontrado", func(t *testing.T) {
		repo := new(perfilRepoMock)
		repo.On("BuscarUserPorID", 999).Return(models.User{}, false)
		_, err := service.NuevoUserService(repo).Obtener(999)
		require.ErrorIs(t, err, service.ErrNoEncontrado)
	})
}

// --- Actualizar ---

func TestUserService_Actualizar(t *testing.T) {
	datos := models.User{Name: "Juan Editado", Email: "a@uleam.edu.ec", Password: "123"}

	t.Run("valido", func(t *testing.T) {
		repo := new(perfilRepoMock)
		actualizado := datos
		actualizado.ID = 1
		repo.On("ActualizarUser", 1, datos).Return(actualizado, true)
		u, err := service.NuevoUserService(repo).Actualizar(1, datos)
		require.NoError(t, err)
		assert.Equal(t, 1, u.ID)
	})
	t.Run("no existe -> ErrNoEncontrado", func(t *testing.T) {
		repo := new(perfilRepoMock)
		repo.On("ActualizarUser", 999, datos).Return(models.User{}, false)
		_, err := service.NuevoUserService(repo).Actualizar(999, datos)
		require.ErrorIs(t, err, service.ErrNoEncontrado)
	})
	t.Run("invalido no toca el repo", func(t *testing.T) {
		repo := new(perfilRepoMock)
		_, err := service.NuevoUserService(repo).Actualizar(1, models.User{Name: ""})
		require.ErrorIs(t, err, service.ErrNombreVacio)
		repo.AssertNotCalled(t, "ActualizarUser")
	})
}

// --- Borrar ---

func TestUserService_Borrar(t *testing.T) {
	t.Run("existe", func(t *testing.T) {
		repo := new(perfilRepoMock)
		repo.On("BorrarUser", 1).Return(true)
		require.NoError(t, service.NuevoUserService(repo).Borrar(1))
	})
	t.Run("no existe -> ErrNoEncontrado", func(t *testing.T) {
		repo := new(perfilRepoMock)
		repo.On("BorrarUser", 999).Return(false)
		require.ErrorIs(t, service.NuevoUserService(repo).Borrar(999), service.ErrNoEncontrado)
	})
}

// --- Listar ---

func TestUserService_Listar(t *testing.T) {
	repo := new(perfilRepoMock)
	repo.On("ListarUsers").Return([]models.User{{ID: 1}, {ID: 2}})
	lista := service.NuevoUserService(repo).Listar()
	assert.Len(t, lista, 2)
	repo.AssertExpectations(t)
}
