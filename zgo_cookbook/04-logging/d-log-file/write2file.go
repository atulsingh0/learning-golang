package main

import (
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

	// Set output of logs to the file
	log.SetOutput(f)
	// Set flags for log
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)

	log.Println("Welcome to the playground")

}
