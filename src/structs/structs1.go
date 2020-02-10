package structs

import "fmt"

type Employee struct {
	FirstName, LastName string
	Salary int
	FullTime bool
}

func StructsIntro() {
	ross := Employee{
		FirstName: "Suresh",
		LastName: "Palemoni",
		FullTime: false,
		Salary: 1100,
	}
	fmt.Println(ross)
}