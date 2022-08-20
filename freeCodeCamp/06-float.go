package main

import "fmt"

func main() {
	n := 3.14
	fmt.Printf("%v, %T\n", n, n)

	n = 1.2e23
	fmt.Printf("%v, %T\n", n, n)

	n = 1.3e12
	fmt.Printf("%v, %T\n", n, n)

	var k float32 = 12.22
	fmt.Printf("%v, %T\n", k, k)

	fmt.Printf("%v\n", n/float64(k))
	fmt.Printf("%0.6v\n", n/float64(k))

}
