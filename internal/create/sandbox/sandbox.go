package sandbox

import (
	"fmt"
	"time"

	"github.com/TadayoshiOtsuka/go-tady/pkg/config"
	"github.com/TadayoshiOtsuka/go-tady/pkg/generator"
	"github.com/TadayoshiOtsuka/go-tady/pkg/gomod"
	"github.com/fatih/color"
)

func Create(target string) error {
	now := time.Now()
	src := fmt.Sprintf("%v%v", "templates", target)
	projectName := config.Config.Name
	userName := config.Config.UserName
	packageName := fmt.Sprintf("github.com/%v/%v", userName, projectName)

	if err := generator.Do(src, projectName); err != nil {
		return err
	}
	if err := gomod.Setup(projectName, packageName); err != nil {
		return err
	}
	fmt.Printf("%s in %d[ms]\n", color.GreenString("SUCCESS"), time.Since(now).Milliseconds())

	return nil
}
