package systems

import (
	ecs "github.com/andygeiss/ecs/core"
	"github.com/andygeiss/engine-example/components"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type inputSystem struct{}

func (a *inputSystem) Process(em ecs.EntityManager) (state int) {
	// Handle player input
	e := em.Get("player")
	velocity := e.Get(components.MaskVelocity).(*components.Velocity)
	// Handle player input
	if key := rl.GetKeyPressed(); key != 0 {
		switch key {
		case 87: // W
			velocity.X, velocity.Y = 0, -100
		case 65: // A
			velocity.X, velocity.Y = -100, 0
		case 83: // S
			velocity.X, velocity.Y = 0, 100
		case 68: // D
			velocity.X, velocity.Y = 100, 0
		}
	}
	return ecs.StateEngineContinue
}

func (a *inputSystem) Setup() {}

func (a *inputSystem) Teardown() {}

func NewInputSystem() ecs.System {
	return &inputSystem{}
}
