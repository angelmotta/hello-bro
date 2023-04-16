package main

import (
	"fmt"
	"github.com/angelmotta/hello-bro/roles/server"
	"log"
	"net"
	"time"
)

func main() {
	log.Println("*** Hello bro started ***")

	s := server.NewServer(":2000")

	// Start client
	go startFastClient(s.Listener.Addr())

	// Test outside this program before stop the server
	time.Sleep(time.Minute * 2)
	log.Println("Calling Server's Stop method")
	s.Stop()
	log.Println("*** Hello bro finished ***")
}

func startFastClient(svrAddr net.Addr) {
	log.Println("start FastClient...")
	// Dial connects to the address specified
	conn, err := net.Dial(svrAddr.Network(), svrAddr.String()) // Dial("tcp", "<ip:port>")
	if err != nil {
		log.Fatal("error net.Dial-op", err)
	}
	defer conn.Close()

	n, err := fmt.Fprintf(conn, "hello bro")
	if err != nil {
		log.Fatal("error client write-op", err)
	}
	log.Printf("FastClient wrote %d bytes", n)
}
