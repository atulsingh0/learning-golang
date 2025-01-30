package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

var fileName = "1000kb.txt"

func timer() func() {
	now := time.Now()
	return func() {
		fmt.Println("Total time to run:", time.Since(now))
	}
}

func main() {

	_ = readFile(fileName)
	_ = readFileChunked10kb(fileName)
	_ = readFileChunked(fileName)
}

func readFile(filename string) string {
	defer timer()()
	data, err := os.ReadFile(filename) // read everything in one go
	if err != nil {
		panic(err)
	}
	return string(data)
}

func readFileChunked10kb(filename string) string {
	defer timer()()
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var data string
	buf := make([]byte, 10240) // 10kb chunks
	for {
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		data += string(buf[:n])
	}
	return data
}

func readFileChunked(filename string) string {
	defer timer()()
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fs, err := f.Stat()
	if err != nil {
		panic(err)
	}
	buf := make([]byte, fs.Size()) // create a buffer of the same size as the file
	// we dont need a for loop here because we are reading the entire file in one go
	n, err := f.Read(buf)
	if err != nil {
		panic(err)
	}
	return string(buf[:n])
}
