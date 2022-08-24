// START OMIT
type Config struct {
	MaxConCurrent *Int // MaxConCurrent the number of concurrent connections
	Prefix *string // Logging Prefix
}
 func NewServer(addr String, config Config) (*Server, error) // HL
 {
	l, err := newServerWithDefaults()
	if err != nil {
		return nil, err
	}
	svr := Server{listener: l}

    if config.Prefix == nil {
        svr.prefix = opt.Prefix
    }
    if option.MaxConCurrent == nil {
        svr.maxconcurrent = opt.MaxConCurrent
    }
	go srv.run()
	return &srv, nil
 }
// END OMIT
