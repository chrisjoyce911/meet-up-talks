package main

import "fmt"

// MAIN_START OMIT
func main() {
	fmt.Printf("%T\t%+v\n", "Hello, Println", "Hello, Println")
	printInterface("Hello, Interface")
	printAny("Hello, Any")
	printAnyWithGenerics("Hello, Generics")

	fmt.Printf("%T\t%+v\n", nil, nil)
	printInterface(nil)
	printAny(nil)
	printAnyWithGenerics[any](nil)
}

// MAIN_END OMIT

// FUNCTIONS_START OMIT
// SUITABLE_START OMIT
func printInterface(foo interface{}) {
	// NOTGENERIC_START OMIT
	fmt.Printf("%T\t%+v\n", foo, foo)
	// NOTGENERIC_END OMIT
}

func printAny(foo any) {
	fmt.Printf("%T\t%v\n", foo, foo)
}

// SUITABLE_END OMIT
func printAnyWithGenerics[T any](foo T) {
	fmt.Printf("%T\t%v\n", foo, foo)
}

// FUNCTIONS_END OMIT
