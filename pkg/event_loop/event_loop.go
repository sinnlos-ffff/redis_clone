package event_loop

import "net"

// Event represents a simple event with a name and data
type Event struct {
	Name string
	Conn net.Conn
}

// EventHandler is a function that handles an event
type EventHandler func(event Event)

type EventLoop struct {
	events   chan Event
	handlers map[string][]EventHandler
}

func NewEventLoop(bufferSize int) *EventLoop {
	return &EventLoop{
		events:   make(chan Event, bufferSize),
		handlers: make(map[string][]EventHandler),
	}
}

func (el *EventLoop) RegisterHandler(eventName string, handler EventHandler) {
	el.handlers[eventName] = append(el.handlers[eventName], handler)
}

func (el *EventLoop) PostEvent(event Event) {
	el.events <- event
}

func (el *EventLoop) Start() {
	go func() {
		for event := range el.events {
			if handlers, found := el.handlers[event.Name]; found {
				for _, handler := range handlers {
					go handler(event) // Run each handler in its own goroutine
				}
			}
		}
	}()
}

func (el *EventLoop) Stop() {
	close(el.events)
}
