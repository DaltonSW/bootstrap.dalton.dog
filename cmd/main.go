// main is the bootstrap entry point
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: setup <config_file_path>")
		os.Exit(1)
	}
	configFilePath := os.Args[1]

	_, err := os.ReadFile(configFilePath)
	if err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
		os.Exit(1)
	}
	return
}
