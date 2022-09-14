// Branching

package main

import "fmt"

func main() {

	if true {
		fmt.Println("Hi, there")
	}

	marks := 77

	if marks >= 75 {
		fmt.Println("Passed with Honor")
	} else if marks >= 60 && marks < 75 {
		fmt.Println("Passed with First Division")
	}

	// logical operator
	fmt.Println(marks > 75 && marks <= 90)
	fmt.Println(marks > 75 || marks > 90)

	fmt.Println(marks > 75 && returnTrue("dummy") && marks > 91)
	fmt.Println("test 1 with returnTrue:", marks < 75 || marks > 91 || returnTrue("test 1"))

	// Shortcircte in OR condition as first condition is true, Go did not evaluated further
	fmt.Println("test 2 with returnTrue:", marks > 75 || returnTrue("test 2") || marks > 91)

}

func returnTrue(inp string) bool {
	fmt.Println("Echo from returnTrue: call from", inp)
	return true
}
