package expressions

import (
	"Better-Language/scanner"
)

type Binary struct {
	Left     Expression
	Operator scanner.Token
	Right    Expression
}

func (b *Binary) ToGrammarString() string {
	return parenthesizeExpression(b.Operator.Lexeme, b.Left, b.Right)
}

func (b *Binary) ToReversePolishNotation() string {
	return reversePolishNotation(b.Operator.Lexeme, b.Left, b.Right)
}

func (b *Binary) Interpret() (any, error) {
	panic("implement me")
}
