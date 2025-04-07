package main

import (
	"fmt"
)

func main() {

	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c) // will prevent the deadlock
	}()

	// if we don't close the channel,
	// the program will stick in deadlock
	// because the channel will be waiting for the data
	for d := range c {
		fmt.Println(d * d)
	}
}
