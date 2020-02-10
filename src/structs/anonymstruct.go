package structs

import "fmt"

func AnonymStruct() {
	suresh := struct {
		firstname, lastname string
		salary int
		fulltime bool
	}{
		firstname: "Suresh",
		lastname: "Palemoni",
		salary: 2199,
	}
	fmt.Printf("%T",suresh)
}