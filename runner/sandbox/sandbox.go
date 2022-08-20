package sandbox

import (
	"fmt"

	"github.com/TadayoshiOtsuka/go-tady/runner/config"
	"github.com/TadayoshiOtsuka/go-tady/runner/generator"
	"github.com/TadayoshiOtsuka/go-tady/runner/gomod"
)

const src = "./templates/sandbox"

func Create() error {
	projectName := config.Config.Name
	if err := generator.Do(src, projectName); err != nil {
		return err
	}
	if err := gomod.Setup(projectName, fmt.Sprintf("github.com/TadayoshiOtsuka/%v", projectName)); err != nil {
		return err
	}

	return nil
}
