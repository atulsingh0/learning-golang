package main

import "fmt"

func main() {
	var n bool = true

	fmt.Printf("%v, %T\n", n, n)

	m := 1 == 2
	fmt.Printf("%v, %T\n", m, m)

	var k bool // default value is false
	fmt.Printf("%v, %T\n", k, k)

}
