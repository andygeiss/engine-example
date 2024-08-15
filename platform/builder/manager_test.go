package builder_test

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/andygeiss/engine-example/platform/builder"
)

type sut struct {
	filename  string
	prefix    string
	suffix    string
	url       string
	urlPrefix string
	version   string
}

func setup(os, fileExt string) *sut {
	version := "5.0"
	urlPrefix := "https://github.com/raysan5/raylib/releases/download/" + version
	prefix := "./testdata/"
	suffix := fmt.Sprintf("raylib-%s_%s", version, os)
	filename := suffix + fileExt
	url := urlPrefix + "/" + filename
	return &sut{
		filename:  filename,
		prefix:    prefix,
		suffix:    suffix,
		url:       url,
		urlPrefix: urlPrefix,
		version:   version,
	}
}

func TestManagerGzip(t *testing.T) {
	sut := setup("macos", ".tar.gz")
	// Act
	mgr := builder.NewManager()
	res := mgr.GetArchive(sut.url)
	defer res.Body.Close()
	gzr := mgr.NewGzipReader(res.Body)
	mgr.UnpackGzip(gzr, sut.prefix, sut.suffix)
	mgr.RemoveUnusedFiles(sut.prefix)
	// Assert
	if err := mgr.Error(); err != nil {
		t.Error(err)
	}
	if _, err := os.Stat(sut.prefix + "include/raylib.h"); errors.Is(err, os.ErrNotExist) {
		t.Error("File should exists")
	}
	if _, err := os.Stat(sut.prefix + "lib/libraylib.a"); errors.Is(err, os.ErrNotExist) {
		t.Error("File should exists")
	}
	if _, err := os.Stat(sut.prefix + "CHANGELOG"); !errors.Is(err, os.ErrNotExist) {
		t.Error("File should not exists")
	}
	if _, err := os.Stat(sut.prefix + "LICENSE"); !errors.Is(err, os.ErrNotExist) {
		t.Error("File should not exists")
	}
	if _, err := os.Stat(sut.prefix + "README.md"); !errors.Is(err, os.ErrNotExist) {
		t.Error("File should not exists")
	}
}

func TestManagerZip(t *testing.T) {
	sut := setup("win64_mingw-w64", ".zip")
	// Act
	mgr := builder.NewManager()
	res := mgr.GetArchive(sut.url)
	defer res.Body.Close()
	zr := mgr.NewZipReader(res.Body, sut.prefix, sut.suffix)
	mgr.UnpackZip(zr, sut.prefix, sut.suffix)
	mgr.RemoveUnusedFiles(sut.prefix)
	// Assert
	if err := mgr.Error(); err != nil {
		t.Error(err)
	}
	if _, err := os.Stat(sut.prefix + "include/raylib.h"); errors.Is(err, os.ErrNotExist) {
		t.Error("File should exists")
	}
	if _, err := os.Stat(sut.prefix + "lib/libraylib.a"); errors.Is(err, os.ErrNotExist) {
		t.Error("File should exists")
	}
	if _, err := os.Stat(sut.prefix + "CHANGELOG"); !errors.Is(err, os.ErrNotExist) {
		t.Error("File should not exists")
	}
	if _, err := os.Stat(sut.prefix + "LICENSE"); !errors.Is(err, os.ErrNotExist) {
		t.Error("File should not exists")
	}
	if _, err := os.Stat(sut.prefix + "README.md"); !errors.Is(err, os.ErrNotExist) {
		t.Error("File should not exists")
	}
}
