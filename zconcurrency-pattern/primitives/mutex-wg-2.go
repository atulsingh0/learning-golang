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

	nums := []int{3, 9, 7, 1, 4, 7, 5, 2, 6, 8}
	res := []int{}

	start := time.Now()
	for _, num := range nums {
		wg.Add(1)
		go processData(&wg, num, &res)
	}
	wg.Wait()
	fmt.Println("Without Mutex:", res, time.Since(start)) // upon multiple run, we will see diff results

	// another  for
	res = []int{} //using Mutex

	start = time.Now()
	for _, num := range nums {
		wg.Add(1)
		go processDataMutex(&wg, num, &res)
	}
	wg.Wait()
	fmt.Println("With Mutex:", res, time.Since(start))

	// avoiding Mutex
	res = make([]int, len(nums))
	start = time.Now()
	for i, num := range nums {
		wg.Add(1)
		go avoidMutexProcess(&wg, num, &res[i])
	}
	wg.Wait()
	fmt.Println("Avoid Mutex:", res, time.Since(start))

}

func processData(wg *sync.WaitGroup, data int, result *[]int) {
	defer wg.Done()
	*result = append(*result, data*2)
}

func processDataMutex(wg *sync.WaitGroup, data int, result *[]int) {
	defer wg.Done()
	mu.Lock()
	defer mu.Unlock()
	*result = append(*result, data*2)
}

func avoidMutexProcess(wg *sync.WaitGroup, data int, resultId *int) {
	defer wg.Done()
	*resultId = data * 2
}
