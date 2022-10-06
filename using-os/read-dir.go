// Go Playground
// https://go.dev/play/p/fObt453GMzE

package main

import (
	"fmt"
	"os"
)

func main() {

	dirName := "/tmp"

	if dirs, err := os.ReadDir(dirName); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Directory:", dirName, "\n")
		fmt.Println("IsDir", "\t\t\t", "Name")
		fmt.Println("------------------------")
		for _, data := range dirs {
			fmt.Println(data.IsDir(), "\t\t\t", data.Name())
		}
	}
}
