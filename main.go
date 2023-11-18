package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/dartt0n/klaus/handlers"
	"github.com/dartt0n/klaus/klaus"
)

func main() {
	k, err := klaus.NewKlaus()
	if err != nil {
		log.Fatal(err)
	}

	handlers.AddStartHandler(k)
	handlers.AddDebugHandler(k)
	handlers.AddAdminHandlers(k)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		k.Run()
	}()

	<-c
	k.Storage.CreateSnapshot()

}
