package list

import (
	"fmt"

	"github.com/TadayoshiOtsuka/go-tady/pkg/errs"
	"github.com/spf13/viper"
)

func Exec() error {
	return showPresets()
}

func showPresets() error {
	err := viper.ReadInConfig()
	if err != nil {
		return errs.ErrInitializeNotComplete
	}
	presets := viper.GetStringMapString("presets")
	fmt.Printf("%v preset found\n", len(presets))
	for k, v := range presets {
		fmt.Printf("%v  at: %v\n", k, v)
	}

	return nil
}
