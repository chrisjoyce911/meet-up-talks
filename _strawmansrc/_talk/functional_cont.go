// START OMIT
// NewServer returns a new Server listening on addr.
func NewServer(addr String, options ...func(*Server)) (*Server, error) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	svr := Server{listener: l}

	for _, option := range options { // HL
		option(&srv) // HL
	} // HL
	return &srv, nil
}

// END OMIT