package runner

import (
	"errors"
	"fmt"

	httpserver "github.com/TadayoshiOtsuka/go-tady/internal/create/http_server"
	"github.com/TadayoshiOtsuka/go-tady/internal/generator"
	"github.com/TadayoshiOtsuka/go-tady/pkg/config"
	"github.com/manifoldco/promptui"
)

func Start() error {
	if err := inputUserName(); err != nil {
		return err
	}
	if err := inputProjectName(); err != nil {
		return err
	}
	if err := selectTemplate(); err != nil {
		return err
	}

	return nil
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

func inputUserName() error {
	p := promptui.Prompt{
		Label: "Your Github user name",
		Validate: func(in string) error {
			if len(in) == 0 {
				return errors.New("user name is must be not empty")
			}
			return nil
		},
	}

	res, err := p.Run()
	if err != nil {
		return err
	}

	config.Config.UserName = res

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
		config.Config.TargetTemplate = "sandbox"
		if err := generator.Create(); err != nil {
			fmt.Println(err)
			return err
		}

	case "http-server":
		config.Config.TargetTemplate = "http-server"
		httpserver.SelectServerTemplate()
	}

	return nil
}
