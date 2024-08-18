package main

import (
	"github.com/andygeiss/ecs"
	"github.com/andygeiss/engine-example/components"
	"github.com/andygeiss/engine-example/systems"
)

//go:generate go run platform/build.go
func main() {
	width, height := 800, 600
	em := ecs.NewEntityManager()
	em.Add(ecs.NewEntity("background", []ecs.Component{
		components.NewPosition().WithX(0).WithY(0),
		components.NewSize().WithWidth(800).WithHeight(600),
		components.NewState(),
		components.NewTexture().WithPath("resources/space.png").WithVisible(true),
	}))
	em.Add(ecs.NewEntity("controls", []ecs.Component{
		components.NewState(),
	}))
	em.Add(ecs.NewEntity("player", []ecs.Component{
		components.NewPosition().WithX(800/2 - 64).WithY(600/2 - 64),
		components.NewSize().WithWidth(128).WithHeight(128),
		components.NewState(),
		components.NewTexture().WithPath("resources/logo.png").WithVisible(true),
		components.NewVelocity().WithX(0).WithY(0),
	}))
	sm := ecs.NewSystemManager()
	sm.Add(
		systems.NewResourceSystem(em),
		systems.NewInputSystem(),
		systems.NewMovementSystem(),
		systems.NewCollisionSystem().WithWidth(width).WithHeight(height),
		systems.NewRenderingSystem().WithWidth(width).WithHeight(height),
		systems.NewStateSystem(),
	)
	de := ecs.NewDefaultEngine(em, sm)
	de.Setup()
	defer de.Teardown()
	de.Run()
}
