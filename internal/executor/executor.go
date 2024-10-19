package executor

type Executer interface {
	Execute(cmd []string) error
}
