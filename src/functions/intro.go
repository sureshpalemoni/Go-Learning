package functions

import "fmt"

func changeValue(a int) int {
	return a + 2
}

// This is named return
func addMultNamed(a, b int) (add int, mul int) {
	add = a + b
	mul = a * b

	return // necessary
}

// returning multiplae values
func addMult(a, b int) (int, int) {
	add := a + b
	mul := a * b	
	return add, mul
}

func Intro() {
	changeValue := changeValue(10)
	fmt.Println(changeValue)

	addResN, multResN := addMultNamed(2, 5)
	fmt.Println(addResN, multResN)

	addRes, multRes := addMult(2, 5)
	fmt.Println(addRes, multRes)
}
