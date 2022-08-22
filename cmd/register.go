package cmd

import (
	"github.com/TadayoshiOtsuka/go-tady/internal/register"
	"github.com/spf13/cobra"
)

var registerCmd = &cobra.Command{
	Use:     "register",
	Aliases: []string{"r"},
	Short:   "Register the current directory to your Preset",
	Run: func(cmd *cobra.Command, args []string) {
		register.Exec()
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
}
