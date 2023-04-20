package main

import (
	"fmt"
	. "github.com/angelmotta/hello-bro/internal/config"
	"github.com/angelmotta/hello-bro/roles/client"
	"github.com/angelmotta/hello-bro/roles/server"
	"log"
	"net"
	"time"
)

func main() {
	log.Println("*** Hello bro started ***")

	// Load Global Configuration
	GlobalConf.Load()
	if GlobalConf.Role == "svr" {
		log.Println("Hello-Bro Server role")
		s := server.NewServer("192.168.1.2:2000")
		log.Println("Hello-Bro Server started")
		log.Println("Waiting for clients...")
		time.Sleep(time.Second * 20)
		log.Println("Hello-Bro server will try to stop")
		s.Stop()
		log.Println("Hello-Bro server stopped")
	} else if GlobalConf.Role == "cli" {
		log.Println("Hello-Bro Client role")
		// Start client
		c := client.NewClient("192.168.1.2:2000")
		log.Println("Hello-Bro Client started")
		c.SendRequest()
		// Client close connection
		c.CloseConn()
	} else {
		log.Println("Looks like no environment variables were given in the input")
	}
	log.Println("*** Hello bro finished ***")
}

// startFastClient starts locally a simple client.
// Sample of first approach to test from the main function how to shut down the server
func startFastClient(svrAddr net.Addr) {
	log.Println("start FastClient...")
	// Dial connects to the address specified
	conn, err := net.Dial(svrAddr.Network(), svrAddr.String()) // Dial("tcp", "<ip:port>")
	if err != nil {
		log.Fatal("error net.Dial-op", err)
	}
	defer conn.Close()
	// Send message to the server
	n, err := fmt.Fprintf(conn, "hello bro")
	if err != nil {
		log.Fatal("error client write-op", err)
	}
	log.Printf("FastClient wrote %d bytes", n)
}
