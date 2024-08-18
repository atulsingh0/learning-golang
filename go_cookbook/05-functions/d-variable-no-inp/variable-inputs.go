package main

import "fmt"

func main() {

	fmt.Println(sum(2, 5))
	fmt.Println(sum(7, 8, 2, 5))

	fmt.Println(sentence("Hello", " ", "World"))
	fmt.Println(sentence("Hello", " ", "World", "!", " ", "How", " ", "are", " ", "you", "?"))

}

// Variadic functions
// Accepts a variable number of arguments
// Variable parameter must be last parameter in function declaration
func sum(a ...int) int {
	s := 0
	for _, v := range a {
		s += v
	}
	return s

}

func sentence(a ...string) string {
	s := ""
	for _, v := range a {
		s += v
	}
	return s
}
