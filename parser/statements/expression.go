package statements

import (
	"Better-Language/parser/expressions"
)

type Expression struct {
	Expression expressions.Expression
}

func (e *Expression) Run() error {
	_, err := e.Expression.Evaluate()
	return err
}
