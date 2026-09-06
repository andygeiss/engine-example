package game

import "embed"

// resourcesFS holds the sprites the game draws. They are embedded so the
// binary runs from any directory — `go install` puts it on the PATH, far away
// from this repository.
//
// The paths keep the directory, so a component names its sprite as
// "resources/logo.png". That is what //go:embed stores.
//
//go:embed resources
var resourcesFS embed.FS
