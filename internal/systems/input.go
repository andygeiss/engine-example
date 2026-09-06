package systems

import (
	"github.com/andygeiss/ecs"
	"github.com/andygeiss/engine-example/internal/components"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// InputSystem writes which keys are down right now into the state of the
// "controls" entity, so no other system has to talk to the keyboard.
type InputSystem struct {
	keyMap map[int32]uint64
}

func (a *InputSystem) Process(em ecs.EntityManager) (engineState int) {
	e := em.Get("controls")
	state := e.Get(components.MaskState).(*components.State)
	// Handle player input
	for key := range a.keyMap {
		if rl.IsKeyDown(key) {
			state.Set(a.keyMap[key], 0)
		} else {
			state.Remove(a.keyMap[key], 0)
		}
	}
	return ecs.StateEngineContinue
}

func (a *InputSystem) Setup() {}

func (a *InputSystem) Teardown() {}

// NewInputSystem creates an InputSystem that reads WASD.
func NewInputSystem() *InputSystem {
	return &InputSystem{
		keyMap: map[int32]uint64{
			rl.KeyW: components.StateControlsW,
			rl.KeyA: components.StateControlsA,
			rl.KeyS: components.StateControlsS,
			rl.KeyD: components.StateControlsD,
		},
	}
}
