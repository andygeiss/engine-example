package components

type State struct {
	Value int `json:"value"`
}

func (a *State) Mask() uint64 {
	return MaskState
}

func NewState() *State {
	return &State{}
}
