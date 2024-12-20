package main

import (
	"fmt"
	"os"
	"wrollup/wtools"

	"wrollup/cmd"
)

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
