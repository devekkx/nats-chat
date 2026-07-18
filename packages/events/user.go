package events

type UserConnectedEvent struct {
	UserID string `json:"userId"`
}

type UserDisconnectedEvent struct {
	UserID string `json:"userId"`
}
