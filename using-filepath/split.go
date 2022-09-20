// Go Playground
// https://go.dev/play/p/9nK9thcROtf

package main

import (
	"fmt"
	"path/filepath"
)

func main() {

	// Split from final seperator
	fmt.Println(filepath.Split("/tmp/test/app.log"))
	fmt.Println(filepath.Split("/tmp/test/"))
	fmt.Println(filepath.Split("/tmp/test"))

}
