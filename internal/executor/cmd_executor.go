package executor

import (
	"log"
	"net"

	"github.com/subrotokumar/rover/internal/command"
)

type commandExecuter struct {
	conn net.Conn
}

func NewExecutor(conn net.Conn) Executer {
	return &commandExecuter{conn}
}

func (e *commandExecuter) Execute(cmd []string) {
	if len(cmd) == 0 {
		e.conn.Write([]byte("-ERR empty cmd"))
		return
	}
	commandRunner := command.CommandFactory(cmd)
	response := commandRunner.Execute(cmd)
	log.Printf("Response => %s", response)
	e.conn.Write([]byte(response))
}
