package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()
	//// seq call
	//echo("hi")
	//echo("bye")

	// making it async
	go echo("hello") // running as routine
	echo("world")

	fmt.Println(time.Since(start))
}

func echo(data string) {
	for i := 0; i < 50; i++ {
		fmt.Println(data, i)
	}
}
