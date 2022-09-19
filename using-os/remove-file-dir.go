// Go Playground
// https://go.dev/play/p/X-u3JgPfW44

package main

import (
	"fmt"
	"os"
)

func main() {

	// Removing file or dir (empty)
	if err := os.Remove("/tmp/training"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Dir/File has been deleted.")
	}

	// Remove dir with files
	// Return No err if file or Dir does not exists
	if err := os.RemoveAll("/tmp/testing"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Dir/File has been deleted.")
	}

}
