package version

import (
	"fmt"
	"runtime/debug"
)

var version = ""

func Exec() {
	fmt.Printf("go-tady version %s\n", getBuildVersion())
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
