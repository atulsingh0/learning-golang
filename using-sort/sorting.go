package main

import (
	"fmt"
	"sort"
)

func main() {

	ages := []int{23, 45, 87, 20, 22, 45, 34, 88, 67, 19, 20}

	// sorting
	sort.Ints(ages)
	fmt.Println(ages)
}
