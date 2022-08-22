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
	Do(src string, conf *config.ProjectConfig) error
}

type Generator struct {
	engine engine.IEngine
}

func NewGenerator(engine engine.IEngine) IGenerator {
	return &Generator{
		engine: engine,
	}
}

func (g *Generator) Do(src string, conf *config.ProjectConfig) error {
	now := time.Now()
	pjn := conf.Name
	un := conf.UserName
	pn := fmt.Sprintf("github.com/%v/%v", un, pjn)
	if err := g.engine.Start(src, pjn, pn); err != nil {
		return err
	}
	utils.PrintWithElapsedMilliSec("Project Generate Done.", now)

	if err := gomod.Setup(pjn); err != nil {
		return err
	}
	utils.PrintWithElapsedMilliSec("Setup go mod Done.", now)
	utils.PrintWithElapsedMilliSec(color.GreenString("Success🎉"), now)

	return nil
}
