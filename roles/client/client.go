package client

import (
	"fmt"
	"log"
	"net"
)

type Client struct {
	ServerAddr string
	conn       *net.Conn
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
	c.conn = &conn
	return c
}

// SendRequest sends a message to the server
func (c *Client) SendRequest() {
	//defer (*c.Conn).Close()
	// Send message to the Server's connection
	reqMessage := "hello bro"
	n, err := fmt.Fprintf(*c.conn, reqMessage)
	if err != nil {
		log.Fatal("error client write-op request", err)
	}
	log.Printf("Client sent request with %d bytes", n)

	// Receive message from the server
	buf := make([]byte, 2048)
	_, err = (*c.conn).Read(buf)
	if err != nil {
		log.Fatal("error client read-op response", err)
	}

	resMessage := string(buf)
	log.Printf("Client received response: %s", resMessage)
}

// CloseConn closes the client's connection with the server
func (c *Client) CloseConn() {
	log.Println("Client is closing connection with the server")
	err := (*c.conn).Close()
	if err != nil {
		log.Fatal("Something went wrong. Client CloseConnection:", err)
	}
}

func (c *Client) DoSimpleWorkload() {
	c.SendRequest()
	c.CloseConn()
}
