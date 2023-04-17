package main

import (
	"fmt"
	"github.com/angelmotta/hello-bro/roles/client"
	"github.com/angelmotta/hello-bro/roles/server"
	"log"
	"net"
	"time"
)

func main() {
	log.Println("*** Hello bro started ***")

	s := server.NewServer("192.168.1.4:2000")

	// Start client
	//go startFastClient(s.Listener.Addr())
	c := client.NewClient(s.ServerAddr)
	c.SendAndBye()

	// Test outside this program before stop the server
	time.Sleep(time.Minute * 2)
	log.Println("Calling Server's Stop method")
	s.Stop()
	log.Println("*** Hello bro finished ***")
}

// startFastClient starts locally a simple client.
// First approach to test from the main function how to shut down the server
func startFastClient(svrAddr net.Addr) {
	log.Println("start FastClient...")
	// Dial connects to the address specified
	conn, err := net.Dial(svrAddr.Network(), svrAddr.String()) // Dial("tcp", "<ip:port>")
	if err != nil {
		log.Fatal("error net.Dial-op", err)
	}
	defer conn.Close()
	// Send message
	n, err := fmt.Fprintf(conn, "hello bro")
	if err != nil {
		log.Fatal("error client write-op", err)
	}
	log.Printf("FastClient wrote %d bytes", n)
}
