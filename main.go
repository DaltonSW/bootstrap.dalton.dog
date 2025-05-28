package main

import (
	"log"

	"go.dalton.dog/setup/cmd"
	"go.dalton.dog/setup/cmd/utils"
)

var _ cmd.RawTaskList
var _ utils.DistroPackageManager

func main() {
	err := cmd.Run()

	// _, err := utils.DeterminePackageManager()

	if err != nil {
		log.Fatal(err)
	}
}
