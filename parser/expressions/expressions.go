package expressions

import (
	"fmt"
)

type Expression interface {
	ToGrammarString() string
	ToReversePolishNotation() string
}

func parenthesizeExpression(name string, expressions ...Expression) (parenthesizedName string) {
	parenthesizedName = fmt.Sprintf("(%s", name)
	for _, expression := range expressions {
		parenthesizedName += fmt.Sprintf(" %s", expression.ToGrammarString())
	}
	parenthesizedName = fmt.Sprintf("%s)", parenthesizedName)
	return parenthesizedName
}

func reversePolishNotation(name string, expressions ...Expression) (reversePolishNotation string) {
	reversePolishNotation = ""
	for _, expression := range expressions {
		reversePolishNotation += fmt.Sprintf("%s ", expression.ToReversePolishNotation())
	}
	reversePolishNotation += name
	return reversePolishNotation
}
