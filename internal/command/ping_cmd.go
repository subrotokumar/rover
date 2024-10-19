package command

import "fmt"

type PingCommand struct {
}

func NewPingCommand() Command {
	return &PingCommand{}
}

func (c *PingCommand) Execute(cmd []string) (string, error) {
	switch len(cmd) {
	case 1:
		return "+PONG\r\n", nil
	case 2:
		return fmt.Sprintf("+%s\r\n", cmd[1]), nil
	default:
		return "-ERR ERR wrong number of arguments for 'ping' command", nil
	}

}
