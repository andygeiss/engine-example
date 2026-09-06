package systems

import (
	"github.com/andygeiss/ecs"
	"github.com/andygeiss/engine-example/internal/components"
)

// playerSpeed is how fast the player moves, in pixels per second.
const playerSpeed = 100

// MovementSystem turns the keys that are down into a velocity, then moves
// everything that has a position and a velocity.
type MovementSystem struct {
	frameTime func() float32
}

func (a *MovementSystem) Process(em ecs.EntityManager) (state int) {
	controls := em.Get("controls")
	controlsState := controls.Get(components.MaskState).(*components.State)
	// Set the velocity based on the player's input.
	velocity := em.Get("player").Get(components.MaskVelocity).(*components.Velocity)
	if controlsState.Value > 0 {
		if controlsState.HasState(components.StateControlsW) {
			velocity.X = 0
			velocity.Y = -playerSpeed
		}
		if controlsState.HasState(components.StateControlsA) {
			velocity.X = -playerSpeed
			velocity.Y = 0
		}
		if controlsState.HasState(components.StateControlsS) {
			velocity.X = 0
			velocity.Y = playerSpeed
		}
		if controlsState.HasState(components.StateControlsD) {
			velocity.X = playerSpeed
			velocity.Y = 0
		}
	} else {
		velocity.X = 0
		velocity.Y = 0
	}
	// Calculate the next position of the sprites.
	delta := a.frameTime()
	for _, e := range em.FilterByMask(components.MaskPosition | components.MaskVelocity) {
		position := e.Get(components.MaskPosition).(*components.Position)
		velocity := e.Get(components.MaskVelocity).(*components.Velocity)
		position.X += velocity.X * delta
		position.Y += velocity.Y * delta
	}
	return ecs.StateEngineContinue
}

func (a *MovementSystem) Setup() {}

func (a *MovementSystem) Teardown() {}

// NewMovementSystem creates a MovementSystem. frameTime returns how many
// seconds the last frame took: the game passes rl.GetFrameTime, a test passes
// a fixed step so the result does not depend on how fast the machine is.
func NewMovementSystem(frameTime func() float32) *MovementSystem {
	return &MovementSystem{frameTime: frameTime}
}
