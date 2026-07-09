package service

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"marketplace-api/internal/models"
)

func TestCategoriaService_NombreVacio_NoLlegaRepositorio(t *testing.T) {
	repo := new(ordenRepoMock)
	svc := NewCategoriaService(repo)

	_, err := svc.Crear(models.Categoria{Name: ""})

	require.ErrorIs(t, err, ErrNombreVacio)
	repo.AssertNotCalled(t, "CrearCategoria", mock.Anything)
	repo.AssertExpectations(t)
}

func TestCategoriaService_CrearValida_LlegaRepositorio(t *testing.T) {
	repo := new(ordenRepoMock)
	svc := NewCategoriaService(repo)
	categoria := models.Categoria{Name: "Libros"}
	guardada := categoria
	guardada.ID = 1

	repo.On("CrearCategoria", categoria).Return(guardada).Once()

	resultado, err := svc.Crear(categoria)

	require.NoError(t, err)
	require.Equal(t, guardada, resultado)
	repo.AssertExpectations(t)
}

func TestCategoriaService_ActualizarNoEncontrado_RetornaError(t *testing.T) {
	repo := new(ordenRepoMock)
	svc := NewCategoriaService(repo)
	categoria := models.Categoria{Name: "Tecnologia"}

	repo.On("ActualizarCategoria", 99, categoria).Return(models.Categoria{}, false).Once()

	_, err := svc.Actualizar(99, categoria)

	require.ErrorIs(t, err, ErrNoEncontrado)
	repo.AssertExpectations(t)
}

func TestCategoriaService_BorrarNoEncontrado_RetornaError(t *testing.T) {
	repo := new(ordenRepoMock)
	svc := NewCategoriaService(repo)

	repo.On("BorrarCategoria", 99).Return(false).Once()

	err := svc.Borrar(99)

	require.ErrorIs(t, err, ErrNoEncontrado)
	repo.AssertExpectations(t)
}
