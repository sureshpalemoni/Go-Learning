package interfaces

import "fmt"

//use the interfaces amd struct from multipleinterfaces.go file

type Skin interface {
	Color() float64
}

func TypeAssert() {
	var s ShapeMulti = Cube{3}
	
	// We are converting the interface from ShapeMulti to Object
	value1, ok := s.(Object)
	fmt.Printf("Dynamic value of Shape %q with value %v implements interface object? %v\n", "s", value1, ok)
	fmt.Println("Volume of Cube is: ", value1.Volume())

	
	//value2, ok := s.(ShapeMulti)
	// fmt.Printf("Dynamic value of Shape %q with value %v implements interface ShapeMulti? %v\n", "s", value2, ok)
	value2 := s
	fmt.Println("Area of c of interface type is", value2.Area())

	// We are converting the interface from ShapeMulti to Skin
	value3, ok := s.(Skin)
	fmt.Printf("Dynamic value of Shape %q with value %v implements interface Skin? %v\n", "s", value3, ok)
}