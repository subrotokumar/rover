package app

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/subrotokumar/rover/internal/executor"
)

func Serve() {
	dir := flag.String("dir", "", "The directory where RDB files are stored")
	dbfilename := flag.String("dbfilename", "", "The name of the RDB file")
	flag.Parse()
	fmt.Println(*dir, *dbfilename)
	app := NewApplication(dir, dbfilename)
	app.Banner()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", app.PORT))
	if err != nil {
		log.Fatalf("Failed to bind to port %s: %v", app.PORT, err)
	}
	defer listener.Close()

	log.Println("* Ready to accept connections tcp")
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
	db := 0
	for {
		n, err := conn.Read(buf)
		start := time.Now()
		if err != nil {
			if err == io.EOF {
				log.Printf("- ERR Client %s disconnected", conn.RemoteAddr())
				break
			}
			log.Printf("- ERR reading from client: %v", err)
			break
		}
		parsedCmd, err := app.parser.Parse(buf[:n])
		if err != nil {
			conn.Write([]byte("- ERR Unable to parse input: %v"))
			continue
		}

		cmd, ok := parsedCmd.([]string)
		if !ok {
			conn.Write([]byte("- ERR Failed to assert the parsed command"))
			continue
		}
		if isSystemFunction(&cmd) {
			if app.HandleSystemCmd(cmd, &db, conn) {
				continue
			}
		}

		response := executor.Execute(db, cmd)
		conn.Write([]byte(response))
		if app.DebugMode {
			log.Printf("* %v : %v", cmd, time.Since(start))
		}
	}
}

var list = []string{"select", "config"}

func isSystemFunction(cmd *[]string) bool {
	for _, val := range list {
		if strings.ToLower((*cmd)[0]) == val {
			return true
		}
	}
	return false
}

func (app *App) HandleSystemCmd(cmd []string, db *int, conn net.Conn) bool {
	if strings.ToUpper(cmd[0]) == "SELECT" && len(cmd) == 2 {
		n, err := strconv.Atoi(cmd[1])
		if err != nil {
			return true
		} else if n < 0 || n > 15 {
			conn.Write([]byte("-ERR DB index is out of range\r\n"))
		}
		*db = n
		conn.Write([]byte("+OK\r\n"))
	} else if strings.ToUpper(cmd[0]) == "CONFIG" && strings.ToUpper(cmd[1]) == "GET" && len(cmd) == 3 {
		if cmd[2] == "dir" {
			conn.Write([]byte(fmt.Sprintf("*2\r\n$3\r\ndir\r\n$%d\r\n%s\r\n", len(app.dir), app.dir)))
		} else if cmd[2] == "dbfilename" {
			conn.Write([]byte(fmt.Sprintf("*2\r\n$10\r\ndbfilename\r\n$%d\r\n%s\r\n", len(app.dbfilename), app.dbfilename)))
		} else {
			conn.Write([]byte("-ERR unsupported CONFIG parameter\r\n"))
		}
		return true
	}
	return false
}
