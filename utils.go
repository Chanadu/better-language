package main

import (
	"errors"
	"fmt"
)

func sendError(line int, message string) error {
	return errors.New(fmt.Sprintf("Error Line: %d, %s", line, message))
}
