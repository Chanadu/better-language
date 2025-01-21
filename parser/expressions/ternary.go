package expressions

type Ternary struct {
	Condition   Expression
	TrueBranch  Expression
	FalseBranch Expression
}

func (t Ternary) ToGrammarString() string {
	return parenthesizeExpression("?", t.Condition, t.TrueBranch, t.FalseBranch)
}

func (t Ternary) ToReversePolishNotation() string {
	return reversePolishNotation("?", t.Condition, t.TrueBranch, t.FalseBranch)
}

func (t Ternary) Evaluate() (any, error) {
	panic("implement me")
}
