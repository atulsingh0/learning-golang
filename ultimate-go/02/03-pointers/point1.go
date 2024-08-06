package main

import "fmt"

func main() {

	count := 10
	fmt.Printf("%d, %v\n", count, &count)

	increment(count)
	fmt.Printf("%d, %v\n", count, &count)

	increment2(&count)
	fmt.Printf("%d, %v\n", count, &count)
}

func increment(x int) {
	x++
	fmt.Printf("now x is inside func: %d, %v\n", x, &x)
}

func increment2(x *int) {
	*x++
	fmt.Printf("now x is inside func: %d, %v\n", *x, x)
}
