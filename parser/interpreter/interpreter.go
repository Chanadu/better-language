package interpreter

import (
	"Better-Language/parser/builtin"
	"Better-Language/parser/environment"
	"Better-Language/parser/statements"
	"Better-Language/utils"
)

type Interpreter interface {
	Interpret(statements []statements.Statement) (ok bool)
}

type interpreter struct {
	globals     environment.Environment
	environment environment.Environment
}

func NewInterpreter() Interpreter {
	globals := environment.NewEnvironment(nil)

	globals.Define("clock", builtin.Clock{})

	return &interpreter{
		globals:     globals,
		environment: globals,
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
