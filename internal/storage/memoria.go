package storage

import (
	"marketplace-api/internal/models"
	"sync"
)

type Memoria struct {
	// MODULO 1
	users        []models.User
	nextUserID   int
	reviews      []models.Review
	nextReviewID int
	badges       []models.Badge
	nextBadgeID  int

	// MODULO 2
	categorias      []models.Categoria
	nextCategoriaID int
	productos       []models.Producto
	nextProductoID  int
	orden           []models.Orden
	nextOrdenID     int

	// MODULO 3
	messages          []models.Message
	nextMessageID     int
	missions          []models.Mission
	nextMissionID     int
	usermissions      []models.UserMission
	nextUserMissionID int

	mu sync.RWMutex
}

func NuevaMemoria() *Memoria {
	return &Memoria{
		// MODULO 1
		users:        []models.User{},
		nextUserID:   1,
		reviews:      []models.Review{},
		nextReviewID: 1,
		badges:       []models.Badge{},
		nextBadgeID:  1,

		// MODULO 2
		categorias:      []models.Categoria{},
		nextCategoriaID: 1,
		productos:       []models.Producto{},
		nextProductoID:  1,
		orden:           []models.Orden{},
		nextOrdenID:     1,

		// MODULO 3
		messages:          []models.Message{},
		nextMessageID:     1,
		missions:          []models.Mission{},
		nextMissionID:     1,
		usermissions:      []models.UserMission{},
		nextUserMissionID: 1,
	}
}

// =========================================================
// MODULO 1
// =========================================================

// USERS

func (m *Memoria) SeedUsers() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.users = []models.User{
		{ID: 1, Name: "Juan Perez", Email: "juan@uleam.edu.ec", Password: "123456", Level: 1, Reputation: 20},
		{ID: 2, Name: "Maria Lopez", Email: "maria@uleam.edu.ec", Password: "123456", Level: 2, Reputation: 75},
		{ID: 3, Name: "Carlos Zambrano", Email: "carlos@uleam.edu.ec", Password: "123456", Level: 3, Reputation: 150},
		{ID: 4, Name: "Andrea Vera", Email: "andrea@uleam.edu.ec", Password: "123456", Level: 1, Reputation: 30},
		{ID: 5, Name: "Luis Mendoza", Email: "luis@uleam.edu.ec", Password: "123456", Level: 4, Reputation: 220},
		{ID: 6, Name: "Sofia Cedeño", Email: "sofia@uleam.edu.ec", Password: "123456", Level: 2, Reputation: 90},
		{ID: 7, Name: "Kevin Alcivar", Email: "kevin@uleam.edu.ec", Password: "123456", Level: 1, Reputation: 15},
	}
	m.nextUserID = 8
}

func (m *Memoria) ListarUsers() []models.User {
	m.mu.Lock()
	defer m.mu.Unlock()

	copia := make([]models.User, len(m.users))
	copy(copia, m.users)
	return copia
}

func (m *Memoria) BuscarUserPorID(id int) (models.User, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, u := range m.users {
		if u.ID == id {
			return u, true
		}
	}
	return models.User{}, false
}

func (m *Memoria) CrearUser(u models.User) models.User {
	m.mu.Lock()
	defer m.mu.Unlock()

	u.ID = m.nextUserID
	m.nextUserID++
	m.users = append(m.users, u)
	return u
}

func (m *Memoria) ActualizarUser(id int, datos models.User) (models.User, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, u := range m.users {
		if u.ID == id {
			datos.ID = id
			m.users[i] = datos
			return datos, true
		}
	}
	return models.User{}, false
}

func (m *Memoria) BorrarUser(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, u := range m.users {
		if u.ID == id {
			m.users = append(m.users[:i], m.users[i+1:]...)
			return true
		}
	}
	return false
}

