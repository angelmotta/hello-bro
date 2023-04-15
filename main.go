package main

import (
	"github.com/angelmotta/hello-bro/server"
	"log"
)

func main() {
	log.Println("*** Hello bro started ***")

	s := server.NewServer(":2000")

	// TODO: start clients
	// go startClient()

	s.Stop()
}
