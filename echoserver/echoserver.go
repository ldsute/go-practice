package main

import (
	"io"
	"log"
	"net"
)

func main() {
	// Listen on port 2000 on all available addresses of the local system
	l, err := net.Listen("tcp", ":2000")
	if err != nil {
		log.Fatal(err)
	}
	// Close the connection when exiting the main() function
	defer l.Close()

	// This is an infinite loop
	for {
		// Wait for a connection
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		// Handle the connection in a new goroutine
		go func(c net.Conn) {
			// Echo the incoming data
			// The first return value is the byteCount copied, but we won't use (read) that value
			// so we will use the "blank identifier" to "ignore" it. Go does not allow defining a
			// variable that will never be read
			// "_ is a special identifier you can assign anything to but never read from"
			_, err := io.Copy(c, c)
			if err != nil {
				log.Println("echo error: " + err.Error())
			}

			// Close the connection
			c.Close()
		}(conn)
	}
}
