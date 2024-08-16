package systems

import (
	"github.com/andygeiss/ecs"
	"github.com/andygeiss/engine-example/components"
	rl "github.com/andygeiss/engine-example/platform/raylib"
)

type resourceSystem struct {
	em ecs.EntityManager
}

func (a *resourceSystem) Process(em ecs.EntityManager) (state int) {
	for _, e := range em.FilterByMask(components.MaskTexture) {
		texture := e.Get(components.MaskTexture).(*components.Texture)
		if texture.Tex == nil {
			tex := rl.LoadTexture(texture.Path)
			texture.Tex = &tex
		}
	}
	return ecs.StateEngineContinue
}

func (a *resourceSystem) Setup() {}

func (a *resourceSystem) Teardown() {
	for _, e := range a.em.FilterByMask(components.MaskTexture) {
		texture := e.Get(components.MaskTexture).(*components.Texture)
		if texture.Tex != nil {
			rl.UnloadTexture(*texture.Tex)
		}
	}
}

func NewResourceSystem(em ecs.EntityManager) ecs.System {
	return &resourceSystem{em: em}
}
