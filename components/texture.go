package components

import rl "github.com/andygeiss/engine-example/platform/raylib"

type Texture struct {
	Path    string `json:"path"`
	Visible bool   `json:"visible"`
	Tex     *rl.Texture
}

func (a *Texture) Mask() uint64 {
	return MaskTexture
}

func (a *Texture) WithPath(path string) *Texture {
	a.Path = path
	return a
}

func (a *Texture) WithVisible(visible bool) *Texture {
	a.Visible = visible
	return a
}

func NewTexture() *Texture {
	return &Texture{}
}
