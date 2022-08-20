package generator

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"
)

func Do(src, rootName string) error {
	fmt.Println("start gen")
	defaultUmask := syscall.Umask(0)
	defer func() {
		syscall.Umask(defaultUmask)
	}()
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
	fs, err := ioutil.ReadDir(src)
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
	if err := os.MkdirAll(filepath.Join(dst, path), 0777); err != nil {
		return err
	}
	subDirSrc, subDirDst := filepath.Join(src, path), filepath.Join(dst, path)
	if err := scan(subDirSrc, subDirDst); err != nil {
		return err
	}

	return nil
}

func genFile(src, dst, path string) error {
	fileSrc := filepath.Join(src, path)
	fileDst := filepath.Join(dst, path)
	buf, err := ioutil.ReadFile(fileSrc)
	if err != nil {
		return err
	}

	file, err := os.Create(fileDst)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err = file.Write(buf); err != nil {
		return err
	}

	return nil
}
