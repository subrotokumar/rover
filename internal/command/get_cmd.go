package command

import (
	"fmt"
	"reflect"

	"github.com/subrotokumar/rover/internal/store"
)

type GetCommand struct {
}

func NewGetCommand() Command {
	return &GetCommand{}
}

func (c *GetCommand) Execute(cmd []string) (string, error) {
	if len(cmd) != 2 {
		return "", fmt.Errorf("-ERR wrong number of arguments for 'get' command")
	}

	store := store.GetInstance()
	key := cmd[1]
	value, err := store.Get(key)
	if err != nil {
		return "", fmt.Errorf("-ERR %v", err)
	}

	if value == nil {
		return "$-1\r\n", nil
	}

	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface || v.Kind() == reflect.Map ||
		v.Kind() == reflect.Slice || v.Kind() == reflect.Chan || v.Kind() == reflect.Func {
		if v.IsNil() {
			return "$-1\r\n", nil
		}
	}

	valueStr := fmt.Sprintf("%v", value)
	return fmt.Sprintf("$%d\r\n%s\r\n", len(valueStr), valueStr), nil
}
