package systems

import (
	"fmt"

	ecs "github.com/andygeiss/ecs/core"
	"github.com/andygeiss/engine-example/components"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// renderingSystem ...
type renderingSystem struct {
	err           error
	title         string
	width, height int32
}

func (a *renderingSystem) Error() error {
	return a.err
}

func (a *renderingSystem) Setup() {
	rl.InitWindow(a.width, a.height, a.title)
}

func (a *renderingSystem) Process(em ecs.EntityManager) (state int) {
	// First check if app should stop.
	if rl.WindowShouldClose() {
		return ecs.StateEngineStop
	}
	// Clear the screen when the window is ready.
	if rl.IsWindowReady() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		// Draw the entities first.
		a.renderEntities(em)
		// Then draw the text over it.
		rl.DrawText(fmt.Sprintf("FPS %d", rl.GetFPS()), 10, 10, 20, rl.Red)
		rl.DrawText("ESC to exit", 670, 10, 20, rl.Red)
		rl.DrawText("WASD to move", 330, 200, 20, rl.Yellow)
		rl.EndDrawing()
	}
	return ecs.StateEngineContinue
}

func (a *renderingSystem) Teardown() {
	rl.CloseWindow()
}

func (a *renderingSystem) WithHeight(height int) *renderingSystem {
	a.height = int32(height)
	return a
}

func (a *renderingSystem) WithTitle(title string) *renderingSystem {
	a.title = title
	return a
}

func (a *renderingSystem) WithWidth(width int) *renderingSystem {
	a.width = int32(width)
	return a
}

func (a *renderingSystem) renderEntities(em ecs.EntityManager) {
	for _, e := range em.FilterByMask(components.MaskPosition | components.MaskSize) {
		position := e.Get(components.MaskPosition).(*components.Position)
		size := e.Get(components.MaskSize).(*components.Size)
		texture := e.Get(components.MaskTexture)
		// Draw a bounding box
		rl.DrawRectangleLines(int32(position.X), int32(position.Y), int32(size.Width), int32(size.Height), rl.Red)
		// Draw a texture, if available
		if texture != nil {
			tx := texture.(*components.Texture)
			if !tx.Visible {
				continue
			}
			rl.DrawTexture(*tx.Tex, int32(position.X), int32(position.Y), rl.White)
		}
	}
}

func NewRenderingSystem() *renderingSystem {
	return &renderingSystem{}
}