// Review
func (m *Memoria) SeedReviews() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.reviews = []models.Review{
		{ID: 1, ReviewerID: 2, ReviewedID: 1, Rating: 5, Comment: "Excelente vendedor, todo fue rapido y seguro"},
		{ID: 2, ReviewerID: 1, ReviewedID: 3, Rating: 4, Comment: "Buena comunicacion y entrega puntual"},
		{ID: 3, ReviewerID: 4, ReviewedID: 2, Rating: 5, Comment: "Muy amable y responsable"},
		{ID: 4, ReviewerID: 5, ReviewedID: 1, Rating: 4, Comment: "El producto estaba en buen estado"},
		{ID: 5, ReviewerID: 6, ReviewedID: 4, Rating: 5, Comment: "Recomendado para futuras compras"},
		{ID: 6, ReviewerID: 7, ReviewedID: 5, Rating: 3, Comment: "La entrega demoro un poco pero llego bien"},
		{ID: 7, ReviewerID: 3, ReviewedID: 6, Rating: 5, Comment: "Excelente experiencia de compra"},
	}
	m.nextReviewID = 8
}

func (m *Memoria) ListarReviews() []models.Review {
	m.mu.Lock()
	defer m.mu.Unlock()

	copia := make([]models.Review, len(m.reviews))
	copy(copia, m.reviews)
	return copia
}

func (m *Memoria) BuscarReviewPorID(id int) (models.Review, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, u := range m.reviews {
		if u.ID == id {
			return u, true
		}
	}
	return models.Review{}, false
}

// Bagdges
func (m *Memoria) SeedBadges() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.badges = []models.Badge{
		{ID: 1, Name: "Novato", Description: "Alcanza 10 puntos de reputacion", RequiredRep: 10},
		{ID: 2, Name: "Colaborador", Description: "Alcanza 50 puntos de reputacion", RequiredRep: 50},
		{ID: 3, Name: "Vendedor Confiable", Description: "Alcanza 100 puntos de reputacion", RequiredRep: 100},
		{ID: 4, Name: "Comerciante Experto", Description: "Alcanza 200 puntos de reputacion", RequiredRep: 200},
		{ID: 5, Name: "Tutor Destacado", Description: "Recibe excelentes calificaciones de otros usuarios", RequiredRep: 300},
		{ID: 6, Name: "Embajador ULEAM", Description: "Mantiene una reputacion sobresaliente en la plataforma", RequiredRep: 500},
		{ID: 7, Name: "Leyenda ULEAM", Description: "Alcanza el maximo reconocimiento dentro del marketplace", RequiredRep: 1000},
	}
	m.nextBadgeID = 8
}

func (m *Memoria) ListarBadges() []models.Badge {
	m.mu.Lock()
	defer m.mu.Unlock()

	copia := make([]models.Badge, len(m.badges))
	copy(copia, m.badges)
	return copia
}

func (m *Memoria) BuscarBadgePorID(id int) (models.Badge, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, u := range m.badges {
		if u.ID == id {
			return u, true
		}
	}
	return models.Badge{}, false
}

// =========================================================
// MODULO 2
// =========================================================
// CATEGORIAS

func (m *Memoria) SeedCategorias() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.categorias = []models.Categoria{
		{ID: 1, Name: "Libros"},
		{ID: 2, Name: "Equipos tecnologicos"},
		{ID: 3, Name: "Material de laboratorio"},
		{ID: 4, Name: "Uniformes"},
		{ID: 5, Name: "Accesorios universitarios"},
		{ID: 6, Name: "Otros"},
	}
	m.nextCategoriaID = 6
}
func (m *Memoria) ListarCategorias() []models.Categoria {
	m.mu.Lock()
	defer m.mu.Unlock()

	copia := make([]models.Categoria, len(m.categorias))
	copy(copia, m.categorias)
	return copia
}

// GET:DEVUELVE UNA CATEGORIA POR ID, SI NO SE ENCUENTRA DEVUELVE UN BOOLEANO FALSE
func (m *Memoria) BuscarCategoriaPorID(id int) (models.Categoria, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, c := range m.categorias {
		if c.ID == id {
			return c, true
		}
	}
	return models.Categoria{}, false
}

