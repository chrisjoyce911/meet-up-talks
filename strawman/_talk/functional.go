// START OMIT
func NewServer(addr String, options ...func(*Server)) (*Server, error)

func main() {
	svr, _ := NewServer("localhost") // defaults

	timeout := func(srv *Server) {
		srv.timeout = 60 * time.Second
	}

	maxconns := func(srv *Server) {
		srv.maxconns = 10
	}

	// timeout after 60 seconds, 10 clients max
	svr2, _ := NewServer("localhost", timeout, maxconns)
}

// END OMIT