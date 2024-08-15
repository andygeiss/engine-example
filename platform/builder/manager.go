package builder

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type Manager struct {
	err error
}

func (a *Manager) Error() error {
	return a.err
}

func (a *Manager) GetArchive(url string) *http.Response {
	res, err := http.Get(url)
	if err != nil {
		a.err = err
		return nil
	}
	return res
}

func (a *Manager) NewGzipReader(reader io.Reader) *gzip.Reader {
	if a.err != nil {
		return nil
	}
	gzr, err := gzip.NewReader(reader)
	if err != nil {
		a.err = err
		return nil
	}
	return gzr
}

func (a *Manager) RemoveUnusedFiles(prefix string) {
	if a.err != nil {
		return
	}
	os.Remove(prefix + "CHANGELOG")
	os.Remove(prefix + "LICENSE")
	os.Remove(prefix + "README.md")
}

func (a *Manager) UnpackGzip(gzr *gzip.Reader, prefix, suffix string) {
	if a.err != nil {
		return
	}
	tarReader := tar.NewReader(gzr)
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		targetFile := prefix + header.Name
		targetFile = strings.ReplaceAll(targetFile, suffix, "")
		switch header.Typeflag {
		case tar.TypeDir:
			os.Mkdir(targetFile, 0755)
		case tar.TypeReg:
			outFile, err := os.Create(targetFile)
			if err != nil {
				log.Fatal(err)
			}
			if _, err := io.Copy(outFile, tarReader); err != nil {
				log.Fatal(err)
			}
			outFile.Close()
		}
	}
}

func NewManager() *Manager {
	return &Manager{}
}
