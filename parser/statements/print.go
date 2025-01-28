package statements

import (
	"fmt"

	"github.com/fatih/color"

	"Better-Language/parser/environment"
	"Better-Language/parser/expressions"
)

type Print struct {
	Expression expressions.Expression
}

func (p *Print) Run(env environment.Environment) (err error) {
	v, err := p.Expression.Evaluate(env)
	if err != nil {
		return
	}

	_, _ = fmt.Println(color.GreenString("%v", v))
	return err
}
