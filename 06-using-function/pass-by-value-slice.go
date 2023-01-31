package main

import "fmt"

func main() {

	a := []int{2, 3, 4}

	fmt.Println("Before Slice:", a)
	change(a) // pass by value
	fmt.Printf("Pointer of Slice Object in main: %p\n", a)
	fmt.Println("After Slice:", a)
}

func change(x []int) {

	// Though, We are passing Slice by name but Slice itself is a list of pointer reference
	// Hence, if we change the value of an element in this function, it will change the original as well

	fmt.Println("Within function, Slice before assignment:", x)
	x[0] = 99
	fmt.Println("Within function, Slice after assignment:", x)
	fmt.Printf("Pointer of Slice Object in function: %p\n", x)
}
