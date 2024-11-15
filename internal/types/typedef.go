package types

import (
	"fmt"
	"time"
)

type StoredValue struct {
	ExpireAt time.Time
	Value    interface{}
}

func NewValue(value interface{}, duration time.Duration) StoredValue {
	return StoredValue{
		ExpireAt: time.Now().Add(duration),
		Value:    value,
	}
}

func (value *StoredValue) IsExpired() bool {
	if value == nil {
		return true
	} else if *value == (StoredValue{}) {
		return true
	} else if time.Now().After(value.ExpireAt) && (value.ExpireAt != time.Time{}) {
		return true
	}
	return false
}

func (value *StoredValue) String() string {
	return fmt.Sprintf("%v", value.Value)
}
