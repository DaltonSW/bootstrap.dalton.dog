package main

import (
	"log"

	"go.dalton.dog/setup/cmd"
)

func main() {
	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}
