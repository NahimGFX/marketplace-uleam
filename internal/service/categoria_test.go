package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"marketplace-api/internal/models"
	"marketplace-api/internal/service"
)

// Usa el mismo ordenRepoMock definido en orden_test.go.

func TestCategoriaService_Crear(t *testing.T) {
	casos := []struct {
		nombre        string
		entrada       models.Categoria
		errEsperado   error
		debePersistir bool
	}{
		{"nombre vacio rechazado", models.Categoria{Name: "  "}, service.ErrNombreVacio, false},
		{"categoria valida se persiste", models.Categoria{Name: "Electronica"}, nil, true},
	}
	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			repo := new(ordenRepoMock)
			if c.debePersistir {
				guardado := c.entrada
				guardado.ID = 1
				repo.On("CrearCategoria", c.entrada).Return(guardado)
			}
			svc := service.NewCategoriaService(repo)

			creado, err := svc.Crear(c.entrada)

			if c.errEsperado != nil {
				require.ErrorIs(t, err, c.errEsperado)
				repo.AssertNotCalled(t, "CrearCategoria")
			} else {
				require.NoError(t, err)
				assert.Equal(t, 1, creado.ID)
				repo.AssertCalled(t, "CrearCategoria", c.entrada)
			}
		})
	}
}

func TestCategoriaService_Obtener(t *testing.T) {
	t.Run("existe", func(t *testing.T) {
		repo := new(ordenRepoMock)
		repo.On("BuscarCategoriaPorID", 1).Return(models.Categoria{ID: 1, Name: "Electronica"}, true)
		c, ok := service.NewCategoriaService(repo).Obtener(1)
		assert.True(t, ok)
		assert.Equal(t, "Electronica", c.Name)
	})
	t.Run("no existe", func(t *testing.T) {
		repo := new(ordenRepoMock)
		repo.On("BuscarCategoriaPorID", 999).Return(models.Categoria{}, false)
		_, ok := service.NewCategoriaService(repo).Obtener(999)
		assert.False(t, ok)
	})
}

func TestCategoriaService_Actualizar(t *testing.T) {
	datos := models.Categoria{Name: "Hogar"}

	t.Run("valido", func(t *testing.T) {
		repo := new(ordenRepoMock)
		actualizado := datos
		actualizado.ID = 1
		repo.On("ActualizarCategoria", 1, datos).Return(actualizado, true)
		c, err := service.NewCategoriaService(repo).Actualizar(1, datos)
		require.NoError(t, err)
		assert.Equal(t, 1, c.ID)
	})
	t.Run("no existe -> ErrNoEncontrado", func(t *testing.T) {
		repo := new(ordenRepoMock)
		repo.On("ActualizarCategoria", 999, datos).Return(models.Categoria{}, false)
		_, err := service.NewCategoriaService(repo).Actualizar(999, datos)
		require.ErrorIs(t, err, service.ErrNoEncontrado)
	})
	t.Run("invalido no toca el repo", func(t *testing.T) {
		repo := new(ordenRepoMock)
		_, err := service.NewCategoriaService(repo).Actualizar(1, models.Categoria{Name: ""})
		require.ErrorIs(t, err, service.ErrNombreVacio)
		repo.AssertNotCalled(t, "ActualizarCategoria")
	})
}

func TestCategoriaService_Borrar(t *testing.T) {
	t.Run("existe", func(t *testing.T) {
		repo := new(ordenRepoMock)
		repo.On("BorrarCategoria", 1).Return(true)
		require.NoError(t, service.NewCategoriaService(repo).Borrar(1))
	})
	t.Run("no existe -> ErrNoEncontrado", func(t *testing.T) {
		repo := new(ordenRepoMock)
		repo.On("BorrarCategoria", 999).Return(false)
		require.ErrorIs(t, service.NewCategoriaService(repo).Borrar(999), service.ErrNoEncontrado)
	})
}

func TestCategoriaService_Listar(t *testing.T) {
	repo := new(ordenRepoMock)
	repo.On("ListarCategorias").Return([]models.Categoria{{ID: 1}, {ID: 2}})
	lista := service.NewCategoriaService(repo).Listar()
	assert.Len(t, lista, 2)
	repo.AssertExpectations(t)
}
