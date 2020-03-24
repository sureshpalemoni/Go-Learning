package pointers

import "fmt"

func changeValue(p *int) {
	*p = *p + 10
} 

func PassingToFunction() {
	a := 2
	//you can pass two ways
	// By create a pointer and pass it
	pa := &a
	changeValue(pa)
	fmt.Println("value of a is: ", a)

	// By directly passing the address of variable
	changeValue(&a)
	fmt.Println("value of a is: ", a)

}