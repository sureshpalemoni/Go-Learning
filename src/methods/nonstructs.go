// Methods for Non structs Type
package methods

import (
	"fmt"
	"strings"
)

type MyString string 
// Since MySring type belongs to this package,
// method toUppercase belongs MyString type

func (s MyString) toUppercase() string {
	normalString := string(s)
	return strings.ToUpper(normalString)
}

func NonStructs() {
	str := MyString("Hello Palemoni")
	fmt.Println(str.toUppercase())
}