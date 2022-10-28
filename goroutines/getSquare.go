package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()
	for i := 0; i < 10; i++ {
		getSquare(i) // it will not print any output
	}
	fmt.Println("Total Time Taken (if process in seq):", time.Since(start))

	// Processing via Go Routine
	start = time.Now()

	for i := 0; i < 10; i++ {
		go getSquare(i) // it will not print any output
	}
	fmt.Println("Total Time Taken (if process via go routines):", time.Since(start))
	time.Sleep(1 * time.Second)

}

func getSquare(no int) int {
	time.Sleep(1 * time.Second)
	return no * no
}
