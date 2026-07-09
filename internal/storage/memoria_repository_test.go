package storage

import (
	"testing"

	"github.com/stretchr/testify/require"

	"marketplace-api/internal/models"
)

func TestMemoria_PerfilRepository_CRUD(t *testing.T) {
	repo := NuevaMemoria()
	repo.SeedUsers()
	repo.SeedReviews()
	repo.SeedBadges()

	user := repo.CrearUser(models.User{Name: "Ana"})
	require.NotZero(t, user.ID)
	require.NotEmpty(t, repo.ListarUsers())
	_, ok := repo.BuscarUserPorID(user.ID)
	require.True(t, ok)
	user, ok = repo.ActualizarUser(user.ID, models.User{Name: "Ana Editada"})
	require.True(t, ok)
	require.Equal(t, "Ana Editada", user.Name)
	require.True(t, repo.BorrarUser(user.ID))
	require.False(t, repo.BorrarUser(999))

	review := repo.CrearReview(models.Review{ReviewerID: 1, ReviewedID: 2, Comment: "Bien"})
	require.NotZero(t, review.ID)
	require.NotEmpty(t, repo.ListarReviews())
	_, ok = repo.BuscarReviewPorID(review.ID)
	require.True(t, ok)
	review, ok = repo.ActualizarReview(review.ID, models.Review{Comment: "Mejor"})
	require.True(t, ok)
	require.Equal(t, "Mejor", review.Comment)
	require.True(t, repo.BorrarReview(review.ID))
	require.False(t, repo.BorrarReview(999))

	badge := repo.CrearBadge(models.Badge{Name: "Novato"})
	require.NotZero(t, badge.ID)
	require.NotEmpty(t, repo.ListarBadges())
	_, ok = repo.BuscarBadgePorID(badge.ID)
	require.True(t, ok)
	badge, ok = repo.ActualizarBadge(badge.ID, models.Badge{Name: "Experto"})
	require.True(t, ok)
	require.Equal(t, "Experto", badge.Name)
	require.True(t, repo.BorrarBadge(badge.ID))
	require.False(t, repo.BorrarBadge(999))
}

func TestMemoria_OrdenRepository_CRUD(t *testing.T) {
	repo := NuevaMemoria()
	repo.SeedCategorias()
	repo.SeedProductos()
	repo.SeedOrdenes()

	categoria := repo.CrearCategoria(models.Categoria{Name: "Nueva"})
	require.NotZero(t, categoria.ID)
	require.NotEmpty(t, repo.ListarCategorias())
	_, ok := repo.BuscarCategoriaPorID(categoria.ID)
	require.True(t, ok)
	categoria, ok = repo.ActualizarCategoria(categoria.ID, models.Categoria{Name: "Editada"})
	require.True(t, ok)
	require.Equal(t, "Editada", categoria.Name)
	require.True(t, repo.BorrarCategoria(categoria.ID))
	require.False(t, repo.BorrarCategoria(999))

	producto := repo.CrearProducto(models.Producto{Nombre: "Calculadora", Precio: 10, CategoriaID: 1})
	require.NotZero(t, producto.ID)
	require.NotEmpty(t, repo.ListarProductos())
	_, ok = repo.BuscarProductoPorID(producto.ID)
	require.True(t, ok)
	producto, ok = repo.ActualizarProducto(producto.ID, models.Producto{Nombre: "Calculadora Pro", Precio: 20})
	require.True(t, ok)
	require.Equal(t, "Calculadora Pro", producto.Nombre)
	require.True(t, repo.BorrarProducto(producto.ID))
	require.False(t, repo.BorrarProducto(999))

	orden := repo.CrearOrden(models.Orden{ProductoID: 1, IDComprador: 2, Estado: "pendiente"})
	require.NotZero(t, orden.ID)
	require.NotEmpty(t, repo.ListarOrdenes())
	_, ok = repo.BuscarOrdenPorID(orden.ID)
	require.True(t, ok)
	orden, ok = repo.ActualizarOrden(orden.ID, models.Orden{ProductoID: 1, IDComprador: 2, Estado: "pagado"})
	require.True(t, ok)
	require.Equal(t, "pagado", orden.Estado)
	require.True(t, repo.BorrarOrden(orden.ID))
	require.False(t, repo.BorrarOrden(999))
}

func TestMemoria_ComunidadRepository_CRUD(t *testing.T) {
	repo := NuevaMemoria()
	repo.SeedMessages()
	repo.SeedMissions()
	repo.SeedUsermissions()

	message := repo.CrearMessage(models.Message{SenderID: 1, ReceiverID: 2, Content: "Hola"})
	require.NotZero(t, message.ID)
	require.NotEmpty(t, repo.ListarMessages())
	_, ok := repo.BuscarMessagePorID(message.ID)
	require.True(t, ok)
	message, ok = repo.ActualizarMessage(message.ID, models.Message{Content: "Editado"})
	require.True(t, ok)
	require.Equal(t, "Editado", message.Content)
	require.True(t, repo.BorrarMessage(message.ID))
	require.False(t, repo.BorrarMessage(999))

	mission := repo.CrearMision(models.Mission{Title: "Primera venta", Description: "Vende algo"})
	require.NotZero(t, mission.ID)
	require.NotEmpty(t, repo.ListarMissions())
	_, ok = repo.BuscarMisionPorID(mission.ID)
	require.True(t, ok)
	mission, ok = repo.ActualizarMision(mission.ID, models.Mission{Title: "Editada"})
	require.True(t, ok)
	require.Equal(t, "Editada", mission.Title)
	require.True(t, repo.BorrarMision(mission.ID))
	require.False(t, repo.BorrarMision(999))

	userMission := repo.CrearUserMission(models.UserMission{UserID: 1, MissionID: 2})
	require.NotZero(t, userMission.ID)
	require.NotEmpty(t, repo.ListarUserMissions())
	_, ok = repo.BuscarUserMissionPorID(userMission.ID)
	require.True(t, ok)
	userMission, ok = repo.ActualizarUserMission(userMission.ID, models.UserMission{UserID: 1, MissionID: 2, Completed: true})
	require.True(t, ok)
	require.True(t, userMission.Completed)
	require.True(t, repo.BorrarUserMission(userMission.ID))
	require.False(t, repo.BorrarUserMission(999))
}
