package interfaces

import "fmt"

type ShapePointer interface {
	AreaPointer() float64
	Perimeter() float64
}

func (r *Rect) AreaPointer() float64 {
	return r.Width * r.Height
}

func PointerReceiver() {
	var s ShapePointer = &Rect{5.0, 4.0}
	fmt.Println("Area of rectangle is: ", s.AreaPointer())

	// Perimeter struct is from interfaceintro.go file
	fmt.Println("Area of rectangle is: ", s.Perimeter())
}
