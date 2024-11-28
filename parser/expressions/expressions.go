package expressions

import (
	"fmt"

	"Better-Language/scanner"
)

type Expression interface {
	ToGrammarString() string
}

func parenthesizeExpression(name string, expressions ...Expression) (parenthesizedName string) {
	parenthesizedName = fmt.Sprintf("(%s", name)
	for _, expression := range expressions {
		parenthesizedName += fmt.Sprintf(" %s", expression.ToGrammarString())
	}
	parenthesizedName = fmt.Sprintf("%s)", parenthesizedName)
	return parenthesizedName
}

type Binary struct {
	Left     Expression
	Operator scanner.Token
	Right    Expression
}

func (b *Binary) ToGrammarString() string {
	return parenthesizeExpression(b.Operator.Lexeme, b.Left, b.Right)
}

type Grouping struct {
	InternalExpression Expression
}

func (g *Grouping) ToGrammarString() string {
	return parenthesizeExpression("group", g.InternalExpression)
}

type Literal struct {
	Value any
}

func (l *Literal) ToGrammarString() string {
	if l.Value == nil {
		return "null"
	}
	return fmt.Sprint(l.Value)
}

type Unary struct {
	Operator scanner.Token
	Right    Expression
}

func (u *Unary) ToGrammarString() string {
	return parenthesizeExpression(u.Operator.Lexeme, u.Right)
}
