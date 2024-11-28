package main

import (
	"os"

	"Better-Language/utils"
)

func main() {
	args := os.Args

	if len(args) > 2 {
		utils.ReportDebugf("Usage: gbpl [script file]")
		os.Exit(2)
	} else if len(args) == 1 {
		LineReader()
	} else {
		FileReader(args[1])
	}

	os.Exit(0)
}
