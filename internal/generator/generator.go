package generator

import (
	"fmt"
	"time"

	"github.com/TadayoshiOtsuka/go-tady/pkg/config"
	"github.com/TadayoshiOtsuka/go-tady/pkg/generator"
	"github.com/TadayoshiOtsuka/go-tady/pkg/gomod"
	"github.com/TadayoshiOtsuka/go-tady/pkg/printutils"
	"github.com/fatih/color"
)

func Create() error {
	now := time.Now()
	src := fmt.Sprintf("%v%v", "templates/", config.Config.TargetTemplate)
	projectName := config.Config.Name
	userName := config.Config.UserName
	packageName := fmt.Sprintf("github.com/%v/%v", userName, projectName)

	if err := generator.Do(src, projectName, packageName); err != nil {
		return err
	}
	printutils.PrintWithElapsed("Generate Template Done.", now)

	if err := gomod.Setup(projectName); err != nil {
		return err
	}
	printutils.PrintWithElapsed("go mod setup Done.", now)
	printutils.PrintWithElapsed(color.GreenString("SUCCESS🎉"), now)

	return nil
}
