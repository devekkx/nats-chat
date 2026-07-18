package events

import "time"

type ChatMessageSendEvent struct {
	UserID  string `json:"userId"`
	RoomID  string `json:"roomId"`
	Content string `json:"content"`
}

type ChatMessageCreatedEvent struct {
	ID        string    `json:"id"`
	UserID    string    `json:"userId"`
	RoomID    string    `json:"roomId"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}
