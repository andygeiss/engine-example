package components_test

import (
	"testing"
	"testing/synctest"
	"time"

	"github.com/andygeiss/engine-example/internal/components"
)

func TestStateSetAndRemove(t *testing.T) {
	tests := []struct {
		name string
		set  []uint64
		drop []uint64
		want uint64
	}{
		{"a fresh state holds nothing", nil, nil, components.StateControlsNo},
		{"one key sets one bit", []uint64{components.StateControlsW}, nil, components.StateControlsW},
		{"two keys set two bits", []uint64{components.StateControlsW, components.StateControlsD}, nil, components.StateControlsW | components.StateControlsD},
		{"setting the same key twice changes nothing", []uint64{components.StateControlsA, components.StateControlsA}, nil, components.StateControlsA},
		{"removing the only key empties the state", []uint64{components.StateControlsS}, []uint64{components.StateControlsS}, components.StateControlsNo},
		{"removing one key leaves the other", []uint64{components.StateControlsW, components.StateControlsD}, []uint64{components.StateControlsW}, components.StateControlsD},
		{"removing a key that was never set changes nothing", []uint64{components.StateControlsW}, []uint64{components.StateControlsA}, components.StateControlsW},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			state := components.NewState()
			for _, s := range tt.set {
				state.Set(s, 0)
			}
			for _, s := range tt.drop {
				state.Remove(s, 0)
			}
			// Next holds what was asked for; Value follows one Tick later.
			if got := state.Next; got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

// Set and Remove only ask for a change. Tick applies it, and not before the
// duration has passed — that is what lets an animation hold a pose.
func TestStateTickWaitsForTheDuration(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		state := components.NewState()
		state.Set(components.StatePlayerMove, 100*time.Millisecond)

		state.Tick()
		if state.HasState(components.StatePlayerMove) {
			t.Error("the state applied straight away, want it to wait 100ms")
		}

		synctest.Sleep(150 * time.Millisecond)
		state.Tick()
		if !state.HasState(components.StatePlayerMove) {
			t.Error("the state never applied, want it applied after 150ms")
		}
	})
}

// HasState asks whether every bit of the argument is set, so a value that
// holds two keys answers yes to each one on its own.
func TestStateHasState(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		state := components.NewState()
		state.Set(components.StateControlsW|components.StateControlsD, 0)
		synctest.Sleep(time.Millisecond)
		state.Tick()

		for _, want := range []uint64{components.StateControlsW, components.StateControlsD} {
			if !state.HasState(want) {
				t.Errorf("HasState(%d) = false, want true", want)
			}
		}
		if state.HasState(components.StateControlsA) {
			t.Error("HasState(A) = true, want false")
		}
	})
}
