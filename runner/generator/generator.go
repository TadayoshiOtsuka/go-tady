package generator

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func MakeRoot(name string) error {
	if err := os.MkdirAll(name, 0700); err != nil {
		return err
	}

	return nil
}

func Scan(src, dst string) error {
	fs, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}

	for _, f := range fs {
		if f.IsDir() {
			if err := os.MkdirAll(filepath.Join(dst, f.Name()), 0700); err != nil {
				return err
			}
			subDirSrc, subDirDst := filepath.Join(src, f.Name()), filepath.Join(dst, f.Name())
			if err := Scan(subDirSrc, subDirDst); err != nil {
				return err
			}
		} else {
			fileSrc := filepath.Join(src, f.Name())
			buf, err := ioutil.ReadFile(fileSrc)
			if err != nil {
				return err
			}

			fileDst := filepath.Join(dst, f.Name())
			file, err := os.Create(fileDst)
			if err != nil {
				return err
			}
			defer file.Close()

			if _, err = file.Write(buf); err != nil {
				return err
			}
		}
	}

	return nil
}
