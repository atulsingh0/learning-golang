// Go Playground
// https://go.dev/play/p/P3eeSqsFWWo

package main

import (
	"fmt"
	"os"
)

func main() {

	if dir, err := os.MkdirTemp("/tmp", "log-*"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Dir has been created", dir)
		defer os.RemoveAll(dir) // remove
	}

}
