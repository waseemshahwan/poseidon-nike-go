package main

import (
	"fmt"
	"net"
	"os"
)

const (
    CONN_HOST = "localhost"
    CONN_PORT = "3333"
    CONN_TYPE = "tcp"
)

func main() {
    // Listen for incoming connections.
    l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
    if err != nil {
        fmt.Println("Error listening:", err.Error())
        os.Exit(1)
    }
    // Close the listener when the application closes.
    defer l.Close()
    fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
    for {
        // Listen for an incoming connection.
        conn, err := l.Accept()
        if err != nil {
            fmt.Println("Error accepting: ", err.Error())
            os.Exit(1)
        }
        // Handle connections in a new goroutine.
        go handleRequest(conn)
    }
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
  
	payload := make(chan []byte)
	exit := make(chan bool)
	
	go func(payload chan []byte, exit chan bool) {
		// Signal an exit when the enclosing function is done
		defer func () { exit <- true }()

		buf := make([]byte, 2048)
		for true {
			read, err := conn.Read(buf)
			if err != nil {
				println(err)
				break
			}

			if read > 0 {
				exit <- false
				payload <- buf[:read]
			}
		}
	}(payload, exit)

	active := false
	for <-exit == false {
		data := <-payload
		print(string(data))
	}

	conn.Close()
}