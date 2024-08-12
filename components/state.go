package components

const (
	StateControlsW = uint64(1 << 0)
	StateControlsA = uint64(1 << 1)
	StateControlsS = uint64(1 << 2)
	StateControlsD = uint64(1 << 3)
)

type State struct {
	Value uint64 `json:"value"`
}

func (a *State) HasState(state uint64) bool {
	return a.Value&state == state
}

func (a *State) Mask() uint64 {
	return MaskState
}

func (a *State) Remove(state uint64) {
	a.Value &= ^state
}

func (a *State) Set(state uint64) {
	a.Value |= state
}

func (a *State) Toggle(state uint64) {
	if a.HasState(state) {
		a.Remove(state)
	} else {
		a.Set(state)
	}
}

func NewState() *State {
	return &State{}
}
