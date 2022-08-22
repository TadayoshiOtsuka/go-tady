/*
Copyright © 2022 TadayoshiOtsuka <ohtukayoshi.yoshi@gmail.com>

*/
package cmd

import (
	"errors"

	c "github.com/TadayoshiOtsuka/go-tady/internal/create"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var createCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"c"},
	Short:   "Create project from your Preset or default Preset",
	Long: `Create project from your Preset or default Preset.
From your preset:
 => Create a project by selecting from the your go-tady configurations. The configuration file is .go-tady.toml.
From default preset:
 => Create a project by selecting from the preset project under https://github.com/TadayoshiOtsuka/go-tady/tree/master/assets/presets
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := viper.ReadInConfig()
		if err != nil {
			return errors.New(color.RedString("can not find config file. please run 'go-tady init'"))
		}
		return c.Exec()
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
