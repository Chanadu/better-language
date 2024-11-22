package main

import (
	"fmt"
	"os"
)

func LineReader() error {
	var line string
	for {
		if _, err := fmt.Scanln(&line); err != nil {
			return err
		}
		if line == "" {
			break
		}
		if err := run(line); err != nil {
			return err
		}
	}
	return nil
}

func FileReader(fileName string) error {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}
	if err := run(string(data)); err != nil {
		return err
	}
	return nil
}

func run(source string) error {
	// Get Tokens And Print Later
	fmt.Println(source)
	return nil
}
