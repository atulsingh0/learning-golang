// Go Playground
// https://go.dev/play/p/Q1aQHlJwwph

package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.ReadFile("/etc/hosts")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(file))
	}
}
