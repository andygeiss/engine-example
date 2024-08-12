package systems

import (
	ecs "github.com/andygeiss/ecs/core"
	"github.com/andygeiss/engine-example/components"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type movementSystem struct{}

func (a *movementSystem) Process(em ecs.EntityManager) (state int) {
	controls := em.Get("controls")
	controlsState := controls.Get(components.MaskState).(*components.State)
	for _, e := range em.FilterByMask(components.MaskPosition | components.MaskVelocity) {
		position := e.Get(components.MaskPosition).(*components.Position)
		velocity := e.Get(components.MaskVelocity).(*components.Velocity)
		// Set the velocity based on the player's input.
		if controlsState.Value > 0 {
			if controlsState.Value&components.StateControlsW > 0 {
				velocity.X = 0
				velocity.Y = -100
			}
			if controlsState.Value&components.StateControlsA > 0 {
				velocity.X = -100
				velocity.Y = 0
			}
			if controlsState.Value&components.StateControlsS > 0 {
				velocity.X = 0
				velocity.Y = 100
			}
			if controlsState.Value&components.StateControlsD > 0 {
				velocity.X = 100
				velocity.Y = 0
			}
		} else {
			velocity.X = 0
			velocity.Y = 0
		}
		// Calculate the next position of the sprite.
		position.X += velocity.X * rl.GetFrameTime()
		position.Y += velocity.Y * rl.GetFrameTime()
	}
	return ecs.StateEngineContinue
}

func (a *movementSystem) Setup() {}

func (a *movementSystem) Teardown() {}

func NewMovementSystem() ecs.System {
	return &movementSystem{}
}
