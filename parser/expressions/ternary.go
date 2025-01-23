package expressions

import (
	"Better-Language/utils"
)

type Ternary struct {
	LineNumber  int
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
	cond, err := t.Condition.Evaluate()
	if err != nil {
		return nil, err
	}

	if cond == nil {
		return nil, utils.CreateRuntimeErrorf(t.LineNumber, "Condition of ternary operator cannot be null")
	}
	if _, ok := cond.(bool); !ok {
		return nil, utils.CreateRuntimeErrorf(t.LineNumber, "Condition of ternary operator must be a boolean")
	}
	condBool := cond.(bool)
	if condBool {
		return t.TrueBranch.Evaluate()
	}

	return t.FalseBranch.Evaluate()
}
