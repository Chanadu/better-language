package parser

import (
	"Better-Language/parser/expressions"
	"Better-Language/utils"
)

func Interpret(expression expressions.Expression) (ok bool) {
	println(expression.ToGrammarString())

	v, err := expression.Evaluate()
	if err != nil {
		utils.ReportError(err)
		return false
	}
	utils.ReportDebugf("Result: %v", v)

	return true
}
