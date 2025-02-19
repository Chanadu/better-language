package functions

import (
	"Better-Language/parser/environment"
	"time"
)

type Callable interface {
	Call(environment.Environment, []any) (any, error)
	Arity() int
}

type Clock struct {
}

func (c *Clock) Call(env environment.Environment, args []any) (any, error) {
	return time.Now().UnixMilli(), nil
}
