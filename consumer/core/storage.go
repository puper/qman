package core

type EventType int

const (
	EVENT_CREATE EventType = 0
	EVENT_UPDATE EventType = 1
	EVENT_DELETE EventType = 2
)

type Event struct {
	Type    EventType
	OldData []byte
	NewData []byte
}

type SubscriptionChangeCallback func(*Event)

type Storage interface {
	WatchSubscriptionChange(SubscriptionChangeCallback)
	Start() error
	Stop()
}
