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
	Short:   "create template",
	Long:    "create template",
	RunE: func(cmd *cobra.Command, args []string) error {
		return c.Exec()
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
