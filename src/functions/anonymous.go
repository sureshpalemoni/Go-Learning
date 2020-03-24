package functions

import "fmt"

// A function can also go as a value.
// We can assign the function to a variable
var addVar = func(a, b int) int {
	return a + b
}

func AnonymousFunc() {
	fmt.Println("Add Result is: ", addVar(5, 3))


	// Anonymous function - Immediate return.
	// return value is yield and used when variable is passed
	subtract := func(a, b int) int {
		return a - b
	}(5, 3)
	fmt.Println("Add Result is: ", subtract)
}