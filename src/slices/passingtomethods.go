package slices

import "fmt"


func makesquares(slice []int ) {
	for i,v := range slice {
		slice[i] = v * v
	}
}

func makesquaresarray(array [5]int ) {
	for i,v := range array {
		array[i] = v * v
	}
}

func PassbyValue() {
	a := [5]int{1,2,3,4,5}
	s := []int{1,2,3,4,5}

	// slices are passed as values.
	makesquares(s)
	fmt.Println(s)

	// Arrays are passed as a copy
	// hence underlying values are not changed
	makesquaresarray(a)
	fmt.Println(a)

}