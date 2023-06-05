package main

import "fmt"

func main() {

	// using goto statement
	loop(5)

}

func loop(number int) {
	// use of goto

	i := 0

outer:
	fmt.Println(i)
	i++

	if i < number {
		goto outer
	}
}
