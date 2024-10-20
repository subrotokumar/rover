package command

import (
	"fmt"
	"strconv"

	"github.com/subrotokumar/rover/internal/store"
)

type IncrByCommand struct {
}

func NewIncrByCommand() Command {
	return &IncrByCommand{}
}

func (c *IncrByCommand) Execute(cmd []string) string {
	if len(cmd) != 3 {
		return "-ERR wrong number of arguments for 'incr' command\r\n"
	}
	incrBy, err := strconv.Atoi(cmd[2])
	if err != nil {
		return "-ERR value is not an integer or out of range\r\n"
	}
	store := store.GetInstance()
	value, err := store.Get(cmd[1])
	if err != nil {
		store.Insert(cmd[1], incrBy)
		return fmt.Sprintf(":%d\r\n", incrBy)
	}

	valueNum, err := strconv.Atoi(fmt.Sprintf("%v", value))
	if err != nil {
		return "-ERR value is not an integer or out of range\r\n"
	}
	valueNum = valueNum + incrBy
	store.Insert(cmd[1], valueNum)
	return fmt.Sprintf(":%d\r\n", valueNum)
}
