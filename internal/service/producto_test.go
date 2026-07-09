package service

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"marketplace-api/internal/models"
)

func TestProductoService_NombreVacio_NoLlegaRepositorio(t *testing.T) {
	repo := new(ordenRepoMock)
	svc := NewPorductoService(repo)

	_, err := svc.Crear(models.Producto{
		Nombre:      "",
		Precio:      10,
		CategoriaID: 1,
	})

	require.ErrorIs(t, err, ErrNombreVacio)
	repo.AssertNotCalled(t, "CrearProducto", mock.Anything)
	repo.AssertExpectations(t)
}

func TestProductoService_PrecioNegativo_NoLlegaRepositorio(t *testing.T) {
	repo := new(ordenRepoMock)
	svc := NewPorductoService(repo)

	_, err := svc.Crear(models.Producto{
		Nombre:      "Calculadora",
		Precio:      -1,
		CategoriaID: 1,
	})

	require.ErrorIs(t, err, ErrPrecioNegativo)
	repo.AssertNotCalled(t, "CrearProducto", mock.Anything)
	repo.AssertExpectations(t)
}

func TestProductoService_CategoriaInvalida_NoLlegaRepositorio(t *testing.T) {
	repo := new(ordenRepoMock)
	svc := NewPorductoService(repo)

	_, err := svc.Crear(models.Producto{
		Nombre:      "Calculadora",
		Precio:      10,
		CategoriaID: 0,
	})

	require.ErrorIs(t, err, ErrCategoriaInvalida)
	repo.AssertNotCalled(t, "CrearProducto", mock.Anything)
	repo.AssertExpectations(t)
}

func TestProductoService_CrearValido_LlegaRepositorio(t *testing.T) {
	repo := new(ordenRepoMock)
	svc := NewPorductoService(repo)
	producto := models.Producto{
		Nombre:      "Calculadora",
		Descripcion: "Calculadora cientifica",
		Precio:      15,
		CategoriaID: 1,
	}
	guardado := producto
	guardado.ID = 2

	repo.On("CrearProducto", producto).Return(guardado).Once()

	resultado, err := svc.Crear(producto)

	require.NoError(t, err)
	require.Equal(t, guardado, resultado)
	repo.AssertExpectations(t)
}

func TestProductoService_ActualizarNoEncontrado_RetornaError(t *testing.T) {
	repo := new(ordenRepoMock)
	svc := NewPorductoService(repo)
	producto := models.Producto{
		Nombre:      "Calculadora",
		Precio:      15,
		CategoriaID: 1,
	}

	repo.On("ActualizarProducto", 99, producto).Return(models.Producto{}, false).Once()

	_, err := svc.Actualizar(99, producto)

	require.ErrorIs(t, err, ErrNoEncontrado)
	repo.AssertExpectations(t)
}

func TestProductoService_BorrarNoEncontrado_RetornaError(t *testing.T) {
	repo := new(ordenRepoMock)
	svc := NewPorductoService(repo)

	repo.On("BorrarProducto", 99).Return(false).Once()

	err := svc.Borrar(99)

	require.ErrorIs(t, err, ErrNoEncontrado)
	repo.AssertExpectations(t)
}
