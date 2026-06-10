package models

type User struct {
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
}
