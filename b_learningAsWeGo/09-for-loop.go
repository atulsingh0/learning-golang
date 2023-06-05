package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println()

	j := 4
	for ; j < 64; j = j * 2 {
		fmt.Println(j)
	}

	// for {
	// 	k := 5
	// 	fmt.Println(k)
	// 	k = k + 2
	// 	if k < 20 {
	// 		continue
	// 	}
	// }
}
