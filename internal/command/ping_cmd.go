package command

import (
	"fmt"
)

type PingCommand struct {
}

func NewPingCommand() Command {
	return &PingCommand{}
}

func (c *PingCommand) Execute(cmd []string) string {
	switch len(cmd) {
	case 1:
		return "+PONG\r\n"
	case 2:
		return fmt.Sprintf("$%d\r\n%s\r\n", len(cmd[1]), cmd[1])
	default:
		fmt.Println("Hello World")
		return "-ERR wrong number of arguments for 'ping' command\r\n"
	}
}
