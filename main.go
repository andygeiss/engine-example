// Command engine-example is a small game that shows how to build an engine
// with the ecs package: entities carry components, systems do the work.
//
// It opens a window. WASD moves the logo, ESC closes it.
package main

import (
	"context"

	"github.com/andygeiss/ecs"
	"github.com/andygeiss/engine-example/internal/game"
)

func main() {
	g := game.New(800, 600, "Example Engine")
	engine := ecs.NewDefaultEngine(g.Entities, g.Systems)
	engine.Setup()
	defer engine.Teardown()
	engine.Run(context.Background())
}
