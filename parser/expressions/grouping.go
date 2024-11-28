package expressions

type Grouping struct {
	InternalExpression Expression
}

func (g *Grouping) ToGrammarString() string {
	return parenthesizeExpression("group", g.InternalExpression)
}

func (g *Grouping) ToReversePolishNotation() string {
	return g.InternalExpression.ToReversePolishNotation()
}
