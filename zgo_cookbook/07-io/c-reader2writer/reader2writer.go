package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

var url string = "https://sample-videos.com/text/Sample-text-file-1000kb.txt"
var filename string = "1000kb.txt"

func timer() func() {
	now := time.Now()
	return func() {
		fmt.Println("Total time to run:", time.Since(now))
	}
}

func main() {
	readWrite(url)
	readWrite2(url)
	readWrite3(url)
}

func readWrite(url string) {
	defer timer()()
	// Open the url
	r, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	// read the data
	data, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	// Create a file
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		panic(err)
	}
}

func readWrite2(url string) {
	defer timer()()
	// Open the url
	r, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	fh, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(fh, r.Body) // use io.Copy to copy from r.Body to fh
	if err != nil {
		panic(err)
	}
	defer fh.Close()
}

func readWrite3(url string) {
	defer timer()()
	// Open the url
	r, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	fh, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	buf := bufio.NewWriter(fh)    // writing into buffer speeds up the writing process
	_, err = io.Copy(buf, r.Body) // use io.Copy to copy from r.Body to fh
	if err != nil {
		panic(err)
	}
	defer fh.Close()
}
