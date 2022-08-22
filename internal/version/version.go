package version

import (
	"runtime/debug"

	"github.com/TadayoshiOtsuka/go-tady/pkg/utils"
)

var version = ""

func Exec() {
	utils.PrintGoTady(getBuildVersion())
}

func getBuildVersion() string {
	if version != "" {
		return version
	}
	i, ok := debug.ReadBuildInfo()
	if !ok {
		return "unknown"
	}
	return i.Main.Version
}
