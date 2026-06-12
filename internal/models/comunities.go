package models

type Message struct {
	ID         int    `json:"id"`
	SenderID   int    `json:"sender_id"`
	ReceiverID int    `json:"receiver_id"`
	Content    string `json:"content"`
}

type Mission struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	RequiredLevel int    `json:"required_level"`
	RewardPoints  int    `json:"reward_points"`
}

type UserMission struct {
	ID        int  `json:"id"`
	UserID    int  `json:"user_id"`
	MissionID int  `json:"mission_id"`
	Completed bool `json:"completed"`
}
