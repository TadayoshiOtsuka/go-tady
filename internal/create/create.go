package runner

import (
	"errors"
	"fmt"

	"github.com/TadayoshiOtsuka/go-tady/internal/create/generator"
	"github.com/TadayoshiOtsuka/go-tady/internal/create/generator/engine"
	"github.com/TadayoshiOtsuka/go-tady/pkg/config"
	"github.com/manifoldco/promptui"
)

var (
	ErrEmptyProjectName  = errors.New("project name is must be not empty")
	ErrEmptyUserName     = errors.New("user name is must be not empty")
	ErrInvalidCreateType = errors.New("invalid create type")
)

type CreateType string

const (
	userPreset = CreateType("userPreset")
	preset     = CreateType("preset")
)

func Exec() error {
	if err := inputUserName(); err != nil {
		return err
	}
	if err := inputProjectName(); err != nil {
		return err
	}
	t, err := selectCreateType()
	if err != nil {
		return err
	}
	switch t {
	case userPreset:
		return genFromUserPreset()

	case preset:
		return genFromPreset()
	default:
		return ErrInvalidCreateType
	}
}

func genFromUserPreset() error {
	if err := selectUserTemplate(); err != nil {
		return err
	}
	e := engine.NewUserPresetEngine()
	g := generator.NewGenerator(e)
	// 設定ファイルから絶対パスを読み込み
	src := fmt.Sprintf("%v%v", "presets/", config.Config.TargetPreset)
	if err := g.Do(src, config.Config); err != nil {
		return err
	}

	return nil
}

func genFromPreset() error {
	if err := selectPresetTemplate(); err != nil {
		return err
	}
	e := engine.NewPresetEngine()
	g := generator.NewGenerator(e)
	src := fmt.Sprintf("%v%v", "presets/", config.Config.TargetPreset)
	if err := g.Do(src, config.Config); err != nil {
		return err
	}

	return nil
}

func selectCreateType() (CreateType, error) {
	p := promptui.Select{
		Label: "select a create type",
		Items: []string{
			string(userPreset),
			string(preset),
		},
	}
	_, res, err := p.Run()
	if err != nil {
		return "", err
	}

	switch res {
	case string(userPreset):
		return userPreset, nil
	case string(preset):
		return preset, nil
	default:
		return "", nil
	}
}

func selectUserTemplate() error {
	return nil
}

func inputProjectName() error {
	p := promptui.Prompt{
		Label: "project name",
		Validate: func(in string) error {
			if len(in) == 0 {
				return ErrEmptyProjectName
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
		Label: "your github user name",
		Validate: func(in string) error {
			if len(in) == 0 {
				return ErrEmptyUserName
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
		Label: "select a preset",
		Items: []string{
			"Sandbox",
		},
	}
	_, res, err := p.Run()
	if err != nil {
		return err
	}

	switch res {
	case "Sandbox":
		config.Config.TargetPreset = "sandbox"
	}

	return nil
}
