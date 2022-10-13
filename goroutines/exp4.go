package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// For loop
	start := time.Now()
	for i := 1; i <= 10; i++ {
		getSquare(i)
	}
	fmt.Println("\nTotal Time Taken is standard processing:", time.Since(start))

	start = time.Now()
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			getSquare(x)
		}(i)
	}
	wg.Wait()
	fmt.Println("\nTotal Time Taken is standard processing:", time.Since(start))

}

func getSquare(n int) {
	time.Sleep(time.Second * 1)
	fmt.Print(n*n, " ")
}
