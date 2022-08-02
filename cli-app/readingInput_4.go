package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("This is an file reader")

	f, err := os.Open("./flagTest.go")

	if err != nil {
		fmt.Println("Err while opening file", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		fmt.Printf("--> You typed: %v\n", scanner.Text())

	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
