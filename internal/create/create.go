package runner

import (
	"fmt"
	"os"

	"github.com/TadayoshiOtsuka/go-tady/internal/create/generator"
	"github.com/TadayoshiOtsuka/go-tady/internal/create/generator/engine"
	"github.com/TadayoshiOtsuka/go-tady/pkg/config"
	"github.com/TadayoshiOtsuka/go-tady/pkg/errs"
	"github.com/manifoldco/promptui"
	"github.com/spf13/viper"
)

type CreateType string

const (
	userPreset    = CreateType("preset")
	defaultPreset = CreateType("default")
)

func Exec(args []string) error {
	if err := setProjectName(args); err != nil {
		return err
	}
	t, err := selectCreateType()
	if err != nil {
		return err
	}
	switch t {
	case userPreset:
		return genFromUserPreset()

	case defaultPreset:
		return genFromPreset()

	default:
		return errs.ErrInvalidCreateType
	}
}

func genFromUserPreset() error {
	if err := selectUserTemplate(viper.GetStringMapString("presets")); err != nil {
		return err
	}
	e := engine.NewUserPresetEngine()
	g := generator.NewGenerator(e)
	src := config.Config.TargetPreset
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
			string(defaultPreset),
		},
	}
	_, res, err := p.Run()
	if err != nil {
		return "", err
	}

	switch res {
	case string(userPreset):
		return userPreset, nil
	case string(defaultPreset):
		return defaultPreset, nil
	default:
		return "", nil
	}
}

func selectUserTemplate(presets map[string]string) error {
	var items []string
	for k := range presets {
		items = append(items, k)
	}
	p := promptui.Select{
		Label: "select a your preset",
		Items: items,
	}
	_, res, err := p.Run()
	if err != nil {
		return err
	}

	for k, v := range presets {
		if k == res {
			config.Config.TargetPreset = v
		}
	}

	return nil
}

func setProjectName(args []string) error {
	if len(args) == 0 {
		return errs.ErrEmptyProjectName
	}

	pn := args[0]
	if isDirExists(pn) {
		return errs.ErrDirIsAlreadyExists(pn)
	}

	config.Config.Name = pn

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

func isDirExists(projectName string) bool {
	fs, err := os.Stat(projectName)
	if err != nil {
		return false
	}
	if !fs.IsDir() {
		return false
	}

	return true
}
