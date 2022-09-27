package main

import "fmt"

func main() {

	x := 4.0

	square(x)
	fmt.Println("Value of x after calling square:", x)

	squareRef(&x)
	fmt.Println("Value of x after calling square:", x)

	// calling addAll func
	addAll(2, 3, 4, 5, 6, 7, 7)

}

// copy by value
func square(n float64) {
	fmt.Printf("Square of %v is: %v\n", n, n*n)
	// setting value of n to its half
	n = n / 2
}

// copy by reference
func squareRef(n *float64) {
	fmt.Printf("Square of %v is: %v\n", *n, *n**n)
	// setting value of x to its half
	*n = *n / 2
}

// variadic parameter
func addAll(data ...float64) {
	var result float64
	for _, d := range data {
		result += d
	}
	fmt.Printf("Total of %v: %v ", data, result)
}
