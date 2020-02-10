package structs

import "fmt"

type FullNameType func(string, string) string

type EmployeeFunctionField struct {
	FirstName, LastName string
	FullName FullNameType
}

func FunctionFields() {
	e := EmployeeFunctionField{
		FirstName: "Suresh",
		LastName: "Palemoni",
		FullName: func(firstname string, lastname string) string {
			return firstname + " " + lastname
		},
	}

	fmt.Println(e.FullName(e.FirstName, e.LastName))
}
