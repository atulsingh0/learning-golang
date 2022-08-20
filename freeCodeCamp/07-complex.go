package main

import "fmt"

func main() {
	var n complex64 = 2i
	fmt.Printf("%v, %T\n", n, n)

	var m complex64 = 2 - 3i
	fmt.Printf("%v, %T\n", n+m, n+m)
	fmt.Printf("%v, %T\n", n-m, n-m)
	fmt.Printf("%v, %T\n", n*m, n*m)
	fmt.Printf("%v, %T\n", n/m, n/m)

	// Split the data into real and imaginory
	fmt.Printf("%v, %T\n", real(m), real(m))
	fmt.Printf("%v, %T\n", imag(m), imag(m))

	// creating a complex no
	var p complex64 = complex(3, 20)
	fmt.Printf("%v, %T\n", p, p)

}
