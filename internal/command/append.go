package command

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/subrotokumar/rover/internal/store"
	"github.com/subrotokumar/rover/internal/types"
)

type AppendCommand struct {
}

func NewAppendCommand() Command {
	return &AppendCommand{}
}

func (c *AppendCommand) Execute(cmd []string) string {
	if len(cmd) != 3 {
		return "-ERR wrong number of arguments for 'append' command\r\n"
	}

	key, newValue := cmd[1], cmd[2]

	store := store.GetInstance()
	value, err := store.Get(key)

	if value == nil {
		return AppendNew(store, key, newValue)
	}
	if err != nil {
		return AppendNew(store, key, newValue)
	}

	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface || v.Kind() == reflect.Map ||
		v.Kind() == reflect.Slice || v.Kind() == reflect.Chan || v.Kind() == reflect.Func {
		if v.IsNil() {
			return AppendNew(store, key, newValue)
		}
	}

	if expirableVal, ok := value.(types.ExpirableValue); ok {
		if time.Now().After(expirableVal.ExpireAt) {
			store.Delete(key)
			return AppendNew(store, key, newValue)
		} else {
			var builder strings.Builder
			builder.WriteString(fmt.Sprintf("%v", expirableVal.Value))
			builder.WriteString(newValue)
			expirableVal.Value = interface{}(builder.String())
			store.Insert(key, expirableVal)
			return fmt.Sprintf(":%d\r\n", builder.Len())
		}
	} else {
		var builder strings.Builder
		builder.WriteString(fmt.Sprintf("%v", value))
		builder.WriteString(newValue)
		return fmt.Sprintf(":%d\r\n", builder.Len())
	}

}

func AppendNew(store *store.SafeMap[string, interface{}], key, value string) string {
	store.Insert(key, value)
	return fmt.Sprintf(":%d\r\n", len(value))
}
