// START OMIT
func NewServer(addr String, config ...Config) (*Server, error)

func main() {
	svr, _ := NewServer("localhost") // defaults

	// timeout after 5 minutes, 10 clients max
	svr2, _ := NewServer("localhost", Config{
		Timeout: 300 * time.Second,
		MaxConns: 10,
	})
	....
}
// END OMIT