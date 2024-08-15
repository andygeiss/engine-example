package builder_test

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/andygeiss/engine-example/platform/builder"
)

func TestManager(t *testing.T) {
	// Arrange
	version := "5.0"
	urlPrefix := "https://github.com/raysan5/raylib/releases/download/" + version
	prefix := "./testdata/"
	suffix := fmt.Sprintf("raylib-%s_%s", version, "macos")
	filename := suffix + ".tar.gz"
	url := urlPrefix + "/" + filename

	// Act
	mgr := builder.NewManager()
	res := mgr.GetArchive(url)
	defer res.Body.Close()
	gzr := mgr.NewGzipReader(res.Body)
	mgr.UnpackGzip(gzr, prefix, suffix)
	mgr.RemoveUnusedFiles(prefix)

	// Assert
	if err := mgr.Error(); err != nil {
		t.Error(err)
	}
	if _, err := os.Stat(prefix + "include/raylib.h"); errors.Is(err, os.ErrNotExist) {
		t.Error("File should exists")
	}
	if _, err := os.Stat(prefix + "lib/libraylib.a"); errors.Is(err, os.ErrNotExist) {
		t.Error("File should exists")
	}
	if _, err := os.Stat(prefix + "CHANGELOG"); !errors.Is(err, os.ErrNotExist) {
		t.Error("File should not exists")
	}
	if _, err := os.Stat(prefix + "LICENSE"); !errors.Is(err, os.ErrNotExist) {
		t.Error("File should not exists")
	}
	if _, err := os.Stat(prefix + "README.md"); !errors.Is(err, os.ErrNotExist) {
		t.Error("File should not exists")
	}
}
