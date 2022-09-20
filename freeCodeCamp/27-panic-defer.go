package main

import "fmt"

func main() {

	// Order of exetion
	// Main func --> Defer --> Panic
	// So resource will be taken case via defer even your app is panicing
	fmt.Println("start")
	defer fmt.Println("step1")
	fmt.Println("step2")
	panic("something bad happened")
	fmt.Println("end")

}
