package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()
	for i := 0; i < 20; i++ {
		somework(i, 10) // it will seq run
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func somework(no int, t int) {
	start := time.Now()
	time.Sleep(time.Duration(t) * 100 * time.Millisecond)
	fmt.Printf("Worker: %d, TIme Taken: %v\n", no, time.Since(start))
}
