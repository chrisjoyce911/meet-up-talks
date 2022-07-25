// START OMIT
func NewServer(addr String, config *Config) (*Server, error)

func main() {
	svr, _ := NewServer("localhost", nil)

	config := Config{Prot: 9000}
	svr2, _ := NewServer("localhost", &config)

	config.Port = 9001 // what happens now ?
}

// END OMIT