package methods

import (
	"fmt"
)

func (e *Employee) changeNameP(newName string) {
	e.FirstName = newName
}

func PointerReceiver() {
	e := Employee{
		FirstName: "Naresh",
		LastName: "Palemoni",
		Salary: 1400,
	}
	// e before the name change
	fmt.Println("e before name change = ", e.FirstName)
	// Create the pointer to e
	ep := &e
	// Here the method belongs to pointer of the type and 
	// copies only the pointer to the object instead copy of the object
	ep.changeNameP("Palemoni")
	fmt.Println("e After name change = ", e.FirstName)
}