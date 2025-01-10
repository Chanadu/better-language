package expressions

import (
	"Better-Language/scanner"
	"Better-Language/scanner/tokentype"
	"Better-Language/utils"
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

func (u *Unary) Interpret() (any, error) {
	right, _ := u.Right.Interpret()
	switch u.Operator.Type {
	case tokentype.Minus:
		d, dOk := right.(float64)
		i, iOk := right.(int)

		if !dOk && !iOk {
			return nil, utils.CreateErrorf("expect a number(double or int) after '-'")
		}
		if dOk && iOk {
			panic("Number is both int and float64")
		}

		if dOk {
			return -d, nil
		}

		return -i, nil
	case tokentype.Not:
		b, ok := right.(bool)
		if !ok {
			return nil, utils.CreateErrorf("expect a boolean after '!'")
		}
		return !b, nil

	case tokentype.BitwiseNOT:
		i, ok := right.(int)
		if !ok {
			return nil, utils.CreateErrorf("expect an integer after '~'")
		}
		return ^i, nil
	default:
		panic("Unknown unary operator")
	}
}
