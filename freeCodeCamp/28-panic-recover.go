package main

import "fmt"

func main() {

	fmt.Println("Hello World!")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error:", err)
			panic(err)
		}
	}()
	panic("Something is not right")

}
