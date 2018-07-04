package mysql

import (
	"time"

	"github.com/puper/qman/consumer/core"
)

type Storage struct {
}

func (this *Storage) WatchSubscriptionChange(cb core.SubscriptionChangeCallback) {
	tk := time.NewTicker(time.Second)
	for range tk.C {
		cb(&core.Event{
			Type: core.EVENT_CREATE,
			Data: core.SubscriptionConfig{
				Name:  "testName",
				Topic: "testTopic",
				Tag:   "testTag",
			},
		})
	}
}
