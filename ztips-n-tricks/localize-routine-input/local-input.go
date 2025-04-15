package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// write to standard output
			time.Sleep(1 * time.Second)
			fmt.Printf("Routine: %d\n", i)

		}()
	}
	wg.Wait()
}
