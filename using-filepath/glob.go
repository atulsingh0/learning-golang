// Go Playground
// https://go.dev/play/p/4tgSJ6X1fLD

package main

import (
	"fmt"
	"path/filepath"
)

func main() {

	// fetch data based on pattern
	if data, err := filepath.Glob("/etc/[a-m]*"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(data)
	}
}
