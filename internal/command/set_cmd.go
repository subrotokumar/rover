package command

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/subrotokumar/rover/internal/store"
	"github.com/subrotokumar/rover/internal/types"
)

type SetCommand struct {
}

func NewSetCommand() Command {
	return &SetCommand{}
}

func (c *SetCommand) Execute(cmd []string) string {
	if len(cmd) < 3 || len(cmd) > 6 {
		return "-ERR wrong number of arguments for 'set' command\r\n"
	}

	store := store.GetInstance()
	key, value := cmd[1], cmd[2]
	var (
		expireDuration time.Duration
		setIfExists    bool
		setIfNotExists bool
		getOldValue    bool
	)

	for i := 3; i < len(cmd); i++ {
		fmt.Printf("%d", i)
		fmt.Printf("%s", cmd[i])
		part := strings.ToUpper(cmd[i])

		switch part {
		case "EX":
			if i+1 <= 6 && i+1 < len(cmd) {
				seconds, err := strconv.Atoi(cmd[i+1])
				if err != nil {
					return "-ERR invalid expire time in 'set' command\r\n"
				}
				expireDuration = time.Duration(seconds) * time.Second
				i++
			} else {
				return "-ERR syntax error\r\n"
			}

		case "PX":
			if i+1 <= 6 && i+1 < len(cmd) {
				milliseconds, err := strconv.Atoi(cmd[i+1])
				if err != nil {
					return "-ERR invalid expire time in 'set' command\r\n"
				}
				expireDuration = time.Duration(milliseconds) * time.Millisecond
				i++
			} else {
				return "-ERR syntax error\r\n"
			}

		case "NX":
			setIfNotExists = true

		case "XX":
			setIfExists = true

		case "GET":
			getOldValue = true

		default:
			return "-ERR syntax error\r\n"
		}
	}
	if setIfExists && setIfNotExists {
		return "-ERR NX and XX options at the same time are not compatible\r\n"
	}

	var oldValue interface{}
	if getOldValue {
		oldValue, _ = store.Get(key)
	}

	existing, _ := store.Get(key)
	if setIfNotExists && existing != nil {
		return "$-1\r\n"
	}
	if setIfExists && existing == nil {
		return "$-1\r\n"
	}

	if expireDuration > 0 {
		store.Insert(key, types.ExpirableValue{
			Value:    value,
			ExpireAt: time.Now().Add(expireDuration),
		})
	} else {
		store.Insert(key, value)
	}

	if getOldValue {
		if expirableVal, ok := oldValue.(types.ExpirableValue); ok {
			if time.Now().After(expirableVal.ExpireAt) {
				go store.Delete(key)
				return "$-1\r\n"
			}
			oldValueStr := fmt.Sprintf("%v", expirableVal.Value)
			return fmt.Sprintf("$%d\r\n%s\r\n", len(oldValueStr), oldValueStr)
		} else {
			oldValueStr := fmt.Sprintf("%v", value)
			return fmt.Sprintf("$%d\r\n%s\r\n", len(oldValueStr), oldValueStr)
		}

	}
	return "+OK\r\n"

}
