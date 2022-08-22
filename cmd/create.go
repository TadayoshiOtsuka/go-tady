/*
Copyright © 2022 TadayoshiOtsuka <ohtukayoshi.yoshi@gmail.com>

*/
package cmd

import (
	c "github.com/TadayoshiOtsuka/go-tady/internal/create"
	"github.com/spf13/cobra"
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
		return c.Exec()
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
