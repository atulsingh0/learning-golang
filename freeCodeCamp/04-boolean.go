package main

import "fmt"

func main() {
	var n bool = true

	fmt.Printf("%v, %T\n", n, n)

	m := 1 == 2
	fmt.Printf("%v, %T\n", m, m)

	var k bool // default value is false
	fmt.Printf("%v, %T\n", k, k)

	var (
		alpha = true
		beta  = false
		gamma = true
	)

	fmt.Println(alpha, beta, gamma)

	fmt.Println(!alpha)        // Negation
	fmt.Println(alpha && beta) // and operation
	fmt.Println(alpha || beta) // Or operation

	fmt.Println(alpha == gamma) // comparison

}
