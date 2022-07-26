package main

import "fmt"

func main() {
	fmt.Println("I am running from main func.")

	test := func() {
		fmt.Println("Hello!\n")
	}
	test()

	// --------
	beta := func(x int) {
		fmt.Printf("Hello, %v\n", x)
	}
	beta(5)

	// ---------
	gamma := func(x int) int {
		return x * x
	}(8)

	fmt.Printf("%v\n", gamma)

}
