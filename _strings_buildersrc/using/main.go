package main

import (
	"fmt"
	"strings"
)

func main() {
	var s = []string{"one", "two", "three", "four", "five"}
	fmt.Printf("%q\n", formatString(s))
}

// START OMIT
func formatString(s []string) string {
	var sb strings.Builder
	for _, str := range s {
		sb.WriteString(str) // HL
	}
	return sb.String()
}

//END OMIT
