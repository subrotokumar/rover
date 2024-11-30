package app

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

var list = []string{"select", "config", "keys"}

func isSystemFunction(cmd *[]string) bool {
	for _, val := range list {
		if strings.ToLower((*cmd)[0]) == val {
			return true
		}
	}
	return false
}

func (app *App) HandleSystemCmd(cmd []string, db *int, conn net.Conn) bool {
	switch strings.ToUpper(cmd[0]) {
	case "SELECT":
		if len(cmd) != 2 {
			return false
		}
		n, err := strconv.Atoi(cmd[1])
		if err != nil {
			return true
		} else if n < 0 || n > 15 {
			conn.Write([]byte("-ERR DB index is out of range\r\n"))
		}
		*db = n
		conn.Write([]byte("+OK\r\n"))
		return true
	case "CONFIG":
		if strings.ToUpper(cmd[1]) != "GET" || len(cmd) != 3 {
			return false
		}
		if cmd[2] == "dir" {
			conn.Write([]byte(fmt.Sprintf("*2\r\n$3\r\ndir\r\n$%d\r\n%s\r\n", len(app.dir), app.dir)))
		} else if cmd[2] == "dbfilename" {
			conn.Write([]byte(fmt.Sprintf("*2\r\n$10\r\ndbfilename\r\n$%d\r\n%s\r\n", len(app.dbfilename), app.dbfilename)))
		} else {
			conn.Write([]byte("-ERR unsupported CONFIG parameter\r\n"))
		}
		return true
	case "KEYS":
		if cmd[1] != "*" {
			return false
		}
		fileContent := readFile(app.dir + "/" + app.dbfilename)
		write := []byte(fmt.Sprintf("*1\r\n$%d\r\n%s\r\n", len(fileContent), fileContent))
		conn.Write(write)
		return true
	default:
		return false
	}
}
