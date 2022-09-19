// Go Playground
// https://go.dev/play/p/MF-dT57HXTM

package main

import (
	"fmt"
	"os"
)

func main() {

	// creating dir /tmp/gopal
	if err := os.Mkdir("/tmp/gopal", 0750); err == nil {
		fmt.Println("Dir has been created.")
	} else {
		fmt.Println(err)
	}

	// create the dir with parent dir
	if err := os.MkdirAll("/tmp/go1/go2/go3", 0750); err == nil {
		fmt.Println("Dir has been created.")
	} else {
		fmt.Println(err)
	}
}
