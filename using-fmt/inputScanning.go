package main

import "fmt"

func main() {

	var fname string
	var lname string

	fmt.Println("What is your name? ")
	fmt.Scanf("%s %s", &fname, &lname) // default sep is " "

	fmt.Println("Your name is", fname, lname)
}
