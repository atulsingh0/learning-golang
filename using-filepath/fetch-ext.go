package main

import (
	"fmt"
	"path/filepath"
)

func main() {

	// Fetch the file extensions
	fmt.Println(filepath.Ext("test.txt"))
	fmt.Println(filepath.Ext("test"))
	fmt.Println(filepath.Ext("test.go"))

}
