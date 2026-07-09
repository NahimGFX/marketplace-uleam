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
// FAKE MISSION REPOSITORY
// =======================

type missionFake struct {
	data []models.Mission
	next int
}

func nuevoMissionFake() *missionFake {

	return &missionFake{
		data: []models.Mission{},
	}
}

func (f *missionFake) CrearMision(m models.Mission) models.Mission {

	f.next++

	m.ID = f.next

	f.data = append(
		f.data,
		m,
	)

	return m
}

func (f *missionFake) ListarMissions() []models.Mission {

	return f.data
}

func (f *missionFake) BuscarMisionPorID(id int) (models.Mission, bool) {

	for _, m := range f.data {

		if m.ID == id {
			return m, true
		}
	}

	return models.Mission{}, false
}

func (f *missionFake) ActualizarMision(id int, m models.Mission) (models.Mission, bool) {

	for i, mission := range f.data {

		if mission.ID == id {

			m.ID = id

			f.data[i] = m

			return m, true
		}
	}

	return models.Mission{}, false
}

func (f *missionFake) BorrarMision(id int) bool {

	for i, m := range f.data {

		if m.ID == id {

			f.data = append(
				f.data[:i],
				f.data[i+1:]...,
			)

			return true
		}
	}

	return false
}

// =======================
// USER MESSAGE METHODS
// PARA CUMPLIR INTERFAZ
// =======================

func (f *missionFake) ListarMessages() []models.Message {
	return nil
}

func (f *missionFake) BuscarMessagePorID(id int) (models.Message, bool) {
	return models.Message{}, false
}

func (f *missionFake) CrearMessage(m models.Message) models.Message {
	return m
}

func (f *missionFake) ActualizarMessage(id int, m models.Message) (models.Message, bool) {
	return m, false
}

func (f *missionFake) BorrarMessage(id int) bool {
	return false
}

func (f *missionFake) ListarUserMissions() []models.UserMission {
	return nil
}

func (f *missionFake) BuscarUserMissionPorID(id int) (models.UserMission, bool) {
	return models.UserMission{}, false
}

func (f *missionFake) CrearUserMission(m models.UserMission) models.UserMission {
	return m
}

func (f *missionFake) ActualizarUserMission(id int, m models.UserMission) (models.UserMission, bool) {
	return m, false
}

func (f *missionFake) BorrarUserMission(id int) bool {
	return false
}

var _ storage.ComunidadRepository = (*missionFake)(nil)

// =======================
// JWT
// =======================

func generarJWTTMission() string {

	secret := []byte("marketplace-uleam-secreto-demo-cambiar-en-S12")

	claims := service.Claims{

		UsuarioID: 1,

		RegisteredClaims: jwt.RegisteredClaims{

			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(time.Hour),
			),

			IssuedAt: jwt.NewNumericDate(
				time.Now(),
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

// =======================
// ENTORNO
// =======================

func construirEntornoMission() http.Handler {

	repo := nuevoMissionFake()

	missionService := service.NuevoMissionService(repo)

	authService := service.NuevoAuthService(nil)

	server := handlers.NewServer(

		nil, // UserService
		nil, // ReviewService
		nil, // BadgeService
		nil, // MessageService

		missionService,

		nil, // UserMissionService

		authService,

		nil, // CategoriaService
		nil, // ProductoService
		nil, // OrdenService
	)

	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router) {

		r.Group(func(r chi.Router) {

			r.Use(
				middleware.Auth(authService),
			)

			// MESSAGES NO SON NECESARIOS AQUÍ

			// MISSIONS CRUD COMPLETO
			r.Get(
				"/missions",
				server.ListarMissions,
			)

			r.Get(
				"/missions/{id}",
				server.ObtenerMision,
			)

			r.Post(
				"/missions",
				server.CrearMision,
			)

			r.Put(
				"/missions/{id}",
				server.ActualizarMision,
			)

			r.Delete(
				"/missions/{id}",
				server.BorrarMision,
			)

		})

	})

	return r
}

// =======================
// TEST 1
// CREAR MISION
// =======================

func TestCrearMission_OK(t *testing.T) {

	h := construirEntornoMission()

	body := `

{
"title":"Publicar producto",
"description":"Sube tu primer producto",
"required_level":1,
"reward_points":20
}

`

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/missions",
		strings.NewReader(body),
	)

	req.Header.Set(
		"Content-Type",
		"application/json",
	)

	req.Header.Set(
		"Authorization",
		"Bearer "+generarJWTTMission(),
	)

	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	require.Equal(
		t,
		http.StatusCreated,
		rec.Code,
	)

	var mission models.Mission

	require.NoError(
		t,
		json.NewDecoder(rec.Body).Decode(&mission),
	)

	assert.Equal(
		t,
		"Publicar producto",
		mission.Title,
	)

}

