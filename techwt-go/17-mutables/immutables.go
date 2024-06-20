package main

import "fmt"

func main() {

	// Immutable: can not change

	var x int = 4
	y := x
	y = 7

	fmt.Println("The value of x, and y:", x, y)

	// Mutuable: which can change

	a := []int{2, 4, 6}
	b := a
	b[0] = 99
	fmt.Println("The value of a, and b:", a, b)

	// Array is Immutable
	p := [3]int{2, 9, 7}
	q := p
	q[0] = 88
	fmt.Println(p, q)

	// Maps are Mutable
	al := map[int]int{
		1: 1,
		2: 4,
		3: 9,
	}

	pl := al
	pl[1] = 77

	fmt.Println("Maps are:", al, pl)

}
