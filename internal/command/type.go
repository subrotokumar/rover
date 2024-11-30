package command

import (
	"fmt"
	"reflect"

	"github.com/subrotokumar/rover/internal/store"
	"github.com/subrotokumar/rover/internal/types"
)

const stringResp = "+string\r\n"
const noneResp = "+none\r\n"

type TypeCommand struct {
}

func NewTypeCommand() Command {
	return &TypeCommand{}
}

func (c *TypeCommand) Execute(db int, cmd []string) string {

	if len(cmd) != 2 {
		return "-ERR wrong number of arguments for 'type' command\r\n"
	}

	store := store.GetInstance()
	key := cmd[1]
	value, err := store.Get(db, key)
	if (value == types.StoredValue{}) {
		return noneResp
	}
	if err != nil {
		return fmt.Sprintf("-ERR %v\r\n", err)
	}

	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface || v.Kind() == reflect.Map ||
		v.Kind() == reflect.Slice || v.Kind() == reflect.Chan || v.Kind() == reflect.Func {
		if v.IsNil() {
			return noneResp
		}
	}

	if value.IsExpired() {
		go store.Delete(db, key)
		return noneResp
	}
	return stringResp
}
