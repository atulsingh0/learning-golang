package main

import (
	"fmt"
	"math"
)

func main() {

	// using

	r := Rectangle{
		3, 4,
	}
	c := Circle{
		Radius: 7,
	}

	//fmt.Println("Area of Rectagle (3,4):", r.Area())
	//fmt.Println("Area of Circle (7):", c.Area())

	// using interface
	fmt.Println("Area of Rectangle (3,4):", CalcArea(r))
	fmt.Println("Area of Circle (7):", CalcArea(c))

}

// creating an interface
// Should match the method signature
type Areaer interface {
	Area() float64
}

func CalcArea(a Areaer) float64 {
	return a.Area()
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
