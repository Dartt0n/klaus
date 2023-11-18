package main

import (
	"log"

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

	k.Run()
}
