package main

import "fmt"

func main() {

	// switch stmt is a cleaner way of multiple ifs

	age := 21

	switch age >= 18 {
	case true:
		fmt.Println("You are adult.")
	case false:
		fmt.Println("You are not adult.")
	}

	switch {
	case age >= 18:
		fmt.Println("You can drive.")
	case age > 14 && age < 18:
		fmt.Println("You can drive with parents.")
	case age <= 14:
		fmt.Println("You can not drive.")
	}

	switch age {
	case 18:
		fmt.Println("You are 18.")
	case 0:
		fmt.Println("You did not born.")
	default:
		fmt.Println("I am not sure if you exists.")
	}

	switch age {
	case 13, 14, 15, 16, 17, 18, 19:
		fmt.Println("You are teenager.")
	case 0:
		fmt.Println("You did not born.")
	default:
		fmt.Println("You are adult.")
	}

}
