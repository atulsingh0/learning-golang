package main

import "fmt"

func main() {

	name := "monday"
	x := 5
	y := 2
	out := 0

	if name == "monday" {
		fmt.Println("You are correct")
	}

	if x%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	// local var
	if out := x / y; out != 0 {
		fmt.Println(out)
	}

	// outer out
	fmt.Println(out)

}
