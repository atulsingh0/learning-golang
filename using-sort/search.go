package main

import (
	"fmt"
	"sort"
)

func main() {

	ages := []int{23, 45, 87, 20, 22, 45, 34, 88, 67, 19, 20}

	// sorting

	for i, v := range ages {
		fmt.Println(i, v)
	}

	fmt.Println()
	// Slice should be sorted for below to work :)
	fmt.Println("Searching 50, index: ", sort.SearchInts(ages, 50))
	fmt.Println("Searching 20, index: ", sort.SearchInts(ages, 20))

	// Let's sort now
	sort.Ints(ages)
	fmt.Println("After sorting")
	fmt.Println("Searching 50, index: ", sort.SearchInts(ages, 50))
	fmt.Println("Searching 20, index: ", sort.SearchInts(ages, 20))

	// searching 19
	fmt.Println("Searching 19, index (0): ", sort.SearchInts(ages, 19), ages)
	// searching 21, which does not exists
	// It tell us where we can insert the item in sorted array
	fmt.Println("Searching 21, index (0): ", sort.SearchInts(ages, 21), ages)

}
