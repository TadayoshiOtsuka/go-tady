package register

import (
	"fmt"
	"os"

	"github.com/TadayoshiOtsuka/go-tady/pkg/errs"
	"github.com/fatih/color"
	"github.com/spf13/viper"
)

func Exec(args []string) error {
	if len(args) == 0 {
		return errs.ErrNotEnoughRegisterArgs
	}

	pn := args[0]
	if err := registerNewPreset(pn); err != nil {
		return err
	}
	fmt.Print(color.GreenString("Success\n"))

	return nil
}

func registerNewPreset(name string) error {
	d, err := os.Getwd()
	if err != nil {
		return err
	}
	presets := viper.GetStringMapString("presets")
	presets[name] = d
	viper.Set("presets", presets)
	if err := viper.WriteConfig(); err != nil {
		return err
	}

	return nil
}
