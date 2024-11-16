package command

import (
	"fmt"
	"strconv"

	"github.com/subrotokumar/rover/internal/store"
	"github.com/subrotokumar/rover/internal/types"
)

type IncrCommand struct {
}

func NewIncrCommand() Command {
	return &IncrCommand{}
}

func (c *IncrCommand) Execute(db int, cmd []string) string {
	if len(cmd) != 2 {
		return "-ERR wrong number of arguments for 'incr' command\r\n"
	}
	store := store.GetInstance()
	value, err := store.Get(db, cmd[1])
	if err != nil {
		store.Insert(db, cmd[1], types.StoredValue{Value: "1"})
		return ":1\r\n"
	}

	valueNum, err := strconv.Atoi(value.String())
	if err != nil {
		return "-ERR value is not an integer or out of range\r\n"
	}
	valueNum++
	store.Insert(db, cmd[1], types.StoredValue{Value: valueNum})
	return fmt.Sprintf(":%d\r\n", valueNum)
}
