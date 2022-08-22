package setup

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func Exec() error {
	return genSettingFile()
}

func genSettingFile() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	viper.Set("presets", map[string]string{})
	if err := viper.WriteConfigAs(fmt.Sprintf("%v/.go-tady.toml", home)); err == nil {
		viper.WatchConfig()
	}

	return nil
}