func (m *Memoria) CrearCategoria(c models.Categoria) models.Categoria {
	m.mu.Lock()
	defer m.mu.Unlock()

	c.ID = m.nextCategoriaID
	m.nextCategoriaID++
	m.categorias = append(m.categorias, c)
	return c
}
func (m *Memoria) ActualizarCategoria(id int, datos models.Categoria) (models.Categoria, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, cat := range m.categorias {
		if cat.ID == id {
			datos.ID = id
			m.categorias[i] = datos
			return datos, true
		}
	}
	return models.Categoria{}, false
}
func (m *Memoria) BorrarCategoria(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, c := range m.categorias {
		if c.ID == id {
			m.categorias = append(m.categorias[:i], m.categorias[i+1:]...)
			return true
		}
	}
	return false
}

// PRODUCTOS

func (m *Memoria) SeedProductos() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.productos = []models.Producto{
		{ID: 1, Nombre: "Libro de Programación", Descripcion: "Un libro completo sobre programación en Go.", Precio: 25.99, CategoriaID: 1},
		{ID: 2, Nombre: "Libro de Física General ", Descripcion: "Libro en buen estado para carreras de ingeniería.", Precio: 20.00, CategoriaID: 2},
		{ID: 3, Nombre: "Laptop Lenovo ThinkPad P15", Descripcion: "Laptop ideal para tareas y proyectos académicos.", Precio: 680.00, CategoriaID: 3},
		{ID: 4, Nombre: "Calculadora Científica Casio", Descripcion: "Calculadora para matemáticas, física y estadística.", Precio: 20.00, CategoriaID: 4},
		{ID: 5, Nombre: "Bata de laboratorio", Descripcion: "Bata protectora para prácticas de laboratorio.", Precio: 15.00, CategoriaID: 5},
		{ID: 6, Nombre: "Kit de disección", Descripcion: "Kit completo para prácticas de biología y química.", Precio: 22.00, CategoriaID: 6},
		{ID: 7, Nombre: "Mochila universitaria", Descripcion: "Mochila resistente y espaciosa para llevar libros y laptop.", Precio: 35.00, CategoriaID: 7},
		{ID: 8, Nombre: "Estuche para laptop", Descripcion: "Estuche protector para laptop.", Precio: 12.00, CategoriaID: 8},
		{ID: 9, Nombre: "Organizador de escritorio", Descripcion: "Organizador para útiles y materiales de estudio.", Precio: 10.00, CategoriaID: 9},
	}
	m.nextProductoID = 9
}

func (m *Memoria) ListarProductos() []models.Producto {
	m.mu.Lock()
	defer m.mu.Unlock()

	copia := make([]models.Producto, len(m.productos))
	copy(copia, m.productos)
	return copia
}

func (m *Memoria) BuscarProductoPorID(id int) (models.Producto, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, p := range m.productos {
		if p.ID == id {
			return p, true
		}
	}
	return models.Producto{}, false
}

func (m *Memoria) CrearProducto(p models.Producto) models.Producto {
	m.mu.Lock()
	defer m.mu.Unlock()

	p.ID = m.nextProductoID
	m.nextProductoID++
	m.productos = append(m.productos, p)
	return p
}

func (m *Memoria) ActualizarProducto(id int, datos models.Producto) (models.Producto, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, p := range m.productos {
		if p.ID == id {
			datos.ID = id
			m.productos[i] = datos
			return datos, true
		}
	}
	return models.Producto{}, false
}

func (m *Memoria) BorrarProducto(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, p := range m.productos {
		if p.ID == id {
			m.productos = append(m.productos[:i], m.productos[i+1:]...)
			return true
		}
	}
	return false
}

//ORDENES

func (m *Memoria) SeedOrdenes() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.orden = []models.Orden{
		{ID: 1, ProductoID: 1, IDComprador: 2, Estado: "Pendiente"},
		{ID: 2, ProductoID: 3, IDComprador: 4, Estado: "Enviado"},
		{ID: 3, ProductoID: 5, IDComprador: 1, Estado: "Entregado"},
		{ID: 4, ProductoID: 9, IDComprador: 2, Estado: "Cancelado"},
	}
	m.nextOrdenID = 4
}

