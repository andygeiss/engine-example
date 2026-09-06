package systems

import (
	"github.com/andygeiss/ecs"
	"github.com/andygeiss/engine-example/internal/components"
)

// StateSystem advances every state that is waiting on a timer, and marks the
// player as moving while a direction key is down.
type StateSystem struct{}

func (a *StateSystem) Process(em ecs.EntityManager) (state int) {
	controls := em.Get("controls")
	controlsState := controls.Get(components.MaskState).(*components.State)
	player := em.Get("player")
	playerState := player.Get(components.MaskState).(*components.State)
	for _, e := range em.FilterByMask(components.MaskState) {
		state := e.Get(components.MaskState).(*components.State)
		state.Tick()
	}
	if controlsState.HasState(components.StateControlsW) ||
		controlsState.HasState(components.StateControlsA) ||
		controlsState.HasState(components.StateControlsS) ||
		controlsState.HasState(components.StateControlsD) {
		playerState.Set(components.StatePlayerMove, 0)
	} else {
		playerState.Remove(components.StatePlayerMove, 0)
	}
	return ecs.StateEngineContinue
}

func (a *StateSystem) Setup() {}

func (a *StateSystem) Teardown() {}

// NewStateSystem creates a StateSystem.
func NewStateSystem() *StateSystem {
	return &StateSystem{}
}
