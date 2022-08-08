package main

import (
	"fmt"
	"os"
	"strings"
)

func main(){

	args := os.Args[1:]

	if len(args) >= 1 {
		// fmt.Printf("Hello!! %v\n", args[0])
		fmt.Printf("Hello!! %v\n", strings.Join(args[:], " "))
	} else {
		fmt.Println("Hello World!!")
	}
}