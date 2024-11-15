package command

import (
	"fmt"
	"reflect"

	"github.com/subrotokumar/rover/internal/store"
	"github.com/subrotokumar/rover/internal/types"
)

type GetCommand struct {
}

func NewGetCommand() Command {
	return &GetCommand{}
}

func (c *GetCommand) Execute(cmd []string) string {
	if len(cmd) != 2 {
		return "-ERR wrong number of arguments for 'get' command\r\n"
	}

	store := store.GetInstance()
	key := cmd[1]
	value, err := store.Get(key)
	if (value == types.StoredValue{}) {
		return "$-1\r\n"
	}
	if err != nil {
		return fmt.Sprintf("-ERR %v\r\n", err)
	}

	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface || v.Kind() == reflect.Map ||
		v.Kind() == reflect.Slice || v.Kind() == reflect.Chan || v.Kind() == reflect.Func {
		if v.IsNil() {
			return "$-1\r\n"
		}
	}

	if value.IsExpired() {
		go store.Delete(key)
		return "$-1\r\n"
	}
	valueStr := fmt.Sprintf("%v", value.String())
	return fmt.Sprintf("$%d\r\n%s\r\n", len(valueStr), valueStr)
}
