package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
)

func main() {
	fmt.Println("Current Go version: ", runtime.Version())
	fmt.Printf("It is running on: %v\n", runtime.GOOS)

	// reading input
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter your name: ")

	msg, _ := reader.ReadString('\n')
	fmt.Printf("Hi! %v", msg)
}
