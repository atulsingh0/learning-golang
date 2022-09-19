// Go Playground
// https://go.dev/play/p/on0MQcWndQY

package main

import (
	"fmt"
	"os"
)

func main() {

	// Get Current Dir
	if dir, err := os.Getwd(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Current Working Directory:", dir)
	}

	// Get Temp Dir
	if dir := os.TempDir(); dir != "" {
		fmt.Println("Temp Directory:", dir)
	}

	// Get User dir
	if dir, err := os.UserHomeDir(); err == nil {
		fmt.Println("User Home Directory:", dir)
	} else {
		fmt.Println(err)
	}

}
