package builder

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"io"
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

func (a *Manager) NewZipReader(reader io.Reader, prefix, suffix string) *zip.ReadCloser {
	if a.err != nil {
		return nil
	}
	filename := prefix + suffix + ".zip"
	os.MkdirAll(prefix, 0755)
	file, _ := os.Create(filename)
	io.Copy(file, reader)
	file.Close()
	zr, err := zip.OpenReader(filename)
	if err != nil {
		a.err = err
		return nil
	}
	return zr
}

func (a *Manager) OpenZipReader(path string) *zip.ReadCloser {
	if a.err != nil {
		return nil
	}
	zr, err := zip.OpenReader(path)
	if err != nil {
		a.err = err
		return nil
	}
	return zr
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
			a.err = err
			break
		}
		targetFile := prefix + header.Name
		targetFile = strings.ReplaceAll(targetFile, suffix, "")
		switch header.Typeflag {
		case tar.TypeDir:
			os.Mkdir(targetFile, 0755)
		case tar.TypeReg:
			outFile, err := os.Create(targetFile)
			if err != nil {
				a.err = err
				break
			}
			if _, err := io.Copy(outFile, tarReader); err != nil {
				a.err = err
				break
			}
			outFile.Close()
		}
	}
}

func (a *Manager) UnpackZip(zr *zip.ReadCloser, prefix, suffix string) {
	if a.err != nil {
		return
	}
	for _, file := range zr.File {
		path := file.Name
		parts := strings.Split(path, string(os.PathSeparator))
		path = strings.Join(parts[1:], string(os.PathSeparator))
		if file.FileInfo().IsDir() {
			os.MkdirAll(prefix+path, 0755)
			continue
		}
		src, err := file.Open()
		if err != nil {
			a.err = err
			return
		}
		defer src.Close()
		dst, err := os.Create(prefix + path)
		if _, err := io.Copy(dst, src); err != nil {
			a.err = err
			return
		}
		dst.Close()
	}
	os.Remove(prefix + suffix + ".zip")
}

func NewManager() *Manager {
	return &Manager{}
}
