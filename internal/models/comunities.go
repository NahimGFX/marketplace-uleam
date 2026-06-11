package models

type Message struct {
	ID         int
	SenderID   int
	ReceiverID int
	Content    string
}

type Mission struct {
	ID            int
	Title         string
	Description   string
	RequiredLevel int
	RewardPoints  int
}

type UserMission struct {
	ID        int
	UserID    int
	MissionID int
	Completed bool
}
