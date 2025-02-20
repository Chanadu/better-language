package statements

import (
	"github.com/Chanadu/better-language/parser/environment"
	"github.com/Chanadu/better-language/scanner"
)

type Function struct {
	Name   scanner.Token
	Params []scanner.Token
	Body   []Statement
}

func (f Function) Run(env environment.Environment) error {
	panic("implement me")
}
