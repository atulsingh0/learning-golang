package main

import "fmt"

func main() {

	var a = []int{2, 3, 4, 5, 7, 2, 1, 9, 4}

	fmt.Println("\nOne way: ")
	for i := 0; i < len(a); i++ {
		fmt.Println(i, a[i])
	}

	fmt.Println("\nAnother way: ")
	for i, v := range a {
		fmt.Println(i, v)
	}

	fmt.Println("Print dups ")
	for _, v := range a {
		for _, b := range a {
			if v == b {
				fmt.Printf("%d ", v)
			}
		}
	}

	fmt.Println("\nPrint dups where both are not at same index ")
	for i, v := range a {
		for j, b := range a {
			if v == b && i != j {
				fmt.Printf("%d ", v)
			}
		}
	}

	fmt.Println("\nPrint dups where both are not at same index and index j should be greater than i ")
	for i, v := range a {
		for j, b := range a {
			if v == b && i < j {
				fmt.Printf("%d ", v)
			}
		}
	}

	fmt.Println("\nAnother way: Print dups where both are not at same index and index j should be greater than i ")
	for i, v := range a {
		for j := i + 1; j < len(a); j++ {
			if v == a[j] {
				fmt.Printf("%d ", v)
			}
		}
	}
}
