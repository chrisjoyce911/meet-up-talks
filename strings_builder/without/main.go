package main

import "fmt"

// START OMIT
func main() {
	var s = []string{"one", "two", "three", "four", "five"}
	fmt.Printf("%q\n", formatString(s))
}

func formatString(s []string) (ret string) {
	for _, str := range s {
		ret += str
	}
	return
}

//END OMIT
