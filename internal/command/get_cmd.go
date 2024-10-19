package command

type GetCommand struct {
}

func NewGetCommand() Command {
	return &GetCommand{}
}

func (c *GetCommand) Execute(cmd []string) (string, error) {
	return "-ERR unknown command\r\n", nil
}
