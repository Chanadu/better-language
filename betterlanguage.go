package main

import (
	"Better-Language/scanner"
	"Better-Language/utils"
	"bufio"
	"fmt"
	"os"
)

func LineReader() error {
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		line := input.Text()
		//fmt.Println(line)
		if err := run(line); err != nil {
			return err
		}
	}

	if err := input.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
	return nil
}

func FileReader(fileName string) error {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}
	if err = run(string(data)); err != nil {
		return err
	}
	return nil
}

func run(source string) error {
	// Get Tokens And Print Later
	sc := scanner.NewScanner(source)
	tokens, err := sc.ScanTokens()
	if err != nil {
		return utils.CreateAndReportErrorf("Run Error: %e", err)
	}

	// Print Tokens
	for _, t := range tokens {
		fmt.Println(t)
	}

	//fmt.Println(source)
	return nil
}
