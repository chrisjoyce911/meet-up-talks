// START OMIT
type Config struct {
	// the port to listen on, if unset defaults to 8080
	Port int
}

func main() {
	svr, _ := NewServer("localhost",
		Config{
			Port: 0, // opps I can't do this
		},
	)
}

// END OMIT
