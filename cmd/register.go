package cmd

import (
	"github.com/TadayoshiOtsuka/go-tady/internal/register"
	"github.com/spf13/cobra"
)

var registerCmd = &cobra.Command{
	Use:     "register",
	Short:   "register",
	Long:    `register`,
	Aliases: []string{"r"},
	Run: func(cmd *cobra.Command, args []string) {
		register.Exec()
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
}
