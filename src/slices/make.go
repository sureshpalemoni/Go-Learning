package slices

import "fmt"


// make is a built-in function that helps you create an empty slice.
// The make function can create many empty composite types.
// func make(t Type, size ...IntegerType) Type

func MakeSlice() {
	s1 := make([]int, 2, 8)
	s2 := []int{1,2,3}

	fmt.Printf("before => s1=%v, s2=%v\n", s1, s2)
	fmt.Println("before => s1 == nil", s1 == nil)

	n1 := copy(s1, s2)
	fmt.Printf("after => n1=%d, s1=%v, s2=%v\n", n1, s1, s2)
	fmt.Println("after => s1 == nil", s1 == nil)
}