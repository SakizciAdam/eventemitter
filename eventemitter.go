package eventemitter

type EventBus struct {
	Listeners map[string][]func(...interface{})

	Emit func(name string, args ...interface{})

	AddListener func(name string, listener func(...interface{}))
}

func New() *EventBus {
	bus := EventBus{}

	bus.Listeners = make(map[string][]func(...interface{}))

	bus.Emit = func(name string, args ...interface{}) {
		for key, val := range bus.Listeners {
			if key == name {
				for _, listener := range val {
					listener(args...)
				}
			}
		}
	}

	bus.AddListener = func(name string, listener func(...interface{})) {
		bus.Listeners[name] = append(bus.Listeners[name], listener)
	}

	return &bus
}
