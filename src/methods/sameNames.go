package methods

import (
	"fmt"
	"interfaces"
	"math"
)

type Rect interfaces.Rect
type Circle interfaces.Circle

func (r Rect) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return 2 * math.Pi * c.Radius
}

func SameNames() {
	rect := Rect{5.0, 4.0}
	circ := Circle{5.0}
	fmt.Printf("Area of rectangle = %0.2f\n", rect.Area())
	fmt.Printf("Area of Circle = %0.2f\n", circ.Area())
}
