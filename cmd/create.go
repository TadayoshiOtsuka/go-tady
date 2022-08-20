/*
Copyright © 2022 TadayoshiOtsuka <ohtukayoshi.yoshi@gmail.com>

*/
package cmd

import (
	"os"

	runner "github.com/TadayoshiOtsuka/go-tady/runner"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create template",
	Long:  "create template create template",
	Run: func(cmd *cobra.Command, args []string) {
		os.Exit(runner.Create())
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
