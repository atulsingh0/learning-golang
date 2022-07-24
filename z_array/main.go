// array
// Ordered collection of values (fixed length)

package main

import "fmt"

func main() {
	var arr [5]string
	var val [7]uint

	fmt.Printf("arr:  %v\n", arr)
	fmt.Printf("val: %v\n", val)

	// ----------------------------
	arr[0] = "us"
	val[1] = 100
	fmt.Println("")

	fmt.Printf("arr:  %v\n", arr)
	fmt.Printf("val: %v\n", val)

	// ----------------------------

	var name = [3]string{"atul", "divya", "priya"}
	fmt.Printf("Names: %v\n\n", name)

	// ----------------------------

	var rand = [4]int{9, 2, 3, 4}
	var sum int = 0

	for _, v := range rand {
		sum += v
	}
	fmt.Printf("Sum of array: %v\n", sum)

	for i := 0; i < len(rand); i++ {
		sq := rand[i] * rand[i]
		rand[i] = sq
	}
	fmt.Printf("Squared arr: %v\n\n", rand)
	// ----------------------------

}
