package methods

import (
	"fmt"
)

func (e Employee) changeName(newName string) {
	e.FirstName = newName
}


func ValueReceiver() {
	e := Employee{
		FirstName: "Naresh",
		Salary: 1400,
	}

	fmt.Println("e before name change = ", e.FirstName)
	// Here the changeName method just received the copy of the object 
	// and doesn't impact the original object.
	e.changeName("Palemoni")
	fmt.Println("e After name change = ", e.FirstName)
}