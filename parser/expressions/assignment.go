package expressions

import (
	"Better-Language/parser/environment"
	"Better-Language/scanner"
)

type Assignment struct {
	Name  scanner.Token
	Value Expression
}

func (a Assignment) ToGrammarString() string {
	return parenthesizeExpression("=", a.Value)
}

func (a Assignment) ToReversePolishNotation() string {
	return reversePolishNotation("=", a.Value)
}

func (a Assignment) Evaluate(env environment.Environment) (any, error) {
	val, err := a.Value.Evaluate(env)
	if err != nil {
		return nil, err
	}

	env.Assign(a.Name, val)
	return val, nil
}
