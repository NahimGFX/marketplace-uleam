package service_test

import (
	"testing"

	"marketplace-api/internal/models"
	"marketplace-api/internal/service"
)

// Mock del repositorio
type mockOrdenRepo struct {
	ordenes []models.Orden
}

func (m *mockOrdenRepo) ListarOrdenes() []models.Orden { return m.ordenes }
func (m *mockOrdenRepo) BuscarOrdenPorID(id int) (models.Orden, bool) {
	return models.Orden{}, false
}
func (m *mockOrdenRepo) CrearOrden(o models.Orden) models.Orden { return o }
func (m *mockOrdenRepo) ActualizarOrden(id int, o models.Orden) (models.Orden, bool) {
	return o, true
}
func (m *mockOrdenRepo) BorrarOrden(id int) bool { return true }

// Métodos que exige OrdenRepository pero no usa OrdenService
func (m *mockOrdenRepo) ListarCategorias() []models.Categoria { return nil }
func (m *mockOrdenRepo) BuscarCategoriaPorID(id int) (models.Categoria, bool) {
	return models.Categoria{}, false
}
func (m *mockOrdenRepo) CrearCategoria(c models.Categoria) models.Categoria { return c }
func (m *mockOrdenRepo) ActualizarCategoria(id int, c models.Categoria) (models.Categoria, bool) {
	return c, true
}
func (m *mockOrdenRepo) BorrarCategoria(id int) bool        { return true }
func (m *mockOrdenRepo) ListarProductos() []models.Producto { return nil }
func (m *mockOrdenRepo) BuscarProductoPorID(id int) (models.Producto, bool) {
	return models.Producto{}, false
}
func (m *mockOrdenRepo) CrearProducto(p models.Producto) models.Producto { return p }
func (m *mockOrdenRepo) ActualizarProducto(id int, p models.Producto) (models.Producto, bool) {
	return p, true
}
func (m *mockOrdenRepo) BorrarProducto(id int) bool { return true }

// =====================================
// TESTS
// =====================================

// Prueba que una orden con estado vacío es rechazada y NO llega al repositorio
func TestCrearOrden_EstadoVacio_Rechazado(t *testing.T) {
	repo := &mockOrdenRepo{}
	svc := service.NewOrdenService(repo)

	_, err := svc.Crear(models.Orden{
		ProductoID:  1,
		IDComprador: 1,
		Estado:      "", // estado vacío — debe ser rechazado
	})

	if err == nil {
		t.Fatal("se esperaba error por estado vacío, pero no hubo error")
	}
	if err != service.ErrEstadoVacio {
		t.Fatalf("error esperado: %v, error obtenido: %v", service.ErrEstadoVacio, err)
	}
	if len(repo.ordenes) != 0 {
		t.Fatal("la orden no debió guardarse en el repositorio")
	}
}

// Prueba que una orden con producto_id inválido es rechazada
func TestCrearOrden_ProductoIDInvalido_Rechazado(t *testing.T) {
	repo := &mockOrdenRepo{}
	svc := service.NewOrdenService(repo)

	_, err := svc.Crear(models.Orden{
		ProductoID:  0, // inválido
		IDComprador: 1,
		Estado:      "pendiente",
	})

	if err == nil {
		t.Fatal("se esperaba error por producto_id inválido, pero no hubo error")
	}
	if err != service.ErrProductoInvalido {
		t.Fatalf("error esperado: %v, error obtenido: %v", service.ErrProductoInvalido, err)
	}
}
