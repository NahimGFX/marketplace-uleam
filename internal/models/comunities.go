package models

// ============================
// MESSAGE
// ============================
type Message struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	SenderID   int    `json:"sender_id"`
	ReceiverID int    `json:"receiver_id"`
	Content    string `json:"content"`

	// Relaciones Belongs-To
	Sender   User `json:"-" gorm:"foreignKey:SenderID"`
	Receiver User `json:"-" gorm:"foreignKey:ReceiverID"`
}

// ============================
// MISSION
// ============================
type Mission struct {
	ID            int    `json:"id" gorm:"primaryKey"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	RequiredLevel int    `json:"required_level"`
	RewardPoints  int    `json:"reward_points"`

	// Relación Has-Many
	UserMissions []UserMission `json:"-" gorm:"foreignKey:MissionID"`
}

// ============================
// USER MISSION
// ============================
type UserMission struct {
	ID        int  `json:"id" gorm:"primaryKey"`
	UserID    int  `json:"user_id"`
	MissionID int  `json:"mission_id"`
	Completed bool `json:"completed"`

	// Relaciones Belongs-To
	User    User    `json:"-" gorm:"foreignKey:UserID"`
	Mission Mission `json:"-" gorm:"foreignKey:MissionID"`
}
