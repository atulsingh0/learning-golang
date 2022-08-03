package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Reading a file family.txx\n")

	// scanner
	file, err := os.Open("family.txt")
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		//fmt.Printf(scanner.Text() + "\n")

		var name string
		var age int

		n, err := fmt.Sscanf(scanner.Text(), "%s is %d years old.", &name, &age)
		if err != nil {
			fmt.Println(err)
		}

		if n == 2 {
			fmt.Printf("%s : %d\n", name, age)
		}
	}
}
