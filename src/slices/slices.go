package slices

import "fmt"

func SlicesNew() {
	var s []int
	fmt.Println(s, "- s is nil: ", s==nil)

	a := [...]int{1,2,3,4,5,6,7,8,9}
	s = a[2:4]
	fmt.Println("Array a is: ", a)
	fmt.Println(s, "- s is nil: ", s==nil)
	fmt.Println("_____________")

	a[2] = 13
	a[3] = 14
	fmt.Println("Array a is: ", a)
	fmt.Println(s, "- s is nil: ", s==nil) //S is a reference to array a
	fmt.Println("_____________")


	//Length and capacity
	fmt.Println(cap(s), len(s)) 
	// Length is two since its refering only 2:4, capacity is 7 becuase starts from 2 and to the end 

	s[0] = 3
	fmt.Println("Array a is: ", a)
	fmt.Println(s, "- s is nil: ", s==nil)
 }