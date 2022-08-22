package engine

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/TadayoshiOtsuka/go-tady/pkg/config"
)

type UserPresetEngine struct{}

func NewUserPresetEngine() IEngine {
	return &UserPresetEngine{}
}

func (g *UserPresetEngine) Start(src, rootName, packageName string) error {
	umask := syscall.Umask(0)
	defer syscall.Umask(umask)
	if err := g.makeRoot(rootName); err != nil {
		return err
	}
	if err := g.readOldPackageName(src); err != nil {
		return err
	}
	if err := g.scan(src, rootName, packageName); err != nil {
		return err
	}

	return nil
}

func (g *UserPresetEngine) makeRoot(name string) error {
	if err := os.MkdirAll(name, 0777); err != nil {
		return err
	}

	return nil
}

func (g *UserPresetEngine) readOldPackageName(src string) error {
	fp, err := os.Open(fmt.Sprintf("%v/go.mod", src))
	if err != nil {
		return err
	}
	defer fp.Close()

	s := bufio.NewScanner(fp)
	s.Scan()
	c := s.Text()
	if !strings.Contains(c, "module") {
		return errors.New("cannot find module definition")
	}
	config.Config.OldPackageName = c[7:]

	return nil
}

func (g *UserPresetEngine) scan(src, dst, packageName string) error {
	fs, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, f := range fs {
		if f.IsDir() {
			if f.Name() == ".git" {
				continue
			}
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

func (g *UserPresetEngine) genDir(src, dst, packageName, path string) error {
	dirSrc, dirDst := filepath.Join(src, path), filepath.Join(dst, path)
	if err := os.MkdirAll(dirDst, 0777); err != nil {
		return err
	}
	if err := g.scan(dirSrc, dirDst, packageName); err != nil {
		return err
	}

	return nil
}

func (g *UserPresetEngine) genFile(src, dst, packageName, name string) error {
	fs, fd := filepath.Join(src, name), filepath.Join(dst, name)
	file, err := os.ReadFile(fs)
	if err != nil {
		return err
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

func (g *UserPresetEngine) replacePackageName(file []byte, packageName string) []byte {
	c := string(file)
	c = strings.ReplaceAll(c, config.Config.OldPackageName, packageName)
	return []byte(c)
}
