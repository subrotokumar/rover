package command

type ErrCommand struct {
}

func NewErrCommand() Command {
	return &ErrCommand{}
}

func (c *ErrCommand) Execute(cmd []string) (string, error) {
	return "-ERR unknown command\r\n", nil
}
