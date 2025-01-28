package statements

import (
	"Better-Language/parser/environment"
	"Better-Language/parser/expressions"
)

type Expression struct {
	Expression expressions.Expression
}

func (e *Expression) Run(env environment.Environment) error {
	_, err := e.Expression.Evaluate(env)
	return err
}
