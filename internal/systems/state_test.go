package systems_test

import (
	"testing"
	"testing/synctest"
	"time"

	"github.com/andygeiss/engine-example/internal/components"
	"github.com/andygeiss/engine-example/internal/systems"
)

// The player is "moving" while any direction key is down, and "idle" again
// once they are all released. Each answer takes two passes: one to ask for the
// change, one for Tick to apply it.
func TestStateSystemMarksThePlayerMovingAndIdle(t *testing.T) {
	tests := []struct {
		name string
		keys uint64
		want bool
	}{
		{"no key is idle", components.StateControlsNo, false},
		{"W is moving", components.StateControlsW, true},
		{"A is moving", components.StateControlsA, true},
		{"S is moving", components.StateControlsS, true},
		{"D is moving", components.StateControlsD, true},
		{"two keys is still moving", components.StateControlsW | components.StateControlsD, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				em := newWorld()
				press(em, tt.keys)
				state := systems.NewStateSystem()

				state.Process(em)
				synctest.Sleep(time.Millisecond)
				state.Process(em)

				playerState := em.Get("player").Get(components.MaskState).(*components.State)
				if got := playerState.HasState(components.StatePlayerMove); got != tt.want {
					t.Errorf("moving = %v, want %v", got, tt.want)
				}
			})
		})
	}
}
