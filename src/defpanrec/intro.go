// This is defer, pani and recover package
// Intro to panic
package defpanrec

import "fmt"

func accessElement(a []int, index int) int{
	if len(a) > index {
		return a[index]
	} else {
		panic("Out of bound condition") // Raising the panic manually
	}
}


func Intro() {
	a := []int{2,3,4}
	fmt.Println(accessElement(a, 2))
	//Accessing out of bound element
	fmt.Println(accessElement(a, 4))
}


