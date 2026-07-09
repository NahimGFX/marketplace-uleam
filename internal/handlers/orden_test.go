package handlers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"

	"marketplace-api/internal/handlers"
	"marketplace-api/internal/middleware"
	"marketplace-api/internal/models"
	"marketplace-api/internal/service"
	"marketplace-api/internal/storage"
)

// =======================
// FAKE EN MEMORIA
// =======================

type fakeOrdenStore struct {
	ordenes []models.Orden
	next    int
}

func nuevoFakeOrdenStore() *fakeOrdenStore {
	return &fakeOrdenStore{
		ordenes: []models.Orden{},
	}
}

// =======================
// ORDENES
// =======================

func (f *fakeOrdenStore) ListarOrdenes() []models.Orden {
	return f.ordenes
}

func (f *fakeOrdenStore) BuscarOrdenPorID(id int) (models.Orden, bool) {
	for _, o := range f.ordenes {
		if o.ID == id {
			return o, true
		}
	}
	return models.Orden{}, false
}

func (f *fakeOrdenStore) CrearOrden(o models.Orden) models.Orden {
	f.next++
	o.ID = f.next
	f.ordenes = append(f.ordenes, o)
	return o
}

func (f *fakeOrdenStore) ActualizarOrden(id int, o models.Orden) (models.Orden, bool) {
	return o, false
}

func (f *fakeOrdenStore) BorrarOrden(id int) bool {
	return false
}

// =======================
// CUMPLIR INTERFAZ
// =======================

func (f *fakeOrdenStore) ListarCategorias() []models.Categoria { return nil }
func (f *fakeOrdenStore) BuscarCategoriaPorID(id int) (models.Categoria, bool) {
	return models.Categoria{}, false
}
func (f *fakeOrdenStore) CrearCategoria(c models.Categoria) models.Categoria { return c }
func (f *fakeOrdenStore) ActualizarCategoria(id int, c models.Categoria) (models.Categoria, bool) {
	return c, false
}
func (f *fakeOrdenStore) BorrarCategoria(id int) bool        { return false }
func (f *fakeOrdenStore) ListarProductos() []models.Producto { return nil }
func (f *fakeOrdenStore) BuscarProductoPorID(id int) (models.Producto, bool) {
	return models.Producto{}, false
}
func (f *fakeOrdenStore) CrearProducto(p models.Producto) models.Producto { return p }
func (f *fakeOrdenStore) ActualizarProducto(id int, p models.Producto) (models.Producto, bool) {
	return p, false
}
func (f *fakeOrdenStore) BorrarProducto(id int) bool { return false }

var _ storage.OrdenRepository = (*fakeOrdenStore)(nil)

// =======================
// JWT REAL
// =======================

func generarJWTValidoOrden() string {
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

func construirEntornoOrden() http.Handler {
	repo := nuevoFakeOrdenStore()
	ordenSvc := service.NewOrdenService(repo)
	authSvc := service.NuevoAuthService(nil)

	server := handlers.NewServer(
		nil,      // UserService
		nil,      // ReviewService
		nil,      // BadgeService
		nil,      // MessageService
		nil,      // MissionService
		nil,      // UserMissionService
		authSvc,  // AuthService
		nil,      // CategoriaService
		nil,      // ProductoService
		ordenSvc, // OrdenService
	)

	r := chi.NewRouter()
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/ordenes", server.ListarOrdenes)
		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(authSvc))
			r.Post("/ordenes", server.CrearOrden)
		})
	})

	return r
}

// =======================
// TEST 1: 201 CREATED
// =======================

func TestCrearOrden_OK(t *testing.T) {
	h := construirEntornoOrden()
	token := generarJWTValidoOrden()

	body := `{
		"producto_id":1,
		"comprador_id":1,
		"estado":"pendiente"
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/ordenes",
		bytes.NewBufferString(body),
	)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
}

// =======================
// TEST 2: LISTAR 200 OK
// =======================

func TestListarOrdenes(t *testing.T) {
	h := construirEntornoOrden()

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/v1/ordenes",
		nil,
	)

	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}

// =======================
// TEST 3: 401 SIN TOKEN
// =======================

func TestCrearOrden_SinToken(t *testing.T) {
	h := construirEntornoOrden()

	body := `{
		"producto_id":1,
		"comprador_id":1,
		"estado":"pendiente"
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/ordenes",
		bytes.NewBufferString(body),
	)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}
