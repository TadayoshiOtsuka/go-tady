package httpserver

import (
	"log"

	"github.com/manifoldco/promptui"
)

func SelectServerTemplate() {
	p := promptui.Select{
		Label: "Select a HTTP package",
		Items: []string{
			"net/http",
		},
	}
	_, res, err := p.Run()
	if err != nil {
		log.Panicln(err)
	}

	switch res {
	case "net/http":
		log.Print("net/http selected")
	}
}
