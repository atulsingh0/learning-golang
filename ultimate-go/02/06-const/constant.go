package main

import "fmt"

func main() {
	const (
		x = 10.0
		y = 3
	)
	// no explicit type casting required for constant
	fmt.Println(x / y)

}
