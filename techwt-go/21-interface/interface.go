package main

import (
	"fmt"
)

type Circle struct {
	Radius float64
}

type Square struct {
	Side float64
}

type Rectangle struct {
	Width  float64
	Breath float64
}

func (c *Circle) Perimeter() float64 {
	return 2 * 22 * c.Radius / 7
}

func (c *Circle) Area() float64 {
	return 22 / 7 * c.Radius * c.Radius
}

func (c *Circle) Diameter() float64 {
	return 2 * c.Radius
}

func (s *Square) Area() float64 {
	return s.Side * s.Side
}

func (s *Square) Perimeter() float64 {
	return 4 * s.Side
}

//func (s *Square) Diagonal() flaat64 {
//	return math.Sqrt(2 * math.Pow(s.Side, 2))
//}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Breath
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Breath)
}

//func (r *Rectangle) Diagonal() flaat64 {
//	return math.Sqrt(math.Pow(r.Breath, 2) + math.Pow(r.Width, 2))
//}

// defining interface
type shape interface {
	Area() float64
	Perimeter() float64
}

func Area(t shape) {
	// type Switching
	switch t.(type) {
	case *Circle:
		fmt.Printf("Area of the object Circle (%T) is: %.2f\n", t, t.Area())
	case *Square:
		fmt.Printf("Area of the object Square (%T) is: %.2f\n", t, t.Area())
	}
}

func Perimeter(t shape) {
	fmt.Printf("Perimeter of the object (%T) is: %.2f\n", t, t.Perimeter())
}

func main() {

	// interface are implements
	// Same behaviour falls under same interface

	c1 := Circle{7}
	s1 := Square{4}
	r1 := Rectangle{3, 4}

	// Let's Call the Area
	fmt.Println("Area of Circle, Square and Rectangles are:", c1.Area(), s1.Area(), r1.Area())

	Area(&c1)
	Area(&s1)

	Perimeter(&r1)
	Perimeter(&c1)

}
