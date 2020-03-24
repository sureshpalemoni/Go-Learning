package slices

import "fmt"

func Append() {
	var s []int
	fmt.Println(s, "- s is nil: ", s==nil)

	a := [...]int{1,2,3,4,5,6,7,8,9}
	s = a[2:4]
	fmt.Println("Array a is: ", a)
	fmt.Println(s, "- s is nil: ", s==nil)

	newS := append(s, 13, 14)
	fmt.Println(a, s, newS) // Value of original slice unchanged but values of arrays gets changed


	newSS := append(s, 13, 14, 15, 16, 17, 18, 19, 20)
	fmt.Println(a, s, newSS) // Value of original slice unchanged but values of arrays gets changed	
	fmt.Println(len(newSS), cap(newSS))

	// Ananymous Slice
	AnS := []int{1,2,3,4}
	fmt.Println(len(AnS), cap(AnS))
	AnS = append(AnS, 7,5)
	fmt.Println(len(AnS), cap(AnS))

	// When the length is increased more than original arrray length, capacity is doubled
	AnS = append(AnS, 8,6,9)
	fmt.Println(len(AnS), cap(AnS))
}