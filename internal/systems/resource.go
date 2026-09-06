package systems

import (
	"fmt"
	"io/fs"
	"path"

	"github.com/andygeiss/ecs"
	"github.com/andygeiss/engine-example/internal/components"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// ResourceSystem uploads the sprite of every entity that has one to the
// graphics card, and frees them all again at teardown.
type ResourceSystem struct {
	em        ecs.EntityManager
	err       error
	resources fs.FS
}

// Err reports the first load that failed, or nil. Process stops the engine on
// such a failure, because a game with no sprites has nothing to show.
func (a *ResourceSystem) Err() error {
	return a.err
}

func (a *ResourceSystem) Process(em ecs.EntityManager) (state int) {
	for _, e := range em.FilterByMask(components.MaskTexture) {
		texture := e.Get(components.MaskTexture).(*components.Texture)
		if texture.Tex != nil {
			continue
		}
		tex, err := loadTexture(a.resources, texture.Path)
		if err != nil {
			a.err = err
			return ecs.StateEngineStop
		}
		texture.Tex = &tex
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

// loadTexture reads a sprite out of the embedded files and uploads it. The
// image itself lives in main memory only long enough to be copied over.
func loadTexture(resources fs.FS, name string) (rl.Texture2D, error) {
	data, err := fs.ReadFile(resources, name)
	if err != nil {
		return rl.Texture2D{}, fmt.Errorf("read texture %q: %w", name, err)
	}
	img := rl.LoadImageFromMemory(path.Ext(name), data, int32(len(data)))
	if img == nil || img.Data == nil {
		return rl.Texture2D{}, fmt.Errorf("decode texture %q: not a %s image raylib can read", name, path.Ext(name))
	}
	defer rl.UnloadImage(img)
	tex := rl.LoadTextureFromImage(img)
	if tex.ID == 0 {
		return rl.Texture2D{}, fmt.Errorf("upload texture %q: no window, or the graphics card refused it", name)
	}
	return tex, nil
}

// NewResourceSystem creates a ResourceSystem. It keeps the entity manager
// because Teardown has to free the textures and is handed nothing.
func NewResourceSystem(em ecs.EntityManager, resources fs.FS) *ResourceSystem {
	return &ResourceSystem{em: em, resources: resources}
}
