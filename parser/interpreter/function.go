package interpreter

import (
	"github.com/Chanadu/better-language/parser/environment"
	"github.com/Chanadu/better-language/parser/statements"
)

type Function struct {
	Declaration statements.Function
}

func (f Function) Call(environment environment.Environment, args []any) (any, error) {
	panic("implement me")
}

func (f Function) Arity() int {
	// TODO implement me
	panic("implement me")
}

func (f Function) String() string {
	// TODO implement me
	panic("implement me")
}
