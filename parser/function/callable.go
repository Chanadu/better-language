package function

import (
	"github.com/Chanadu/better-language/parser/interpreter"
)

type Callable interface {
	Call(interpreter.Interpreter, []any) (any, error)
	Arity() int
	String() string
}
