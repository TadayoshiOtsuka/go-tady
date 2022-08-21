package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Print("running...")
	cs := make(chan os.Signal, 1)
	signal.Notify(cs, syscall.SIGINT)
	<-cs
	os.Exit(0)
}
