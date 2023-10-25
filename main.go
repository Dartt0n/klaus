package main

import (
	"log"

	"github.com/dartt0n/klaus/klaus"
)

func main() {
	k, err := klaus.NewKlaus()
	if err != nil {
		log.Fatal(err)
	}

	k.Run()
}
