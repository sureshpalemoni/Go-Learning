package pointers

import "fmt"

func Intro() {
	var pa *int
	// A pointer has a nil value because it is not pointing to any data (value) in the RAM at the moment.
	fmt.Printf("Pointer pa of type %T and value %v\n", pa, pa)

	a := 1
	pb := &a
	fmt.Printf("Pointer pb of type %T and value %v\n", pb, pb)

	// Dereferencing, you have to prefix "*" to the pointer.
	b := 2
	pc := &b
	fmt.Printf("Pointer pc of type %T and value %v\n", pc , *pc)

	// Changing variable value using pointers
	d := 3
	fmt.Println("Value of d before change is: ", d)
	pd := &d
	*pd = 4
	fmt.Println("Value of d after change is: ", d)
	fmt.Printf("Pointer pd of type %T and value %v\n", pd , *pd)

	// new function
	// built-in function which returns pointer with address
	// it will have underlying value as 0. syntax: func new(Type) *Type
	pn := new(int)
	fmt.Printf("Pointer pn of type %T and value %v\n", pn , *pn)
}