package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan struct{}, 3) // limit 3 routine at a time
	start := time.Now()
	for i := 0; i < 40; i++ {
		ch <- struct{}{}
		j := i
		go somework(j, 10, ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func somework(no int, t int, ch chan struct{}) {
	start := time.Now()
	time.Sleep(time.Duration(t) * 100 * time.Millisecond)
	<-ch
	fmt.Printf("Worker: %d, TIme Taken: %v\n", no, time.Since(start))
}
