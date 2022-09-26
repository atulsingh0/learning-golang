package main

import "fmt"

func main() {

	a := 43
	b := a
	fmt.Println(a, b)
	a = 27
	fmt.Println(a, b)

	var x int = 21
	var y *int = &x
	fmt.Println(x, y)
	x = 30
	fmt.Println(x, y) // y store the address of x
	fmt.Println(x, *y)
	*y = 14
	fmt.Println(x, *y)

	// Adding 2 to var via pointer
	fmt.Printf("Value of X: %v\n", x)
	addTwo(&x)
	fmt.Printf("Value of X after calling addTwo func: %v\n", x)
}


func addTwo( x *int) {
	*x += 2
}