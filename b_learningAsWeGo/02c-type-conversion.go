package main

import "fmt"

func main() {

	amt := 5
	interest := 2.75

	total := float64(amt) + float64(amt)*interest/100
	fmt.Println("Total amt:", total)

	// converting float64 to int
	// it will drop the decimal values
	fmt.Println(int(interest))
}
