package expressions

import (
	"testing"

	"Better-Language/scanner"
	"Better-Language/scanner/tokentype"
	"Better-Language/utils"
)

func TestExpressions(t *testing.T) {
	e := (&Binary{
		Left: &Unary{
			Operator: scanner.Token{
				Type:    tokentype.Minus,
				Lexeme:  "-",
				Literal: nil,
				Line:    1,
			},
			Right: &Literal{
				Value: 123,
			},
		},
		Operator: scanner.Token{
			Type:    tokentype.Star,
			Lexeme:  "*",
			Literal: nil,
			Line:    1,
		},
		Right: &Grouping{
			InternalExpression: &Literal{
				Value: 45.67,
			},
		},
	}).ToGrammarString()

	t.Logf("Expression: %s", e)
	utils.AssertEqual(t, "(* (- 123) (group 45.67))", e)
}
