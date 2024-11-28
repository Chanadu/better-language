package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"

	"Better-Language/scanner"
	"Better-Language/utils"
)

func LineReader() {
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		line := input.Text()
		run(line)
	}

	if err := input.Err(); err != nil {
		utils.CreateAndReportErrorf("Error reading input: %e", err)
	}
}

func FileReader(fileName string) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		utils.CreateAndReportErrorf("File Reading Error: %e", err)
		return
	}
	run(string(data))
}

func run(source string) {
	sc := scanner.NewScanner(source)
	tokens, err := sc.ScanTokens()

	if err != nil {
		utils.CreateAndReportErrorf("Token Scanning Error: %e", err)
		return
	}

	for _, t := range tokens {
		formattedToken := strings.Join(strings.Split(fmt.Sprintf("%#v", t.String()), " "), "\t")
		utils.ReportDebugf("Token: %s", color.CyanString(formattedToken))
	}
}
