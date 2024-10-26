package command

import (
	"fmt"

	"github.com/subrotokumar/rover/internal/store"
)

type DelCommand struct {
}

func NewDelCommand() Command {
	return &DelCommand{}
}

func (c *DelCommand) Execute(cmd []string) string {
	if len(cmd) < 2 {
		return "-ERR wrong number of arguments for 'del' command\r\n"
	}
	store := store.GetInstance()
	count := 0
	for i := 1; i < len(cmd); i++ {
		err := store.Delete(cmd[i])
		if err == nil {
			count++
		}
	}
	return fmt.Sprintf(":%d\r\n", count)
}
