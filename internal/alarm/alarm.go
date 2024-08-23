package alarm

import (
	"fmt"
	"github.com/kafka-embracetheday/goResourceWatcher/internal/logger"
)

// Observer
type Observer interface {
	Notify(message string)
}

// Subject
type Subject interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers(message string)
}

// Alarm
type Alarm struct {
	observers map[Observer]struct{}
}

// NewAlarm
func NewAlarm() *Alarm {
	return &Alarm{
		observers: make(map[Observer]struct{}),
	}
}

func (a *Alarm) RegisterObservers(observers ...Observer) {
	for _, observer := range observers {
		a.observers[observer] = struct{}{}
	}
}

func (a *Alarm) RemoveObserver(observer Observer) {
	delete(a.observers, observer)
}

func (a *Alarm) NotifyObservers(message string) {
	for observer := range a.observers {
		observer.Notify(message)
	}
}

// EmailNotifier
type EmailNotifier struct {
	Recipient string
}

func (e *EmailNotifier) Notify(message string) {
	fmt.Printf("发送电子邮件到 %s: %s\n", e.Recipient, message)
}

// LogNotifier
type LogNotifier struct{}

func (l *LogNotifier) Notify(message string) {
	logger.Logger.Errorf("日志告警: %s\n", message)
}
