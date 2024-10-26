package types

import "time"

type ExpirableValue struct {
	ExpireAt time.Time
	Value    interface{}
}

func NewExpirableValue(value interface{}, duration time.Duration) ExpirableValue {
	return ExpirableValue{
		ExpireAt: time.Now().Add(duration),
		Value:    value,
	}
}
