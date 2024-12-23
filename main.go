package main

import (
	"fmt"
	"os"
	"wrollup/wtools"

	"wrollup/cmd"
)

// go run main.go create --job sensor --indice vsd
// go run main.go delete --job sensor
// go run main.go clean --indice vsd --duration 2M
// go run main.go get all
// go run main.go get all
// go run main.go get --job sensor

func main() {
	err := wtools.LogToFile()
	if err != nil {
		fmt.Printf("Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}

	if err := cmd.Execute(); err != nil {
		wtools.Error("Error executing command: " + err.Error())
		os.Exit(1)
	}
}
