package storage

import (
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	"marketplace-api/internal/models"
)

func nuevoAlmacenSQLiteTest(t *testing.T) *AlmacenSQLite {
	t.Helper()

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)

	err = db.AutoMigrate(
		&models.User{},
		&models.Review{},
		&models.Badge{},
		&models.Categoria{},
		&models.Producto{},
		&models.Orden{},
		&models.Message{},
		&models.Mission{},
		&models.UserMission{},
	)
	require.NoError(t, err)

	return NuevoAlmacenSQLite(db)
}

func TestAlmacenSQLite_UserRepository_CRUD(t *testing.T) {
	repo := nuevoAlmacenSQLiteTest(t)

	creado := repo.CrearUser(models.User{Name: "Ana", Email: "ana@uleam.edu.ec", Password: "123456"})
	require.NotZero(t, creado.ID)

	encontrado, ok := repo.BuscarUserPorID(creado.ID)
	require.True(t, ok)
	require.Equal(t, "Ana", encontrado.Name)
	require.Len(t, repo.ListarUsers(), 1)

	actualizado, ok := repo.ActualizarUser(creado.ID, models.User{Name: "Ana Editada", Email: "ana@uleam.edu.ec", Password: "abc"})
	require.True(t, ok)
	require.Equal(t, "Ana Editada", actualizado.Name)

	_, ok = repo.ActualizarUser(999, models.User{Name: "Nadie"})
	require.False(t, ok)
	require.True(t, repo.BorrarUser(creado.ID))
	require.False(t, repo.BorrarUser(999))
}

func TestAlmacenSQLite_ReviewRepository_CRUD(t *testing.T) {
	repo := nuevoAlmacenSQLiteTest(t)

	creado := repo.CrearReview(models.Review{ReviewerID: 1, ReviewedID: 2, Rating: 5, Comment: "Excelente"})
	require.NotZero(t, creado.ID)

	encontrado, ok := repo.BuscarReviewPorID(creado.ID)
	require.True(t, ok)
	require.Equal(t, "Excelente", encontrado.Comment)
	require.Len(t, repo.ListarReviews(), 1)

	actualizado, ok := repo.ActualizarReview(creado.ID, models.Review{ReviewerID: 1, ReviewedID: 2, Rating: 4, Comment: "Bueno"})
	require.True(t, ok)
	require.Equal(t, "Bueno", actualizado.Comment)

	_, ok = repo.ActualizarReview(999, models.Review{Comment: "Nada"})
	require.False(t, ok)
	require.True(t, repo.BorrarReview(creado.ID))
	require.False(t, repo.BorrarReview(999))
}

func TestAlmacenSQLite_BadgeRepository_CRUD(t *testing.T) {
	repo := nuevoAlmacenSQLiteTest(t)

	creado := repo.CrearBadge(models.Badge{Name: "Novato", Description: "Primer logro", RequiredRep: 10})
	require.NotZero(t, creado.ID)

	encontrado, ok := repo.BuscarBadgePorID(creado.ID)
	require.True(t, ok)
	require.Equal(t, "Novato", encontrado.Name)
	require.Len(t, repo.ListarBadges(), 1)

	actualizado, ok := repo.ActualizarBadge(creado.ID, models.Badge{Name: "Experto", Description: "Mejorado", RequiredRep: 100})
	require.True(t, ok)
	require.Equal(t, "Experto", actualizado.Name)

	_, ok = repo.ActualizarBadge(999, models.Badge{Name: "Nadie"})
	require.False(t, ok)
	require.True(t, repo.BorrarBadge(creado.ID))
	require.False(t, repo.BorrarBadge(999))
}

func TestAlmacenSQLite_CategoriaRepository_CRUD(t *testing.T) {
	repo := nuevoAlmacenSQLiteTest(t)

	creada := repo.CrearCategoria(models.Categoria{Name: "Libros"})
	require.NotZero(t, creada.ID)

	encontrada, ok := repo.BuscarCategoriaPorID(creada.ID)
	require.True(t, ok)
	require.Equal(t, "Libros", encontrada.Name)
	require.Len(t, repo.ListarCategorias(), 1)

	actualizada, ok := repo.ActualizarCategoria(creada.ID, models.Categoria{Name: "Tecnologia"})
	require.True(t, ok)
	require.Equal(t, "Tecnologia", actualizada.Name)

	_, ok = repo.ActualizarCategoria(999, models.Categoria{Name: "Nadie"})
	require.False(t, ok)
	require.True(t, repo.BorrarCategoria(creada.ID))
	require.False(t, repo.BorrarCategoria(999))
}

func TestAlmacenSQLite_ProductoRepository_CRUD(t *testing.T) {
	repo := nuevoAlmacenSQLiteTest(t)

	creado := repo.CrearProducto(models.Producto{Nombre: "Calculadora", Descripcion: "Cientifica", Precio: 15, CategoriaID: 1})
	require.NotZero(t, creado.ID)

	encontrado, ok := repo.BuscarProductoPorID(creado.ID)
	require.True(t, ok)
	require.Equal(t, "Calculadora", encontrado.Nombre)
	require.Len(t, repo.ListarProductos(), 1)

	actualizado, ok := repo.ActualizarProducto(creado.ID, models.Producto{Nombre: "Calculadora Pro", Precio: 20, CategoriaID: 1})
	require.True(t, ok)
	require.Equal(t, "Calculadora Pro", actualizado.Nombre)

	_, ok = repo.ActualizarProducto(999, models.Producto{Nombre: "Nadie"})
	require.False(t, ok)
	require.True(t, repo.BorrarProducto(creado.ID))
	require.False(t, repo.BorrarProducto(999))
}

