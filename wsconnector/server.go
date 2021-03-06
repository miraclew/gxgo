package push

import (
	"log"
)

type Server struct {
	clients map[int64]*Client
	addCh   chan *Client
	delCh   chan *Client
	doneCh chan bool
	errCh  chan error
}

var server *Server

func NewServer() *Server {
	if server == nil {
		clients := make(map[int64]*Client)
		addCh := make(chan *Client)
		delCh := make(chan *Client)
		doneCh := make(chan bool)
		errCh := make(chan error)

		server = &Server{
			clients,
			addCh,
			delCh,
			doneCh,
			errCh,
		}
	}

	return server
}

func (s *Server) Add(c *Client) {
	s.addCh <- c
}

func (s *Server) Del(c *Client) {
	s.delCh <- c
}

func (s *Server) Done() {
	s.doneCh <- true
}

func (s *Server) Err(err error) {
	s.errCh <- err
}

// Listen and serve.
// It serves client connection and broadcast request.
func (s *Server) Listen() {

	log.Println("Listening server...")

	for {
		select {

		// Add new a client
		case c := <-s.addCh:
			log.Println("Added new client")
			s.clients[c.id] = c
			log.Println("Now", len(s.clients), "clients connected.")

		// del a client
		case c := <-s.delCh:
			log.Println("Delete client")
			delete(s.clients, c.id)

		case err := <-s.errCh:
			log.Println("Error:", err.Error())

		case <-s.doneCh:
			return
		}
	}
}
