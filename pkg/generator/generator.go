package generator

import (
	"os"
	"path/filepath"
	"strings"
	"syscall"

	"embed"
)

//go:embed templates
var templateFS embed.FS

func Do(src, rootName, packageName string) error {
	umask := syscall.Umask(0)
	defer syscall.Umask(umask)
	if err := makeRoot(rootName); err != nil {
		return err
	}
	if err := scan(src, rootName, packageName); err != nil {
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

func scan(src, dst, packageName string) error {
	fs, err := templateFS.ReadDir(src)
	if err != nil {
		return err
	}

	for _, f := range fs {
		if f.IsDir() {
			if err := genDir(src, dst, packageName, f.Name()); err != nil {
				return err
			}
		} else {
			if err := genFile(src, dst, packageName, f.Name()); err != nil {
				return err
			}
		}
	}

	return nil
}

func genDir(src, dst, packageName, path string) error {
	dirSrc, dirDst := filepath.Join(src, path), filepath.Join(dst, path)
	if err := os.MkdirAll(dirDst, 0777); err != nil {
		return err
	}
	if err := scan(dirSrc, dirDst, packageName); err != nil {
		return err
	}

	return nil
}

func genFile(src, dst, packageName, name string) error {
	fs, fd := filepath.Join(src, name), filepath.Join(dst, name)
	file, err := templateFS.ReadFile(fs)
	if err != nil {
		return err
	}
	if name == "go.mod.td" {
		fd = filepath.Join(dst, "go.mod")
	}

	f, err := os.Create(fd)
	if err != nil {
		return err
	}
	defer f.Close()

	file = replacePackageName(file, packageName)
	if _, err = f.Write(file); err != nil {
		return err
	}

	return nil
}

func replacePackageName(file []byte, packageName string) []byte {
	c := string(file)
	c = strings.ReplaceAll(c, "GO_TADY_PACKAGE_NAME", packageName)
	return []byte(c)
}
