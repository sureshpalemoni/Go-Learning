package interfaces

import "fmt"

type ShapeMulti interface {
	Area() float64
}

type Object interface {
	Volume() float64
}

type Cube struct {
	side float64
}

func (c Cube) Area() float64 {
	return 6 * (c.side * c.side)
}

func (c Cube) Volume() float64 {
	return (c.side * c.side * c.side)
}

func MultipleInterfaces() {
	c := Cube{3}
	// You may declare variables for Shape and Object Interfaces like
	// var s ShapeMulti = c
	// var o object = c
	// And use s.Area() , o.Volume() ==> This is static type 
	// We are using c since struct Cube implements both the interfaces.
	fmt.Println("Volume of c of interface type shape is", c.Volume())
	fmt.Println("Area of c of interface type is", c.Area())
}