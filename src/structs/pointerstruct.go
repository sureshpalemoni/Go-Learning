package structs

import "fmt"

// Using the Struct type Employee from structs1.go file
func PointerStruct() {
	suresh := &Employee{
		FirstName: "Suresh",
	}
	fmt.Println("firstname: ",(*suresh).FirstName)
	fmt.Println("firstname: ",(suresh).FirstName)
	fmt.Println("firstname: ",suresh.FirstName)
}