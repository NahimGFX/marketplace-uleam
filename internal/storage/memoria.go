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

// =========================================================
// MODULO 2
// =========================================================

// =========================================================
// MODULO 3
// =========================================================
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
