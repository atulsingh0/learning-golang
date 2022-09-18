// Go Playground
// https://go.dev/play/p/YrmSmstpN_u

package main

import (
	"fmt"
	"os"
)

func main() {

	// creating a dummy file in /tmp dir
	file, err := os.Create("/tmp/test")
	if err != nil {
		fmt.Println(err)
	} else {
		// Fetching file Stats
		stat, _ := os.Stat(file.Name())
		fmt.Printf("file %v has created and permission are as :%v\n", file.Name(), stat.Mode())
	}
}
