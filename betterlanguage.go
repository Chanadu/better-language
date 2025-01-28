package main

import (
	"bufio"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/fatih/color"

	"Better-Language/globals"
	"Better-Language/parser"
	"Better-Language/parser/statements"
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
	globals.HasErrors = false
	tokens, ok := runScanner(source)
	if !ok {
		utils.ReportDebugf("Errors found in scanner, exiting")
		return
	}
	printTokens(tokens)

	statement, ok := runParser(tokens)
	if !ok {
		utils.ReportDebugf("Errors found in parsing, exiting")
		return
	}

	ok = parser.Interpret(statement)
	if !ok {
		utils.ReportDebugf("Errors found in runtime, exiting")
		return
	}
}

func runScanner(source string) (tokens []scanner.Token, ok bool) {
	sc := scanner.NewScanner(source)
	tokens, err := sc.ScanTokens()

	if err != nil {
		utils.CreateAndReportErrorf("Token Scanning Error: %e", err)
		return nil, false
	}
	if globals.HasErrors {
		utils.ReportDebugf("Errors found in scanning, exiting")
		return nil, false
	}
	return tokens, true
}

func runParser(tokens []scanner.Token) (stmts []statements.Statement, ok bool) {
	p := parser.NewParser(tokens)

	stmts, err := p.Parse()

	if err != nil {
		utils.CreateAndReportParsingErrorf("%s", err.Error())
		return nil, false
	}

	if globals.HasErrors {
		utils.ReportDebugf("Errors found in parsing, exiting")
		return nil, false
	}

	return stmts, true
}

// func printExpressions(statements expressions.Expression) {
// 	utils.ReportDebugf("Parsed: %v", statements.ToGrammarString())
// }

//goland:noinspection GoUnusedFunction
func printTokens(tokens []scanner.Token) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)
	_, _ = fmt.Fprintln(w, color.CyanString("Type\tLexeme\tLiteral\tLine"))
	for _, t := range tokens {
		_, _ = fmt.Fprintln(w, color.CyanString(fmt.Sprintf("%s\t%#v\t%#v\t%d", t.Type.String(), t.Lexeme, t.Literal, t.Line)))
	}
	if err := w.Flush(); err != nil {
		utils.CreateAndReportErrorf("Error printing tokens: %e", err)
	}
}
