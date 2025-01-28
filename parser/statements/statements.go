package statements

import (
	"Better-Language/parser/environment"
)

type Statement interface {
	Run(env environment.Environment) error
}
