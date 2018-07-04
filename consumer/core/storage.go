package core

import "fmt"

type EventType int

const (
	EVENT_CREATE EventType = 0
	EVENT_UPDATE EventType = 1
	EVENT_DELETE EventType = 2
)

var (
	storages = make(map[string]NewStorageFunc)
)

type NewStorageFunc func(interface{}) (Storage, error)

func RegisterStorage(name string, f NewStorageFunc) {
	storages[name] = f
}

type StorageConfig struct {
	Name   string      `json:"name"`
	Config interface{} `json:"config"`
}

func NewStorage(config *StorageConfig) (Storage, error) {
	f, ok := storages[config.Name]
	if ok {
		return f(config.Config)
	}
	return nil, fmt.Errorf("storage `%v` not found")
}

type Event struct {
	Type EventType
	Data SubscriptionConfig
}

type SubscriptionChangeCallback func(*Event)

type Storage interface {
	WatchSubscriptionChange(SubscriptionChangeCallback)
}
