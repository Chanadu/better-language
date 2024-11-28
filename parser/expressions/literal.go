package expressions

import (
	"fmt"
)

type Literal struct {
	Value any
}

func (l *Literal) ToGrammarString() string {
	if l.Value == nil {
		return "null"
	}
	return fmt.Sprint(l.Value)
}

func (l *Literal) ToReversePolishNotation() string {
	if l.Value == nil {
		return "null"
	}
	return fmt.Sprint(l.Value)
}
