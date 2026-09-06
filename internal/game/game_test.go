package game_test

import (
	"testing"

	"github.com/andygeiss/engine-example/internal/components"
	"github.com/andygeiss/engine-example/internal/game"
)

func TestNewEntityManagerBuildsTheWorld(t *testing.T) {
	t.Parallel()
	em := game.NewEntityManager(800, 600)

	tests := []struct {
		id   string
		want uint64
	}{
		{"background", components.MaskPosition | components.MaskSize | components.MaskState | components.MaskTexture},
		{"controls", components.MaskState},
		{"player", components.MaskPosition | components.MaskSize | components.MaskState | components.MaskTexture | components.MaskVelocity},
	}
	for _, tt := range tests {
		t.Run(tt.id, func(t *testing.T) {
			e := em.Get(tt.id)
			if e == nil {
				t.Fatalf("no entity %q", tt.id)
			}
			if got := e.Mask(); got != tt.want {
				t.Errorf("mask = %d, want %d", got, tt.want)
			}
		})
	}
	if got := len(em.Entities()); got != len(tests) {
		t.Errorf("got %d entities, want %d", got, len(tests))
	}
}

// The player starts in the middle whatever size the window is, so -width and
// -height do not drop it off the screen.
func TestNewEntityManagerCentresThePlayer(t *testing.T) {
	tests := []struct {
		name          string
		width, height int
		wantX, wantY  float32
	}{
		{"the default window", 800, 600, 336, 236},
		{"a wide window", 1920, 1080, 896, 476},
		{"a window smaller than the player", 100, 100, -14, -14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			em := game.NewEntityManager(tt.width, tt.height)
			got := em.Get("player").Get(components.MaskPosition).(*components.Position)
			if got.X != tt.wantX || got.Y != tt.wantY {
				t.Errorf("position = (%v, %v), want (%v, %v)", got.X, got.Y, tt.wantX, tt.wantY)
			}
		})
	}
}

// The background fills the window it was built for.
func TestNewEntityManagerSizesTheBackgroundToTheWindow(t *testing.T) {
	t.Parallel()
	em := game.NewEntityManager(1024, 768)
	got := em.Get("background").Get(components.MaskSize).(*components.Size)
	if got.Width != 1024 || got.Height != 768 {
		t.Errorf("size = (%v, %v), want (1024, 768)", got.Width, got.Height)
	}
}

// New wires the world and the systems without opening a window, so the whole
// thing can be built in a test. Nothing reaches the graphics card until the
// engine's Setup runs.
func TestNewBuildsAGameWithoutAWindow(t *testing.T) {
	t.Parallel()
	g := game.New(800, 600, "Example Engine", false)

	if got := len(g.Entities.Entities()); got != 3 {
		t.Errorf("got %d entities, want 3", got)
	}
	if got := len(g.Systems.Systems()); got != 6 {
		t.Errorf("got %d systems, want 6", got)
	}
	if err := g.Err(); err != nil {
		t.Errorf("Err() = %v, want nil before anything ran", err)
	}
}
