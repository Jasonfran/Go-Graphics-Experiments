package engine

type EventHandler func(EventType, interface{})

type Dispatcher struct {
	handlers map[EventType][]EventHandler
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{handlers: map[EventType][]EventHandler{}}
}

func (d *Dispatcher) Subscribe(t EventType, h EventHandler) {
	handlers, ok := d.handlers[t]
	if !ok {
		d.handlers[t] = []EventHandler{}
	}

	d.handlers[t] = append(handlers, h)
}

func (d *Dispatcher) Trigger(t EventType, data interface{}) {
	handlers, ok := d.handlers[t]
	if !ok {
		return
	}

	for _, handler := range handlers {
		handler(t, data)
	}
}
