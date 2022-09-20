// Go Playground
// https://go.dev/play/p/JeKjB3L6MPF

package main

import (
	"fmt"
	"path/filepath"
)

func main() {

	// return true if match else false

	fmt.Println(filepath.Match("/etc/a*", "/etc/abs"))
	fmt.Println(filepath.Match("/home/\\*", "/home/*"))
	fmt.Println(filepath.Match("/home/?opher", "/home/gopher"))
	fmt.Println(filepath.Match("/home/catch/*", "/home/catche/foo/bar"))

}
