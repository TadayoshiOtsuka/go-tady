package httpserver

import (
	"log"

	"github.com/manifoldco/promptui"
)

func SelectServerTemplate() {
	p := promptui.Select{
		Label: "Choose Create Template",
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
