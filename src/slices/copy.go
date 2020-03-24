package slices

import "fmt"

func Copy() {
	var s []int
	s2 := []int{5,6,7}
	fmt.Println(s, "- s is nil: ", s==nil)

	a := [...]int{1,2,3,4,5,6,7,8,9}
	s = a[2:6]
	fmt.Println("Array a is: ", a)
	fmt.Println(s, "- s is nil: ", s==nil)

	// Copying elements from s1 to s2
	newS := copy(s2, s) // returns number of elements copied or replaces
	fmt.Println(a, s, newS) 

}