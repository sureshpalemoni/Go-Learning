package defpanrec

import "fmt"

func defFooStart() {
	fmt.Println("defFooStart() executed")
}

func defFooEnd() {
	fmt.Println("defFooEnd() executed")
}

func defMain() {
	fmt.Println("defMain() executed")
}

func foo() {
	fmt.Println("foo() executed")
	
	defer defFooStart() // defer defFooStart call
	
	panic("panic from foo()")
	
	defer defFooEnd() // defer defFooEnd call
	
	fmt.Println("foo() done")
}

func DeferIntro() {
	fmt.Println("main() started")
	
	defer defMain() // defer defMain call
	
	foo() // call foo function
	
	fmt.Println("main() done")
}
