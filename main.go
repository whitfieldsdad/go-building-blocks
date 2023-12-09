package main

import (
	"log"

	"github.com/whitfieldsdad/go-building-blocks/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
