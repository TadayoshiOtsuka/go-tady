package generator

import (
	"fmt"
	"time"

	"github.com/TadayoshiOtsuka/go-tady/pkg/config"
	"github.com/TadayoshiOtsuka/go-tady/pkg/generator"
	"github.com/TadayoshiOtsuka/go-tady/pkg/gomod"
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
	fmt.Printf("%s in %d[ms]\n", "Generate Template Done.", time.Since(now).Milliseconds())
	if err := gomod.Setup(projectName, packageName); err != nil {
		return err
	}
	fmt.Printf("%s in %d[ms]\n", "go mod setup Done.", time.Since(now).Milliseconds())
	fmt.Printf("%s in %d[ms]\n", color.GreenString("SUCCESS🎉"), time.Since(now).Milliseconds())

	return nil
}
