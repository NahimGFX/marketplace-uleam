package models

type User struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Level      int    `json:"level"`
	Reputation int    `json:"reputation"`
}

type Review struct {
	ID         int    `json:"id"`
	ReviewerID int    `json:"reviewer_id"`
	ReviewedID int    `json:"reviewed_id"`
	Rating     int    `json:"rating"`
	Comment    string `json:"comment"`
}

type Badge struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	RequiredRep int    `json:"required_rep"`
}
