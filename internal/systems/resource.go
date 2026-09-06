package systems

import (
	"github.com/andygeiss/ecs"
	"github.com/andygeiss/engine-example/internal/components"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// ResourceSystem uploads the sprite of every entity that has one to the
// graphics card, and frees them all again at teardown.
type ResourceSystem struct {
	em ecs.EntityManager
}

func (a *ResourceSystem) Process(em ecs.EntityManager) (state int) {
	for _, e := range em.FilterByMask(components.MaskTexture) {
		texture := e.Get(components.MaskTexture).(*components.Texture)
		if texture.Tex == nil {
			tex := rl.LoadTexture(texture.Path)
			texture.Tex = &tex
		}
	}
	return ecs.StateEngineContinue
}

func (a *ResourceSystem) Setup() {}

func (a *ResourceSystem) Teardown() {
	for _, e := range a.em.FilterByMask(components.MaskTexture) {
		texture := e.Get(components.MaskTexture).(*components.Texture)
		if texture.Tex != nil {
			rl.UnloadTexture(*texture.Tex)
		}
	}
}

// NewResourceSystem creates a ResourceSystem. It keeps the entity manager
// because Teardown has to free the textures and is handed nothing.
func NewResourceSystem(em ecs.EntityManager) *ResourceSystem {
	return &ResourceSystem{em: em}
}
