package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println("print", i)
	}

	// another way
	j := 0
	for j < 5 {
		fmt.Println("Another print", j)
		j++
	}

	// another way 2
	k := 0
	for ; k < 5; k++ {
		fmt.Println("Another way print", k)
	}

	// sample infinite loop
	for {
		if k > 10 {
			break
		}
		fmt.Println("Printing from infinite loop", k)
		k++
	}

	// print all nos b/w 20 to 30 which are odd and not divide by 3

	for i := 20; i <= 30; i++ {
		if i%2 == 0 || i%3 == 0 {
			continue
		}
		fmt.Println("data -", i)
	}

	// implementing, while, do, done in for
	// print square of no

}
