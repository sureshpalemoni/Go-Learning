package interfaces

import (
	"fmt"
	"strings"
)

func explainType(i interface{}) {
	switch i.(type) {
	case string:
		// we are conevrting to interface of type string to perfrom string ops
		fmt.Println("i stored string:", strings.ToUpper(i.(string)))
		// you can read the i value without conversion
		fmt.Println("i stored string:", i)
	case int:
		fmt.Println("i stored int:", i)
	default:
		fmt.Println("i stored something else:", i)
	}
}

func TypeSwitch() {
	explainType("Hello World")
	explainType(52)
	explainType(false)
}
