package main

import "fmt"

// slice == dymanic array

func main() {

	var beta []string

	beta = append(beta, "is")
	beta = append(beta, "me")

	fmt.Printf("Slice: %v\n", beta)

	// ---------------------------

	var arr = [4]int{1, 2, 3, 4}

	var alpha []int = arr[:3]

	fmt.Printf("arr : %v\n", arr)
	fmt.Printf("length arr : %v\n", len(arr))
	fmt.Printf("capacity arr : %v\n", cap(arr))

	fmt.Printf("\nslice: %v\n", alpha)
	fmt.Printf("length slice : %v\n", len(alpha))
	fmt.Printf("capacity slice : %v\n\n", cap(alpha))

	// ---------------------------

	var gamma []int = []int{2, 3, 4, 5, 6, 7}
	fmt.Printf("gamma: %v\n", gamma)

	delta := append(gamma, 2)
	fmt.Printf("\ndelta: %v\n", delta)

	// ---------------------------

	geta := make([]int, 0)
	fmt.Printf("\ngeta: %v", geta)
	fmt.Printf("\ngeta length: %v", len(geta))
	fmt.Printf("\ngeta capacity: %v\n", cap(geta))

	geta = append(geta, 3, 4, 5, 6)
	fmt.Printf("\ngeta: %v", geta)
	fmt.Printf("\ngeta length: %v", len(geta))
	fmt.Printf("\ngeta capacity: %v", cap(geta))
}
