package parser

import (
	"Better-Language/parser/statements"
	"Better-Language/utils"
)

func Interpret(statements []statements.Statement) (ok bool) {
	for _, statement := range statements {
		err := statement.Run()

		if err != nil {
			utils.ReportError(err)
			return false
		}
	}

	return true
}
