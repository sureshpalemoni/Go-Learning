package interfaces

import "fmt"

type Material interface {
	ShapeMulti
	Object
}

func EmbeddedInterface() {
	c := Cube{3}
	var m Material = c
	var s ShapeMulti = c
	var o Object = c
	fmt.Printf("Dynamic type and value of interface %q of static type Material is %T and %v\n", "m", m, m)
	fmt.Printf("Dynamic type and value of interface %q of static type Material is %T and %v\n", "s", s, s)
	fmt.Printf("Dynamic type and value of interface %q of static type Material is %T and %v\n", "o", o, o)
}