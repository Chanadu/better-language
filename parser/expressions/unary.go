package expressions

import (
	"Better-Language/scanner"
)

type Unary struct {
	Operator scanner.Token
	Right    Expression
}

func (u *Unary) ToGrammarString() string {
	return parenthesizeExpression(u.Operator.Lexeme, u.Right)
}

func (u *Unary) ToReversePolishNotation() string {
	return reversePolishNotation(u.Operator.Lexeme, u.Right)
}
