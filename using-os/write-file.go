// Go Playground
// https://go.dev/play/p/wdQoY4H9pKW

package main

import (
	"fmt"
	"os"
)

func main() {

	str := "this is my learning go programming"

	file, err := os.Create("/tmp/test-file")
	defer file.Close()

	if err != nil {
		fmt.Println(err)
	} else {
		os.WriteFile(file.Name(), []byte(str), 0750)
	}
}
