package cmd

import (
	"github.com/TadayoshiOtsuka/go-tady/internal/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:          "version",
	Aliases:      []string{"v"},
	Short:        "Show go-tady version info.",
	SilenceUsage: true,
	Run: func(cmd *cobra.Command, args []string) {
		version.Exec()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
