// Go Playground
// https://go.dev/play/p/XW-OvdxwF6c

package main

import (
	"fmt"
	"os"
)

func main() {

	// renaming a file /tmp/test, source file should exists for rename operation
	err := os.Rename("/tmp/test", "/tmp/testing")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("File has been renamed.")
	}
}
