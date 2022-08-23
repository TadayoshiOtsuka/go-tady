package cmd

import (
	c "github.com/TadayoshiOtsuka/go-tady/internal/create"
	"github.com/TadayoshiOtsuka/go-tady/pkg/errs"
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
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := viper.ReadInConfig()
		if err != nil {
			return errs.ErrInitializeNotComplete
		}
		return c.Exec(args)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
