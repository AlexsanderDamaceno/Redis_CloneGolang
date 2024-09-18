package main

import (
     "redis_clone/server"
    "log"
    "net"
	
)

func main() {

    listener, err := net.Listen("tcp", ":1234")

    if err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }

    log.Println("Redis server is running on :1234")
	

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Printf("Failed to accept connection: %v", err)
            continue
        }
		go server.HandleConnection(conn)
    }

}
