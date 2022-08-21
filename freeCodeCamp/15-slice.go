package main

import "fmt"

func main() {

	// Slice is a projection of underlying array
	// Slices are by default referecing the data when creating a copy

	a := []int{1, 2, 4}
	fmt.Printf("Slice: %v\n", a)
	fmt.Printf("Slice Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))

	b := a
	a[0] = 99
	fmt.Printf("Slice a: %v\n", a)
	fmt.Printf("Slice b: %v\n", b)

	a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	p := a
	q := a[:]
	r := a[3:]
	s := a[4:8]
	t := a[:7]
	fmt.Printf("%v, %v, %v, %v, %v\n", p, q, r, s, t)

	a = make([]int, 3)
	fmt.Printf("%v, %v\n", a, len(a))

	a = make([]int, 3, 100)
	fmt.Printf("%v, %v, %v\n", a, len(a), cap(a))

	a = []int{}
	fmt.Printf("a: %v, %v, %v\n", a, len(a), cap(a))
	a = append(a, 12) // creating a copy and storing in the same var
	fmt.Printf("a: %v, %v, %v\n", a, len(a), cap(a))
	// If slice is small, copy operation is cheap but for bigger slices, it's better to have a size
	a = append(a, 23, 11, 15, 17)
	fmt.Printf("a: %v, %v, %v\n", a, len(a), cap(a))
	a = append(a, 75, 29)
	fmt.Printf("a: %v, %v, %v\n", a, len(a), cap(a))

	// appending an array via spread operator ...
	fmt.Println("\nAppending an Array via Spread Operator ... ")
	a = append(a, []int{99, 99, 99}...)
	fmt.Printf("a: %v, %v, %v\n", a, len(a), cap(a))
	b = []int{88, 88}
	a = append(a, b...)
	fmt.Printf("a: %v, %v, %v\n", a, len(a), cap(a))

	// Using Slice as a Stack
	fmt.Println("\nUsing Slice as Stack.... ")
	a = []int{1, 2, 3}
	fmt.Printf("a: %v, %v, %v\n", a, len(a), cap(a))
	// Pop an item
	b = a[:len(a)-1]
	fmt.Printf("b: %v, %v, %v\n", b, len(b), cap(b))
	fmt.Printf("a: %v, %v, %v\n", a, len(a), cap(a))

	// B'cuz Slice is by default uses reference
	// Append operation change the original value
	fmt.Printf("a: %v, %v, %v\n", a, len(a), cap(a))
	b = append(a, 1, 2, 3)
	fmt.Printf("b: %v, %v, %v\n", b, len(b), cap(b))
	fmt.Printf("a: %v, %v, %v\n", a, len(a), cap(a))
	c := append(a[1:2], 2)
	fmt.Printf("c: %v, %v, %v\n", c, len(c), cap(c))
	fmt.Printf("a: %v, %v, %v\n", a, len(a), cap(a))
}
