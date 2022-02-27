package eventemitter

type EventBus struct {
	Listeners map[string][]func(...interface{})

	emit func(name string, async bool, args ...interface{})

	EmitSync func(name string, args ...interface{})

	EmitAsync func(name string, args ...interface{})

	AddListener func(name string, listener func(...interface{}))
}

func New() *EventBus {
	bus := EventBus{}

	bus.Listeners = make(map[string][]func(...interface{}))

	bus.emit = func(name string, async bool, args ...interface{}) {
		for key, val := range bus.Listeners {
			if key == name {
				for _, listener := range val {
					if async {
						go listener(args...)
					} else {
						listener(args...)
					}
				}
			}
		}
	}

	bus.EmitSync = func(name string, args ...interface{}) {
		bus.emit(name, false, args...)
	}

	bus.EmitAsync = func(name string, args ...interface{}) {
		bus.emit(name, true, args...)
	}

	bus.AddListener = func(name string, listener func(...interface{})) {
		bus.Listeners[name] = append(bus.Listeners[name], listener)
	}

	return &bus
}