// =======================
// TEST 2
// LISTAR MISIONES
// =======================

func TestListarMission_OK(t *testing.T) {

	h := construirEntornoMission()

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/v1/missions",
		nil,
	)

	req.Header.Set(
		"Authorization",
		"Bearer "+generarJWTTMission(),
	)

	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(
		t,
		http.StatusOK,
		rec.Code,
	)

}

// =======================
// TEST 3
// SIN TOKEN
// =======================

func TestMission_SinToken(t *testing.T) {

	h := construirEntornoMission()

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/v1/missions",
		nil,
	)

	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(
		t,
		http.StatusUnauthorized,
		rec.Code,
	)

}

// =======================
// TEST 4
// JSON INVALIDO
// =======================

func TestCrearMission_JSONInvalido(t *testing.T) {

	h := construirEntornoMission()

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/missions",
		strings.NewReader(`{"title":`),
	)

	req.Header.Set(
		"Authorization",
		"Bearer "+generarJWTTMission(),
	)

	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(
		t,
		http.StatusBadRequest,
		rec.Code,
	)

}

// =======================
// TEST 5
// LISTAR CON DATOS
// =======================

func TestListarMission_ConDatos(t *testing.T) {

	h := construirEntornoMission()

	token := generarJWTTMission()

	reqCrear := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/missions",
		strings.NewReader(`
		{
		"title":"Mision prueba",
		"description":"descripcion",
		"required_level":1,
		"reward_points":10
		}`),
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

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/v1/missions",
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

	var missions []models.Mission

	require.NoError(
		t,
		json.NewDecoder(rec.Body).Decode(&missions),
	)

	assert.Len(
		t,
		missions,
		1,
	)

}

// =======================
// TEST 6
// OBTENER MISION POR ID
// =======================

func TestObtenerMission_OK(t *testing.T) {

	h := construirEntornoMission()

	token := generarJWTTMission()

	// Crear primero
	reqCrear := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/missions",
		strings.NewReader(`{
			"title":"Mision buscar",
			"description":"descripcion",
			"required_level":1,
			"reward_points":20
		}`),
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

	h.ServeHTTP(recCrear, reqCrear)

	// Buscar ID 1
	req := httptest.NewRequest(
		http.MethodGet,
		"/api/v1/missions/1",
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

	var mission models.Mission

	require.NoError(
		t,
		json.NewDecoder(rec.Body).Decode(&mission),
	)

	assert.Equal(
		t,
		"Mision buscar",
		mission.Title,
	)

}

// =======================
// TEST 7
// ACTUALIZAR MISION
// =======================

func TestActualizarMission_OK(t *testing.T) {

	h := construirEntornoMission()

	token := generarJWTTMission()

	// Crear
	reqCrear := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/missions",
		strings.NewReader(`{
			"title":"Original",
			"description":"desc",
			"required_level":1,
			"reward_points":10
		}`),
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

	// Actualizar
	req := httptest.NewRequest(
		http.MethodPut,
		"/api/v1/missions/1",
		strings.NewReader(`{
			"title":"Actualizada",
			"description":"nuevo",
			"required_level":2,
			"reward_points":50
		}`),
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

// =======================
// TEST 8
// BORRAR MISION
// =======================

func TestBorrarMission_OK(t *testing.T) {

	h := construirEntornoMission()

	token := generarJWTTMission()

	// Crear misión
	reqCrear := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/missions",
		strings.NewReader(`{
			"title":"Eliminar",
			"description":"desc",
			"required_level":1,
			"reward_points":5
		}`),
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

	// DELETE
	req := httptest.NewRequest(
		http.MethodDelete,
		"/api/v1/missions/1",
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
