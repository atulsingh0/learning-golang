package main

import "fmt"

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(add_float(1.2, 2.3))
}

func add(x int, y int) int {
	return x + y
}

func add_float(x float32, y float32) float32 {
	return x + y
}
