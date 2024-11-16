package command

import (
	"github.com/subrotokumar/rover/internal/store"
)

type FlushAllCommand struct {
}

func NewFlushAllCommand() Command {
	return &FlushAllCommand{}
}

func (c *FlushAllCommand) Execute(db int, cmd []string) string {
	if len(cmd) > 1 {
		return "-ERR syntax error\r\n"
	}
	store := store.GetInstance()
	store.DeleteAll(db)
	return "+OK\r\n"
}
