package main

import (
	"embed"

	"github.com/andygeiss/ecs-example/components"
	"github.com/andygeiss/ecs-example/systems"
	ecs "github.com/andygeiss/ecs/core"
)

//go:embed resources/**
var efs embed.FS

func main() {
	width, height := 800, 600
	em := ecs.NewEntityManager()
	em.Add(ecs.NewEntity("player", []ecs.Component{
		components.NewPosition().
			WithX(10).WithY(10),
		components.NewSize().
			WithWidth(128).WithHeight(128),
		components.NewTexture().
			WithPath("resources/logo.png").WithVisible(true),
		components.NewVelocity().
			WithX(100).WithY(100),
	}))
	sm := ecs.NewSystemManager()
	sm.Add(
		systems.NewResourceSystem(em),
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
