package components

import (
	"time"
)

const (
	StateControlsNo = uint64(0)
	StateControlsW  = uint64(1 << 0)
	StateControlsA  = uint64(1 << 1)
	StateControlsS  = uint64(1 << 2)
	StateControlsD  = uint64(1 << 3)
)

const (
	StatePlayerIdle = uint64(0)
	StatePlayerMove = uint64(1 << 0)
)

type State struct {
	Duration time.Duration `json:"duration"`
	Next     uint64        `json:"next"`
	Start    time.Time     `json:"start"`
	Value    uint64        `json:"value"`
}

func (a *State) HasState(state uint64) bool {
	return a.Value&state == state
}

func (a *State) Mask() uint64 {
	return MaskState
}

func (a *State) Remove(state uint64, duration time.Duration) {
	a.Duration = duration
	a.Next &= ^state
	a.Start = time.Now()
}

func (a *State) Set(state uint64, duration time.Duration) {
	a.Duration = duration
	a.Next |= state
	a.Start = time.Now()
}

func (a *State) Tick() {
	if a.Value == a.Next {
		return
	}
	if a.Start.Add(a.Duration).Before(time.Now()) {
		a.Value = a.Next
	}
}

func NewState() *State {
	return &State{}
}
