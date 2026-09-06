package systems

import (
	"fmt"

	"github.com/andygeiss/ecs"
	"github.com/andygeiss/engine-example/internal/components"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// RenderingSystem owns the window: it opens one at Setup, draws a frame per
// pass, and closes it at Teardown. It stops the engine when the window closes.
type RenderingSystem struct {
	title         string
	width, height int32
}

func (a *RenderingSystem) Setup() {
	rl.InitWindow(a.width, a.height, a.title)
}

func (a *RenderingSystem) Process(em ecs.EntityManager) (state int) {
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
		controls := em.Get("controls")
		controlsState := controls.Get(components.MaskState).(*components.State)
		player := em.Get("player")
		playerPosition := player.Get(components.MaskPosition).(*components.Position)
		playerState := player.Get(components.MaskState).(*components.State)
		switch {
		case controlsState.HasState(components.StateControlsW):
			rl.DrawText("UP", 70, 510, 20, rl.Red)
		case controlsState.HasState(components.StateControlsA):
			rl.DrawText("LEFT", 10, 540, 20, rl.Red)
		case controlsState.HasState(components.StateControlsS):
			rl.DrawText("DOWN", 55, 570, 20, rl.Red)
		case controlsState.HasState(components.StateControlsD):
			rl.DrawText("RIGHT", 100, 540, 20, rl.Red)
		}
		switch {
		case playerState.HasState(components.StatePlayerMove):
			rl.DrawText("MOVING", int32(playerPosition.X), int32(playerPosition.Y)-20, 20, rl.Red)
		case playerState.HasState(components.StatePlayerIdle):
			rl.DrawText("IDLE", int32(playerPosition.X), int32(playerPosition.Y)-20, 20, rl.Red)
		}
		rl.DrawText(fmt.Sprintf("FPS %d", rl.GetFPS()), 10, 10, 20, rl.Red)
		rl.DrawText("ESC to exit", 670, 10, 20, rl.Red)
		rl.DrawText("WASD to move", 330, 10, 20, rl.Yellow)
		rl.EndDrawing()
	}
	return ecs.StateEngineContinue
}

func (a *RenderingSystem) Teardown() {
	rl.CloseWindow()
}

func (a *RenderingSystem) WithHeight(height int) *RenderingSystem {
	a.height = int32(height)
	return a
}

func (a *RenderingSystem) WithTitle(title string) *RenderingSystem {
	a.title = title
	return a
}

func (a *RenderingSystem) WithWidth(width int) *RenderingSystem {
	a.width = int32(width)
	return a
}

func (a *RenderingSystem) renderEntities(em ecs.EntityManager) {
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
			rl.DrawTextureEx(*tx.Tex, rl.Vector2{X: position.X, Y: position.Y}, 0, 1.0, rl.White)
		}
	}
}

func NewRenderingSystem() *RenderingSystem {
	return &RenderingSystem{}
}
