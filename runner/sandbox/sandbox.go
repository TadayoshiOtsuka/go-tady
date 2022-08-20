package sandbox

import (
	"github.com/TadayoshiOtsuka/go-tady/runner/config"
	"github.com/TadayoshiOtsuka/go-tady/runner/generator"
)

func Create() error {
	projectName := config.Config.Name
	if err := generator.MakeDir(projectName); err != nil {
		return err
	}

	return nil
}
