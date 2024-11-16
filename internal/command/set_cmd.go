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

func (c *SetCommand) Execute(db int, cmd []string) string {
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

	var oldValue types.StoredValue
	if getOldValue {
		oldValue, _ = store.Get(db, key)
	}

	existing, _ := store.Get(db, key)
	if (setIfNotExists && existing != types.StoredValue{}) {
		return "$-1\r\n"
	}
	if (setIfExists && existing == types.StoredValue{}) {
		return "$-1\r\n"
	}

	if expireDuration > 0 {
		store.Insert(db, key, types.StoredValue{
			Value:    value,
			ExpireAt: time.Now().Add(expireDuration),
		})
	} else {
		store.Insert(db, key, types.StoredValue{
			Value:    value,
			ExpireAt: time.Time{},
		})
	}

	if getOldValue {
		if oldValue.IsExpired() {
			go store.Delete(db, key)
			return "$-1\r\n"
		}
		oldValueStr := oldValue.String()
		return fmt.Sprintf("$%d\r\n%s\r\n", len(oldValueStr), oldValueStr)
	}
	return "+OK\r\n"

}
