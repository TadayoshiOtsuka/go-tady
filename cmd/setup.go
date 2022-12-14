package cmd

import (
	"github.com/TadayoshiOtsuka/go-tady/internal/setup"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:          "init",
	Short:        "Initialize .go-tady.toml",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return setup.Exec()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
