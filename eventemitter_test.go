package eventemitter_test

import (
	"testing"

	"github.com/SakizciAdam/eventemitter"
)

func TestEventBus(t *testing.T) {
	bus := eventemitter.New()

	bus.AddListener("test", func(args ...interface{}) {
		if len(args) != 1 {
			t.Errorf("Expected 1 argument, got %d", len(args))
		}

		if args[0] != "hello" {
			t.Errorf("Expected 'test', got '%s'", args[0])
		}
	})

	bus.EmitSync("test", "hello")
}
