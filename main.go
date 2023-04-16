package main

import (
	"github.com/angelmotta/hello-bro/server"
	"log"
	"time"
)

func main() {
	log.Println("*** Hello bro started ***")

	s := server.NewServer(":2000")

	// TODO: start clients
	// go startClient()

	// Test outside this program before stop the server
	time.Sleep(time.Minute * 10)
	log.Println("Calling Server's Stop method")
	s.Stop()
}
