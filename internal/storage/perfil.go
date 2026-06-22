package storage

import (
	"marketplace-api/internal/models"
)

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

func (m *Memoria) CrearReview(u models.Review) models.Review {
	m.mu.Lock()
	defer m.mu.Unlock()

	u.ID = m.nextReviewID
	m.nextReviewID++
	m.reviews = append(m.reviews, u)
	return u
}

func (m *Memoria) ActualizarReview(id int, datos models.Review) (models.Review, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, u := range m.reviews {
		if u.ID == id {
			datos.ID = id
			m.reviews[i] = datos
			return datos, true
		}
	}
	return models.Review{}, false

}

func (m *Memoria) BorrarReview(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, u := range m.reviews {
		if u.ID == id {
			m.reviews = append(m.reviews[:i], m.reviews[i+1:]...)
			return true
		}
	}
	return false
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

func (m *Memoria) CrearBadge(u models.Badge) models.Badge {
	m.mu.Lock()
	defer m.mu.Unlock()

	u.ID = m.nextBadgeID
	m.nextBadgeID++
	m.badges = append(m.badges, u)
	return u
}

func (m *Memoria) ActualizarBadge(id int, datos models.Badge) (models.Badge, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, u := range m.badges {
		if u.ID == id {
			datos.ID = id
			m.badges[i] = datos
			return datos, true
		}
	}
	return models.Badge{}, false
}

func (m *Memoria) BorrarBadge(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, u := range m.badges {
		if u.ID == id {
			m.badges = append(m.badges[:i], m.badges[i+1:]...)
			return true
		}
	}
	return false
}
