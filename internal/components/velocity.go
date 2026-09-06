package components

type Velocity struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

func (a *Velocity) Mask() uint64 {
	return MaskVelocity
}

func (a *Velocity) WithX(x float32) *Velocity {
	a.X = x
	return a
}

func (a *Velocity) WithY(y float32) *Velocity {
	a.Y = y
	return a
}

func NewVelocity() *Velocity {
	return &Velocity{}
}
