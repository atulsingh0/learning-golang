package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// creating a file
	fh, err := os.Create("data-temp.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	// defer the close when work is done (execute just before main exit)
	defer fh.Close()

	names := []string{"A", "B", "C", "E", "F", "G", "H"}

	// wrtiting to file
	for _, data := range names {
		fh.WriteString(data + "\n")
	}

	// Reading data in one go
	fi, _ := os.ReadFile("data-temp.txt")
	defer fh.Close()

	fmt.Println(string(fi))
}