func (m *Memoria) ListarOrdenes() []models.Orden {
	m.mu.Lock()
	defer m.mu.Unlock()

	copia := make([]models.Orden, len(m.orden))
	copy(copia, m.orden)
	return copia
}

func (m *Memoria) BuscarOrdenPorID(id int) (models.Orden, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, o := range m.orden {
		if o.ID == id {
			return o, true
		}
	}
	return models.Orden{}, false
}

func (m *Memoria) CrearOrden(o models.Orden) models.Orden {
	m.mu.Lock()
	defer m.mu.Unlock()

	o.ID = m.nextOrdenID
	m.nextOrdenID++
	m.orden = append(m.orden, o)
	return o
}

func (m *Memoria) ActualizarOrden(id int, datos models.Orden) (models.Orden, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, o := range m.orden {
		if o.ID == id {
			datos.ID = id
			m.orden[i] = datos
			return datos, true
		}
	}
	return models.Orden{}, false
}

func (m *Memoria) BorrarOrden(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, o := range m.orden {
		if o.ID == id {
			m.orden = append(m.orden[:i], m.orden[i+1:]...)
			return true
		}
	}
	return false
}

// =========================================================
// MODULO 3
// =========================================================
// ///////entidad Message
func (m *Memoria) SeedMessages() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.messages = []models.Message{
		{ID: 1, SenderID: 1, ReceiverID: 2, Content: "Hola María, ¿aún tienes disponible el libro de Redes?"},
		{ID: 2, SenderID: 2, ReceiverID: 1, Content: "Sí, todavía está disponible."},
		{ID: 3, SenderID: 3, ReceiverID: 4, Content: "Buenas, ¿podrías compartir los apuntes de Base de Datos?"},
		{ID: 4, SenderID: 4, ReceiverID: 3, Content: "Claro, te los envío en un momento."},
		{ID: 5, SenderID: 5, ReceiverID: 6, Content: "¿Sigues ofreciendo tutorías de Programación Web?"},
		{ID: 6, SenderID: 6, ReceiverID: 5, Content: "Sí, tengo horarios disponibles esta semana."},
		{ID: 7, SenderID: 7, ReceiverID: 1, Content: "Hola, estoy interesado en el libro de Cálculo."},
	}
	m.nextMessageID = 8
}

func (m *Memoria) ListarMessages() []models.Message {
	m.mu.Lock()
	defer m.mu.Unlock()

	copia := make([]models.Message, len(m.messages))
	copy(copia, m.messages)
	return copia
}

// BuscarProductoPorID devuelve el producto con el ID dado (patrón comma-ok).
func (m *Memoria) BuscarMessagePorID(id int) (models.Message, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, p := range m.messages {
		if p.ID == id {
			return p, true
		}
	}
	return models.Message{}, false
}

// CrearProducto agrega un Mensaje nuevo y devuelve el producto con ID asignado.
func (m *Memoria) CrearMessage(p models.Message) models.Message {
	m.mu.Lock()
	defer m.mu.Unlock()

	p.ID = m.nextMessageID
	m.nextMessageID++
	m.messages = append(m.messages, p)
	return p
}

// ActualizarMensaje reemplaza el mensaje con el ID dado.
func (m *Memoria) ActualizarMessage(id int, datos models.Message) (models.Message, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, p := range m.messages {
		if p.ID == id {
			datos.ID = id
			m.messages[i] = datos
			return datos, true
		}
	}
	return models.Message{}, false
}

// BorrarMensaje elimina el mensaje con el ID dado.
func (m *Memoria) BorrarMessage(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, p := range m.messages {
		if p.ID == id {
			m.messages = append(m.messages[:i], m.messages[i+1:]...)
			return true
		}
	}
	return false
}

