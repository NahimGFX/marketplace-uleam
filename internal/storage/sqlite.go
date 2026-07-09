package storage

import (
	"marketplace-api/internal/models"

	"gorm.io/gorm"
)

type AlmacenSQLite struct {
	db *gorm.DB
}

func NuevoAlmacenSQLite(db *gorm.DB) *AlmacenSQLite {
	return &AlmacenSQLite{
		db: db,
	}
}

// =========================
// USERS
// =========================

func (s *AlmacenSQLite) ListarUsers() []models.User {
	var users []models.User
	s.db.Find(&users)
	return users
}

func (s *AlmacenSQLite) BuscarUserPorID(id int) (models.User, bool) {
	var user models.User

	if err := s.db.First(&user, id).Error; err != nil {
		return models.User{}, false
	}

	return user, true
}

func (s *AlmacenSQLite) CrearUser(u models.User) models.User {
	s.db.Create(&u)
	return u
}

func (s *AlmacenSQLite) ActualizarUser(id int, datos models.User) (models.User, bool) {
	var user models.User

	if err := s.db.First(&user, id).Error; err != nil {
		return models.User{}, false
	}

	datos.ID = id
	s.db.Save(&datos)

	return datos, true
}

func (s *AlmacenSQLite) BorrarUser(id int) bool {
	result := s.db.Delete(&models.User{}, id)
	return result.RowsAffected > 0
}

//
// =========================
// REVIEWS
// =========================
//

func (s *AlmacenSQLite) ListarReviews() []models.Review {
	var reviews []models.Review
	s.db.Find(&reviews)
	return reviews
}

func (s *AlmacenSQLite) BuscarReviewPorID(id int) (models.Review, bool) {
	var review models.Review

	if err := s.db.First(&review, id).Error; err != nil {
		return models.Review{}, false
	}

	return review, true
}

func (s *AlmacenSQLite) CrearReview(r models.Review) models.Review {
	s.db.Create(&r)
	return r
}

func (s *AlmacenSQLite) ActualizarReview(id int, datos models.Review) (models.Review, bool) {
	var review models.Review

	if err := s.db.First(&review, id).Error; err != nil {
		return models.Review{}, false
	}

	datos.ID = id
	s.db.Save(&datos)

	return datos, true
}

func (s *AlmacenSQLite) BorrarReview(id int) bool {
	result := s.db.Delete(&models.Review{}, id)
	return result.RowsAffected > 0
}

//
// =========================
// BADGES
// =========================
//

func (s *AlmacenSQLite) ListarBadges() []models.Badge {
	var badges []models.Badge
	s.db.Find(&badges)
	return badges
}

func (s *AlmacenSQLite) BuscarBadgePorID(id int) (models.Badge, bool) {
	var badge models.Badge

	if err := s.db.First(&badge, id).Error; err != nil {
		return models.Badge{}, false
	}

	return badge, true
}

func (s *AlmacenSQLite) CrearBadge(b models.Badge) models.Badge {
	s.db.Create(&b)
	return b
}

func (s *AlmacenSQLite) ActualizarBadge(id int, datos models.Badge) (models.Badge, bool) {
	var badge models.Badge

	if err := s.db.First(&badge, id).Error; err != nil {
		return models.Badge{}, false
	}

	datos.ID = id
	s.db.Save(&datos)

	return datos, true
}

func (s *AlmacenSQLite) BorrarBadge(id int) bool {
	result := s.db.Delete(&models.Badge{}, id)
	return result.RowsAffected > 0
}

//
// MESSAGES
//

func (s *AlmacenSQLite) ListarMessages() []models.Message {
	var messages []models.Message
	s.db.Find(&messages)
	return messages
}

func (s *AlmacenSQLite) BuscarMessagePorID(id int) (models.Message, bool) {
	var message models.Message

	if err := s.db.First(&message, id).Error; err != nil {
		return models.Message{}, false
	}

	return message, true
}

func (s *AlmacenSQLite) CrearMessage(message models.Message) models.Message {
	s.db.Create(&message)
	return message
}

func (s *AlmacenSQLite) ActualizarMessage(id int, datos models.Message) (models.Message, bool) {
	var message models.Message

	if err := s.db.First(&message, id).Error; err != nil {
		return models.Message{}, false
	}

	datos.ID = id
	s.db.Save(&datos)

	return datos, true
}

func (s *AlmacenSQLite) BorrarMessage(id int) bool {
	result := s.db.Delete(&models.Message{}, id)
	return result.RowsAffected > 0
}

//
// MISIONES
//

func (s *AlmacenSQLite) ListarMissions() []models.Mission {
	var missions []models.Mission
	s.db.Find(&missions)
	return missions
}

