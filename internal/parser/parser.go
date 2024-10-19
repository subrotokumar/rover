package parser

import (
	"errors"
)

var (
	//*Error
	ErrEmptyString       error = errors.New("empty string")
	ErrInvalidBulkString error = errors.New("invalid bulk string format")
)

type Parser interface {
	Parse(cmd []byte) (interface{}, error)
	parseSimpleString(data []byte) (string, error)
	parseError(data []byte) (string, error)
	parseInteger(data []byte) (int64, error)
	parseBulkString(data []byte) (string, error)
	parseArray(data []byte) ([]string, error)
}
