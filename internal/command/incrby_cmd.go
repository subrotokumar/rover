package command

import (
	"fmt"
	"strconv"

	"github.com/subrotokumar/rover/internal/store"
	"github.com/subrotokumar/rover/internal/types"
)

type IncrByCommand struct {
}

func NewIncrByCommand() Command {
	return &IncrByCommand{}
}

func (c *IncrByCommand) Execute(db int, cmd []string) string {
	if len(cmd) != 3 {
		return "-ERR wrong number of arguments for 'incrby' command\r\n"
	}
	incrBy, err := strconv.Atoi(cmd[2])
	if err != nil {
		return "-ERR value is not an integer or out of range\r\n"
	}
	store := store.GetInstance()
	value, err := store.Get(db, cmd[1])
	if err != nil {
		store.Insert(db, cmd[1], types.StoredValue{Value: incrBy})
		return fmt.Sprintf(":%d\r\n", incrBy)
	}

	valueNum, err := strconv.Atoi(value.String())
	if err != nil {
		return "-ERR value is not an integer or out of range\r\n"
	}
	valueNum = valueNum + incrBy
	store.Insert(db, cmd[1], types.StoredValue{Value: valueNum})
	return fmt.Sprintf(":%d\r\n", valueNum)
}
