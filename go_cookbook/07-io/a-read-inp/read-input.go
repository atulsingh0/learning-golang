package main

import (
	"fmt"
	"io"
	"os"
)

// use io.Reader interface to read from an input

func main() {

	bytes := make([]byte, 5) // read only 5 bytes

	n, err := os.Stdin.Read(bytes)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes[:n]), n)

	// or
	char, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(char), len(char))
}
