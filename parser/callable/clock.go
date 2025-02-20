package callable

import (
	"time"

	"github.com/Chanadu/better-language/parser/environment"
)

type Clock struct{}

func (c *Clock) Arity() int {
	return 0
}

func (c *Clock) Call(env environment.Environment, args []any) (any, error) {
	return time.Now().UnixMilli(), nil
}

func (c *Clock) String() string {
	return "<clock native callable>"
}
