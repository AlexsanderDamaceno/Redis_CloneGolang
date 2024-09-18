package server

import (
    "bufio"
    "log"
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
            log.Println("Connection closed")
            return
        }

        command, args := ParseRequest(data)  

        response := handleCommand(command, args)

        data = strings.TrimSpace(data)

        conn.Write([]byte(response + "\n"))

    }
}

func handleCommand(command string, args []string) string {

    switch command {

    case "SET":

        if len(args) != 2 {
            return "Wrong number of arguments for set command"
        }

        return "Command successful executed!"

    default:
        return "Unknown command"
    }
}


