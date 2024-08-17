package main

import (
	"log"
	"os"
)

var (
	info  *log.Logger
	err   *log.Logger
	debug *log.Logger
)

func main() {
	info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	err = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	debug = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)

	info.Println("Welcome to the playground")
	err.Println("Welcome to the playground")
	debug.Println("Welcome to the playground")

	// conditional logging
	info.Println("Writing some logs in standard format")
	if isProd() {
		debug.Println("Some additional info abt the error")
	}

}

// isProd returns true if the app is running in production
func isProd() bool {
	return os.Getenv("ENV") == "prod"
}