// ///////////entidad Mission
func (m *Memoria) SeedMissions() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.missions = []models.Mission{
		{ID: 1, Title: "Primera Publicación", Description: "Publica tu primer producto en el marketplace", RequiredLevel: 1, RewardPoints: 20},
		{ID: 2, Title: "Primera Compra", Description: "Realiza tu primera compra dentro de la plataforma", RequiredLevel: 1, RewardPoints: 30},
		{ID: 3, Title: "Vendedor Activo", Description: "Publica 5 productos diferentes", RequiredLevel: 2, RewardPoints: 50},
		{ID: 4, Title: "Comprador Frecuente", Description: "Completa 5 compras exitosas", RequiredLevel: 2, RewardPoints: 50},
		{ID: 5, Title: "Miembro Confiable", Description: "Alcanza 100 puntos de reputación", RequiredLevel: 3, RewardPoints: 75},
		{ID: 6, Title: "Comerciante Experto", Description: "Realiza 10 ventas exitosas", RequiredLevel: 4, RewardPoints: 100},
		{ID: 7, Title: "Leyenda ULEAM", Description: "Alcanza el nivel 5 y mantén una reputación positiva", RequiredLevel: 5, RewardPoints: 150},
	}
	m.nextMissionID = 8
}

func (m *Memoria) ListarMissions() []models.Mission {
	m.mu.Lock()
	defer m.mu.Unlock()

	copia := make([]models.Mission, len(m.missions))
	copy(copia, m.missions)
	return copia
}

func (m *Memoria) BuscarMisionPorID(id int) (models.Mission, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, p := range m.missions {
		if p.ID == id {
			return p, true
		}
	}
	return models.Mission{}, false
}

func (m *Memoria) CrearMision(p models.Mission) models.Mission {
	m.mu.Lock()
	defer m.mu.Unlock()

	p.ID = m.nextMissionID
	m.nextMissionID++
	m.missions = append(m.missions, p)
	return p
}

// ActualizarMision reemplaza la misión con el ID dado.
func (m *Memoria) ActualizarMision(id int, datos models.Mission) (models.Mission, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, p := range m.missions {
		if p.ID == id {
			datos.ID = id
			m.missions[i] = datos
			return datos, true
		}
	}
	return models.Mission{}, false
}

// BorrarMision elimina la misión con el ID dado.
func (m *Memoria) BorrarMision(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, p := range m.missions {
		if p.ID == id {
			m.missions = append(m.missions[:i], m.missions[i+1:]...)
			return true
		}
	}
	return false
}

// ///////entidad UserMission

func (m *Memoria) SeedUsermissions() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.usermissions = []models.UserMission{
		{ID: 1, UserID: 1, MissionID: 1, Completed: true},
		{ID: 2, UserID: 2, MissionID: 1, Completed: true},
		{ID: 3, UserID: 3, MissionID: 2, Completed: false},
		{ID: 4, UserID: 4, MissionID: 3, Completed: true},
		{ID: 5, UserID: 5, MissionID: 4, Completed: false},
		{ID: 6, UserID: 6, MissionID: 5, Completed: true},
		{ID: 7, UserID: 7, MissionID: 2, Completed: false},
	}
	m.nextUserMissionID = 8
}
func (m *Memoria) ListarUsermissions() []models.UserMission {
	m.mu.Lock()
	defer m.mu.Unlock()

	copia := make([]models.UserMission, len(m.usermissions))
	copy(copia, m.usermissions)
	return copia
}

func (m *Memoria) BuscarUserMissionPorID(id int) (models.UserMission, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, p := range m.usermissions {
		if p.ID == id {
			return p, true
		}
	}
	return models.UserMission{}, false
}

func (m *Memoria) CrearUserMission(p models.UserMission) models.UserMission {
	m.mu.Lock()
	defer m.mu.Unlock()

	p.ID = m.nextUserMissionID
	m.nextUserMissionID++
	m.usermissions = append(m.usermissions, p)
	return p
}

func (m *Memoria) ActualizarUserMission(id int, datos models.UserMission) (models.UserMission, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, p := range m.usermissions {
		if p.ID == id {
			datos.ID = id
			m.usermissions[i] = datos
			return datos, true
		}
	}
	return models.UserMission{}, false
}

// BorrarUserMission elimina la misión con el ID dado.
func (m *Memoria) BorrarUserMission(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, p := range m.usermissions {
		if p.ID == id {
			m.usermissions = append(m.usermissions[:i], m.usermissions[i+1:]...)
			return true
		}
	}
	return false
}
