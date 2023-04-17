package client

import (
	"fmt"
	"log"
	"net"
)

type Client struct {
	ServerAddr string
	Conn       *net.Conn
}

func NewClient(svrAddr string) *Client {
	c := &Client{
		ServerAddr: svrAddr,
	}
	// Connect to Server
	// Dial connects to the address specified
	conn, err := net.Dial("tcp", svrAddr) // Dial("tcp", "<ip:port>")
	if err != nil {
		log.Fatal("error net.Dial-op", err)
	}
	c.Conn = &conn
	return c
}

// SendAndBye sends a message to the server and close the tcp connection
func (c *Client) SendAndBye() {
	defer (*c.Conn).Close()
	n, err := fmt.Fprintf(*c.Conn, "hello bro")
	if err != nil {
		log.Fatal("error client write-op", err)
	}
	log.Printf("FastClient wrote %d bytes", n)
}
