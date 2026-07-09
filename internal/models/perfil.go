package models

type User struct {
<<<<<<< HEAD:internal/models/users.go
	ID         int
	Name       string
	Email      string
	Password   string
	Level      int
	Reputation int
}

type Review struct {
	ID         int
	ReviewerID int
	ReviewedID int
	Rating     int
	Comment    string
}

type Badge struct {
	ID          int
	Name        string
	Description string
	RequiredRep int
=======
	ID         int    `json:"id" gorm:"primaryKey"`
	Name       string `json:"name"`
	Email      string `json:"email" gorm:"unique" `
	Password   string `json:"password"`
	Level      int    `json:"level"`
	Reputation int    `json:"reputation"`
}

type Review struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	ReviewerID uint   `json:"reviewer_id"`
	ReviewedID uint   `json:"reviewed_id"`
	Rating     int    `json:"rating"`
	Comment    string `json:"comment"`

	Reviewer User `gorm:"foreignKey:ReviewerID" json:"-"`
	Reviewed User `gorm:"foreignKey:ReviewedID" json:"-"`
}

type Badge struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	RequiredRep int    `json:"required_rep"`
>>>>>>> b13f60a (Guarda cambios antes de cambiar de rama):internal/models/perfil.go
}
