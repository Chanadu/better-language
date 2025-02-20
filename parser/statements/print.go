package statements

import (
	"fmt"

	"github.com/fatih/color"

	"github.com/Chanadu/better-language/parser/environment"
	"github.com/Chanadu/better-language/parser/expressions"
)

type Print struct {
	Expression expressions.Expression
}

func (p *Print) Run(env environment.Environment) (err error) {
	v, err := p.Expression.Evaluate(env)
	if err != nil {
		return err
	}

	if v == nil {
		_, _ = fmt.Println(color.GreenString("null"))
		return nil
	}

	_, _ = fmt.Println(color.GreenString("%v", v))

	return nil
}
