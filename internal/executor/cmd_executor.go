package executor

import (
	"errors"
	"net"

	"github.com/subrotokumar/rover/internal/command"
)

type commandExecuter struct {
	conn net.Conn
}

func NewExecutor(conn net.Conn) Executer {
	return &commandExecuter{conn}
}

func (e *commandExecuter) Execute(cmd []string) error {
	if len(cmd) == 0 {
		return errors.New("empty command")
	}
	commandRunner := command.CommandFactory(cmd)
	response, err := commandRunner.Execute(cmd)
	if err != nil {
		return err
	}
	e.conn.Write([]byte(response))
	return nil
}
