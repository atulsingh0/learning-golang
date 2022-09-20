package main

import "fmt"

func main() {

	a, b := 1, 0

	ans := a / b

	// Go itself will raise a panic here
	fmt.Print("Ans is:", ans)

}
