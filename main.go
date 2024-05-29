package main

import (
	"fmt"
	"net"
	"os"

	"github.com/sinnlos-ffff/redis_clone/pkg/event_loop"
)

func main() {
	// 6379: redis port number
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	eventLoop := event_loop.NewEventLoop(10)

	eventLoop.RegisterHandler("ping", func(event event_loop.Event) {
		_, err := event.Conn.Write([]byte("+PONG\r\n"))
		if err != nil {
			fmt.Println("Error writing on connection: ", err.Error())
			os.Exit(1)
		}
	})

	eventLoop.Start()

	for {
		con, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		defer con.Close()

		input := make([]byte, 128)
		_, err = con.Read(input)

		if err != nil {
			fmt.Println("Error reading on connection: ", err.Error())
			continue
		}

		eventLoop.PostEvent(event_loop.Event{Name: "ping", Conn: con})
	}
}