func (s *AlmacenSQLite) BuscarMisionPorID(id int) (models.Mission, bool) {
	var mission models.Mission

	if err := s.db.First(&mission, id).Error; err != nil {
		return models.Mission{}, false
	}

	return mission, true
}

func (s *AlmacenSQLite) CrearMision(mission models.Mission) models.Mission {
	s.db.Create(&mission)
	return mission
}

func (s *AlmacenSQLite) ActualizarMision(id int, datos models.Mission) (models.Mission, bool) {
	var mission models.Mission

	if err := s.db.First(&mission, id).Error; err != nil {
		return models.Mission{}, false
	}

	datos.ID = id
	s.db.Save(&datos)

	return datos, true
}

func (s *AlmacenSQLite) BorrarMision(id int) bool {
	result := s.db.Delete(&models.Mission{}, id)
	return result.RowsAffected > 0
}

//
// USER MISSIONS
//

func (s *AlmacenSQLite) ListarUserMissions() []models.UserMission {
	var userMissions []models.UserMission
	s.db.Find(&userMissions)
	return userMissions
}

func (s *AlmacenSQLite) BuscarUserMissionPorID(id int) (models.UserMission, bool) {
	var userMission models.UserMission

	if err := s.db.First(&userMission, id).Error; err != nil {
		return models.UserMission{}, false
	}

	return userMission, true
}

func (s *AlmacenSQLite) CrearUserMission(userMission models.UserMission) models.UserMission {
	s.db.Create(&userMission)
	return userMission
}

func (s *AlmacenSQLite) ActualizarUserMission(id int, datos models.UserMission) (models.UserMission, bool) {
	var userMission models.UserMission

	if err := s.db.First(&userMission, id).Error; err != nil {
		return models.UserMission{}, false
	}

	datos.ID = id
	s.db.Save(&datos)

	return datos, true
}

func (s *AlmacenSQLite) BorrarUserMission(id int) bool {
	result := s.db.Delete(&models.UserMission{}, id)
	return result.RowsAffected > 0
}

// =========================
// CATEGORIAS
// =========================

func (s *AlmacenSQLite) ListarCategorias() []models.Categoria {
	var categorias []models.Categoria
	s.db.Find(&categorias)
	return categorias
}

func (s *AlmacenSQLite) BuscarCategoriaPorID(id int) (models.Categoria, bool) {
	var categoria models.Categoria

	if err := s.db.First(&categoria, id).Error; err != nil {
		return models.Categoria{}, false
	}

	return categoria, true
}

func (s *AlmacenSQLite) CrearCategoria(c models.Categoria) models.Categoria {
	s.db.Create(&c)
	return c
}

func (s *AlmacenSQLite) ActualizarCategoria(id int, datos models.Categoria) (models.Categoria, bool) {
	var categoria models.Categoria

	if err := s.db.First(&categoria, id).Error; err != nil {
		return models.Categoria{}, false
	}

	datos.ID = id
	s.db.Save(&datos)

	return datos, true
}

func (s *AlmacenSQLite) BorrarCategoria(id int) bool {
	result := s.db.Delete(&models.Categoria{}, id)
	return result.RowsAffected > 0
}

// =========================
// PRODUCTOS
// =========================

func (s *AlmacenSQLite) ListarProductos() []models.Producto {
	var productos []models.Producto
	s.db.Find(&productos)
	return productos
}

func (s *AlmacenSQLite) BuscarProductoPorID(id int) (models.Producto, bool) {
	var producto models.Producto

	if err := s.db.First(&producto, id).Error; err != nil {
		return models.Producto{}, false
	}

	return producto, true
}

func (s *AlmacenSQLite) CrearProducto(p models.Producto) models.Producto {
	s.db.Create(&p)
	return p
}

func (s *AlmacenSQLite) ActualizarProducto(id int, datos models.Producto) (models.Producto, bool) {
	var producto models.Producto

	if err := s.db.First(&producto, id).Error; err != nil {
		return models.Producto{}, false
	}

	datos.ID = id
	s.db.Save(&datos)

	return datos, true
}

func (s *AlmacenSQLite) BorrarProducto(id int) bool {
	result := s.db.Delete(&models.Producto{}, id)
	return result.RowsAffected > 0
}

// =========================
// ORDENES
// =========================

func (s *AlmacenSQLite) ListarOrdenes() []models.Orden {
	var ordenes []models.Orden
	s.db.Find(&ordenes)
	return ordenes
}

func (s *AlmacenSQLite) BuscarOrdenPorID(id int) (models.Orden, bool) {
	var orden models.Orden

	if err := s.db.First(&orden, id).Error; err != nil {
		return models.Orden{}, false
	}

	return orden, true
}

func (s *AlmacenSQLite) CrearOrden(o models.Orden) models.Orden {
	s.db.Create(&o)
	return o
}

