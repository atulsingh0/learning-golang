package main

import (
	"fmt"
	"os"
)

func main() {

	f, err := os.OpenFile(os.TempDir()+"test.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	WriteFile(f, "Hello, World!\n")
	fmt.Println("File Name:", f.Name())

}

func WriteFile(f *os.File, msg string) {
	f.Write([]byte(msg))
}
