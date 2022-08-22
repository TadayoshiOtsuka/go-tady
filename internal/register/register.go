package register

import (
	"os"

	"github.com/spf13/viper"
)

func Exec(presetName string) error {
	if err := registerNewPreset(presetName); err != nil {
		return err
	}

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
