package main

import "fmt"

func main() {

	// break - will break the current loop
	for i := 1; i < 10; i = i + 1 {
		fmt.Printf("%d ", i)

		if i > 7 {
			break
		}
	}

	fmt.Println()

	// double loop - break
	for i := 1; i < 5; i = i + 1 {
		for j := 2; j < 5; j = j + 1 {

			fmt.Printf("%d %d \n", i, j)
			if j > 3 {
				// this will break the only current loop (var j)
				break
			}
		}
	}
	fmt.Println()

	// brak outer loop by using label
outer:
	for i := 1; i < 5; i = i + 1 {
		for j := 2; j < 5; j = j + 1 {

			fmt.Printf("%d %d \n", i, j)
			if j > 3 {
				// this will break the only current loop (var j)
				break outer
			}
		}
	}
}
