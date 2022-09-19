// Go Playground
// https://go.dev/play/p/3Jn8cADQMrU

package main

import (
	"fmt"
	"os"
)

func main() {

	if dir, err := os.Getwd(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(dir)
	}
}
