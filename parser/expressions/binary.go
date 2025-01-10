package expressions

import (
	"Better-Language/scanner"
	"Better-Language/scanner/tokentype"
	"Better-Language/utils"
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
	left, err := b.Left.Interpret()
	if err != nil {
		return nil, err
	}
	right, err := b.Right.Interpret()
	if err != nil {
		return nil, err
	}

	switch b.Operator.Type {
	case tokentype.NotEqual:
	case tokentype.EqualEqual:

	case tokentype.Greater:
	case tokentype.GreaterEqual:
	case tokentype.Less:
	case tokentype.LessEqual:

	case tokentype.BitwiseOR:

	case tokentype.BitwiseXOR:

	case tokentype.BitwiseAND:

	case tokentype.BitwiseLeftShift:
	case tokentype.BitwiseRightShift:

	case tokentype.Minus:
		if !utils.IsNumber(left) || !utils.IsNumber(right) {
			return nil, utils.CreateErrorf("expect a number after '-'")
		}
	case tokentype.Plus:
		
	case tokentype.Star:
	case tokentype.Slash:
	case tokentype.Percent:
	default:
		panic("Unknown binary operator")
	}
	panic("implement me")
}
