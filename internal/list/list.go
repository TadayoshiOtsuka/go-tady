package list

import (
	"fmt"

	"github.com/spf13/viper"
)

func Exec() error {
	showPresets()

	return nil
}

func showPresets() {
	presets := viper.GetStringMapString("presets")
	fmt.Printf("%v preset found\n", len(presets))
	for v := range presets {
		fmt.Print(v)
	}
}
