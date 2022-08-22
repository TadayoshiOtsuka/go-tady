package utils

import (
	"fmt"
	"time"
)

func PrintWithElapsed(text string, now time.Time) {
	fmt.Printf("%s in %d[ms]\n", text, time.Since(now).Milliseconds())

}
