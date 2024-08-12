package systems

import (
	ecs "github.com/andygeiss/ecs/core"
	"github.com/andygeiss/engine-example/components"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type inputSystem struct {
	keyMap map[int32]uint64
}

func (a *inputSystem) Process(em ecs.EntityManager) (engineState int) {
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

func (a *inputSystem) Setup() {}

func (a *inputSystem) Teardown() {}

func NewInputSystem() ecs.System {
	return &inputSystem{
		keyMap: map[int32]uint64{
			87: components.StateControlsW,
			65: components.StateControlsA,
			83: components.StateControlsS,
			68: components.StateControlsD,
		},
	}
}
