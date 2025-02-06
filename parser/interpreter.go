package parser

import (
	"Better-Language/parser/environment"
	"Better-Language/parser/statements"
	"Better-Language/utils"
)

type Interpreter interface {
	Interpret(statements []statements.Statement) (ok bool)
}

type interpreter struct {
	environment environment.Environment
}

func NewInterpreter() Interpreter {
	return &interpreter{
		environment: environment.NewEnvironment(nil),
	}
}

func (i *interpreter) Interpret(statements []statements.Statement) (ok bool) {
	if len(statements) == 0 {
		return false
	}
	for _, statement := range statements {
		err := statement.Run(i.environment)

		if err != nil {
			utils.ReportError(err)
			return false
		}
	}

	return true
}
