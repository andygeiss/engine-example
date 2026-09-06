package systems_test

import (
	"testing"

	"github.com/andygeiss/engine-example/internal/components"
	"github.com/andygeiss/engine-example/internal/systems"
)

// fixedStep is one tenth of a second per frame, so a speed of 100 pixels per
// second moves exactly 10 pixels and the arithmetic stays readable.
func fixedStep() float32 { return 0.1 }

func TestMovementSystemSetsVelocityFromTheKeys(t *testing.T) {
	tests := []struct {
		name  string
		keys  uint64
		wantX float32
		wantY float32
	}{
		{"no key stands still", components.StateControlsNo, 0, 0},
		{"W moves up", components.StateControlsW, 0, -100},
		{"A moves left", components.StateControlsA, -100, 0},
		{"S moves down", components.StateControlsS, 0, 100},
		{"D moves right", components.StateControlsD, 100, 0},
		// The system checks the keys in W, A, S, D order and each check
		// overwrites the one before, so the last one wins outright.
		{"W and D together give D alone, not a diagonal", components.StateControlsW | components.StateControlsD, 100, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			em := newWorld()
			press(em, tt.keys)

			systems.NewMovementSystem(fixedStep).Process(em)

			got := velocity(em, "player")
			if got.X != tt.wantX || got.Y != tt.wantY {
				t.Errorf("velocity = (%v, %v), want (%v, %v)", got.X, got.Y, tt.wantX, tt.wantY)
			}
		})
	}
}

// The position moves by velocity × frame time, so the same key held for the
// same number of seconds moves the same distance on any machine.
func TestMovementSystemMovesByVelocityTimesFrameTime(t *testing.T) {
	t.Parallel()
	em := newWorld()
	press(em, components.StateControlsD)
	movement := systems.NewMovementSystem(fixedStep)

	for range 3 {
		movement.Process(em)
	}

	got := position(em, "player")
	if want := float32(130); got.X != want {
		t.Errorf("x = %v, want %v (100 + 3 frames × 100px/s × 0.1s)", got.X, want)
	}
	if want := float32(100); got.Y != want {
		t.Errorf("y = %v, want %v — moving right must not change y", got.Y, want)
	}
}

// Letting go of every key stops the player rather than letting it drift.
func TestMovementSystemStopsWhenTheKeysAreReleased(t *testing.T) {
	t.Parallel()
	em := newWorld()
	movement := systems.NewMovementSystem(fixedStep)

	press(em, components.StateControlsD)
	movement.Process(em)
	press(em, components.StateControlsNo)
	movement.Process(em)

	if got := velocity(em, "player"); got.X != 0 || got.Y != 0 {
		t.Errorf("velocity = (%v, %v), want (0, 0)", got.X, got.Y)
	}
}
