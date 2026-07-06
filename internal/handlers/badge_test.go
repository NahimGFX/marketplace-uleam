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
// FAKE REPOSITORY (BADGES)
// ======================================================

type badgeFake struct {
	data   []models.Badge
	nextID int
}

func nuevoBadgeFake() *badgeFake {
	return &badgeFake{data: []models.Badge{}}
}

// ---------------- BADGES CRUD ----------------

func (f *badgeFake) ListarBadges() []models.Badge {
	return f.data
}

func (f *badgeFake) BuscarBadgePorID(id int) (models.Badge, bool) {
	for _, b := range f.data {
		if b.ID == id {
			return b, true
		}
	}
	return models.Badge{}, false
}

func (f *badgeFake) CrearBadge(b models.Badge) models.Badge {
	f.nextID++
	b.ID = f.nextID
	f.data = append(f.data, b)
	return b
}

func (f *badgeFake) ActualizarBadge(id int, b models.Badge) (models.Badge, bool) {
	for i := range f.data {
		if f.data[i].ID == id {
			b.ID = id
			f.data[i] = b
			return b, true
		}
	}
	return models.Badge{}, false
}

func (f *badgeFake) BorrarBadge(id int) bool {
	for i := range f.data {
		if f.data[i].ID == id {
			f.data = append(f.data[:i], f.data[i+1:]...)
			return true
		}
	}
	return false
}

// ======================================================
// IMPLEMENTAR INTERFAZ COMPLETA (OBLIGATORIO)
// ======================================================

func (f *badgeFake) ListarUsers() []models.User { return nil }
func (f *badgeFake) BuscarUserPorID(id int) (models.User, bool) {
	return models.User{}, false
}
func (f *badgeFake) CrearUser(u models.User) models.User { return u }
func (f *badgeFake) ActualizarUser(id int, u models.User) (models.User, bool) {
	return u, false
}
func (f *badgeFake) BorrarUser(id int) bool { return false }

func (f *badgeFake) ListarReviews() []models.Review { return nil }
func (f *badgeFake) BuscarReviewPorID(id int) (models.Review, bool) {
	return models.Review{}, false
}
func (f *badgeFake) CrearReview(r models.Review) models.Review { return r }
func (f *badgeFake) ActualizarReview(id int, r models.Review) (models.Review, bool) {
	return r, false
}
func (f *badgeFake) BorrarReview(id int) bool { return false }

var _ storage.PerfilRepository = (*badgeFake)(nil)

// ======================================================
// JWT
// ======================================================

func generarJWTBadge() string {
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
// ENTORNO DE TEST
// ======================================================

func construirEntornoBadge() http.Handler {

	repo := nuevoBadgeFake()

	badgeSvc := service.NuevoBadgeService(repo)
	authSvc := service.NuevoAuthService(nil)

	server := handlers.NewServer(
		nil,      // UserService
		nil,      // ReviewService
		badgeSvc, // BadgeService
		nil,      // MessageService
		nil,      // MissionService
		nil,      // UserMissionService
		authSvc,  // AuthService
		nil,      // CategoriaService
		nil,      // ProductoService
		nil,      // OrdenService
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
// TEST 1: CREAR BADGE
// ======================================================

func TestCrearBadge_OK(t *testing.T) {

	h := construirEntornoBadge()
	token := generarJWTBadge()

	body := `{
		"name":"Experto",
		"description":"100 puntos",
		"required_rep":100
	}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/badges", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	require.Equal(t, http.StatusTeapot, rec.Code)

	var resp models.Badge
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&resp))

	assert.Equal(t, "Experto", resp.Name)
}

// ======================================================
// TEST 2: LISTAR BADGES
// ======================================================

func TestListarBadges_OK(t *testing.T) {

	h := construirEntornoBadge()
	token := generarJWTBadge()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/badges", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}

// ======================================================
// TEST 3: SIN TOKEN (401)
// ======================================================

func TestBadge_SinToken(t *testing.T) {

	h := construirEntornoBadge()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/badges", strings.NewReader(`{"name":"Experto"}`))

	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}
