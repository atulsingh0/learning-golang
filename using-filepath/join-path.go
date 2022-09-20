// Go Playground
// https://go.dev/play/p/awprzHOM1Aj

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	if cur, _ := os.Getwd(); cur != "" {
		// join current dir with data
		fmt.Println(filepath.Join(cur, "data"))
		fmt.Println(filepath.Join(cur, "test", "app.log"))

	}
}
