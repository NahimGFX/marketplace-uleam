package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"marketplace-api/internal/handlers"
	"marketplace-api/internal/middleware"
	"marketplace-api/internal/models"
	"marketplace-api/internal/service"
	"marketplace-api/internal/storage"
)

// ======================================================
// FAKE
// ======================================================

type profileFake struct {
	badges []models.Badge
	nextID int
}

// ---------------- BADGES ----------------

func (f *profileFake) ListarBadges() []models.Badge {
	return f.badges
}

func (f *profileFake) BuscarBadgePorID(id int) (models.Badge, bool) {
	for _, b := range f.badges {
		if b.ID == id {
			return b, true
		}
	}
	return models.Badge{}, false
}

func (f *profileFake) CrearBadge(b models.Badge) models.Badge {
	f.nextID++
	b.ID = f.nextID
	f.badges = append(f.badges, b)
	return b
}

func (f *profileFake) ActualizarBadge(id int, b models.Badge) (models.Badge, bool) {
	for i := range f.badges {
		if f.badges[i].ID == id {
			b.ID = id
			f.badges[i] = b
			return b, true
		}
	}
	return models.Badge{}, false
}

func (f *profileFake) BorrarBadge(id int) bool {
	return true
}

// ---------------- USERS ----------------

func (f *profileFake) ListarUsers() []models.User { return nil }

func (f *profileFake) BuscarUserPorID(id int) (models.User, bool) {
	return models.User{}, false
}

func (f *profileFake) CrearUser(u models.User) models.User {
	return u
}

func (f *profileFake) ActualizarUser(id int, u models.User) (models.User, bool) {
	return u, false
}

func (f *profileFake) BorrarUser(id int) bool {
	return false
}

// ---------------- REVIEWS ----------------

func (f *profileFake) ListarReviews() []models.Review {
	return nil
}

func (f *profileFake) BuscarReviewPorID(id int) (models.Review, bool) {
	return models.Review{}, false
}

func (f *profileFake) CrearReview(r models.Review) models.Review {
	return r
}

func (f *profileFake) ActualizarReview(id int, r models.Review) (models.Review, bool) {
	return r, false
}

func (f *profileFake) BorrarReview(id int) bool {
	return false
}

var _ storage.PerfilRepository = (*profileFake)(nil)

// ======================================================
// JWT
// ======================================================

func genJWTValido() string {

	secret := []byte("marketplace-uleam-secreto-demo-cambiar-en-S12")

	claims := service.Claims{
		UsuarioID: 1,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	s, _ := token.SignedString(secret)

	return s
}

// ======================================================
// ROUTER
// ======================================================

func constEntorno() http.Handler {

	repo := &profileFake{}

	badgeSvc := service.NuevoBadgeService(repo)
	authSvc := service.NuevoAuthService(nil)

	server := handlers.NewServer(
		nil,
		nil,
		badgeSvc,
		nil,
		nil,
		nil,
		authSvc,
	)

	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router) {

		r.Group(func(r chi.Router) {

			r.Use(middleware.Auth(authSvc))

			r.Post("/badges", server.CrearBadge)
			r.Get("/badges", server.ListarBadges)

		})

	})

	return r
}

// ======================================================
// TEST CREAR
// ======================================================

func TestCrearBadge_OK(t *testing.T) {

	h := construirEntorno()

	token := generarJWTValido()

	body := `{
		"name":"Experto",
		"description":"100 puntos",
		"required_rep":100
	}`

	req := httptest.NewRequest(http.MethodPost,
		"/api/v1/badges",
		strings.NewReader(body))

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	require.Equal(t, http.StatusCreated, rec.Code)

	var badge models.Badge

	require.NoError(t,
		json.NewDecoder(rec.Body).Decode(&badge))

	assert.Equal(t, "Experto", badge.Name)
}

// ======================================================
// TEST LISTAR
// ======================================================

func TestListarBadges(t *testing.T) {

	h := construirEntorno()

	token := generarJWTValido()

	req := httptest.NewRequest(http.MethodGet,
		"/api/v1/badges",
		nil)

	req.Header.Set("Authorization", "Bearer "+token)

	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}

// ======================================================
// TEST 401
// ======================================================

func Test_RutaProtegida_SinToken(t *testing.T) {

	h := construirEntorno()

	req := httptest.NewRequest(http.MethodPost,
		"/api/v1/badges",
		strings.NewReader(`{"name":"Experto"}`))

	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}
