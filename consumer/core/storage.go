package core

type EventType int

const (
	EVENT_CREATE EventType = 0
	EVENT_UPDATE EventType = 1
	EVENT_DELETE EventType = 2
)

type Event struct {
	Type EventType
	Old  *SubscriptionConfig
	New  *SubscriptionConfig
}

type SubscriptionChangeCallback func(*Event)

type Storage interface {
	WatchSubscriptionChange(SubscriptionChangeCallback)
	Start() error
	Stop() error
}
