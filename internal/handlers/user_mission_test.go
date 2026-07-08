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

// =========================
// FAKE USERMISSION
// =========================

type userMissionFake struct {
	data []models.UserMission
	next int
}

func nuevoUserMissionFake() *userMissionFake {
	return &userMissionFake{
		data: []models.UserMission{},
	}
}

func (f *userMissionFake) CrearUserMission(
	m models.UserMission,
) models.UserMission {

	f.next++
	m.ID = f.next

	f.data = append(
		f.data,
		m,
	)

	return m
}

func (f *userMissionFake) ListarUserMissions() []models.UserMission {
	return f.data
}

func (f *userMissionFake) BuscarUserMissionPorID(
	id int,
) (models.UserMission, bool) {

	for _, m := range f.data {

		if m.ID == id {
			return m, true
		}
	}

	return models.UserMission{}, false
}

func (f *userMissionFake) ActualizarUserMission(
	id int,
	m models.UserMission,
) (models.UserMission, bool) {

	for i, item := range f.data {

		if item.ID == id {

			m.ID = id
			f.data[i] = m

			return m, true
		}
	}

	return models.UserMission{}, false
}

func (f *userMissionFake) BorrarUserMission(id int) bool {

	for i, item := range f.data {

		if item.ID == id {

			f.data = append(
				f.data[:i],
				f.data[i+1:]...,
			)

			return true
		}
	}

	return false
}

// =========================
// METODOS PARA INTERFAZ
// =========================

func (f *userMissionFake) ListarMessages() []models.Message {
	return nil
}

func (f *userMissionFake) BuscarMessagePorID(id int) (models.Message, bool) {
	return models.Message{}, false
}

func (f *userMissionFake) CrearMessage(m models.Message) models.Message {
	return m
}

func (f *userMissionFake) ActualizarMessage(id int, m models.Message) (models.Message, bool) {
	return m, false
}

func (f *userMissionFake) BorrarMessage(id int) bool {
	return false
}

func (f *userMissionFake) ListarMissions() []models.Mission {
	return nil
}

func (f *userMissionFake) BuscarMisionPorID(id int) (models.Mission, bool) {
	return models.Mission{}, false
}

func (f *userMissionFake) CrearMision(m models.Mission) models.Mission {
	return m
}

func (f *userMissionFake) ActualizarMision(id int, m models.Mission) (models.Mission, bool) {
	return m, false
}

func (f *userMissionFake) BorrarMision(id int) bool {
	return false
}

var _ storage.ComunidadRepository = (*userMissionFake)(nil)

// =========================
// JWT
// =========================

func tokenUserMission() string {

	secret := []byte(
		"marketplace-uleam-secreto-demo-cambiar-en-S12",
	)

	claims := service.Claims{

		UsuarioID: 1,

		RegisteredClaims: jwt.RegisteredClaims{

			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(time.Hour),
			),
		},
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	s, _ := token.SignedString(secret)

	return s
}

// =========================
// ENTORNO
// =========================

func entornoUserMission() http.Handler {

	repo := nuevoUserMissionFake()

	serviceUM := service.NuevoUserMissionService(repo)

	auth := service.NuevoAuthService(nil)

	server := handlers.NewServer(

		nil,
		nil,
		nil,
		nil,
		nil,
		serviceUM,
		auth,
		nil,
		nil,
		nil,
	)

	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router) {

		r.Use(middleware.Auth(auth))

		r.Post(
			"/usermissions",
			server.CrearUserMission,
		)

		r.Get(
			"/usermissions",
			server.ListarUsermissions,
		)

		r.Get(
			"/usermissions/{id}",
			server.ObtenerUserMission,
		)

		r.Put(
			"/usermissions/{id}",
			server.ActualizarUserMission,
		)

		r.Delete(
			"/usermissions/{id}",
			server.BorrarUserMission,
		)

	})

	return r
}

// =========================
// TEST CREAR
// =========================

