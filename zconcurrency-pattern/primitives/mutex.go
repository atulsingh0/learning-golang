package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var mu sync.Mutex

func main() {

	amount := 0
	start := time.Now()
	for i := 0; i < 500000; i++ {
		go update(&amount, 1)
	}

	time.Sleep(1 * time.Second)
	fmt.Printf("%d, %v\n", amount, time.Since(start))

	amount = 0
	start = time.Now()
	for i := 0; i < 500000; i++ {
		go updateWithMutex(&amount, 1)
	}

	// withMutex, 1 sec limit is not enough
	time.Sleep(1 * time.Second)
	fmt.Printf("%d, %v\n", amount, time.Since(start))

}

func update(amount *int, delta int) {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	*amount += delta // this will create race condition
}

func updateWithMutex(amount *int, delta int) {
	// The mutex ensures only one goroutine can access amount at a time
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	mu.Lock()
	defer mu.Unlock()
	*amount += delta
}
