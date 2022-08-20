package runner

import (
	"errors"

	httpserver "github.com/TadayoshiOtsuka/go-tady/internal/http_server"
	"github.com/TadayoshiOtsuka/go-tady/internal/sandbox"
	"github.com/TadayoshiOtsuka/go-tady/pkg/config"
	"github.com/manifoldco/promptui"
)

func Create() int {
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
