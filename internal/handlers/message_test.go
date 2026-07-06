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

// =======================
// FAKE EN MEMORIA
// =======================

type messageFake struct {
	data []models.Message
	next int
}

func nuevoMessageFake() *messageFake {
	return &messageFake{data: []models.Message{}}
}

func (f *messageFake) CrearMessage(m models.Message) models.Message {
	f.next++
	m.ID = f.next
	f.data = append(f.data, m)
	return m
}

func (f *messageFake) ListarMessages() []models.Message {
	return f.data
}

func (f *messageFake) BuscarMessagePorID(id int) (models.Message, bool) {
	for _, m := range f.data {
		if m.ID == id {
			return m, true
		}
	}
	return models.Message{}, false
}

func (f *messageFake) ActualizarMessage(id int, m models.Message) (models.Message, bool) {
	for i, msg := range f.data {
		if msg.ID == id {
			m.ID = id
			f.data[i] = m
			return m, true
		}
	}
	return models.Message{}, false
}

func (f *messageFake) BorrarMessage(id int) bool {
	for i, msg := range f.data {
		if msg.ID == id {
			f.data = append(f.data[:i], f.data[i+1:]...)
			return true
		}
	}
	return false
}

// --- cumplir interfaz ---
func (f *messageFake) ListarMissions() []models.Mission { return nil }
func (f *messageFake) BuscarMisionPorID(id int) (models.Mission, bool) {
	return models.Mission{}, false
}
func (f *messageFake) CrearMision(m models.Mission) models.Mission { return m }
func (f *messageFake) ActualizarMision(id int, m models.Mission) (models.Mission, bool) {
	return m, false
}
func (f *messageFake) BorrarMision(id int) bool { return false }

func (f *messageFake) ListarUserMissions() []models.UserMission { return nil }
func (f *messageFake) BuscarUserMissionPorID(id int) (models.UserMission, bool) {
	return models.UserMission{}, false
}
func (f *messageFake) CrearUserMission(m models.UserMission) models.UserMission { return m }
func (f *messageFake) ActualizarUserMission(id int, m models.UserMission) (models.UserMission, bool) {
	return m, false
}
func (f *messageFake) BorrarUserMission(id int) bool { return false }

var _ storage.ComunidadRepository = (*messageFake)(nil)

// =======================
// JWT REAL (CLAVE DEL FIX)
// =======================

func generarJWTValido() string {
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

// =======================
// ENTORNO DE TEST
// =======================

func construirEntorno() http.Handler {

	repo := nuevoMessageFake()

	messageSvc := service.NuevoMessageService(repo)
	authSvc := service.NuevoAuthService(nil)

	server := handlers.NewServer(
		nil,        // UserService
		nil,        // ReviewService
		nil,        // BadgeService
		messageSvc, // MessageService
		nil,        // MissionService
		nil,        // UserMissionService
		authSvc,    // AuthService
		nil,        // CategoriaService
		nil,        // ProductoService
		nil,        // OrdenService

	)

	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router) {

		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(authSvc))

			r.Post("/messages", server.CrearMessage)
			r.Get("/messages", server.ListarMessages)
		})
	})

	return r
}

// =======================
// TEST 1: 201 CREATED
// =======================

func TestCrearMessage_OK(t *testing.T) {

	h := construirEntorno()
	token := generarJWTValido()

	body := `{"content":"hola mundo"}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/messages", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	require.Equal(t, http.StatusCreated, rec.Code)

	var resp models.Message
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&resp))

	assert.Equal(t, "hola mundo", resp.Content)
}

// =======================
// TEST 2: LISTAR 200 OK
// =======================

func TestListarMessages(t *testing.T) {

	h := construirEntorno()
	token := generarJWTValido()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/messages", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}

// =======================
// TEST 3: 401 SIN TOKEN
// =======================

func TestRutaProtegida_SinToken(t *testing.T) {

	h := construirEntorno()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/messages", strings.NewReader(`{"content":"hola"}`))
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}
