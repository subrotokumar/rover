package command

import (
	"fmt"
	"strconv"

	"github.com/subrotokumar/rover/internal/store"
)

type DecrByCommand struct {
}

func NewDecrByCommand() Command {
	return &DecrByCommand{}
}

func (c *DecrByCommand) Execute(cmd []string) string {
	if len(cmd) != 3 {
		return "-ERR wrong number of arguments for 'decrby' command\r\n"
	}
	decrBy, err := strconv.Atoi(cmd[2])
	if err != nil {
		return "-ERR value is not an integer or out of range\r\n"
	}
	decrBy = decrBy * -1
	store := store.GetInstance()
	value, err := store.Get(cmd[1])
	if err != nil {
		store.Insert(cmd[1], decrBy)
		return fmt.Sprintf(":%d\r\n", decrBy)
	}

	valueNum, err := strconv.Atoi(fmt.Sprintf("%v", value))
	if err != nil {
		return "-ERR value is not an integer or out of range\r\n"
	}
	valueNum = valueNum + decrBy
	store.Insert(cmd[1], valueNum)
	return fmt.Sprintf(":%d\r\n", valueNum)
}
