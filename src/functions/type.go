// Functions can act as a type
// Used when a function is passed as an argument to another function
// Returning function from the other function.

package functions

import "fmt"

func add(a, b int) int {
	return a + b
} 

func subtract(a, b int) int {
	return a - b
}

// Passing function as third argument, which is of same type int
func calc(a, b int, f func(int, int) int) int {
	r := f(a, b)
	return r
}


// The above function can also be written as 
type calcFunc func(int, int) int
func calculation(a, b int, f calcFunc) int {
	r := f(a, b)
	return r
}



func FuncTypes() {
	// This is to show the type of function.
	fmt.Printf("The function type of add is %T\n", add)
	fmt.Printf("The function type of subtract is %T\n", subtract)

	// passing function as an argument.
	addResult := calc(5, 3, add)
	subtractResult := calc(5, 3, subtract)
	fmt.Println("Add Result is: ", addResult)
	fmt.Println("Subtract result is: ", subtractResult)

	// Created function as a derived type : line 24
	addResultC := calculation(15, 3, add)
	subtractResultC := calculation(15, 3, subtract)
	fmt.Println("Add Result is: ", addResultC)
	fmt.Println("Subtract result is: ", subtractResultC)
}

