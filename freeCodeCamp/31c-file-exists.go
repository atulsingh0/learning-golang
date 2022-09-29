package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {

	filename := "z_test.log"

	_, err := os.Stat(filename)

	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("File does not exist.")
	}

	fh, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer fh.Close()

	fh.WriteString("testing")

}
