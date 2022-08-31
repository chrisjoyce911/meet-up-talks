//go:build ignore
// +build ignore

package trash

// TYPICAL_FUNC_START OMIT

func main() {
	fmt.Println("Hello, nil")
	printInterface(nil)
	printAny(nil)
}

// FUNCTIONS_START OMIT
func printInterface(foo interface{}) {
	fmt.Printf("%v\n", foo)
}

func printAny[T any](foo T) {
	fmt.Printf("%v\n", foo)
}

// FUNCTIONS_END OMIT
//TYPICAL_FUNC_END OMIT
