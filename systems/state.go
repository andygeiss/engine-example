package systems

import (
	ecs "github.com/andygeiss/ecs/core"
	"github.com/andygeiss/engine-example/components"
)

type stateSystem struct{}

func (a *stateSystem) Process(em ecs.EntityManager) (state int) {
	for _, e := range em.FilterByMask(components.MaskState) {
		// Read the state of an entity.
		_ = e.Get(components.MaskState).(*components.State)
		// TODO: Modify the state.
	}
	return ecs.StateEngineContinue
}

func (a *stateSystem) Setup() {}

func (a *stateSystem) Teardown() {}

func NewStateSystem() ecs.System {
	return &stateSystem{}
}
