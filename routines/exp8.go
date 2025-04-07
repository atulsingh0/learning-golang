package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// making it async
	// Adding a waitgroup
	start := time.Now()

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		echo("hello") // running as routine
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		echo("world") // running as routine
	}()

	wg.Wait()

	fmt.Println(time.Since(start))

}

func echo(data string) {
	for i := 0; i < 30; i++ {
		fmt.Println(data, i)
	}
}
