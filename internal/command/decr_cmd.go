package command

import (
	"fmt"
	"strconv"

	"github.com/subrotokumar/rover/internal/store"
)

type DecrCommand struct {
}

func NewDecrCommand() Command {
	return &DecrCommand{}
}

func (c *DecrCommand) Execute(cmd []string) string {
	if len(cmd) != 2 {
		return "-ERR wrong number of arguments for 'decr' command\r\n"
	}
	store := store.GetInstance()
	value, err := store.Get(cmd[1])
	if err != nil {
		store.Insert(cmd[1], -1)
		return ":-1\r\n"
	}

	valueNum, err := strconv.Atoi(fmt.Sprintf("%v", value))
	if err != nil {
		return "-ERR value is not an integer or out of range\r\n"
	}
	valueNum--
	store.Insert(cmd[1], valueNum)
	return fmt.Sprintf(":%d\r\n", valueNum)
}
