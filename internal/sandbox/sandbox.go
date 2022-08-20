package sandbox

import (
	"fmt"

	"github.com/TadayoshiOtsuka/go-tady/pkg/config"
	"github.com/TadayoshiOtsuka/go-tady/pkg/generator"
	"github.com/TadayoshiOtsuka/go-tady/pkg/gomod"
)

const src = "templates/sandbox"

func Create() error {
	projectName := config.Config.Name
	userName := config.Config.UserName
	packageName := fmt.Sprintf("github.com/%v/%v", userName, projectName)

	if err := generator.Do(src, projectName); err != nil {
		return err
	}
	if err := gomod.Setup(projectName, packageName); err != nil {
		return err
	}

	return nil
}
