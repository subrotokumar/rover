package command

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/subrotokumar/rover/internal/store"
	"github.com/subrotokumar/rover/internal/types"
)

type AppendCommand struct {
}

func NewAppendCommand() Command {
	return &AppendCommand{}
}

func (c *AppendCommand) Execute(db int, cmd []string) string {
	if len(cmd) != 3 {
		return "-ERR wrong number of arguments for 'append' command\r\n"
	}

	key, newValue := cmd[1], cmd[2]

	store := store.GetInstance()
	value, err := store.Get(db, key)

	if (value == types.StoredValue{}) {
		return AppendNew(db, store, key, newValue)
	}
	if err != nil {
		return AppendNew(db, store, key, newValue)
	}

	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface || v.Kind() == reflect.Map ||
		v.Kind() == reflect.Slice || v.Kind() == reflect.Chan || v.Kind() == reflect.Func {
		if v.IsNil() {
			return AppendNew(db, store, key, newValue)
		}
	}

	if value.IsExpired() {
		store.Delete(db, key)
		return AppendNew(db, store, key, newValue)
	} else {
		var builder strings.Builder
		builder.WriteString(value.String())
		builder.WriteString(newValue)
		value.Value = interface{}(builder.String())
		store.Insert(db, key, value)
		return fmt.Sprintf(":%d\r\n", builder.Len())
	}

}

func AppendNew(db int, store *store.SafeMap[string, types.StoredValue], key, value string) string {
	store.Insert(db, key, types.StoredValue{
		Value: value,
	})
	return fmt.Sprintf(":%d\r\n", len(value))
}
