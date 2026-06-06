package types

type EventType string

const (
	EventUserCreated EventType = "user.created"
	EventUserDeleted EventType = "user.deleted"
)

type Event struct {
	Type    EventType `json:"type"`
	Payload any       `json:"payload"`
}
