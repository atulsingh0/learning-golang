package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	/*
		INFO
		WARNING
		ERROR
		FATAL
	*/

	fmt.Println("Logging")
	log.Println("This is informational log")

	// Wrting into a log file
	createLogFile("learning.log")
	log.Println("this is my log msg")
	log.Fatal("Ah! I am hungry")

}

func createLogFile(fileName string) {
	file, err := os.Create(fileName)

	if err != nil {
		log.Fatal("Unable to create log file.", err)
	}

	log.SetOutput(file)
	file.Chmod(0700) // changing permission to 700
}