func TestCrearUserMission_OK(t *testing.T) {

	h := entornoUserMission()

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/usermissions",
		strings.NewReader(
			`{
			"user_id":1,
			"mission_id":1,
			"completed":false
			}`,
		),
	)

	req.Header.Set(
		"Authorization",
		"Bearer "+tokenUserMission(),
	)

	req.Header.Set(
		"Content-Type",
		"application/json",
	)

	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(
		t,
		http.StatusCreated,
		rec.Code,
	)

}

// =========================
// LISTAR
// =========================

func TestListarUserMission_OK(t *testing.T) {

	h := entornoUserMission()

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/v1/usermissions",
		nil,
	)

	req.Header.Set(
		"Authorization",
		"Bearer "+tokenUserMission(),
	)

	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(
		t,
		http.StatusOK,
		rec.Code,
	)

}

// =========================
// OBTENER
// =========================

func TestObtenerUserMission_OK(t *testing.T) {

	h := entornoUserMission()

	token := tokenUserMission()

	crear := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/usermissions",
		strings.NewReader(
			`{
			"user_id":1,
			"mission_id":2,
			"completed":true
			}`,
		),
	)

	crear.Header.Set(
		"Authorization",
		"Bearer "+token,
	)

	crear.Header.Set(
		"Content-Type",
		"application/json",
	)

	recCrear := httptest.NewRecorder()

	h.ServeHTTP(recCrear, crear)

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/v1/usermissions/1",
		nil,
	)

	req.Header.Set(
		"Authorization",
		"Bearer "+token,
	)

	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(
		t,
		http.StatusOK,
		rec.Code,
	)

	var um models.UserMission

	json.NewDecoder(rec.Body).Decode(&um)

	assert.Equal(
		t,
		1,
		um.ID,
	)

}

// =========================
// ACTUALIZAR
// =========================

func TestActualizarUserMission_OK(t *testing.T) {

	h := entornoUserMission()

	token := tokenUserMission()

	// primero crear
	reqCrear := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/usermissions",
		strings.NewReader(
			`{
			"user_id":1,
			"mission_id":1,
			"completed":false
			}`,
		),
	)

	reqCrear.Header.Set(
		"Authorization",
		"Bearer "+token,
	)

	reqCrear.Header.Set(
		"Content-Type",
		"application/json",
	)

	recCrear := httptest.NewRecorder()

	h.ServeHTTP(
		recCrear,
		reqCrear,
	)

	require.Equal(
		t,
		http.StatusCreated,
		recCrear.Code,
	)

	// actualizar ID 1
	req := httptest.NewRequest(
		http.MethodPut,
		"/api/v1/usermissions/1",
		strings.NewReader(
			`{
			"user_id":1,
			"mission_id":2,
			"completed":true
			}`,
		),
	)

	req.Header.Set(
		"Authorization",
		"Bearer "+token,
	)

	req.Header.Set(
		"Content-Type",
		"application/json",
	)

	rec := httptest.NewRecorder()

	h.ServeHTTP(
		rec,
		req,
	)

	assert.Equal(
		t,
		http.StatusOK,
		rec.Code,
	)

}

// =========================
// BORRAR
// =========================

func TestBorrarUserMission_OK(t *testing.T) {

	h := entornoUserMission()

	token := tokenUserMission()

	// crear antes de borrar
	reqCrear := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/usermissions",
		strings.NewReader(
			`{
			"user_id":1,
			"mission_id":1,
			"completed":false
			}`,
		),
	)

	reqCrear.Header.Set(
		"Authorization",
		"Bearer "+token,
	)

	reqCrear.Header.Set(
		"Content-Type",
		"application/json",
	)

	recCrear := httptest.NewRecorder()

	h.ServeHTTP(
		recCrear,
		reqCrear,
	)

	require.Equal(
		t,
		http.StatusCreated,
		recCrear.Code,
	)

	// borrar
	req := httptest.NewRequest(
		http.MethodDelete,
		"/api/v1/usermissions/1",
		nil,
	)

	req.Header.Set(
		"Authorization",
		"Bearer "+token,
	)

	rec := httptest.NewRecorder()

	h.ServeHTTP(
		rec,
		req,
	)

	assert.Equal(
		t,
		http.StatusNoContent,
		rec.Code,
	)

}
