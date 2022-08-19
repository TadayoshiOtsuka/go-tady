package sandbox

import (
	"fmt"

	"github.com/TadayoshiOtsuka/go-tady/runner/config"
)

func Create() {
	fmt.Print(config.Config.Name)
}
