package environment

import (
	"Better-Language/scanner"
	"Better-Language/utils"
)

type Environment interface {
	Define(name string, value any) (ok bool)
	Get(name scanner.Token) (value any, ok bool)
	Assign(name scanner.Token, value any) (ok bool)
}

type environment struct {
	values map[string]any
}

func NewEnvironment() Environment {
	return &environment{
		values: map[string]any{},
	}
}

func (e *environment) Define(name string, value any) (ok bool) {
	_, found := e.values[name]
	if found {
		utils.CreateAndReportParsingErrorf("Variable with name '%s' already defined", name)
		return false
	}
	e.values[name] = value
	return true
}

func (e *environment) Assign(name scanner.Token, value any) (ok bool) {
	_, found := e.values[name.Lexeme]
	if !found {
		utils.CreateAndReportParsingErrorf("Undefined variable '%s'", name.Lexeme)
		return false
	}

	e.values[name.Lexeme] = value
	return true
}

func (e *environment) Get(name scanner.Token) (value any, ok bool) {
	v, ok := e.values[name.Lexeme]
	if !ok {
		utils.CreateAndReportParsingErrorf("Undefined variable '%s'", name.Lexeme)
		return nil, false
	}
	return v, true
}
