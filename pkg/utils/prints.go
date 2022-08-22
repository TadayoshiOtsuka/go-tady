package utils

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func PrintWithStartPrefix(text string) {
	fmt.Printf("[%s] %s\n", color.YellowString("START"), text)
}

func PrintDoneWithElapsedMillSec(now time.Time) {
	fmt.Printf("[%s] in %d[ms]\n", color.GreenString("DONE"), time.Since(now).Milliseconds())
}

func PrintWithElapsedMilliSec(text string, now time.Time) {
	fmt.Printf("%s in %d[ms]\n", text, time.Since(now).Milliseconds())

}

func PrintGoTady(arg string) {
	fmt.Printf(`	             _              _
    ____  ___       | |_ __ _  ____| |  _
   / _  |/ _ \ _____| __/ _ | / _  | | | |
  | (_| | (_) |_____| || (_| | (_| | |_| |
   \__, |\___/       \__\__,_|\__,_|\__, |
   |___/                            |___/  %v%v`, arg, "\n")
}
