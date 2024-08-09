package main

import (
	"github.com/andygeiss/ecs-example/components"
	"github.com/andygeiss/ecs-example/systems"
	ecs "github.com/andygeiss/ecs/core"
)

func main() {
	width, height := 800, 600
	em := ecs.NewEntityManager()
	em.Add(ecs.NewEntity("player", []ecs.Component{
		components.NewPosition().
			WithX(10).WithY(10),
		components.NewSize().
			WithWidth(10).WithHeight(10),
		components.NewVelocity().
			WithX(1).WithY(1),
	}))
	sm := ecs.NewSystemManager()
	sm.Add(
		systems.NewMovementSystem(),
		systems.NewCollisionSystem().
			WithWidth(width).WithHeight(height),
		systems.NewRenderingSystem().
			WithWidth(width).WithHeight(height),
	)
	de := ecs.NewDefaultEngine(em, sm)
	de.Setup()
	defer de.Teardown()
	de.Run()
}
