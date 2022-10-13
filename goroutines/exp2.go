package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		go fmt.Print(" S", Square(i))
	}

	for i := 0; i < 10; i++ {
		go fmt.Print(" C", Cube(i))
	}
	// Wait for an input
	fmt.Scanln()
}

func Square(n int) int {
	return n * n
}

func Cube(n int) int {
	return n * n * n
}
