package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()
	for i := 0; i < 10; i++ {
		go func() {
			x := i
			time.Sleep(time.Second * 1)
			fmt.Println("Hi", i, x)
		}()
	}
	fmt.Println("Total Time:", time.Since(start))
	// Waiting for output to print
	time.Sleep(time.Second * 2)

}
