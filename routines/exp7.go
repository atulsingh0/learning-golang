package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()

	// making it async
	go echo("hello") // running as routine
	go echo("world") // this 01 also running as routine

	// this does not print any result
	// as main exit first, so child routines are killed.
	fmt.Println(time.Since(start))
}

func echo(data string) {
	for i := 0; i < 30; i++ {
		fmt.Println(data, i)
	}
}
