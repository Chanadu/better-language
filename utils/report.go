package utils

import (
	"errors"
	"fmt"
	"os"

	"github.com/fatih/color"
)

func reportError(e error) {
	_, _ = fmt.Fprintf(os.Stderr, "Error: %s", color.RedString(e.Error()))
}

func ReportDebugf(format string, args ...any) {
	_, _ = fmt.Println(color.CyanString(format, args...))
}

func CreateErrorf(format string, args ...any) error {
	return errors.New(fmt.Sprintf("%s", fmt.Sprintf(format, args...)))
}

func CreateAndReportErrorf(format string, args ...any) {
	errorMessage := CreateErrorf(fmt.Sprintf(format, args...))
	reportError(errorMessage)
}

func CreateScannerErrorf(line int, format string, args ...any) error {
	return errors.New(fmt.Sprintf("Line: %d, %s", line, fmt.Sprintf(format, args...)))
}

func CreateAndReportScannerErrorf(line int, format string, args ...any) {
	errorMessage := CreateScannerErrorf(line, format, args...)
	reportError(errorMessage)
}

func CreateAndReportParsingError(format string, args ...any) {
	errorMessage := fmt.Sprintf("Parsing: %s", fmt.Sprintf(format, args...))
	reportError(errors.New(errorMessage))
}
