package parser

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"strconv"
)

type RespType rune

const (
	STRING  RespType = '+'
	ERROR   RespType = '-'
	INTEGER RespType = ':'
	BULK    RespType = '$'
	ARRAY   RespType = '*'
)

type respParser struct{}

func NewRespParser() Parser {
	return &respParser{}
}

func (p *respParser) Parse(cmd []byte) (interface{}, error) {
	if len(cmd) == 0 {
		return nil, errors.New("empty input")
	}
	switch RespType(cmd[0]) {
	case STRING:
		return p.parseSimpleString(cmd)
	case ERROR:
		return p.parseError(cmd)
	case INTEGER:
		return p.parseInteger(cmd)
	case BULK:
		return p.parseBulkString(cmd)
	case ARRAY:
		return p.parseArray(cmd)
	default:
		log.Fatalln("Error => \n", string(cmd))
		return nil, errors.New("unknown data type")
	}
}

// parseSimpleString parses a data Simple String
func (p *respParser) parseSimpleString(data []byte) (string, error) {
	return string(bytes.TrimSuffix(data[1:], []byte("\r\n"))), nil
}

// parseError parses a data Error message
func (p *respParser) parseError(data []byte) (string, error) {
	return string(bytes.TrimSuffix(data[1:], []byte("\r\n"))), nil
}

// parseInteger parses a data Integer
func (p *respParser) parseInteger(data []byte) (int64, error) {
	value := bytes.TrimSuffix(data[1:], []byte("\r\n"))
	return strconv.ParseInt(string(value), 10, 64)
}

// parseBulkString parses a data Bulk String
func (p *respParser) parseBulkString(data []byte) (string, error) {
	dataLen := bytes.Index(data, []byte("\r\n"))
	length, err := strconv.Atoi(string(data[1:dataLen]))
	if err != nil {
		return "", err
	}
	if length == -1 {
		return "", nil // Null bulk string
	}

	start := dataLen + 2
	end := start + length
	if end+2 > len(data) || !bytes.Equal(data[end:end+2], []byte("\r\n")) {
		return "", errors.New("invalid bulk string format")
	}
	return string(data[start:end]), nil
}

// parseArray parses a data Array
func (p *respParser) parseArray(data []byte) ([]string, error) {
	i := 0

	// Helper to read the next line (delimited by \r\n)
	readLine := func() (string, error) {
		start := i
		for i < len(data) && data[i] != '\r' {
			i++
		}
		if i+1 >= len(data) || data[i+1] != '\n' {
			return "", fmt.Errorf("invalid data format")
		}
		line := string(data[start:i])
		i += 2 // Skip over \r\n
		return line, nil
	}

	// Read the first line which should start with '*'
	firstLine, err := readLine()
	if err != nil {
		return nil, err
	}

	if firstLine[0] != '*' {
		return nil, fmt.Errorf("invalid array indicator")
	}

	// Parse the array length
	arrayLength, err := strconv.Atoi(firstLine[1:])
	if err != nil {
		return nil, err
	}

	result := make([]string, 0, arrayLength)

	// Process each bulk string
	for j := 0; j < arrayLength; j++ {
		// Read the length indicator (starts with '$')
		lengthLine, err := readLine()
		if err != nil {
			return nil, err
		}

		if lengthLine[0] != '$' {
			return nil, fmt.Errorf("invalid bulk string indicator")
		}

		// Parse the bulk string length
		bulkLength, err := strconv.Atoi(lengthLine[1:])
		if err != nil {
			return nil, err
		}

		// Read the actual bulk string
		bulkString := string(data[i : i+bulkLength])
		result = append(result, bulkString)

		// Skip over the bulk string and \r\n
		i += bulkLength + 2
	}

	return result, nil
}
