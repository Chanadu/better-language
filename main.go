package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) > 2 {
		fmt.Println("Usage: gbpl [script file]")
		os.Exit(2)
	}
	if len(args) == 1 { // Line by Line Reader
		if err := LineReader(); err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}
	} else { // File Reader
		if err := FileReader(args[1]); err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}
	}

	os.Exit(0) // Success
}
