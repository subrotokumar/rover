package command

import "strings"

func CommandFactory(cmd []string) Command {
	command := strings.ToUpper(cmd[0])
	switch command {
	case "PING":
		return NewPingCommand()
	case "SET":
		return NewSetCommand()
	case "GET":
		return NewGetCommand()
	case "DEL":
		return NewDelCommand()
	case "EXISTS":
		return NewExistsCommand()
	case "INCR":
		return NewIncrCommand()
	case "INCRBY":
		return NewIncrByCommand()
	case "DECR":
		return NewDecrCommand()
	case "DECRBY":
		return NewDecrByCommand()
	case "STRLEN":
		return NewStrLenCommand()
	case "FLUSHALL":
		return NewFlushAllCommand()
	case "APPEND":
		return NewAppendCommand()
	default:
		return NewErrCommand()
	}
}
