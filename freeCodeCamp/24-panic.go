package main

import "fmt"

func main() {

	fmt.Println("start")
	// raise panic when program should stop execution due to any error
	panic("There is an error, I can not execute")
	fmt.Println("end")
}
