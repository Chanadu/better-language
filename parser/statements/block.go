package statements

import (
	"Better-Language/parser/environment"
)

type Block struct {
	Statements []Statement
}

func (b Block) Run(env environment.Environment) error {
	return b.executeBlock(b.Statements, environment.NewEnvironment(env))
}

func (b Block) executeBlock(statements []Statement, env environment.Environment) error {
	for _, statement := range statements {
		err := statement.Run(env)
		if err != nil {
			return err
		}
	}
	return nil

}
