package main

import (
	"bufio"
	"fmt"
	"os"
	"text/tabwriter"

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
	tokens, done := runScanner(source)
	if done {
		os.Exit(1)
	}
	printTokens(tokens)

	statement, ok := runParser(tokens)
	if ok {
		os.Exit(1)
	}
	printExpressions(statement)

	ok = parser.Interpret(statement)
	if !ok {
		os.Exit(1)
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

func runParser(tokens []scanner.Token) (statements expressions.Expression, done bool) {
	p := parser.NewParser(tokens)
	statements, err := p.Parse()
	if err != nil {
		utils.CreateAndReportParsingErrorf("%s", err.Error())
		return nil, true
	}
	if globals.HasErrors {
		utils.ReportDebugf("Errors found in parsing, exiting")
		return nil, true
	}
	return statements, false
}

func printExpressions(statements expressions.Expression) {
	utils.ReportDebugf("Parsed: %v", statements.ToGrammarString())
}

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
