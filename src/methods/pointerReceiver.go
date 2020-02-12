package methods

import (
	"fmt"
)

func (e *Employee) changeNameP(newName string) {
	// (*e).FirstName = newName // Not necessary to dereference to the pointer
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
	// ep := &e | We dont need to derefernce pointer if the method has pointer receiver
	// Here the method belongs to pointer of the type and 
	// copies only the pointer to the object instead copy of the object
	e.changeNameP("Palemoni")
	fmt.Println("e After name change = ", e.FirstName)
}