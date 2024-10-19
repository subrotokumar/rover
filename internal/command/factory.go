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
		return NewErrCommand()
	case "EXISTS":
		return NewErrCommand()
	case "INCR":
		return NewErrCommand()
	case "INCRBY":
		return NewErrCommand()
	case "FLUSHALL":
		return NewErrCommand()
	default:
		return NewErrCommand()
	}
}
