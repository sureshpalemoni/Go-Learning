package interfaces

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

func (r Rect) Area() float64 {
	return r.Width * r.Height
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}


func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func Intro() {
	var s Shape
	s = Rect{5.0, 4.0}
	fmt.Println("value of s is: ", s)
	fmt.Printf("Type of s is %T\n", s)
	fmt.Println("Area of Rectangle is: ", s.Area())
	fmt.Println("Perimeter of Rectangle is: ", s.Perimeter())

	s = Circle{5.0}
	fmt.Println("value of s is: ", s)
	fmt.Printf("Type of s is %T\n", s)
	fmt.Println("Area of Circle is: ", s.Area())
	fmt.Println("Perimeter of Circle is: ", s.Perimeter())	
}