package main

import "log"

func main() {

	log.Println("Welcome to the playground")

	log.SetFlags(log.Ldate)
	log.Println("Welcome to the playground")

	log.SetFlags(log.Ldate | log.Ltime)
	log.Println("Welcome to the playground")

	log.SetFlags(log.LstdFlags)
	log.Println("Welcome to the playground")

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("Welcome to the playground")

	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	log.Println("Welcome to the playground")

	// most common flags, if accuracy is not needed for log time
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Welcome to the playground")
}
