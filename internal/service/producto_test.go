package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"marketplace-api/internal/models"
	"marketplace-api/internal/service"
)

// Usa el mismo ordenRepoMock definido en orden_test.go.

func TestProductoService_Crear(t *testing.T) {
	casos := []struct {
		nombre        string
		entrada       models.Producto
		errEsperado   error
		debePersistir bool
	}{
		{"nombre vacio rechazado", models.Producto{Nombre: "  ", Precio: 10, CategoriaID: 1}, service.ErrNombreVacio, false},
		{"precio negativo rechazado", models.Producto{Nombre: "Mouse", Precio: -5, CategoriaID: 1}, service.ErrPrecioNegativo, false},
		{"categoria invalida rechazada", models.Producto{Nombre: "Mouse", Precio: 10, CategoriaID: 0}, service.ErrCategoriaInvalida, false},
		{"producto valido se persiste", models.Producto{Nombre: "Mouse", Precio: 10, CategoriaID: 1}, nil, true},
	}
	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			repo := new(ordenRepoMock)
			if c.debePersistir {
				guardado := c.entrada
				guardado.ID = 1
				repo.On("CrearProducto", c.entrada).Return(guardado)
			}
			svc := service.NewProductoService(repo)

			creado, err := svc.Crear(c.entrada)

			if c.errEsperado != nil {
				require.ErrorIs(t, err, c.errEsperado)
				repo.AssertNotCalled(t, "CrearProducto")
			} else {
				require.NoError(t, err)
				assert.Equal(t, 1, creado.ID)
				repo.AssertCalled(t, "CrearProducto", c.entrada)
			}
		})
	}
}

func TestProductoService_Obtener(t *testing.T) {
	t.Run("existe", func(t *testing.T) {
		repo := new(ordenRepoMock)
		repo.On("BuscarProductoPorID", 1).Return(models.Producto{ID: 1, Nombre: "Mouse"}, true)
		p, ok := service.NewProductoService(repo).Obtener(1)
		assert.True(t, ok)
		assert.Equal(t, "Mouse", p.Nombre)
	})
	t.Run("no existe", func(t *testing.T) {
		repo := new(ordenRepoMock)
		repo.On("BuscarProductoPorID", 999).Return(models.Producto{}, false)
		_, ok := service.NewProductoService(repo).Obtener(999)
		assert.False(t, ok)
	})
}

func TestProductoService_Actualizar(t *testing.T) {
	datos := models.Producto{Nombre: "Teclado", Precio: 20, CategoriaID: 1}

	t.Run("valido", func(t *testing.T) {
		repo := new(ordenRepoMock)
		actualizado := datos
		actualizado.ID = 1
		repo.On("ActualizarProducto", 1, datos).Return(actualizado, true)
		p, err := service.NewProductoService(repo).Actualizar(1, datos)
		require.NoError(t, err)
		assert.Equal(t, 1, p.ID)
	})
	t.Run("no existe -> ErrNoEncontrado", func(t *testing.T) {
		repo := new(ordenRepoMock)
		repo.On("ActualizarProducto", 999, datos).Return(models.Producto{}, false)
		_, err := service.NewProductoService(repo).Actualizar(999, datos)
		require.ErrorIs(t, err, service.ErrNoEncontrado)
	})
	t.Run("invalido no toca el repo", func(t *testing.T) {
		repo := new(ordenRepoMock)
		_, err := service.NewProductoService(repo).Actualizar(1, models.Producto{Nombre: ""})
		require.ErrorIs(t, err, service.ErrNombreVacio)
		repo.AssertNotCalled(t, "ActualizarProducto")
	})
}

func TestProductoService_Borrar(t *testing.T) {
	t.Run("existe", func(t *testing.T) {
		repo := new(ordenRepoMock)
		repo.On("BorrarProducto", 1).Return(true)
		require.NoError(t, service.NewProductoService(repo).Borrar(1))
	})
	t.Run("no existe -> ErrNoEncontrado", func(t *testing.T) {
		repo := new(ordenRepoMock)
		repo.On("BorrarProducto", 999).Return(false)
		require.ErrorIs(t, service.NewProductoService(repo).Borrar(999), service.ErrNoEncontrado)
	})
}

func TestProductoService_Listar(t *testing.T) {
	repo := new(ordenRepoMock)
	repo.On("ListarProductos").Return([]models.Producto{{ID: 1}, {ID: 2}})
	lista := service.NewProductoService(repo).Listar()
	assert.Len(t, lista, 2)
	repo.AssertExpectations(t)
}