func TestAlmacenSQLite_OrdenRepository_CRUD(t *testing.T) {
	repo := nuevoAlmacenSQLiteTest(t)

	creada := repo.CrearOrden(models.Orden{ProductoID: 1, IDComprador: 2, Estado: "pendiente"})
	require.NotZero(t, creada.ID)

	encontrada, ok := repo.BuscarOrdenPorID(creada.ID)
	require.True(t, ok)
	require.Equal(t, "pendiente", encontrada.Estado)
	require.Len(t, repo.ListarOrdenes(), 1)

	actualizada, ok := repo.ActualizarOrden(creada.ID, models.Orden{ProductoID: 1, IDComprador: 2, Estado: "pagado"})
	require.True(t, ok)
	require.Equal(t, "pagado", actualizada.Estado)

	_, ok = repo.ActualizarOrden(999, models.Orden{Estado: "nada"})
	require.False(t, ok)
	require.True(t, repo.BorrarOrden(creada.ID))
	require.False(t, repo.BorrarOrden(999))
}

func TestAlmacenSQLite_MessageRepository_CRUD(t *testing.T) {
	repo := nuevoAlmacenSQLiteTest(t)

	creado := repo.CrearMessage(models.Message{SenderID: 1, ReceiverID: 2, Content: "Hola"})
	require.NotZero(t, creado.ID)

	encontrado, ok := repo.BuscarMessagePorID(creado.ID)
	require.True(t, ok)
	require.Equal(t, "Hola", encontrado.Content)
	require.Len(t, repo.ListarMessages(), 1)

	actualizado, ok := repo.ActualizarMessage(creado.ID, models.Message{SenderID: 1, ReceiverID: 2, Content: "Editado"})
	require.True(t, ok)
	require.Equal(t, "Editado", actualizado.Content)

	_, ok = repo.ActualizarMessage(999, models.Message{Content: "Nada"})
	require.False(t, ok)
	require.True(t, repo.BorrarMessage(creado.ID))
	require.False(t, repo.BorrarMessage(999))
}

func TestAlmacenSQLite_MissionRepository_CRUD(t *testing.T) {
	repo := nuevoAlmacenSQLiteTest(t)

	creada := repo.CrearMision(models.Mission{Title: "Primera venta", Description: "Vende algo", RequiredLevel: 1, RewardPoints: 20})
	require.NotZero(t, creada.ID)

	encontrada, ok := repo.BuscarMisionPorID(creada.ID)
	require.True(t, ok)
	require.Equal(t, "Primera venta", encontrada.Title)
	require.Len(t, repo.ListarMissions(), 1)

	actualizada, ok := repo.ActualizarMision(creada.ID, models.Mission{Title: "Venta pro", Description: "Editada", RequiredLevel: 2, RewardPoints: 50})
	require.True(t, ok)
	require.Equal(t, "Venta pro", actualizada.Title)

	_, ok = repo.ActualizarMision(999, models.Mission{Title: "Nada"})
	require.False(t, ok)
	require.True(t, repo.BorrarMision(creada.ID))
	require.False(t, repo.BorrarMision(999))
}

func TestAlmacenSQLite_UserMissionRepository_CRUD(t *testing.T) {
	repo := nuevoAlmacenSQLiteTest(t)

	creada := repo.CrearUserMission(models.UserMission{UserID: 1, MissionID: 2, Completed: false})
	require.NotZero(t, creada.ID)

	encontrada, ok := repo.BuscarUserMissionPorID(creada.ID)
	require.True(t, ok)
	require.False(t, encontrada.Completed)
	require.Len(t, repo.ListarUserMissions(), 1)

	actualizada, ok := repo.ActualizarUserMission(creada.ID, models.UserMission{UserID: 1, MissionID: 2, Completed: true})
	require.True(t, ok)
	require.True(t, actualizada.Completed)

	_, ok = repo.ActualizarUserMission(999, models.UserMission{Completed: true})
	require.False(t, ok)
	require.True(t, repo.BorrarUserMission(creada.ID))
	require.False(t, repo.BorrarUserMission(999))
}

func TestUsuarioGORM_CRUDPorEmailYPassword(t *testing.T) {
	almacen := nuevoAlmacenSQLiteTest(t)
	repo := NuevoUsuarioGORM(almacen.db)

	creado, err := repo.CrearUsuario(models.User{Name: "Luis", Email: "luis@uleam.edu.ec", Password: "vieja"})
	require.NoError(t, err)
	require.NotZero(t, creado.ID)

	encontrado, ok := repo.BuscarUsuarioPorEmail("luis@uleam.edu.ec")
	require.True(t, ok)
	require.Equal(t, creado.ID, encontrado.ID)

	err = repo.ActualizarPasswordUsuario(creado.ID, "nueva")
	require.NoError(t, err)

	actualizado, ok := repo.BuscarUsuarioPorEmail("luis@uleam.edu.ec")
	require.True(t, ok)
	require.Equal(t, "nueva", actualizado.Password)

	_, ok = repo.BuscarUsuarioPorEmail("nadie@uleam.edu.ec")
	require.False(t, ok)
}
