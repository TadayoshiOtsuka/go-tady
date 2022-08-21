package generator

import (
	"os"
	"path/filepath"
	"syscall"

	"embed"
)

//go:embed templates
var templateFS embed.FS

func Do(src, rootName string) error {
	umask := syscall.Umask(0)
	defer syscall.Umask(umask)
	if err := makeRoot(rootName); err != nil {
		return err
	}
	if err := scan(src, rootName); err != nil {
		return err
	}

	return nil
}

func makeRoot(name string) error {
	if err := os.MkdirAll(name, 0777); err != nil {
		return err
	}

	return nil
}

func scan(src, dst string) error {
	fs, err := templateFS.ReadDir(src)
	if err != nil {
		return err
	}

	for _, f := range fs {
		if f.IsDir() {
			if err := genDir(src, dst, f.Name()); err != nil {
				return err
			}
		} else {
			if err := genFile(src, dst, f.Name()); err != nil {
				return err
			}
		}
	}

	return nil
}

func genDir(src, dst, path string) error {
	dirSrc, dirDst := filepath.Join(src, path), filepath.Join(dst, path)
	if err := os.MkdirAll(dirDst, 0777); err != nil {
		return err
	}
	if err := scan(dirSrc, dirDst); err != nil {
		return err
	}

	return nil
}

func genFile(src, dst, name string) error {
	fs, fd := filepath.Join(src, name), filepath.Join(dst, name)
	buf, err := templateFS.ReadFile(fs)
	if err != nil {
		return err
	}

	f, err := os.Create(fd)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = f.Write(buf); err != nil {
		return err
	}

	return nil
}
