package slices

import "fmt"

func New() {
	var s []int
	fmt.Println(s)
	fmt.Println(s == nil)
}