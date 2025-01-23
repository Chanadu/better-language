package statements

import (
	"fmt"

	"github.com/fatih/color"

	"Better-Language/parser/expressions"
)

type Print struct {
	Expression expressions.Expression
}

func (p *Print) Run() (err error) {
	v, err := p.Expression.Evaluate()
	if err != nil {
		return
	}

	_, _ = fmt.Println(color.GreenString("%v", v))
	return

}
