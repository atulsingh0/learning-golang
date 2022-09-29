package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args[1:]

	fmt.Println("Passed args:", args)

	if len(args) > 0 {
		for _, val := range args {
			fmt.Println(val)
		}

	} else {
		fmt.Println("No arguments are passed.")
	}

}
