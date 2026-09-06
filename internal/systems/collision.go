package systems

import (
	"github.com/andygeiss/ecs"
	"github.com/andygeiss/engine-example/internal/components"
)

// CollisionSystem keeps the sprites inside the window by wrapping them around
// its right and bottom edge.
type CollisionSystem struct {
	width, height int32
}

func (a *CollisionSystem) Process(em ecs.EntityManager) (state int) {
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

func (a *CollisionSystem) Setup() {}

func (a *CollisionSystem) Teardown() {}

// WithHeight sets the window height the system wraps at.
func (a *CollisionSystem) WithHeight(height int) *CollisionSystem {
	a.height = int32(height)
	return a
}

// WithWidth sets the window width the system wraps at.
func (a *CollisionSystem) WithWidth(width int) *CollisionSystem {
	a.width = int32(width)
	return a
}

// NewCollisionSystem creates a CollisionSystem.
func NewCollisionSystem() *CollisionSystem {
	return &CollisionSystem{}
}
