package functions

import "fmt"

func deferfunc() {
	fmt.Println("Defer func program")
}

// Since there is return in the parent function
// defer is executed when it hits return statement
func squares(i int) (square int) {
	square = i * i

	defer deferfunc()

	return
}

// Here parent function doesn't return anything, 
// defer is executed at the end of the function
func DeferIntro() {
	fmt.Println("Main program started")

	defer deferfunc()

	fmt.Println((squares(2)))

	fmt.Println("Main program stopped")
}