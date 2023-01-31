package main

import "fmt"

func main() {

	a := "this"
	b := "that"

	fmt.Println("Before:", a, b)
	change(a, b) // pass by value
	fmt.Println("After:", a, b)

}

func change(x, y string) {

	fmt.Println("Within function, before assignment:", x, y)

	x = "they"
	y = "them"

	fmt.Println("Within function, after assignment:", x, y)
}
