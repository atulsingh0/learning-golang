package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("This is an infinite scanner")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() { // scan infinitely
		fmt.Printf("--> You typed: %v\n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
