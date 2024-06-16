package main

import "fmt"

func main() {
	// arithmetic

	var num1 = 4
	var num2 = 5.7

	no := float64(num1) + num2
	fmt.Println("num1 + num2:", no)

	var num3 = 9
	no1 := num3 / num1 // both no are int, so no is also int
	fmt.Println("int --> num3 / num1:", no1)

	// so to perform divide, we need to convert it to float
	no2 := float64(num3) / float64(num1)
	fmt.Println("float --> num3 / num1:", no2)

}
