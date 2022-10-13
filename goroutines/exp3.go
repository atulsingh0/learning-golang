package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			x := i
			time.Sleep(time.Second * 1)
			fmt.Println("Hi", i, x)
			wg.Done()
		}()
	}
	wg.Wait()
	// time.Sleep(time.Second * 2)

}
