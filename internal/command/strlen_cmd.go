package command

import (
	"fmt"
	"reflect"
	"time"

	"github.com/subrotokumar/rover/internal/store"
	"github.com/subrotokumar/rover/internal/types"
)

type StrLenCommand struct {
}

func NewStrLenCommand() Command {
	return &StrLenCommand{}
}

func (c *StrLenCommand) Execute(cmd []string) string {
	if len(cmd) != 2 {
		return "-ERR wrong number of arguments for 'strlen' command\r\n"
	}

	store := store.GetInstance()
	key := cmd[1]
	value, err := store.Get(key)
	if value == nil {
		return ":0\r\n"
	}
	if err != nil {
		return fmt.Sprintf("-ERR %v\r\n", err)
	}

	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface || v.Kind() == reflect.Map ||
		v.Kind() == reflect.Slice || v.Kind() == reflect.Chan || v.Kind() == reflect.Func {
		if v.IsNil() {
			return ":0\r\n"
		}
	}

	var valueStr string
	if expirableVal, ok := value.(types.ExpirableValue); ok {
		if time.Now().After(expirableVal.ExpireAt) {
			go store.Delete(key)
			return ":0\r\n"
		}
		valueStr = fmt.Sprintf("%v", expirableVal.Value)
	} else {
		valueStr = fmt.Sprintf("%v", value)
	}
	return fmt.Sprintf(":%d\r\n", len(valueStr))
}
