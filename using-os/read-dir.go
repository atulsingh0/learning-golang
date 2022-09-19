// Go Playground
// https://go.dev/play/p/fObt453GMzE

package main

import (
	"fmt"
	"os"
)

func main() {

	if dirs, err := os.ReadDir("."); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Name", "IsDir")
		for _, data := range dirs {
			fmt.Println(data.Name(), data.IsDir())
		}
	}
}
