package statements

import (
	"Better-Language/parser/environment"
	"Better-Language/scanner"
)

type Function struct {
	Name   scanner.Token
	Params []scanner.Token
	Body   []Statement
}

func (f Function) Run(env environment.Environment) error {
	panic("implement me")
}
