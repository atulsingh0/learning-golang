package main

import "fmt"

func main() {

	// logical operators
	// && and
	// || or
	// ! not

	val := true
	fmt.Println("this is val:", val)

	val = true || false && false
	fmt.Println("this is val:", val)

	val = (true || false) && false
	fmt.Println("this is val:", val)

	val = (true || false) && false && 7 > 5
	fmt.Println("this is val:", val)

	val = true || (false && false) && !(7 > 5)
	fmt.Println("this is val:", val)

}
