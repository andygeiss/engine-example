package systems_test

import (
	"testing"

	"github.com/andygeiss/engine-example/internal/systems"
)

func TestCollisionSystemWrapsAtTheEdges(t *testing.T) {
	tests := []struct {
		name  string
		x, y  float32
		wantX float32
		wantY float32
	}{
		{"inside the window nothing moves", 400, 300, 400, 300},
		{"past the right edge wraps to the left", 801, 300, 0, 300},
		{"past the bottom edge wraps to the top", 400, 601, 400, 0},
		{"exactly on the edge still counts as inside", 800, 600, 800, 600},
		// Only the right and bottom edges wrap today. A player walking left
		// or up keeps going and is gone; see the README's next steps.
		{"past the left edge keeps going", -50, 300, -50, 300},
		{"past the top edge keeps going", 400, -50, 400, -50},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			em := newWorld()
			p := position(em, "player")
			p.X, p.Y = tt.x, tt.y

			systems.NewCollisionSystem().WithWidth(800).WithHeight(600).Process(em)

			if p.X != tt.wantX || p.Y != tt.wantY {
				t.Errorf("position = (%v, %v), want (%v, %v)", p.X, p.Y, tt.wantX, tt.wantY)
			}
		})
	}
}

// The system only looks at things that can move, so scenery stays put even
// when it sits outside the window.
func TestCollisionSystemLeavesEntitiesWithoutVelocityAlone(t *testing.T) {
	t.Parallel()
	em := newWorld()
	em.Remove(em.Get("player"))
	em.Add(newScenery(9000, 9000))

	systems.NewCollisionSystem().WithWidth(800).WithHeight(600).Process(em)

	if got := position(em, "scenery"); got.X != 9000 || got.Y != 9000 {
		t.Errorf("position = (%v, %v), want (9000, 9000)", got.X, got.Y)
	}
}
