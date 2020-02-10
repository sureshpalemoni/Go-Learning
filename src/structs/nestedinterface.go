package structs

import "fmt"

type salaried interface {
	getSalary() int
}

// Get the salary struct from nestedstruct.go

type EmployeeInterface struct {
	firstname, lastname string
	salaried
}

func (s Salary ) getSalary() int {
	return s.basic + s.insurance + s.allowance
}

func NestedInterface() {
	suresh := EmployeeInterface{
		firstname: "Suresh",
		lastname: "palemoni",
		salaried: Salary{100, 200, 400},
	}
	
	// Only Methods are promoted when anonymous interface is used as field
	fmt.Println("Suresh's salary is: ", suresh.getSalary()) 
	
	// if you use the below line, it doesn't work as fields are not promoted.
	// fmt.Println("Suresh's salary is: ", suresh.basic )
}