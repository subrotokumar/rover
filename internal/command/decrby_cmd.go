package command

import (
	"fmt"
	"strconv"

	"github.com/subrotokumar/rover/internal/store"
	"github.com/subrotokumar/rover/internal/types"
)

type DecrByCommand struct {
}

func NewDecrByCommand() Command {
	return &DecrByCommand{}
}

func (c *DecrByCommand) Execute(db int, cmd []string) string {
	if len(cmd) != 3 {
		return "-ERR wrong number of arguments for 'decrby' command\r\n"
	}
	decrBy, err := strconv.Atoi(cmd[2])
	if err != nil {
		return "-ERR value is not an integer or out of range\r\n"
	}
	decrBy = decrBy * -1
	store := store.GetInstance()
	value, err := store.Get(db, cmd[1])
	if err != nil {
		store.Insert(db, cmd[1], types.StoredValue{Value: value})
		return fmt.Sprintf(":%d\r\n", decrBy)
	}

	valueNum, err := strconv.Atoi(value.String())
	if err != nil {
		return "-ERR value is not an integer or out of range\r\n"
	}
	valueNum = valueNum + decrBy
	value.Value = valueNum
	store.Insert(db, cmd[1], value)
	return fmt.Sprintf(":%d\r\n", valueNum)
}
