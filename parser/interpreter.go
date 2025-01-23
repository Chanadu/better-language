package parser

import (
	"Better-Language/parser/statements"
)

func Interpret(statements []statements.Statement) (ok bool) {
	for _, statement := range statements {
		err := statement.Run()
		if err != nil {
			return false
		}
	}

	return true
}
