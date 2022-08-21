package engine

import (
	"os"
	"path/filepath"
	"strings"
	"syscall"

	"embed"
)

type PresetEngine struct {
	FS *embed.FS
}

func NewPresetEngine(templateFS *embed.FS) IEngine {
	return &PresetEngine{
		FS: templateFS,
	}
}

func (g *PresetEngine) Start(src, rootName, packageName string) error {
	umask := syscall.Umask(0)
	defer syscall.Umask(umask)
	if err := g.makeRoot(rootName); err != nil {
		return err
	}
	if err := g.scan(src, rootName, packageName); err != nil {
		return err
	}

	return nil
}

func (g *PresetEngine) makeRoot(name string) error {
	if err := os.MkdirAll(name, 0777); err != nil {
		return err
	}

	return nil
}

func (g *PresetEngine) scan(src, dst, packageName string) error {
	fs, err := g.FS.ReadDir(src)
	if err != nil {
		return err
	}

	for _, f := range fs {
		if f.IsDir() {
			if err := g.genDir(src, dst, packageName, f.Name()); err != nil {
				return err
			}
		} else {
			if err := g.genFile(src, dst, packageName, f.Name()); err != nil {
				return err
			}
		}
	}

	return nil
}

func (g *PresetEngine) genDir(src, dst, packageName, path string) error {
	dirSrc, dirDst := filepath.Join(src, path), filepath.Join(dst, path)
	if err := os.MkdirAll(dirDst, 0777); err != nil {
		return err
	}
	if err := g.scan(dirSrc, dirDst, packageName); err != nil {
		return err
	}

	return nil
}

func (g *PresetEngine) genFile(src, dst, packageName, name string) error {
	fs, fd := filepath.Join(src, name), filepath.Join(dst, name)
	file, err := g.FS.ReadFile(fs)
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

	file = g.replacePackageName(file, packageName)
	if _, err = f.Write(file); err != nil {
		return err
	}

	return nil
}

func (g *PresetEngine) replacePackageName(file []byte, packageName string) []byte {
	c := string(file)
	c = strings.ReplaceAll(c, "GO_TADY_PACKAGE_NAME", packageName)
	return []byte(c)
}
