package pointers

import "fmt"

func changeArrayValue(p *[3]int) {
	//*p == original array `a`
	// *p[0] != (*p)[0]
	(*p)[0] *= 2 //multiply the value with 2
	(*p)[1] *= 3

	//or

	p[2] *= 4
}

func CompositeData() {
	a := [3]int{1, 2, 3}
	changeArrayValue(&a)
	
	fmt.Printf("a = %v\n", a)
}
