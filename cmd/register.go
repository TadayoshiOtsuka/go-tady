package cmd

import (
	"errors"

	"github.com/TadayoshiOtsuka/go-tady/internal/register"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var registerCmd = &cobra.Command{
	Use:     "register",
	Aliases: []string{"r"},
	Short:   "Register the current directory to your Preset",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := viper.ReadInConfig()
		if err != nil {
			return errors.New(color.RedString("can not find config file. please run 'go-tady init'"))
		}
		if len(args) == 0 {
			return errors.New("please enter preset name")
		}

		return register.Exec(args[0])
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
}
