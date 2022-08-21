package gomod

import (
	"os"
	"os/exec"
)

func Setup(projectName string) error {
	if err := os.Chdir(projectName); err != nil {
		return err
	}
	if err := tidy(); err != nil {
		return err
	}

	return nil
}

func tidy() error {
	if err := exec.Command("go", "mod", "tidy").Run(); err != nil {
		return err
	}

	return nil
}
