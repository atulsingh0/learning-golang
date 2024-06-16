package main

import "fmt"

func main() {

	num1 := 8
	num2 := 7

	if num1 != num2 {
		fmt.Println("num1 and num2 is not same")
	}

	out := num1 < num2
	fmt.Println("Let's see if this is:", out)

	fmt.Println("Check if nos are equal:", num1-1 == num2)

	if num1%2 == 1 {
		fmt.Println("Num is odd.")
	} else {
		fmt.Println("Num is even")
	}

	// string comparison
	// ASCII A-Za-z == 65 -

	fmt.Println("Ascii value of A, Z, a and z:", 'A', 'Z', 'a', 'z') // runes are represent as ASCII

	if "ABc" > "abc" {
		fmt.Println("This will never execute")
	}

}