func (s *AlmacenSQLite) ActualizarOrden(id int, datos models.Orden) (models.Orden, bool) {
	var orden models.Orden

	if err := s.db.First(&orden, id).Error; err != nil {
		return models.Orden{}, false
	}

	datos.ID = id
	s.db.Save(&datos)

	return datos, true
}

func (s *AlmacenSQLite) BorrarOrden(id int) bool {
	result := s.db.Delete(&models.Orden{}, id)
	return result.RowsAffected > 0
}

//
// SEED (DATOS INICIALES)
//

func (a *AlmacenSQLite) SembrarSiVacio() {
	var n int64

	// =========================
	// USERS
	// =========================
	a.db.Model(&models.User{}).Count(&n)
	if n == 0 {
		users := []models.User{
			{Name: "Juan Perez", Email: "juan@uleam.edu.ec", Password: "123456", Role: "admin", Level: 1, Reputation: 20},
			{Name: "Maria Lopez", Email: "maria@uleam.edu.ec", Password: "123456", Role: "estudiante", Level: 2, Reputation: 75},
			{Name: "Carlos Zambrano", Email: "carlos@uleam.edu.ec", Password: "123456", Role: "estudiante", Level: 3, Reputation: 150},
			{Name: "Andrea Vera", Email: "andrea@uleam.edu.ec", Password: "123456", Role: "estudiante", Level: 1, Reputation: 30},
			{Name: "Luis Mendoza", Email: "luis@uleam.edu.ec", Password: "123456", Role: "estudiante", Level: 4, Reputation: 220},
			{Name: "Sofia Cedeño", Email: "sofia@uleam.edu.ec", Password: "123456", Role: "estudiante", Level: 2, Reputation: 90},
			{Name: "Kevin Alcivar", Email: "kevin@uleam.edu.ec", Password: "123456", Role: "estudiante", Level: 1, Reputation: 15},
		}

		a.db.Create(&users)
	}
	a.db.Model(&models.User{}).Where("role = '' OR role IS NULL").Update("role", "estudiante")
	a.db.Model(&models.User{}).Where("email = ?", "juan@uleam.edu.ec").Update("role", "admin")

	// =========================
	// REVIEWS
	// =========================
	if n == 0 {
		reviews := []models.Review{
			{ReviewerID: 2, ReviewedID: 1, Rating: 5, Comment: "Excelente vendedor, todo fue rapido y seguro"},
			{ReviewerID: 1, ReviewedID: 3, Rating: 4, Comment: "Buena comunicacion y entrega puntual"},
			{ReviewerID: 4, ReviewedID: 2, Rating: 5, Comment: "Muy amable y responsable"},
			{ReviewerID: 5, ReviewedID: 1, Rating: 4, Comment: "El producto estaba en buen estado"},
			{ReviewerID: 6, ReviewedID: 4, Rating: 5, Comment: "Recomendado para futuras compras"},
			{ReviewerID: 7, ReviewedID: 5, Rating: 3, Comment: "La entrega demoro un poco pero llego bien"},
			{ReviewerID: 3, ReviewedID: 6, Rating: 5, Comment: "Excelente experiencia de compra"},
		}
		a.db.Create(&reviews)
	}

	// =========================
	// BADGES
	// =========================
	if n == 0 {
		b := []models.Badge{
			{Name: "Novato", Description: "Alcanza 10 puntos de reputacion", RequiredRep: 10},
			{Name: "Colaborador", Description: "Alcanza 50 puntos de reputacion", RequiredRep: 50},
			{Name: "Vendedor Confiable", Description: "Alcanza 100 puntos de reputacion", RequiredRep: 100},
			{Name: "Comerciante Experto", Description: "Alcanza 200 puntos de reputacion", RequiredRep: 200},
			{Name: "Tutor Destacado", Description: "Recibe excelentes calificaciones de otros usuarios", RequiredRep: 300},
			{Name: "Embajador ULEAM", Description: "Mantiene una reputacion sobresaliente en la plataforma", RequiredRep: 500},
			{Name: "Leyenda ULEAM", Description: "Alcanza el maximo reconocimiento dentro del marketplace", RequiredRep: 1000},
		}
		a.db.Create(&b)
	}

	// =========================
	// REVIEWS
	// =========================
	// Si ya hay mensajes, asumimos que ya está sembrado
	a.db.Model(&models.Categoria{}).Count(&n)
	if n > 0 {
		return
	}

	categorias := []models.Categoria{
		{Name: "Libros"},
		{Name: "Equipos tecnológicos"},
		{Name: "Material de laboratorio"},
		{Name: "Uniformes"},
		{Name: "Accesorios universitarios"},
		{Name: "Otros"},
	}

	a.db.Create(&categorias)

	productos := []models.Producto{
		{Nombre: "Libro de Programación", Descripcion: "Un libro completo sobre programación en Go.", Precio: 25.99, CategoriaID: 1},
		{Nombre: "Libro de Física General", Descripcion: "Libro en buen estado para ingeniería.", Precio: 20.00, CategoriaID: 1},
		{Nombre: "Laptop Lenovo ThinkPad P15", Descripcion: "Laptop ideal para proyectos académicos.", Precio: 680.00, CategoriaID: 3},
		{Nombre: "Calculadora Científica Casio", Descripcion: "Para matemáticas y física.", Precio: 20.00, CategoriaID: 4},
		{Nombre: "Bata de laboratorio", Descripcion: "Protección para prácticas.", Precio: 15.00, CategoriaID: 3},
		{Nombre: "Kit de disección", Descripcion: "Prácticas de biología.", Precio: 22.00, CategoriaID: 3},
		{Nombre: "Mochila universitaria", Descripcion: "Resistente para libros y laptop.", Precio: 35.00, CategoriaID: 5},
		{Nombre: "Estuche para laptop", Descripcion: "Protección portátil.", Precio: 12.00, CategoriaID: 5},
		{Nombre: "Organizador de escritorio", Descripcion: "Orden para materiales.", Precio: 10.00, CategoriaID: 6},
	}

	a.db.Create(&productos)

	ordenes := []models.Orden{
		{ID: 1, ProductoID: 1, IDComprador: 2, Estado: "Pendiente"},
		{ID: 2, ProductoID: 3, IDComprador: 4, Estado: "Enviado"},
		{ID: 3, ProductoID: 5, IDComprador: 1, Estado: "Entregado"},
		{ID: 4, ProductoID: 9, IDComprador: 2, Estado: "Cancelado"},
	}
	a.db.Create(&ordenes)

	// Si ya hay mensajes, asumimos que ya está sembrado
	a.db.Model(&models.Message{}).Count(&n)
	if n > 0 {
		return
	}

	// =========================
	// MESSAGES (opcional seed)
	// =========================
	messages := []models.Message{
		{ID: 1, SenderID: 1, ReceiverID: 2, Content: "Hola María, ¿aún tienes disponible el libro de Redes?"},
		{ID: 2, SenderID: 2, ReceiverID: 1, Content: "Sí, todavía está disponible."},
		{ID: 3, SenderID: 3, ReceiverID: 4, Content: "Buenas, ¿podrías compartir los apuntes de Base de Datos?"},
		{ID: 4, SenderID: 4, ReceiverID: 3, Content: "Claro, te los envío en un momento."},
		{ID: 5, SenderID: 5, ReceiverID: 6, Content: "¿Sigues ofreciendo tutorías de Programación Web?"},
		{ID: 6, SenderID: 6, ReceiverID: 5, Content: "Sí, tengo horarios disponibles esta semana."},
		{ID: 7, SenderID: 7, ReceiverID: 1, Content: "Hola, estoy interesado en el libro de Cálculo."},
	}
	a.db.Create(&messages)

	// =========================
	// MISSIONS
	// =========================
	missions := []models.Mission{
		{ID: 1, Title: "Primera Publicación", Description: "Publica tu primer producto en el marketplace", RequiredLevel: 1, RewardPoints: 20},
		{ID: 2, Title: "Primera Compra", Description: "Realiza tu primera compra dentro de la plataforma", RequiredLevel: 1, RewardPoints: 30},
		{ID: 3, Title: "Vendedor Activo", Description: "Publica 5 productos diferentes", RequiredLevel: 2, RewardPoints: 50},
		{ID: 4, Title: "Comprador Frecuente", Description: "Completa 5 compras exitosas", RequiredLevel: 2, RewardPoints: 50},
		{ID: 5, Title: "Miembro Confiable", Description: "Alcanza 100 puntos de reputación", RequiredLevel: 3, RewardPoints: 75},
		{ID: 6, Title: "Comerciante Experto", Description: "Realiza 10 ventas exitosas", RequiredLevel: 4, RewardPoints: 100},
		{ID: 7, Title: "Leyenda ULEAM", Description: "Alcanza el nivel 5 y mantén una reputación positiva", RequiredLevel: 5, RewardPoints: 150},
	}
	a.db.Create(&missions)

	// =========================
	// USER MISSIONS
	// =========================
	userMissions := []models.UserMission{
		{ID: 1, UserID: 1, MissionID: 1, Completed: true},
		{ID: 2, UserID: 2, MissionID: 1, Completed: true},
		{ID: 3, UserID: 3, MissionID: 2, Completed: false},
		{ID: 4, UserID: 4, MissionID: 3, Completed: true},
		{ID: 5, UserID: 5, MissionID: 4, Completed: false},
		{ID: 6, UserID: 6, MissionID: 5, Completed: true},
		{ID: 7, UserID: 7, MissionID: 2, Completed: false},
	}
	a.db.Create(&userMissions)
}
