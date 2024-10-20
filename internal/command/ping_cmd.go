package command

import (
	"fmt"
	"log"
)

type PingCommand struct {
}

func NewPingCommand() Command {
	return &PingCommand{}
}

func (c *PingCommand) Execute(cmd []string) (string, error) {
	log.Printf("Length %d", len(cmd))
	switch len(cmd) {
	case 1:
		return "+PONG\r\n", nil
	case 2:
		return fmt.Sprintf("$%d\r\n%s\r\n", len(cmd[1]), cmd[1]), nil
	default:
		fmt.Println("Hello World")
		return "-ERR wrong number of arguments for 'ping' command", nil
	}
}
