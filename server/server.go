package server

import (
	"log"
	"net"
	"sync"
)

type Server struct {
	listener net.Listener
	quitChan chan interface{} // channel that's used to signal shutdown
	wg       sync.WaitGroup   // WaitGroup to wait until all the server's goroutines are actually done.
}

// NewServer works a Server's constructor
func NewServer(addr string) *Server {
	s := &Server{
		quitChan: make(chan interface{}),
	}
	// Initialize and assign listener in Server
	l, err := net.Listen("tcp", addr) // addr: "<ip>:<port>"
	if err != nil {
		log.Fatal(err)
	}

	s.listener = l
	s.wg.Add(1)
	// Start main goroutine for the new Server
	go s.serve() // Server listens for new connections in a background goroutine
	return s
}
func (s *Server) serve() {
	defer s.wg.Done()

	for {
		// Wait for a connection
		log.Println("Waiting for a new connection...")
		// Listener's Accept waits for and return the next connection to the listener
		conn, err := s.listener.Accept() // blocking operation
		if err != nil {
			// check quitChan channel in a non-blocking way
			select {
			case <-s.quitChan: // this means the error is caused by the Stop() method
				return // With this return serve() notifies the WaitGroup that it's done
			default:
				log.Println("listener.Accept() error", err)
			}
		} else {
			log.Println("New connection accepted! Let's handle connection!")
			s.wg.Add(1)
			go func() {
				s.handleConnection(conn)
				s.wg.Done()
			}()
		}
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()
	// TODO
}

// Stop tells the server to shut down gracefully: until all the handlers have returned
func (s *Server) Stop() {
	close(s.quitChan) // any subsequent receive from a closed channel (<-s.quitChan) will succeed
	// Stop accepting new clients.
	s.listener.Close() // This will cause the listener.Accept() throws an error and Serve method return quietly
	s.wg.Wait()        // This operation will block until all the handlers have returned
}
