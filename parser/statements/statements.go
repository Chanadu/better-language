package statements

type Statement interface {
	Run() error
}
