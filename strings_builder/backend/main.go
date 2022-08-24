package main

import (
	"fmt"
	"strings"
)

func main() {
	var s = []string{"one", "two", "three", "four", "five"}
	fmt.Printf("\n%q\n", formatString(s))
}

// START OMIT
func formatString(s []string) string {
	var sb strings.Builder
	fmt.Printf("Inital %16s Len : %3d, Cap : %3d\n", " ", sb.Len(), sb.Cap()) // HL
	for _, str := range s {
		sb.WriteString(str)
		fmt.Printf("%5s String Len : %3d, Len : %3d, Cap : %3d\n", // HL
			str, len(str), sb.Len(), sb.Cap()) // HL
	}
	return sb.String()
}

//END OMIT
