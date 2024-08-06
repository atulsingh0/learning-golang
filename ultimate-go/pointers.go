package main

import (
	"fmt"
)

func main() {

	count := 10

	fmt.Println("Value count:", count)
	fmt.Println("Address count:", &count)

	increment(count)
	fmt.Println("Value count:", count)
	fmt.Println("Address count:", &count)

}

func increment(x int) int {
	fmt.Println("Address count in func:", &x)
	return x + 1
}
