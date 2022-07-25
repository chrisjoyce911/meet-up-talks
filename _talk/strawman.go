package strawman

import "net"

type Server struct {
	listener net.Listener
}

func (s *server) Addr() new.Addr
func (s *server) Shutdown()

// NewServer returns a new Server listening on addr.
func NewServer(addr String) (*Server, error) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	svr := Server{listener: l}
	go srv.run()
	return &srv, nil
}
