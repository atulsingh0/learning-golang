package main

import (
	"fmt"
	"path/filepath"
)

func main() {

	// Fetch the Dir
	fmt.Println(filepath.Dir("/tmp/data/test/"))
	fmt.Println(filepath.Dir("/tmp/data/err.log"))

}
