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

	fmt.Println("Cube of 3 is:", calcCube(3))

	fmt.Println("Multiply 3 and 4: ", multiply(3, 4))

	// calling divide func
	if result, err := divide(3, 2); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Result of 3 / 2 is :", result)
	}

	if result, err := divide(3, 0); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Result of 3 / 0 is :", result)
	}
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
// variadic paramter should be used only once and as a last parameter
func addAll(data ...float64) {
	var result float64
	for _, d := range data {
		result += d
	}
	fmt.Printf("Total of %v: %v\n", data, result)
}

func calcCube(n float64) float64 {
	return n * n * n
}

// named return variabled
func multiply(m, n float64) (result float64) {
	result = m * n
	return
}

// multiple return value
func divide(m, n float64) (float64, error) {

	if n == 0 {
		return 0.0, fmt.Errorf("Can not divide by 0.\n")
	} else {
		return m / n, nil
	}

}
