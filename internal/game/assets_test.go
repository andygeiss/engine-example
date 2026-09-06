package game

import "testing"

// The sprites are embedded, so the binary draws something wherever it runs.
// The names are the ones NewEntityManager gives its textures.
func TestResourcesAreEmbedded(t *testing.T) {
	t.Parallel()
	for _, name := range []string{"resources/logo.png", "resources/space.png"} {
		data, err := resourcesFS.ReadFile(name)
		if err != nil {
			t.Errorf("read %q: %v", name, err)
			continue
		}
		if len(data) == 0 {
			t.Errorf("%q is empty", name)
		}
	}
}
