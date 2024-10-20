package app

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/subrotokumar/rover/internal/executor"
)

func Serve() {
	app := NewApplication()
	app.Banner()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", app.PORT))
	if err != nil {
		log.Fatalf("Failed to bind to port %s: %v", app.PORT, err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		go app.handleConnection(conn)
	}
}

func (app *App) handleConnection(conn net.Conn) {
	defer conn.Close()
	log.Printf("New connection from %s", conn.RemoteAddr())

	app.handleRequest(conn)
}

func (app *App) handleRequest(conn net.Conn) {
	buf := make([]byte, 1024)
	executor := executor.NewExecutor(conn)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				log.Printf("Client %s disconnected", conn.RemoteAddr())
				break
			}
			log.Printf("Error reading from client: %v", err)
			break
		}

		parsedCmd, err := app.parser.Parse(buf[:n])
		if err != nil {
			log.Printf("Unable to parse: %v", err)
			continue
		}

		cmd, ok := parsedCmd.([]string)
		if !ok {
			log.Println("Failed to assert the parsed command to []string")
			continue
		}

		log.Printf("Received command => %v", cmd)
		executor.Execute(cmd)
	}
}
