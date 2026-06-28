package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"marketplace-api/internal/handlers"
	"marketplace-api/internal/middleware"
	"marketplace-api/internal/models"

	"github.com/go-chi/chi/v5"
)

// Fake en memoria para los handlers
type fakeOrdenStore struct {
	ordenes []models.Orden
}

func (f *fakeOrdenStore) ListarUsers() []models.User                               { return nil }
func (f *fakeOrdenStore) BuscarUserPorID(id int) (models.User, bool)               { return models.User{}, false }
func (f *fakeOrdenStore) CrearUser(u models.User) models.User                      { return u }
func (f *fakeOrdenStore) ActualizarUser(id int, u models.User) (models.User, bool) { return u, true }
func (f *fakeOrdenStore) BorrarUser(id int) bool                                   { return true }
func (f *fakeOrdenStore) ListarReviews() []models.Review                           { return nil }
func (f *fakeOrdenStore) BuscarReviewPorID(id int) (models.Review, bool) {
	return models.Review{}, false
}
func (f *fakeOrdenStore) CrearReview(r models.Review) models.Review { return r }
func (f *fakeOrdenStore) ActualizarReview(id int, r models.Review) (models.Review, bool) {
	return r, true
}
func (f *fakeOrdenStore) BorrarReview(id int) bool                                    { return true }
func (f *fakeOrdenStore) ListarBadges() []models.Badge                                { return nil }
func (f *fakeOrdenStore) BuscarBadgePorID(id int) (models.Badge, bool)                { return models.Badge{}, false }
func (f *fakeOrdenStore) CrearBadge(b models.Badge) models.Badge                      { return b }
func (f *fakeOrdenStore) ActualizarBadge(id int, b models.Badge) (models.Badge, bool) { return b, true }
func (f *fakeOrdenStore) BorrarBadge(id int) bool                                     { return true }
func (f *fakeOrdenStore) ListarCategorias() []models.Categoria                        { return nil }
func (f *fakeOrdenStore) BuscarCategoriaPorID(id int) (models.Categoria, bool) {
	return models.Categoria{}, false
}
func (f *fakeOrdenStore) CrearCategoria(c models.Categoria) models.Categoria { return c }
func (f *fakeOrdenStore) ActualizarCategoria(id int, c models.Categoria) (models.Categoria, bool) {
	return c, true
}
func (f *fakeOrdenStore) BorrarCategoria(id int) bool        { return true }
func (f *fakeOrdenStore) ListarProductos() []models.Producto { return nil }
func (f *fakeOrdenStore) BuscarProductoPorID(id int) (models.Producto, bool) {
	return models.Producto{}, false
}
func (f *fakeOrdenStore) CrearProducto(p models.Producto) models.Producto { return p }
func (f *fakeOrdenStore) ActualizarProducto(id int, p models.Producto) (models.Producto, bool) {
	return p, true
}
func (f *fakeOrdenStore) BorrarProducto(id int) bool                   { return true }
func (f *fakeOrdenStore) ListarOrdenes() []models.Orden                { return f.ordenes }
func (f *fakeOrdenStore) BuscarOrdenPorID(id int) (models.Orden, bool) { return models.Orden{}, false }
func (f *fakeOrdenStore) CrearOrden(o models.Orden) models.Orden {
	f.ordenes = append(f.ordenes, o)
	return o
}
func (f *fakeOrdenStore) ActualizarOrden(id int, o models.Orden) (models.Orden, bool) { return o, true }
func (f *fakeOrdenStore) BorrarOrden(id int) bool                                     { return true }
func (f *fakeOrdenStore) ListarMessages() []models.Message                            { return nil }
func (f *fakeOrdenStore) BuscarMessagePorID(id int) (models.Message, bool) {
	return models.Message{}, false
}
func (f *fakeOrdenStore) CrearMessage(m models.Message) models.Message { return m }
func (f *fakeOrdenStore) ActualizarMessage(id int, m models.Message) (models.Message, bool) {
	return m, true
}
func (f *fakeOrdenStore) BorrarMessage(id int) bool        { return true }
func (f *fakeOrdenStore) ListarMissions() []models.Mission { return nil }
func (f *fakeOrdenStore) BuscarMisionPorID(id int) (models.Mission, bool) {
	return models.Mission{}, false
}
func (f *fakeOrdenStore) CrearMision(m models.Mission) models.Mission { return m }
func (f *fakeOrdenStore) ActualizarMision(id int, m models.Mission) (models.Mission, bool) {
	return m, true
}
func (f *fakeOrdenStore) BorrarMision(id int) bool                 { return true }
func (f *fakeOrdenStore) ListarUsermissions() []models.UserMission { return nil }
func (f *fakeOrdenStore) BuscarUserMissionPorID(id int) (models.UserMission, bool) {
	return models.UserMission{}, false
}
func (f *fakeOrdenStore) CrearUserMission(m models.UserMission) models.UserMission { return m }
func (f *fakeOrdenStore) ActualizarUserMission(id int, m models.UserMission) (models.UserMission, bool) {
	return m, true
}
func (f *fakeOrdenStore) BorrarUserMission(id int) bool { return true }

func setupRouter(s *handlers.Server) http.Handler {
	r := chi.NewRouter()
	r.Get("/api/v1/ordenes", s.ListarOrdenes)
	r.Group(func(r chi.Router) {
		r.Use(middleware.RequiereToken)
		r.Post("/api/v1/ordenes", s.CrearOrden)
	})
	return r
}

// Test 1: POST /ordenes sin token debe responder 401
func TestCrearOrden_SinToken_Retorna401(t *testing.T) {
	fake := &fakeOrdenStore{}
	srv := handlers.NewServer(fake)
	r := setupRouter(srv)

	body, _ := json.Marshal(models.Orden{
		ProductoID:  1,
		IDComprador: 1,
		Estado:      "pendiente",
	})

	req := httptest.NewRequest(http.MethodPost, "/api/v1/ordenes", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	// Sin header Authorization

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Fatalf("esperado 401, obtenido %d", w.Code)
	}
}

// Test 2: GET /ordenes debe responder 200 con lista vacía
func TestListarOrdenes_Retorna200(t *testing.T) {
	fake := &fakeOrdenStore{}
	srv := handlers.NewServer(fake)
	r := setupRouter(srv)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/ordenes", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("esperado 200, obtenido %d", w.Code)
	}
}
