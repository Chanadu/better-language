package function

import (
	"Better-Language/parser/interpreter"
)

type Callable interface {
	Call(interpreter.Interpreter, []any) (any, error)
	Arity() int
	String() string
}
