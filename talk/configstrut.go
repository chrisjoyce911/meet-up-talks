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
// A Config structure isused to configure the Server.
type Config struct {
	// Timeout sets theamount of time befirs closing
	// idle connections, or forever if not provided.
	Timeout time.Duration

	// The server will accept TLS connections if
	// the certificate is provided
	Cert *tls.Cert
}

// NewServer returns a Server listening on addr
 NewServer(addr String, config Config) (*Server, error) // HL
// END OMIT
