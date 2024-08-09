package systems

import (
	"github.com/andygeiss/ecs-example/components"
	ecs "github.com/andygeiss/ecs/core"
)

type collisionSystem struct {
	width, height int32
}

func (a *collisionSystem) Process(em ecs.EntityManager) (state int) {
	for _, e := range em.FilterByMask(components.MaskPosition | components.MaskVelocity) {
		position := e.Get(components.MaskPosition).(*components.Position)
		if position.X > float32(a.width) {
			position.X = 0
		}
		if position.Y > float32(a.height) {
			position.Y = 0
		}
	}
	return ecs.StateEngineContinue
}

func (a *collisionSystem) Setup() {}

func (a *collisionSystem) Teardown() {}

func (a *collisionSystem) WithHeight(height int) *collisionSystem {
	a.height = int32(height)
	return a
}

func (a *collisionSystem) WithWidth(width int) *collisionSystem {
	a.width = int32(width)
	return a
}
func NewCollisionSystem() *collisionSystem {
	return &collisionSystem{}
}
