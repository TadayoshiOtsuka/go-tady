/*
Copyright © 2022 NAME HERE <ohtukayoshi.yoshi@gmail.com>

*/
package cmd

import (
	"errors"
	"os"

	"github.com/TadayoshiOtsuka/go-tady/runner/config"
	httpserver "github.com/TadayoshiOtsuka/go-tady/runner/http_server"
	"github.com/TadayoshiOtsuka/go-tady/runner/sandbox"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create template",
	Long:  "create template create template",
	Run: func(cmd *cobra.Command, args []string) {
		os.Exit(create())
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}

func create() int {
	inputProjectName()
	selectTemplate()

	return 0
}

func inputProjectName() error {
	p := promptui.Prompt{
		Label: "Project name",
		Validate: func(in string) error {
			if len(in) == 0 {
				return errors.New("project name is must be not empty")
			}
			return nil
		},
	}

	res, err := p.Run()
	if err != nil {
		return err
	}

	config.Config.Name = res

	return nil
}

func selectTemplate() error {
	p := promptui.Select{
		Label: "Select a project type",
		Items: []string{
			"sandbox",
			"http-server",
		},
	}
	_, res, err := p.Run()
	if err != nil {
		return err
	}

	switch res {
	case "sandbox":
		if err := sandbox.Create(); err != nil {
			return err
		}

	case "http-server":
		httpserver.SelectServerTemplate()
	}

	return nil
}
