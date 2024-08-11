package main

import (
	"embed"

	ecs "github.com/andygeiss/ecs/core"
	"github.com/andygeiss/engine-example/components"
	"github.com/andygeiss/engine-example/systems"
)

//go:embed resources/**
var efs embed.FS

func main() {
	width, height := 800, 600
	em := ecs.NewEntityManager()
	em.Add(ecs.NewEntity("player", []ecs.Component{
		components.NewPosition().
			WithX(800/2 - 64).WithY(600/2 - 64),
		components.NewSize().
			WithWidth(128).WithHeight(128),
		components.NewTexture().
			WithPath("resources/logo.png").WithVisible(true),
		components.NewVelocity().
			WithX(0).WithY(0),
	}))
	sm := ecs.NewSystemManager()
	sm.Add(
		systems.NewResourceSystem(em),
		systems.NewInputSystem(),
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
