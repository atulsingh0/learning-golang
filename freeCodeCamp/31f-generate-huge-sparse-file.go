package main

import (
	"fmt"
	"log"
	"os"
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
)

func main() {
	f, err := os.Create("foo.bar")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Generating Sparse file
	if err := f.Truncate(1e7); err != nil {
		log.Fatal(err)
	}

	fh, err := os.Stat("foo.bar")
	fmt.Printf("File Size (MB): %0.2f\n", float64(fh.Size())/MB)

}
