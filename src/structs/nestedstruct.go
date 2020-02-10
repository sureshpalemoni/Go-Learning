package structs

import "fmt"

type Salary struct {
	basic, insurance, allowance int
}

type EmployeeNested struct {
	firstname, lastname string
	salary 				Salary
	fulltime bool
}

func NestedStruct() {
	suresh := EmployeeNested{
		firstname: "Suresh",
		lastname: "Palemoni",
		salary: Salary{100, 200, 400},
		fulltime: true,
	}
	fmt.Println(suresh)
}