package main

import (
	"fmt"
	"os"
)

func main(){

	args := os.Args[1:]

	if len(args) >= 1 {
		fmt.Printf("Hello!! %v\n", args[0])
	} else {
		fmt.Println("Hello World!!")
	}
}