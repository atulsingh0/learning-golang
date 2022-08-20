package main

import (
	"fmt"
)

// global var
var (
	i float64 = 32.2
	k int64   = 234
)

func main() {
	var i int64 = 23
	fmt.Printf("%v, %T\n", i, i)

	// conversion to float32
	j := float32(i)
	fmt.Printf("%v, %T\n", j, j)

	fmt.Printf("%v, %T\n", k, k)

	k := 2.34 // shadowing the var k
	fmt.Printf("%v, %T\n", int16(k), int16(k))

	// Value of var k in another function
	test()

	// convering value of k to string
	// kStr := string(k) // thows an error
	kStr := fmt.Sprintf("%f", k)
	fmt.Printf("%v, %T\n", kStr, kStr)

	kStr = fmt.Sprintf("%0.2f", k)
	fmt.Printf("%v, %T\n", kStr, kStr)
}

func test() {
	fmt.Printf("%v, %T\n", k, k)
}
