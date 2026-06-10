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
