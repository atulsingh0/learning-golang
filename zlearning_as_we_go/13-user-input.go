package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Input: ", input)
	}

	fmt.Print("Enter another input: ")
	another, err := reader.ReadBytes('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("another Input: ", another)
	}
}
