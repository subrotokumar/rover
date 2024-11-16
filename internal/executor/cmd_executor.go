package executor

import (
	"net"

	"github.com/subrotokumar/rover/internal/command"
)

type commandExecuter struct {
	conn net.Conn
}

func NewExecutor(conn net.Conn) Executer {
	return &commandExecuter{conn}
}

func (e *commandExecuter) Execute(db int, cmd []string) string {
	if len(cmd) == 0 {
		return "-ERR empty cmd"
	}
	commandRunner := command.CommandFactory(cmd)
	response := commandRunner.Execute(db, cmd)
	return response
}
