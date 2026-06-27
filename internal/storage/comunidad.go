package storage

import (
	"marketplace-api/internal/models"
)

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
