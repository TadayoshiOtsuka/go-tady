package cmd

import (
	"github.com/TadayoshiOtsuka/go-tady/internal/register"
	"github.com/TadayoshiOtsuka/go-tady/pkg/errs"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var registerCmd = &cobra.Command{
	Use:          "register",
	Aliases:      []string{"r"},
	Short:        "Register the current directory to your Preset",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := viper.ReadInConfig()
		if err != nil {
			return errs.ErrInitializeNotComplete
		}

		return register.Exec(args)
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
}
