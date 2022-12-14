package cmd

import (
	"github.com/TadayoshiOtsuka/go-tady/internal/list"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "Show your presets",
	Long:    `Show your presets`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return list.Exec()
	},
	SilenceUsage: true,
}

func init() {
	rootCmd.AddCommand(listCmd)
}
