package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	c := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	// Process the channel while the goroutine is running
	for d := range c {
		fmt.Println(d)
	}
	// Wait for goroutine to finish
	//(though it's already done by this point)
	wg.Wait()
}
