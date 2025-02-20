package expressions

import (
	"fmt"

	"github.com/Chanadu/better-language/parser/environment"
	"github.com/Chanadu/better-language/parser/function"
	"github.com/Chanadu/better-language/scanner"
)

type Call struct {
	Callee Expression
	Para   scanner.Token
	Args   []Expression
}

func (c *Call) ToGrammarString() string {
	return parenthesizeExpression(c.Para.Lexeme, c.Args...)
}

func (c *Call) ToReversePolishNotation() string {
	return reversePolishNotation(c.Para.Lexeme, c.Args...)
}

func (c *Call) Evaluate(env environment.Environment) (any, error) {
	// panic("implement me")

	callee, err := c.Callee.Evaluate(env)
	if err != nil {
		return nil, err
	}

	var args []any
	for _, arg := range c.Args {
		value, err := arg.Evaluate(env)
		if err != nil {
			return nil, err
		}
		args = append(args, value)
	}
	var function function.Callable
	var ok bool

	if function, ok = callee.(function.Callable); !ok {
		return nil, fmt.Errorf("can only call functions and classes, %v", c.Para)
	}

	if len(args) != function.Arity() {
		return nil, fmt.Errorf("expected %d arguments but got %d", function.Arity(), len(args))
	}

	return function.Call(env, args)
}
