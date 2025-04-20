package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var mu sync.Mutex

func main() {
	// waitGroup helps to proper goroutine execution
	var wg sync.WaitGroup
	// Initialize the random number generator
	rand.New(rand.NewSource(time.Now().UnixNano()))

	amount := 0
	start := time.Now()
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go update(&amount, 1, &wg)
	} // total should be 10000
	wg.Wait()
	fmt.Printf("%d, %v\n", amount, time.Since(start))

	amount = 0
	start = time.Now()
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go updateWithMutex(&amount, 1, &wg)
	}
	wg.Wait()
	fmt.Printf("%d, %v\n", amount, time.Since(start))

}

func update(amount *int, delta int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	*amount += delta // this will create race condition

}

func updateWithMutex(amount *int, delta int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	// mutex will make sure only one routine access the amount data
	mu.Lock()
	defer mu.Unlock()
	*amount += delta
}
