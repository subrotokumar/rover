package command

type Command interface {
	Execute(db int, cmd []string) string
}
