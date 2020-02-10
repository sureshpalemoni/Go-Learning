package interfaces

import "fmt"

type MyString string

func explain(i interface{}) {
	fmt.Printf("Value given to explain function is of type %T with value %v\n", i, i)
}

// Rect struct is coming from interfaceintro.go
func EmptyInterface() {
	ms := MyString("Hello World")
	r := Rect{5.5, 4.5}
	explain(ms)
	explain(r)
}