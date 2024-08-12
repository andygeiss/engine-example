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

func (a *State) Mask() uint64 {
	return MaskState
}

func NewState() *State {
	return &State{}
}
