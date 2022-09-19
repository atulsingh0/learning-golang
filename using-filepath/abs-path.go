package main

import (
	"fmt"
	"path/filepath"
)

func main() {

	// Fetch Absolute path by adding current working directory where program is running
	path, _ := filepath.Abs("test")
	fmt.Println("Absolute path:", path)
}
