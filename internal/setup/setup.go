package setup

import (
	"errors"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/viper"
)

var (
	ErrEmptyUserName = errors.New("user name is must be not empty")
)

func Exec() error {
	return genSettingFile()
}

func genSettingFile() error {
	un, err := inputUserName()
	if err != nil {
		return err
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	viper.Set("username", un)
	viper.Set("presets", map[string]string{})
	if err := viper.WriteConfigAs(fmt.Sprintf("%v/.go-tady.toml", home)); err != nil {
		return err
	}
	fmt.Print(color.GreenString("Success\n"))

	return nil
}

func inputUserName() (string, error) {
	p := promptui.Prompt{
		Label: "your github user name",
		Validate: func(in string) error {
			if len(in) == 0 {
				return ErrEmptyUserName
			}
			return nil
		},
	}
	res, err := p.Run()
	if err != nil {
		return "", err
	}

	return res, nil
}
