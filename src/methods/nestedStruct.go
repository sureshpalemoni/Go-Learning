package methods

import "fmt"

type Contact struct {
	phone, address string
}

type EmployeeMethod struct {
	name string
	salary int
	Contact // Since this is Anonymous filed its field gets promoted to EmployeeMethod Struct
}

func (e *EmployeeMethod) changePhone(newPhone string) {
	e.phone = newPhone // Field Promotion
}

/* Methods implemented by the anonymous struct will also
gets promoted to the parent struct
The above method can also be written as:

func (e *Contact) changePhone(newPhone string) {
	c.phone = newPhone
}

 */


func NestedStruct() {
	e := EmployeeMethod{
		name: "Suresh",
		salary: 20,
		Contact: Contact{"0551", "Abu Dhbai,UAE"},
	}
	// Phone Number before changing
	fmt.Println("e before change =", e.phone) // Field Promotion
	e.changePhone("0231")
	// if the method  has value receiver 
	// use (&e).changePhone("0231")
	// Go will take care under the hood
	// Phone number after changing
	fmt.Println("e after change = ", e.phone) //Field Promotion
}

