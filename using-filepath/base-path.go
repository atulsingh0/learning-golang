package main

import (
	"fmt"
	"path/filepath"
)

func main() {

	// fetch base path from path string
	fmt.Println("Base path of /tmp/test1/d2 is:", filepath.Base("/tmp/test1/d2"))
	fmt.Println("Base path of /tmp/test1/success.log:is ", filepath.Base("/tmp/test1/success.log"))

}
