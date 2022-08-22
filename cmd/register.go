package cmd

import (
	"errors"

	"github.com/TadayoshiOtsuka/go-tady/internal/register"
	"github.com/spf13/cobra"
)

var registerCmd = &cobra.Command{
	Use:     "register",
	Aliases: []string{"r"},
	Short:   "Register the current directory to your Preset",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("please enter preset name")
		}
		return register.Exec(args[0])
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
}
