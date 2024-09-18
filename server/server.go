package server

import (
    "bufio"
    "log"
    "fmt"
    "net"
    "strings"
)



var db_cache = NewStorageCache()

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

        db_cache.Set(args[0] , args[1]);
        return "\nOK\r\n"

    case "GET":
        
        if len(args) != 1 {
            return "Wrong number of arguments for set command"
        }

        value, found_value := db_cache.Get(args[0])

        if found_value {
            return fmt.Sprintf("\r\n%s\r\n", value)
        }

        return fmt.Sprintf("\n-1\r\n")


    default:
        return "Unknown command"
    }
}


