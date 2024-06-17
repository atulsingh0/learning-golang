package main

import "fmt"

func main() {

	beta := []string{"hi", "how", "are", "you", "?"}

	// looping over beta
	for i := 0; i < len(beta); i++ {
		fmt.Print(beta[i], " ")
	}
	fmt.Println()

	// looping via range
	for i, v := range beta {
		fmt.Println(i, v)
	}
}
