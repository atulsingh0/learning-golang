package main

import "fmt"

func main() {
	var i int = 23
	fmt.Printf("%v, %T\n", i, i)

	var j uint16 = 11
	fmt.Printf("%v, %T\n", j, j)

	k := 2
	fmt.Printf("%v, %T\n", k, k)

	fmt.Printf("%v\n", i/k)
	fmt.Printf("%v\n", i%k)

	fmt.Printf("%v\n", i/int(j)) // data conversion is required
	fmt.Printf("%v\n", i+int(j))

	fmt.Printf("%b,%b,%b, %v\n", i, k, i&k, i&k) // binary operation  AND
	fmt.Printf("%b,%b,%b, %v\n", i, k, i|k, i|k) // OR
	fmt.Printf("%b,%b,%b, %v\n", i, k, i^k, i^k)
	fmt.Printf("%b,%b,%b, %v\n", i, k, i&^k, i&^k)

	a := 8

	fmt.Printf("%b, %v\n", a, a)
	fmt.Printf("%b, %v\n", a<<3, a<<3)
	fmt.Printf("%b, %v\n", a>>3, a>>3)
	fmt.Printf("%b, %v\n", a>>5, a>>5)

}
