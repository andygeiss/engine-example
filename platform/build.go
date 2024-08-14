package main

import (
	"fmt"
	"os"

	"github.com/andygeiss/engine-example/platform/builder"
)

func main() {
	version := "5.0"
	extensions := []string{".tar.gz"}
	urlPrefix := "https://github.com/raysan5/raylib/releases/download/" + version + "/"
	osMap := map[string]string{
		"darwin":  "macos",
		"linux":   "linux_amd64",
		"windows": "win64_mingw-w64",
	}
	mgr := builder.NewManager()
	for _, fileExt := range extensions {
		suffix := fmt.Sprintf("raylib-%s_%s", version, osMap[os.Getenv("GOOS")])
		filename := suffix + fileExt
		url := urlPrefix + "/" + filename
		res := mgr.GetArchive(url)
		defer res.Body.Close()
		switch fileExt {
		case ".tar.gz":
			gzr := mgr.NewGzipReader(res.Body)
			mgr.UnpackGzip(gzr, suffix)
		}
		mgr.RemoveUnusedFiles()
	}
}