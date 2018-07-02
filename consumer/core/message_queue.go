package core

import (
	"container/list"
	"sync"
)

const (
	STATE_NOT_STATED = 0
	STATE_STARTED    = 1
	STATE_STOPPED    = 2
)

type MessageQueue struct {
	mutex          sync.Mutex
	data           *list.List
	state          int
	processHandler func()
	stopHandler    func()
}

func NewMessageQueue(ph func(msg *Message), sh func()) *MessageQueue {
	return &Queue{
		data:           list.New(),
		state:          STATE_NOT_STATED,
		processHandler: ph,
		stopHandler:    sh,
	}
}

func (this *MessageQueue) loop() {
	for {
		this.mutex.Lock()
		e := this.data.Back()
		if e == nil {
			this.state = STATE_STOPPED
			if this.stopHandler != nil {
				this.stopHandler()
			}
			this.mutex.Unlock()
			return
		}
		this.mutex.Unlock()
		msg := e.Value.(*Message)
		this.processHandler(msg)
		this.mutex.Lock()
		this.data.Remove(e)
		this.mutex.Unlock()
	}
}

func (this *MessageQueue) Put(msg *Message) bool {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	if this.state == STATE_NOT_STATED {
		this.data.PushFront(msg)
		go this.loop()
		return true
	} else if this.state == STATE_STARTED {
		this.data.PushFront(msg)
		return true
	}
	return false
}
