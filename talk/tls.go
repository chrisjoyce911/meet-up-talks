package main

import (
	"_improved/strawman"
)

// START OMIT
func main() {
	myserver := strawman.NewServer(addr, clientTimeout, maxcons, maxconcurrent, nil) // HL
	myserver.Shutdown()
}

// END OMIT
