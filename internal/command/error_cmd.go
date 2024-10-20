package command

import "fmt"

type ErrCommand struct {
}

func NewErrCommand() Command {
	return &ErrCommand{}
}

func (c *ErrCommand) Execute(cmd []string) string {
	switch len(cmd) {
	case 1:
		return "-ERR unknown command 'exist', with args beginning with:\r\n"
	default:
		return fmt.Sprintf("-ERR unknown command 'exist', with args beginning with: '%s'\r\n", cmd[1])
	}
}
