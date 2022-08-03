package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	var fileName string
	var choice int
	fmt.Println("Welcome to File Reader utili")

	fmt.Printf("Enter file name: ")

	_, err := fmt.Scanf("%s", &fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	fmt.Println("Enter your choice to print:")
	fmt.Println("Enter 1 for Upper Case")
	fmt.Println("Enter 2 for Lower Case")

	_, err = fmt.Scanf("%d", &choice)

	if err != nil {
		fmt.Println(err)
		os.Exit(4)
	}

	// reading file in one go
	contents, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	switch choice {
	case 1:
		fmt.Println(strings.ToUpper(string(contents)))
	case 2:
		fmt.Println(strings.ToLower(string(contents)))
	}

}
