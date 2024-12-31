package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"

	"Better-Language/globals"
	"Better-Language/parser"
	"Better-Language/parser/expressions"
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
	if globals.HasErrors {
		utils.ReportDebugf("Errors found in scanning, exiting")
		return
	}

	// printTokens(tokens)
	statements, done := printExpressions(tokens, err)
	if done {
		return
	}

	// fmt.Println(statements)
	utils.ReportDebugf("Parsed: %v", statements.ToGrammarString())
	if globals.HasErrors {
		utils.ReportDebugf("Errors found in parsing, exiting")
		return
	}
}

func printExpressions(tokens []scanner.Token, err error) (expressions.Expression, bool) {
	par := parser.NewParser(tokens)
	statements, err := par.Parse()
	if err != nil {
		utils.CreateAndReportParsingErrorf("%s", err.Error())
		return nil, true
	}
	return statements, false
}

//goland:noinspection GoUnusedFunction
func printTokens(tokens []scanner.Token) {
	for _, t := range tokens {
		formattedToken := fmt.Sprintf("Type: %20s \t\t Lexeme: %s \t\t Literal: %20v \t\t Line: %d", t.Type.String(), t.Lexeme, t.Literal, t.Line)
		utils.ReportDebugf("Token: %s", color.CyanString(formattedToken))
	}
}
