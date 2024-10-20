package app

import (
	"fmt"
	"log"

	"github.com/subrotokumar/rover/internal/parser"
	"github.com/subrotokumar/rover/internal/store"
)

type App struct {
	PORT    string
	VERSION string

	store  *store.SafeMap[string, interface{}]
	parser parser.Parser
}

func NewApplication() *App {
	return &App{
		PORT:    "8989",
		VERSION: "v0.1.0",
		store:   store.GetInstance(),
		parser:  parser.NewRespParser(),
	}
}

func (c *App) Banner() {
	fmt.Printf("%s", LOGO)
	log.Printf("Rover Server running")
	log.Printf("Port:\t\t%s", c.PORT)
	log.Printf("Documentation:\thttps://rover.subrotokumar.dev")
	log.Printf("Source code:\thttps://github.com/subrotokumar/rover")
}
