package structs

import "fmt"

// Salary is a structure promoted here from nestedstruct.go file
type PromotedEMployeeField struct {
	firstname, lastname string
	fulltime bool
	Salary
}

func PromotedEMployee() {
	suresh := PromotedEMployeeField{
		firstname: "Suresh",
		lastname: "Palemoni",
		fulltime: true,
		Salary: Salary{100,200,2003},
	}
	fmt.Println("Suresh's Basic Salary", suresh.Salary.basic)
	fmt.Println("Suresh's Basic Salary ==> Field Promotion:", suresh.basic) //field promotion
}