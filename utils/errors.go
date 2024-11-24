package utils

import (
	"errors"
	"fmt"
	"os"

	"github.com/fatih/color"
)

func CreateError(message string) error {
	return errors.New(fmt.Sprintf("Error: %s", message))
}

func CreateErrorf(format string, args ...interface{}) error {
	return errors.New(fmt.Sprintf("Error: %s", fmt.Sprintf(format, args...)))
}

func CreateScannerError(line int, message string) error {
	return errors.New(fmt.Sprintf("Error Line: %d, %s", line, message))
}

func CreateScannerErrorf(line int, format string, args ...interface{}) error {
	return errors.New(fmt.Sprintf("Error Line: %d, %s", line, fmt.Sprintf(format, args...)))
}

func reportError(e error) error {
	if _, err := fmt.Fprintln(os.Stderr, color.RedString(e.Error())); err != nil {
		return err
	}
	return nil
}

func CreateAndReportScannerError(line int, message string) error {
	e := CreateScannerError(line, message)
	return reportError(e)
}

func CreateAndReportScannerErrorf(line int, format string, args ...interface{}) error {
	e := CreateScannerErrorf(line, format, args...)
	return reportError(e)
}

func CreateAndReportError(message string) error {
	e := CreateError(message)
	return reportError(e)
}

func CreateAndReportErrorf(format string, args ...interface{}) error {
	e := CreateErrorf(fmt.Sprintf(format, args...))
	return reportError(e)
}
