package app

import (
	"log"
	"net"

	"github.com/subrotokumar/rover/internal/executor"
	"github.com/subrotokumar/rover/internal/parser"
)

var (
	printf  = log.Printf
	println = log.Println
)

type App struct {
	conn     net.Conn
	parser   parser.Parser
	executor executor.Executer
}

func NewApplication(conn net.Conn) *App {
	return &App{
		conn:     conn,
		parser:   parser.NewRespParser(),
		executor: executor.NewExecutor(conn),
	}
}
