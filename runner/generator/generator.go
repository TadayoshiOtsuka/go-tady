package generator

import (
	"os"
	"path/filepath"
)

func MakeDir(name string) error {
	dp := filepath.Join(name)
	if err := os.Mkdir(dp, 0700); err != nil {
		return err
	}

	return nil
}

func copyFile() {}
