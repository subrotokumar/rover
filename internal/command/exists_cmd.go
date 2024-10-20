package command

import (
	"fmt"

	"github.com/subrotokumar/rover/internal/store"
)

type ExistsCommand struct {
}

func NewExistsCommand() Command {
	return &ExistsCommand{}
}

func (c *ExistsCommand) Execute(cmd []string) string {
	if len(cmd) < 2 {
		return "-ERR wrong number of arguments for 'exists' command\r\n"
	}
	store := store.GetInstance()
	count := 0
	for i := 1; i < len(cmd); i++ {
		if ok := store.ContainsKey(cmd[i]); ok {
			count++
		}
	}
	return fmt.Sprintf(":%d\r\n", count)
}
