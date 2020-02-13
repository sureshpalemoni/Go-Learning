// Methods for Non structs Type
package methods

import (
	"fmt"
	"strings"
)

func (s string) toUppercase() string {
	return strings.ToUpper(s)
}

func NonStructs() {
	str := "Hello Palemoni"
	fmt.Println(str.toUppercase())
}