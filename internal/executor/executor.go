package executor

type Executer interface {
	Execute(db int, cmd []string) string
}
