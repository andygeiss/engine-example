package systems

import (
	"github.com/andygeiss/ecs-example/components"
	ecs "github.com/andygeiss/ecs/core"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type movementSystem struct{}

func (a *movementSystem) Process(em ecs.EntityManager) (state int) {
	for _, e := range em.FilterByMask(components.MaskPosition | components.MaskVelocity) {
		position := e.Get(components.MaskPosition).(*components.Position)
		velocity := e.Get(components.MaskVelocity).(*components.Velocity)
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
