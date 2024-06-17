package main

import (
	"fmt"
)

func main() {

	// Ordered collection of same type of data
	// Fixed length
	// based on data type, value will be defaulted

	var arr [4]int
	fmt.Println(arr)

	arr[3] = 8
	arr[1] = 2
	fmt.Println(arr)

	var sent [8]string
	sent[0] = "Hi"
	sent[1] = "How"

	fmt.Println(sent)

	no := [5]float64{2.3, 1, 4, 0, 3.1}
	fmt.Println(no)

	// length and capacity
	fmt.Println("len and cap of arr", len(arr), cap(arr))
	var sum1 int
	for i := 0; i < len(arr); i++ {
		sum1 = sum1 + arr[i]
	}
	fmt.Println("total:", sum1)

}
