package main

import "fmt"

// closure func is a function that can refer to variables from the scope in which it was defined.

func outer() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	next := outer()
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
}

// closure func can be used as a callback function
// widely used in event-driven programming
// and web development
