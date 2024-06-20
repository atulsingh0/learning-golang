package main

import "fmt"

func main() {
	// calling func
	greet()

	// square of 4
	square(4)

	// with Return value
	add(3, 5) // will not print anything
	fmt.Println("Sum of 3 and 5:", add(3, 5))

	// multiple return values
	a, b := sqcu(7, 3)
	fmt.Println("Multiple return:", a, b)

	// checking defer
	fmt.Println("Let's see defer:", defe(3, 5))
}

// absolute func
func greet() {
	fmt.Println("Hi, How are you?")
}

// accept input
func square(x int) {
	fmt.Printf("Square of %d: %d\n", x, x*x)
}

// return
func add(x, y int) int {
	return x + y
}

func sqcu(x, y int) (int, int) {
	return x + y, x - y
}

// using defer
func defe(x, y int) int {
	defer fmt.Println("This is call just before the exit of said func")
	fmt.Println("I am in defe func")
	return x * y
}
