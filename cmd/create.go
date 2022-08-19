/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"

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
		inputProjectName()
		selectTemplate()
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func inputProjectName() {
	p := promptui.Prompt{
		Label: "project name",
	}

	res, err := p.Run()
	if err != nil {
		log.Panicln(err)
	}

	config.Config.Name = res
}

func selectTemplate() {
	p := promptui.Select{
		Label: "Choose Create Template",
		Items: []string{
			"sandbox",
			"http-server",
		},
	}
	_, res, err := p.Run()
	if err != nil {
		log.Panicln(err)
	}

	switch res {
	case "sandbox":
		sandbox.Create()

	case "http-server":
		httpserver.SelectServerTemplate()
	}
}
