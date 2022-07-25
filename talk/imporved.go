package strawman

import (
	"crypto/tls"
	"net"
	"time"
)

type Server struct {
	listener net.Listener
}

func (s *server) Addr() new.Addr
func (s *server) Shutdown()

// START OMIT
// NewServer returns a new Server listening on addr.
// clientTimeout defines the max length of an idle connection
// maxcons limites the number of concurrent connections
// maxconcurrent limites the number of concurrent connections from a single IP
// cert is the TLS cerfificate for the connection
func NewServer(addr String, clientTimeout time.Duration, maxcons, maxconcurrent int, // HL
	cert *tls.Cert) (*Server, error) // HL
	{
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	svr := Server{listener: l}
	go srv.run()
	return &srv, nil
}

// END OMIT
