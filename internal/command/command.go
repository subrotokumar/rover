package command

type Command interface {
	Execute(cmd []string) string
}
