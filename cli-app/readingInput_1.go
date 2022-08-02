package main

import "fmt"

func main() {
	var name string
	fmt.Printf("Enter your Name: ")
	inp, _ := fmt.Scanf("%s", &name) // this scan will not entertain spaced input

	fmt.Printf("Hi %s, Entered input is %v\n", name, inp)

	fmt.Printf("Enter your full name: ")
	inp, _ = fmt.Scanf("%q", &name)
	fmt.Printf("Hi %s, Entered input is %v\n", name, inp)

}
