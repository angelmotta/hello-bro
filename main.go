package main

import (
	"io"
	"log"
	"net"
)

func main() {
	log.Println("*** Hello bro started ***")

	l, err := net.Listen("tcp", ":2000")
	if err != nil {
		log.Fatal(err)
	}

	defer l.Close()
	for {
		// Wait for a connection
		log.Println("Waiting for a new connection...")
		conn, err := l.Accept() // Accept Waits for and return the next connection to the listener (blocking operation)
		log.Println("New connection accepted!")
		if err != nil {
			log.Fatal(err)
		}

		go func(c net.Conn) {
			// Echo all incoming Data as a response for the client
			// io.Copy copies from src to dst until EOF is reached on src (e.g. Until 'quit' is sent via telnet)
			if _, err := io.Copy(c, c); err != nil {
				log.Fatal("issue with io.Copy echo", err)
			}
			// Shut down the connection
			log.Println("...client disconnected")
			c.Close()
		}(conn)
	}
}
