package components

type Size struct {
	Width  float32 `json:"width"`
	Height float32 `json:"height"`
}

func (a *Size) Mask() uint64 {
	return MaskSize
}

func (a *Size) WithWidth(width float32) *Size {
	a.Width = width
	return a
}

func (a *Size) WithHeight(height float32) *Size {
	a.Height = height
	return a
}

func NewSize() *Size {
	return &Size{}
}
