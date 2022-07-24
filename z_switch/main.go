package main

import (
	"fmt"
)

func main() {

	var ans uint

	fmt.Printf("Enter a number between 0 to 10: ")
	fmt.Scan(&ans)

	// fmt.Println(reflect.TypeOf(ans))

	switch {
	case ans > 8:
		fmt.Println("Entered no is greater than 8")
	case ans > 5:
		fmt.Println("Entered no is greater than 5")
	case ans > 1:
		fmt.Println("Entered no is greater than 1")
	case ans == 1:
		fmt.Println("Entered no is equal to 1")
	default:
		fmt.Println("Switch1: Something is not right with your input.")
	}

	switch ans {
	case 8:
		fmt.Println("Entered no is 8")
	case 7:
		fmt.Println("Entered no is 7")
	case 6:
		fmt.Println("Entered no is 6")
	case 5:
		fmt.Println("Entered no is 5")
	default:
		fmt.Println("Switch2: Something is not right with your input.")
	}
}
