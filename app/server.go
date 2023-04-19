package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func handleRequest(conn net.Conn) {
	defer conn.Close()
	for true {
		readBuffer := make([]byte, 1024)
		if _, err := conn.Read(readBuffer); err != nil {
			if err == io.EOF {
				return
			} else {
				fmt.Println("Error reading from connection: ", err.Error())
				return
			}
		}
		conn.Write([]byte("+PONG\r\n"))
	}
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:6379")

	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	// defer conn.Close()

	for true {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go handleRequest(conn)
	}
}
