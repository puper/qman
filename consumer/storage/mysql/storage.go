package mysql

import (
    "github.com/puper/qman/consumer/core"
)

type Storage struct {
}

fun (this *Storage) WatchSubscriptionChange(cb core.SubscriptionChangeCallback) {
    cb(&core.Event{
        Type: core.EVENT_CREATE,
        OldDatas: []byte(""),
        NewData: []byte(""),
    })
}