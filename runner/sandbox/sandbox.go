package sandbox

import (
	"github.com/TadayoshiOtsuka/go-tady/runner/config"
	"github.com/TadayoshiOtsuka/go-tady/runner/generator"
)

const src = "./templates/sandbox"

func Create() error {
	projectName := config.Config.Name
	if err := generator.Do(src, projectName); err != nil {
		return err
	}
	// if err := generator.MakeRoot(projectName); err != nil {
	// 	return err
	// }

	// if err := generator.Scan(src, projectName); err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }

	return nil
}
