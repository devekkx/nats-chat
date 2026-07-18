package events

type PresenceUpdatedEvent struct {
	UserID string `json:"userId"`
	Status string `json:"status"`
}
