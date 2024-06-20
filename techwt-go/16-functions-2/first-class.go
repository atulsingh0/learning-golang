package main

import (
	"fmt"
	"math"
)

func main() {
	// functions are first class citizen in golang
	// it have types as like any other variables

	alpha := test
	alpha() // calling alpha

	// inline func
	beta := func() {
		fmt.Println("Hey this is beta")
	}
	beta()

	// another example of inline func
	addTest := func(x, y int) int {
		return x + y
	}

	fmt.Println("Sum of 7 and 8:", addTest(7, 8))

	// execute the inline func inline
	minus4 := func(x int) int {
		return x - 4
	}(12)
	fmt.Println("Output of 12 - 4 is:", minus4)

	// generate func for cube
	cube := generateFunc(3)

	fmt.Println("Cube of 2:", cube(2.0))
	fmt.Println("Cube of 4:", cube(4.0))

	// generate func for square
	sqr := generateFunc(2)
	fmt.Println("Cube of 5:", sqr(5.0))
	fmt.Println("Cube of 6:", sqr(6.0))

	// another way to call and use the func
	fmt.Println("2 ^ 5 is:", generateFunc(5)(2))

	// another way of closure
	func() {
		fmt.Println("Headless Function")
	}()

	// another closure
	fmt.Println("Nested Closure func:", returnFunc(3)())
}

func test() {
	fmt.Println("Call from test func")
}

// closure
func generateFunc(y float64) func(x float64) float64 {
	return func(x float64) float64 {
		return math.Pow(x, y)
	}
}

func returnFunc(x int) func() int {
	return func() int {
		return x * x * 4
	}
}
