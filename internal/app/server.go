package app

import (
	"fmt"
	"io"
	"net"
	"os"

	"github.com/subrotokumar/rover/internal/config"
)

func Serve() {
	cfg := config.NewDefaultConfig()
	cfg.Banner()
	l, err := net.Listen("tcp", "0.0.0.0:"+cfg.PORT)
	if err != nil {
		printf("Failed to bind to port %s", cfg.PORT)
		os.Exit(1)
	}

	c, err := l.Accept()
	cfg.Banner()
	if err != nil {
		println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	defer c.Close()
	println("Sdsdd")
	cfg.Banner()
	app := NewApplication(c)
	app.handleRequest(c)
}

func (app *App) handleRequest(c net.Conn) {
	for {
		buf := make([]byte, 128)
		_, err := c.Read(buf)

		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("error reading from client: ", err.Error())
			os.Exit(1)
		}
		parsedCmd, err := app.parser.Parse(buf)
		if err != nil {
			println("Unable to parse ", err)
		}
		cmd, ok := parsedCmd.([]string)
		if !ok {
			fmt.Println("Failed to assert the parsed command to []string")
			return
		} else {
			printf("CMD => %v\n", cmd)
		}
		app.ExecuteCmd(cmd)
	}
}

func (app *App) ExecuteCmd(cmd []string) {
	err := app.executor.Execute(cmd)
	if err != nil {
		printf("%s", err.Error())
	}
}
