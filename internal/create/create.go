package runner

import (
	"embed"
	"errors"
	"fmt"

	"github.com/TadayoshiOtsuka/go-tady/internal/create/generator"
	"github.com/TadayoshiOtsuka/go-tady/internal/create/generator/engine"
	"github.com/TadayoshiOtsuka/go-tady/pkg/config"
	"github.com/manifoldco/promptui"
)

//go:embed all:presets
var presetFS embed.FS

func Start() error {
	if err := inputUserName(); err != nil {
		return err
	}
	if err := inputProjectName(); err != nil {
		return err
	}
	if t := selectTemplateType(); t == "user" {
		if err := selectUserTemplate(); err != nil {
			return err
		}
		e := engine.NewUserPresetEngine()
		g := generator.NewGenerator(e)
		// 設定ファイルから絶対パスを読み込み
		src := fmt.Sprintf("%v%v", "presets/", config.Config.TargetTemplate)
		if err := g.Do(src, config.Config); err != nil {
			return err
		}
	} else {
		if err := selectPresetTemplate(); err != nil {
			return err
		}
		e := engine.NewPresetEngine(&presetFS)
		g := generator.NewGenerator(e)
		src := fmt.Sprintf("%v%v", "presets/", config.Config.TargetTemplate)
		if err := g.Do(src, config.Config); err != nil {
			return err
		}
	}

	return nil
}

func selectTemplateType() string {
	return ""
}

func selectUserTemplate() error {
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

func selectPresetTemplate() error {
	p := promptui.Select{
		Label: "Select a project type",
		Items: []string{
			"Sandbox",
			"HTTP Server",
		},
	}
	_, res, err := p.Run()
	if err != nil {
		return err
	}

	switch res {
	case "Sandbox":
		config.Config.TargetTemplate = "sandbox"

	case "HTTP Server":
		config.Config.TargetTemplate = "http_server/rest/nethttp"
	}

	return nil
}
