package engine

type EventType uint32
type EventHandler func(EventType, interface{})

type EventDispatcher struct {
	handlers map[EventType][]EventHandler
}

func NewDispatcher() *EventDispatcher {
	return &EventDispatcher{handlers: map[EventType][]EventHandler{}}
}

func (d *EventDispatcher) Subscribe(t EventType, h EventHandler) {
	handlers, ok := d.handlers[t]
	if !ok {
		d.handlers[t] = []EventHandler{}
	}

	d.handlers[t] = append(handlers, h)
}

func (d *EventDispatcher) Trigger(t EventType, data interface{}) {
	handlers, ok := d.handlers[t]
	if !ok {
		return
	}

	for _, handler := range handlers {
		handler(t, data)
	}
}
