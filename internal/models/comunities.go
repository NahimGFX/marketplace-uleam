package models

// MESSAGE
type Message struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	SenderID   int    `json:"sender_id"`
	ReceiverID int    `json:"receiver_id"`
	Content    string `json:"content"`
}

// MISSION
type Mission struct {
	ID            int    `json:"id" gorm:"primaryKey"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	RequiredLevel int    `json:"required_level"`
	RewardPoints  int    `json:"reward_points"`
}

// USER MISSION
type UserMission struct {
	ID        int  `json:"id" gorm:"primaryKey"`
	UserID    int  `json:"user_id"`
	MissionID int  `json:"mission_id"`
	Completed bool `json:"completed"`
}
