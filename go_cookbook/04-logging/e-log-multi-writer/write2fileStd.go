package main

import (
	"io"
	"log"
	"os"
)

func main() {

	// Opening a log file
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	writer := io.MultiWriter(f, os.Stdout)
	log.SetOutput(writer)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)

	log.Println("Welcome to the playground")
	log.Println("Writing 2 places")

}
