package talk

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
// NewServer returns a Server listening on addr
 NewServer(addr String) (*Server, error) // HL

 // NewTLSServer returns a secure Server listening on addr
 NewTLSServer(addr String, cert *tls.Cert) (*Server, error) // HL

 // NewServerWithTimeout returns a Server listening on addr that disconnects idel clients
 NewServerWithTimeout(addr String, clientTimeout time.Duration, maxcons, maxconcurrent int) // HL
  (*Server, error) // HL

// NewTLSSServerWithTimeout returns a secure Server listening on addr 
// that disconnects idel clients
 NewServerWithTimeout(addr String, clientTimeout time.Duration, maxcons, // HL
	 maxconcurrent int, cert *tls.Cert) (*Server, error) // HL
// END OMIT
