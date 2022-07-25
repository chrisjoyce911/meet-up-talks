package main

import (
	"_improved/strawman"
)

// START OMIT
func main() {
	myserver := strawman.NewServer(addr, clientTimeout, 0, 0, nil) // HL
	myserver.Shutdown()
}

// END OMIT
