// Go Playground
// https://go.dev/play/p/9jkHB_d1qZ-

package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	if _, err := os.Stat("/tmp/testing"); err == nil {
		fmt.Println("File/Dir exists.")

	} else if errors.Is(err, os.ErrNotExist) {
		fmt.Println("File/Dir does not exists.")

	} else {
		fmt.Println("Err:", err)
	}
}
