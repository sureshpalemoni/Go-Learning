package arrays

import "fmt"

func MultiDimensional() {
	m := [...][2]int {
		[...]int{1, 2},
		[...]int{3, 4},
	}

	//or can be written as

	n := [...][2]int{{1,2}, {3,4}}

	fmt.Println(m)
	fmt.Println(n)
}