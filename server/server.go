package server

import (
    "bufio"
    "fmt"
    "net"
    "strings"
)

func ParseRequest(data string) (string, []string) {
    parts := strings.Fields(data)
    if len(parts) == 0 {
        return "", nil
    }
    command := strings.ToUpper(parts[0])
    args := parts[1:]
    return command, args
}

func HandleConnection(conn net.Conn) {
    defer conn.Close()
    
    reader := bufio.NewReader(conn)
    for {
        data, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("Connection closed")
            return
        }

        data = strings.TrimSpace(data)

    }
}

