package command

type SetCommand struct {
}

func NewSetCommand() Command {
	return &SetCommand{}
}

func (c *SetCommand) Execute(cmd []string) (string, error) {
	return "+OK\r\n", nil
}
