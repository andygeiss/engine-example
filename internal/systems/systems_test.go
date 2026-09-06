package systems_test

import (
	"github.com/andygeiss/ecs"
	"github.com/andygeiss/engine-example/internal/components"
)

// newWorld builds the smallest world the systems under test need: a player
// that can move, and the "controls" entity holding the keys that are down.
func newWorld() ecs.EntityManager {
	em := ecs.NewEntityManager()
	em.Add(ecs.NewEntity("controls", []ecs.Component{
		components.NewState(),
	}))
	em.Add(ecs.NewEntity("player", []ecs.Component{
		components.NewPosition().WithX(100).WithY(100),
		components.NewSize().WithWidth(128).WithHeight(128),
		components.NewState(),
		components.NewVelocity(),
	}))
	return em
}

// press puts keys down without going through the keyboard, the way the input
// system would. The duration is 0 and Tick is skipped, so Value is set here
// directly — a test should not have to wait for a clock to see its own input.
func press(em ecs.EntityManager, keys uint64) {
	state := em.Get("controls").Get(components.MaskState).(*components.State)
	state.Value = keys
	state.Next = keys
}

func position(em ecs.EntityManager, id string) *components.Position {
	return em.Get(id).Get(components.MaskPosition).(*components.Position)
}

func velocity(em ecs.EntityManager, id string) *components.Velocity {
	return em.Get(id).Get(components.MaskVelocity).(*components.Velocity)
}

// newScenery is an entity with a position but no velocity: it can be drawn,
// but no system may move it.
func newScenery(x, y float32) *ecs.Entity {
	return ecs.NewEntity("scenery", []ecs.Component{
		components.NewPosition().WithX(x).WithY(y),
		components.NewSize().WithWidth(64).WithHeight(64),
	})
}
