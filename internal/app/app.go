package app

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/subrotokumar/rover/internal/parser"
	"github.com/subrotokumar/rover/internal/store"
	"github.com/subrotokumar/rover/internal/types"
)

type App struct {
	PORT       string
	VERSION    string
	DebugMode  bool
	store      *store.SafeMap[string, types.StoredValue]
	parser     parser.Parser
	dir        string
	dbfilename string
}

func NewApplication(dir, dbfilename *string) *App {
	mode := os.Getenv("DEBUG_MODE")
	debug_mode := false
	if strings.ToLower(mode) == "true" {
		debug_mode = true
	}
	return &App{
		PORT:       "8989",
		DebugMode:  debug_mode,
		VERSION:    "v0.1.0",
		store:      store.GetInstance(),
		parser:     parser.NewRespParser(),
		dir:        *dir,
		dbfilename: *dbfilename,
	}
}

func (c *App) Banner() {
	fmt.Printf("%s", LOGO)
	log.Printf("* Rover Server running")
	log.Printf("* Port:\t\t%s", c.PORT)
	log.Printf("* Documentation:\thttps://rover.subrotokumar.dev")
	log.Printf("* Source code:\thttps://github.com/subrotokumar/rover")
}
