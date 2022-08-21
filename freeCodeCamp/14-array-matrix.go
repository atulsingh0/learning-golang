package main

import "fmt"

func main() {
	var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}

	fmt.Printf("identityMatrix: %v\n", identityMatrix)
}
