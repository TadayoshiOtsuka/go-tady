package generator

import (
	"fmt"
	"time"

	"github.com/TadayoshiOtsuka/go-tady/internal/create/generator/engine"
	"github.com/TadayoshiOtsuka/go-tady/pkg/config"
	"github.com/TadayoshiOtsuka/go-tady/pkg/gomod"
	"github.com/TadayoshiOtsuka/go-tady/pkg/utils"
	"github.com/fatih/color"
)

type IGenerator interface {
	Do(src string, config *config.ProjectConfig) error
}

type Generator struct {
	engine engine.IEngine
}

func NewGenerator(engine engine.IEngine) IGenerator {
	return &Generator{
		engine: engine,
	}
}

func (g *Generator) Do(src string, config *config.ProjectConfig) error {
	now := time.Now()
	projectName := config.Name
	userName := config.UserName
	packageName := fmt.Sprintf("github.com/%v/%v", userName, projectName)
	if err := g.engine.Start(src, projectName, packageName); err != nil {
		return err
	}
	utils.PrintWithElapsed("Project Generate Done.", now)

	if err := gomod.Setup(projectName); err != nil {
		return err
	}
	utils.PrintWithElapsed("Setup go mod Done.", now)
	utils.PrintWithElapsed(color.GreenString("Success🎉"), now)

	return nil
}
