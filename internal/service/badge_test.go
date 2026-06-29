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

type perfilRepoFake struct {
	llamado bool
}

// =======================
// BADGES
// =======================

func (f *perfilRepoFake) ListarBadges() []models.Badge {
	return nil
}

func (f *perfilRepoFake) BuscarBadgePorID(id int) (models.Badge, bool) {
	return models.Badge{}, false
}

func (f *perfilRepoFake) CrearBadge(b models.Badge) models.Badge {
	f.llamado = true
	return b
}

func (f *perfilRepoFake) ActualizarBadge(id int, datos models.Badge) (models.Badge, bool) {
	return models.Badge{}, true
}

func (f *perfilRepoFake) BorrarBadge(id int) bool {
	return true
}

// =======================
// USERS
// =======================

func (f *perfilRepoFake) ListarUsers() []models.User {
	return nil
}

func (f *perfilRepoFake) BuscarUserPorID(id int) (models.User, bool) {
	return models.User{}, false
}

func (f *perfilRepoFake) CrearUser(u models.User) models.User {
	return u
}

func (f *perfilRepoFake) ActualizarUser(id int, datos models.User) (models.User, bool) {
	return models.User{}, true
}

func (f *perfilRepoFake) BorrarUser(id int) bool {
	return true
}

// =======================
// REVIEWS
// =======================

func (f *perfilRepoFake) ListarReviews() []models.Review {
	return nil
}

func (f *perfilRepoFake) BuscarReviewPorID(id int) (models.Review, bool) {
	return models.Review{}, false
}

func (f *perfilRepoFake) CrearReview(r models.Review) models.Review {
	return r
}

func (f *perfilRepoFake) ActualizarReview(id int, datos models.Review) (models.Review, bool) {
	return models.Review{}, true
}

func (f *perfilRepoFake) BorrarReview(id int) bool {
	return true
}

// Verificación de interfaz
var _ storage.PerfilRepository = (*perfilRepoFake)(nil)

// =======================
// TEST
// =======================

func TestBadgeService_NameVacio(t *testing.T) {
	repo := &perfilRepoFake{}
	svc := NuevoBadgeService(repo)

	_, err := svc.Crear(models.Badge{
		Name:        "",
		Description: "Badge de prueba",
		RequiredRep: 100,
	})

	require.ErrorIs(t, err, ErrNombreVacio)
	require.False(t, repo.llamado, "el repositorio no debía ser llamado")
}
