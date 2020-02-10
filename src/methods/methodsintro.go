package methods

import (
	"fmt"
	"structs"
)

// Imported the struct fields from structs package
type Employee structs.Employee


/* reguar Function
func fullName(firstname string, lastname string) (fullName string) {
	fullName = firstname + " " + lastname
	return
}
*/

func (e Employee) fullName() string {
	return e.FirstName + " " + e.LastName
}

func MethodsIntro() {
	e := Employee{
		FirstName: "Suresh",
		LastName: "Palemoni",
	}
	fmt.Println(e.fullName())
}