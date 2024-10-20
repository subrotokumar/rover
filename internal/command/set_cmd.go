package command

import (
	"github.com/subrotokumar/rover/internal/store"
)

type SetCommand struct {
}

func NewSetCommand() Command {
	return &SetCommand{}
}

func (c *SetCommand) Execute(cmd []string) (string, error) {
	store := store.GetInstance()
	key, value := cmd[1], cmd[2]
	store.Insert(key, value)
	return "+OK\r\n", nil
}
