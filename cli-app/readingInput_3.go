package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("This is an infinite scanner until you entered 'quit'")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() { // scan infinitely
		fmt.Printf("--> You typed: %v\n", scanner.Text())

		if scanner.Text() == "quit" {
			fmt.Println("We are quitting...")
			os.Exit(3)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
