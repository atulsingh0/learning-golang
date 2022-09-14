// Branching

package main

import (
	"fmt"
)

func main() {

	// switch have default break b/w cases
	marks := 10

	switch {
	case marks > 75:
		fmt.Println("Passed with Honor")
	case marks > 60:
		fmt.Println("Passed with First Class")
	case marks > 50:
		fmt.Println("Passed with Second Class")
	case marks > 35:
		fmt.Println("Passed")
	default:
		fmt.Println("Fail")

	}

	no := 81

	switch no {
	case 1, 3, 5:
		fmt.Println("one, three, five")
	case 7, 9, 0:
		fmt.Println("seven, nine, zero")
	case 2, 4, 6:
		fmt.Println("two, four, fix")
	case 8:
		fmt.Println("eight")

	default:
		fmt.Println("No not in b/w 1 or 10")
	}

	// making switch fallthrough

	i := 10
	switch {
	case i <= 10:
		fmt.Println("No is <= 10")
		fallthrough // this will simply fallthough and execute the next stmt without checking the condition
	case i <= 20:
		fmt.Println("No is <= 20")
		fallthrough
	case i > 20:
		fmt.Println("No is > 20")
	case i > 30:
		fmt.Println("No is > 30")
	}

	// Type switch

	var k interface{} = "al"

	switch k.(type) {
	case int:
		fmt.Println("I am an Int")
	case float64:
		fmt.Println("I am a float64")
	case string:
		fmt.Println("I am a string")
	default:
		fmt.Println("I dont know who I am ")
	}

}
