package errs

import (
	"errors"
	"fmt"
)

var (
	ErrInitializeNotComplete = errors.New("can not find config file. please run 'go-tady init'")
	ErrEmptyProjectName      = errors.New("should be enter new project name. e.g.) 'go-tady create myApp'")
	ErrNotEnoughRegisterArgs = errors.New("should be enter preset name. e.g.) 'go-tady register presetA'")
	ErrEmptyUserName         = errors.New("user name is must be not empty")
	ErrInvalidCreateType     = errors.New("invalid create type")
	ErrDirIsAlreadyExists    = func(dir string) error {
		return fmt.Errorf("'%v' is already exists. should be use unique name in current directory", dir)
	}
)
