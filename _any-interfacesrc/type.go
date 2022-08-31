//go:build ignore
// +build ignore

package main()

import (
	"fmt"

)

// TYPICAL_FUNC_START OMIT
func main() {
	printAny((*int)(nil))
	printAny[*int](nil)
	printInterface((*int)(nil))
}
//TYPICAL_FUNC_END OMIT

func printInterface(foo interface{}) {
	fmt.Printf("%v\n", foo)
}

func printAny[T any](foo T) {
	fmt.Printf("%v\n", foo)
}
