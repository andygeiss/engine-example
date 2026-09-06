// Package game defines the world this example runs: the entities it starts
// with, and the systems that move them.
package game

import (
	"github.com/andygeiss/ecs"
	"github.com/andygeiss/engine-example/internal/components"
	"github.com/andygeiss/engine-example/internal/systems"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// playerSize is the edge length of the logo sprite in pixels. The player
// starts in the middle of the window, so half of it is the offset.
const playerSize = 128

// Game is a world and the systems that run it. Build one with New and hand
// both halves to ecs.NewDefaultEngine.
type Game struct {
	Entities ecs.EntityManager
	Systems  ecs.SystemManager
}

// New builds the world for a window of the given size. Nothing here talks to
// the graphics card yet — that waits for the engine's Setup.
func New(width, height int, title string) *Game {
	em := NewEntityManager(width, height)
	sm := ecs.NewSystemManager()
	sm.Add(
		systems.NewResourceSystem(em),
		systems.NewInputSystem(),
		systems.NewMovementSystem(rl.GetFrameTime),
		systems.NewCollisionSystem().WithWidth(width).WithHeight(height),
		systems.NewRenderingSystem().WithWidth(width).WithHeight(height).WithTitle(title),
		systems.NewStateSystem(),
	)
	return &Game{Entities: em, Systems: sm}
}

// NewEntityManager fills a world with the three entities this example has: a
// background, the player, and the "controls" entity that holds which keys are
// down right now.
func NewEntityManager(width, height int) ecs.EntityManager {
	em := ecs.NewEntityManager()
	em.Add(ecs.NewEntity("background", []ecs.Component{
		components.NewPosition().WithX(0).WithY(0),
		components.NewSize().WithWidth(float32(width)).WithHeight(float32(height)),
		components.NewState(),
		components.NewTexture().WithPath("resources/space.png").WithVisible(true),
	}))
	em.Add(ecs.NewEntity("controls", []ecs.Component{
		components.NewState(),
	}))
	em.Add(ecs.NewEntity("player", []ecs.Component{
		components.NewPosition().WithX(float32(width/2 - playerSize/2)).WithY(float32(height/2 - playerSize/2)),
		components.NewSize().WithWidth(playerSize).WithHeight(playerSize),
		components.NewState(),
		components.NewTexture().WithPath("resources/logo.png").WithVisible(true),
		components.NewVelocity().WithX(0).WithY(0),
	}))
	return em
}
