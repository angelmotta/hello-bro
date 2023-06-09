package server

import (
	"io"
	"log"
	"net"
	"strings"
	"sync"
)

// Server is a TCP Server
type Server struct {
	Listener   net.Listener
	quitChan   chan interface{} // channel that's used to signal shutdown
	wg         sync.WaitGroup   // WaitGroup to wait until all the server's goroutines are actually done.
	ServerAddr string
}

// NewServer function works as Server's constructor
func NewServer(addr string) *Server {
	s := &Server{
		quitChan:   make(chan interface{}),
		ServerAddr: addr,
	}
	// Initialize a listener for the new Server
	l, err := net.Listen("tcp", s.ServerAddr) // addr: "<ip>:<port>"
	if err != nil {
		log.Fatal(err)
	}
	// Assign listener to the new server
	s.Listener = l
	s.wg.Add(1)
	// Start main goroutine for the new Server
	go s.serve() // Server listens for new connections in a background goroutine
	return s
}

// serve method is the main goroutine of a Server (It accepts new connections and handle requests)
func (s *Server) serve() {
	defer s.wg.Done()

	for {
		// Wait for a connection
		log.Println("Waiting for a new connection...")
		// Listener's Accept waits for and return the next connection to the listener
		conn, err := s.Listener.Accept() // blocking operation
		if err != nil {
			// check quitChan channel in a non-blocking way
			select {
			case <-s.quitChan: // this means the error is intentionally caused by the Stop() method
				return // With this return, serve() notifies the WaitGroup that it's done
			default:
				log.Println("Server listener.Accept() error", err)
			}
		} else {
			log.Println("New connection accepted! Let's handle connection!")
			s.wg.Add(1)
			go func() {
				s.handleConnection(conn)
				s.wg.Done()
				log.Println("Server - connection to client closed")
			}()
		}
	}
}

// handleConnection process client requests
func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 2048)
	hello := "Hello bro, did you say: "
	var b strings.Builder
	for {
		log.Println("Waiting for new incoming data...")
		n, err := conn.Read(buf) // block operation
		if err != nil && err != io.EOF {
			log.Println("Server conn.Read error:", err)
		}
		if n == 0 {
			log.Println("handleConnection completed")
			return // handleConnection() method is done
		}
		log.Println("Read-op done")
		// Prepare response for the client
		b.WriteString(hello)
		b.Write(buf)
		dataPayload := b.String()
		response := []byte(dataPayload)
		// Send response & reset buffer
		_, err = conn.Write(response)
		if err != nil {
			log.Fatal("error at con.Write")
		}
		log.Println("data to be send to the client", dataPayload)
		b.Reset()
	}
}

func (s *Server) handleConnectionEcho(c net.Conn) {
	// Echo all incoming Data as a response for the client
	// io.Copy(dst, src) copies from src to dst until EOF is reached on src (e.g. Until 'quit' is sent via telnet)
	if _, err := io.Copy(c, c); err != nil {
		log.Fatal("issue with io.Copy echo", err)
	}
	// Shut down the connection
	c.Close()
	log.Println("...client disconnected")
}

// Stop tells the server to shut down gracefully: until all the handlers have returned
func (s *Server) Stop() {
	log.Println("Server is stopping the service...")
	close(s.quitChan) // any subsequent receive from a closed channel (<-s.quitChan) will succeed
	// Stop accepting new clients.
	s.Listener.Close() // This will cause the listener.Accept() throws an error and Serve method return quietly
	s.wg.Wait()        // This operation will block until all the handlers have returned
}
