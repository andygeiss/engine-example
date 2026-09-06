package systems

import (
	"fmt"

	"github.com/andygeiss/ecs"
	"github.com/andygeiss/engine-example/internal/components"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// hudSize is the font size of the text drawn over the game, in pixels.
const hudSize = 20

// RenderingSystem owns the window: it opens one at Setup, draws a frame per
// pass, and closes it at Teardown. It stops the engine when the window closes.
type RenderingSystem struct {
	title         string
	verbose       bool
	width, height int32
}

func (a *RenderingSystem) Setup() {
	// raylib writes its log to stdout and gives no way to redirect it, so
	// the log stays off unless the reader asks for it with -v. stdout
	// belongs to the data.
	level := rl.LogWarning
	if a.verbose {
		level = rl.LogInfo
	}
	rl.SetTraceLogLevel(level)
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
		a.renderHUD(em)
		rl.EndDrawing()
	}
	return ecs.StateEngineContinue
}

func (a *RenderingSystem) Teardown() {
	rl.CloseWindow()
}

// WithHeight sets the height of the window in pixels.
func (a *RenderingSystem) WithHeight(height int) *RenderingSystem {
	a.height = int32(height)
	return a
}

// WithTitle sets the text in the window's title bar.
func (a *RenderingSystem) WithTitle(title string) *RenderingSystem {
	a.title = title
	return a
}

// WithVerbose turns raylib's own log on. It goes to stdout, which is why it
// is off by default — see the README's waivers.
func (a *RenderingSystem) WithVerbose(verbose bool) *RenderingSystem {
	a.verbose = verbose
	return a
}

// WithWidth sets the width of the window in pixels.
func (a *RenderingSystem) WithWidth(width int) *RenderingSystem {
	a.width = int32(width)
	return a
}

// renderHUD draws the text over the game. Every position is measured from an
// edge, so the window can be any size the -width and -height flags allow.
func (a *RenderingSystem) renderHUD(em ecs.EntityManager) {
	controlsState := em.Get("controls").Get(components.MaskState).(*components.State)
	player := em.Get("player")
	playerPosition := player.Get(components.MaskPosition).(*components.Position)
	playerState := player.Get(components.MaskState).(*components.State)
	switch {
	case controlsState.HasState(components.StateControlsW):
		rl.DrawText("UP", 70, a.height-90, hudSize, rl.Red)
	case controlsState.HasState(components.StateControlsA):
		rl.DrawText("LEFT", 10, a.height-60, hudSize, rl.Red)
	case controlsState.HasState(components.StateControlsS):
		rl.DrawText("DOWN", 55, a.height-30, hudSize, rl.Red)
	case controlsState.HasState(components.StateControlsD):
		rl.DrawText("RIGHT", 100, a.height-60, hudSize, rl.Red)
	}
	switch {
	case playerState.HasState(components.StatePlayerMove):
		rl.DrawText("MOVING", int32(playerPosition.X), int32(playerPosition.Y)-hudSize, hudSize, rl.Red)
	case playerState.HasState(components.StatePlayerIdle):
		rl.DrawText("IDLE", int32(playerPosition.X), int32(playerPosition.Y)-hudSize, hudSize, rl.Red)
	}
	fps := fmt.Sprintf("FPS %d", rl.GetFPS())
	rl.DrawText(fps, 10, 10, hudSize, rl.Red)
	esc := "ESC to exit"
	rl.DrawText(esc, a.width-rl.MeasureText(esc, hudSize)-10, 10, hudSize, rl.Red)
	wasd := "WASD to move"
	rl.DrawText(wasd, (a.width-rl.MeasureText(wasd, hudSize))/2, 10, hudSize, rl.Yellow)
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

// NewRenderingSystem creates a RenderingSystem.
func NewRenderingSystem() *RenderingSystem {
	return &RenderingSystem{}
}
