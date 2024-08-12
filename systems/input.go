package systems

import (
	ecs "github.com/andygeiss/ecs/core"
	"github.com/andygeiss/engine-example/components"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type inputSystem struct {
	keyMap map[int32]uint64
}

func (a *inputSystem) Process(em ecs.EntityManager) (state int) {
	controls := em.Get("controls")
	controlsState := controls.Get(components.MaskState).(*components.State)
	// Handle player input
	a.handleKeyDown(87, controlsState) // W
	a.handleKeyDown(65, controlsState) // A
	a.handleKeyDown(83, controlsState) // S
	a.handleKeyDown(68, controlsState) // D
	return ecs.StateEngineContinue
}

func (a *inputSystem) Setup() {}

func (a *inputSystem) Teardown() {}

func (a *inputSystem) handleKeyDown(key int32, state *components.State) {
	if rl.IsKeyDown(key) {
		state.Set(a.keyMap[key])
	} else {
		state.Remove(a.keyMap[key])
	}
}

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
