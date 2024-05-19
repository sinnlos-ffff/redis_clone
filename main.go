package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// 6379: redis port number
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	con, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	defer con.Close()

	for {
		input := make([]byte, 1024)
		_, err := con.Read(input)
		if err != nil {
			fmt.Println("Error reading on connection: ", err.Error())
			os.Exit(1)
		}

		con.Write([]byte("+PONG\r\n"))
		if err != nil {
			fmt.Println("Error writing on connection: ", err.Error())
			os.Exit(1)
		}
	}
}
