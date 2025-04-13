package main

import (
	"fmt"
	"math"
)

func main() {

	r := Rectangle{
		3, 4,
	}
	c := Circle{
		Radius: 7,
	}

	s := Square{
		Rectangle{
			4, 4,
		},
	}

	// using interface
	fmt.Println("Area of Rectangle (3,4):", CalcArea(r))
	fmt.Println("Area of Circle (7):", CalcArea(c))
	fmt.Println("Area of Square (4):", CalcArea(s))

}

// creating an interface
// Should match the method signature
type Areaer interface {
	Area() float64
}

func CalcArea(a Areaer) float64 {
	return a.Area()
}

type Square struct {
	Rectangle
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Length, Width float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}
